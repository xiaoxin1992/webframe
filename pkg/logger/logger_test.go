package logger

import (
	"testing"
	"webframe/conf"
)

func TestName(t *testing.T) {
	err := conf.NewConfig("etc/config.toml")
	if err != nil {
		t.Fatalf("init config error %s", err.Error())
	}
	NewLogger("/tmp/")
	//for {
	//	Logger.S("test").Info("this is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test logger")
	//}
	Logger.S("test").Info("this is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test loggerthis is test logger")

	//fmt.Println(time.Format("2006-01-02 15:04:05"))
	Logger.Sync()
}
