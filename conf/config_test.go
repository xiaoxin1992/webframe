package conf

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	err := NewConfig("etc/config.toml")
	if err != nil {
		t.Fatalf("init config error %s", err.Error())
	}
	fmt.Println(GetConfig().App)
}
