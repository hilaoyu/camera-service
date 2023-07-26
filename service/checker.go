package service

import (
	"fmt"
	"gitee.com/hilaoyu/go-basic-utils/utils"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"path/filepath"
	"time"
)

const MotionCheckMinimumArea = 3000

type Checker struct {
	MinimumArea        float64
	motionCheckImgPrev gocv.Mat
	imgBlackWhite      gocv.Mat
	motionCheckMog2    gocv.BackgroundSubtractorMOG2
	motionCheckKernel  gocv.Mat
	classifier         *gocv.CascadeClassifier

	CheckFace        bool
	CheckMotion      bool
	DrawMark         bool
	TrueWithoutCheck bool
}

var (
	RecordChecker = NewChecker()
)

func NewChecker() *Checker {

	checker := &Checker{
		MinimumArea:        MotionCheckMinimumArea,
		motionCheckImgPrev: gocv.NewMat(),
		imgBlackWhite:      gocv.NewMat(),
		motionCheckMog2:    gocv.NewBackgroundSubtractorMOG2(),
		motionCheckKernel:  gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3)),
		CheckFace:          false,
		CheckMotion:        false,
		DrawMark:           false,
		TrueWithoutCheck:   false,
	}
	classifier := gocv.NewCascadeClassifier()
	xmlFile := filepath.Join(utils.GetSelfPath(), "face-check.xml")
	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
	} else {
		checker.classifier = &classifier
	}

	return checker

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

func (c *Checker) SetDrawMark(v bool) *Checker {
	c.DrawMark = v
	return c
}

func (c *Checker) Check(img gocv.Mat) (checkPassed bool, imgRe gocv.Mat) {
	if c.TrueWithoutCheck {
		return true, img
	}

	if c.CheckFace {
		checkPassed, imgRe = c.CheckHasFace(img)
		if checkPassed {
			//fmt.Println("checker face true")
			return true, imgRe
		}

	}

	if c.CheckMotion {
		checkPassed, imgRe = c.CheckHasMotion(img)
		if checkPassed {
			//fmt.Println("checker Motion true")
			return true, imgRe
		}
	}

	return false, img
}
func (c *Checker) CheckHasFace(img gocv.Mat) (bool, gocv.Mat) {
	if !c.CheckFace {
		return c.TrueWithoutCheck, img
	}
	if nil == c.classifier {
		return false, img
	}

	nowName := time.Now().Format("2006-01-02_150405")
	rects := c.classifier.DetectMultiScale(img)
	//fmt.Sprintf("face rects %d \n", len(rects))
	blue := color.RGBA{0, 0, 255, 0}
	checkPassed := false
	if len(rects) > 0 {
		//
		if c.DrawMark {
			for _, r := range rects {
				gocv.Rectangle(&img, r, blue, 3)
			}
		}

		checkPassed = true
	}
	if checkPassed {
		fmt.Println("Face detected", nowName)
		if c.DrawMark {
			gocv.PutText(&img, "Face detected", image.Pt(10, 40), gocv.FontHersheyPlain, 1.2, blue, 2)
			gocv.IMWrite("face-"+nowName+".jpg", img)
		}

	}
	return checkPassed, img
}
func (c *Checker) CheckHasMotion(img gocv.Mat) (bool, gocv.Mat) {
	if !c.CheckMotion {
		return c.TrueWithoutCheck, img
	}

	gocv.CvtColor(img, &c.imgBlackWhite, gocv.ColorBGRAToGray)
	if c.motionCheckImgPrev.Empty() {
		c.motionCheckMog2.Apply(c.imgBlackWhite, &c.motionCheckImgPrev)
	}
	nowName := time.Now().Format("2006-01-02_150405")
	//godump.Dump(motionCheckImgPrev)
	//godump.Dump(img)
	imgThresh := gocv.NewMat()
	defer imgThresh.Close()
	c.motionCheckMog2.Apply(c.imgBlackWhite, &c.motionCheckImgPrev)

	gocv.Threshold(c.motionCheckImgPrev, &imgThresh, 50, 255, gocv.ThresholdBinary)

	// then dilate
	gocv.Dilate(imgThresh, &imgThresh, c.motionCheckKernel)

	// now find contours
	contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()
	//fmt.Println(fmt.Sprintf("motion contours %d", contours.Size()))

	checkPassed := false
	color1 := color.RGBA{255, 0, 0, 0}
	color2 := color.RGBA{0, 255, 0, 0}
	for i := 0; i < contours.Size(); i++ {
		area := gocv.ContourArea(contours.At(i))
		if area < c.MinimumArea {
			continue
		}

		if c.DrawMark {
			gocv.DrawContours(&img, contours, i, color1, 2)
			rect := gocv.BoundingRect(contours.At(i))
			gocv.Rectangle(&img, rect, color2, 2)
		}

		checkPassed = true
	}
	if checkPassed {
		fmt.Println("Motion detected", nowName)

		if c.DrawMark {
			gocv.PutText(&img, "Motion detected", image.Pt(10, 60), gocv.FontHersheyPlain, 1.2, color1, 2)
			gocv.IMWrite("motion-"+nowName+".jpg", img)
			gocv.IMWrite("motion-t-"+nowName+".jpg", imgThresh)
		}

	}
	return checkPassed, img
}

func (c *Checker) Close() {
	c.motionCheckImgPrev.Close()
	c.imgBlackWhite.Close()
	c.motionCheckMog2.Close()
	c.motionCheckKernel.Close()
	c.classifier.Close()
}
