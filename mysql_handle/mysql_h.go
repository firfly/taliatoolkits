package mysql_handle

import (
	"github.com/firfly/taliatoolkits/logging"
)


var log = logging.GetDefaultLogger()
func MysqlHandleTest(astr string)  {
	log.Warn("passing a str:::", astr)

}
