package config

import (
	"path"
	"runtime"
	"testing"
)

func init() {
	NewConfig()
}

func TestRootPath(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(filename)))
	t.Error(rootPath)
}

func TestConfig(t *testing.T) {
	t.Log(conf.Get("disease"))
}
