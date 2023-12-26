# Camera Services

摄像头服务，依赖opencv,用于实时查看usb摄像头画面
支持录像（可配置为检测到人脸时或检测到运动时录像）

## 已测试系统 
ubuntu 22.04  
## 使用方法  
把bin目录放到系统任意目录下，
给 camera_service_linux 和 regAndStartCameraService.sh 可执行权限

chmod +x ./camera_service_linux  
chmod +x ./regAndStartCameraService.sh 
然后运行 regAndStartCameraService.sh 

## 配置
按需要修改 conf/app.conf 文件

## 实时在线查看
  
http://your_ip:8080/console  





