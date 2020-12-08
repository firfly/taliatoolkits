package main

import (
	"go.uber.org/zap"
	"github.com/firfly/taliatoolkits/logging"
	"github.com/firfly/taliatoolkits/mysql_handle"
)

const LogFileFullPath = "/tmp/migrator.log"
var log *zap.SugaredLogger

func init()  {
	logging.InitLogger(LogFileFullPath)
	log = logging.GetDefaultLogger()
}


func main() {

	log.Info("this is a info:::", LogFileFullPath)
	mysql_handle.MysqlHandleTest("str from main.....")

}
