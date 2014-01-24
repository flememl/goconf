package conf

import (
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	oldEnv := os.Getenv("GO_CONFIG_PATH")
	_, filename, _ := runtime.Caller(1)
	os.Setenv("GO_CONFIG_PATH", path.Join(path.Dir(filename), "config_test.json"))
	defer os.Setenv("GO_CONFIG_PATH", oldEnv)

	if Config.Database.Port != "3306" {
		t.Fatal("Configuration database port doesn't match", Config.Database.Port, "!= 3306")
	}
	if Config.Database.Host != "127.0.0.1" {
		t.Fatal("Configuration database host doesn't match", Config.Database.Host, "!= 127.0.0.1")
	}
}
