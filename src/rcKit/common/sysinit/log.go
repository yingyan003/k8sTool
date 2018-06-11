package sysinit

import (
	mylog "github.com/maxwell92/gokits/log"
	"scaler/common/constant"
	"scaler/common/utils"
)

func InitLog() {
	logLevel := utils.LoadEnvVar(constant.ENV_LOGLEVEL, constant.LOGLEVEL)
	mylog.Log.SetLevelByName(logLevel)
	mylog.Log.Infof("Log Level=%s", logLevel)
}
