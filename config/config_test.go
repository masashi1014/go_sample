package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	conf, err := NewConfig()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if conf.DB.User == "" {
		t.Fatal("no user")
	}

	if conf.DB.Password == "" {
		t.Fatal("no password")
	}

	if conf.DB.Host == "" {
		t.Fatal("no host")
	}

	if conf.DB.Port == 0 {
		t.Fatal("no port")
	}

	if conf.DB.Db == "" {
		t.Fatal("no db")
	}
}
