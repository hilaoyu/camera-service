package service

import (
	"fmt"
	"gitee.com/hilaoyu/go-basic-utils/utils"
	"gocv.io/x/gocv"
	"image"
	"path/filepath"
)

const MotionCheckMinimumArea = 3000

type Checker struct {
	MinimumArea        float64
	motionCheckImgPrev gocv.Mat
	motionCheckMog2    gocv.BackgroundSubtractorMOG2
	motionCheckKernel  gocv.Mat

	CheckFace        bool
	CheckMotion      bool
	TrueWithoutCheck bool
}

var (
	RecordChecker = NewChecker()
)

func NewChecker() *Checker {
	return &Checker{
		MinimumArea:        MotionCheckMinimumArea,
		motionCheckImgPrev: gocv.NewMat(),
		motionCheckMog2:    gocv.NewBackgroundSubtractorMOG2(),
		motionCheckKernel:  gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3)),
		CheckFace:          false,
		CheckMotion:        false,
		TrueWithoutCheck:   false,
	}
}
func (c *Checker) SetTrueWithoutCheck(v bool) *Checker {
	c.TrueWithoutCheck = v
	return c
}
func (c *Checker) SetCheckFace(v bool) *Checker {
	c.CheckFace = v
	return c
}

func (c *Checker) SetCheckMotion(v bool) *Checker {
	c.CheckMotion = v
	return c
}

func (c *Checker) Check(img gocv.Mat) bool {
	if c.TrueWithoutCheck {
		return true
	}

	if c.CheckFace && c.CheckHasFace(img) {
		//fmt.Println("checker face true")
		return true
	}

	if c.CheckMotion && c.CheckHasMotion(img) {
		//fmt.Println("checker Motion true")
		return true
	}

	return false
}
func (c *Checker) CheckHasFace(img gocv.Mat) bool {
	if !c.CheckFace {
		return c.TrueWithoutCheck
	}
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	xmlFile := filepath.Join(utils.GetSelfPath(), "frontalface.xml")
	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return false
	}
	rects := classifier.DetectMultiScale(img)
	//fmt.Sprintf("face rects %d \n", len(rects))
	return len(rects) > 0
}
func (c *Checker) CheckHasMotion(img gocv.Mat) bool {
	if !c.CheckMotion {
		return c.TrueWithoutCheck
	}

	if c.motionCheckImgPrev.Empty() {
		c.motionCheckMog2.Apply(img, &c.motionCheckImgPrev)
	}
	//godump.Dump(motionCheckImgPrev)
	//godump.Dump(img)
	imgThresh := gocv.NewMat()
	defer imgThresh.Close()
	c.motionCheckMog2.Apply(img, &c.motionCheckImgPrev)

	gocv.Threshold(c.motionCheckImgPrev, &imgThresh, 25, 255, gocv.ThresholdBinary)

	// then dilate
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	gocv.Dilate(imgThresh, &imgThresh, kernel)
	kernel.Close()

	// now find contours
	contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()
	//fmt.Println(fmt.Sprintf("motion contours %d", contours.Size()))

	for i := 0; i < contours.Size(); i++ {
		area := gocv.ContourArea(contours.At(i))
		if area > c.MinimumArea {
			fmt.Println(fmt.Sprintf("motion find %f ", area))
			return true
		}
	}

	return false
}

func (c *Checker) Close() {
	c.motionCheckImgPrev.Close()
	c.motionCheckMog2.Close()
	c.motionCheckKernel.Close()
}
