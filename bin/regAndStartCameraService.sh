#!/usr/bin/env sh
## 

#当前脚本所在目录
workDir=$(cd `dirname $0`; pwd)
#cd $workDir

chmod +x ${workDir}/camera_service_linux
sed -r -i 's#^WorkingDirectory=.*$#WorkingDirectory='${workDir}'#' ${workDir}/camera.service
sed -r -i 's#^ExecStart=.*$#ExecStart='${workDir}'\/camera_service_linux#' ${workDir}/camera.service


cp -f ${workDir}/camera.service /lib/systemd/system/
systemctl enable camera.service && systemctl restart camera.service

echo "success"