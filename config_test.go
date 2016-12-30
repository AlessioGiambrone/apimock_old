package main

import (
	"os"
	"testing"
)

var envKey = "mykey"
var envVars = []string{"uno", "due", "tre"}
var defEnv = "mydefault"

func TestGetenv(t *testing.T) {
	err := os.Unsetenv(envKey)
	for _, e := range envVars {
		if err != nil {
			panic(err)
		}
		if getenv(envKey, defEnv) != defEnv {
			t.Errorf("should have used default: %v %v\n", e, getenv(envKey, defEnv))
		}
	}
	for _, e := range envVars {
		err := os.Setenv(envKey, e)
		if err != nil {
			panic(err)
		}
		if getenv(envKey, defEnv) != e {
			t.Errorf("env key mismatch: %v %v\n", e, getenv(envKey, defEnv))
		}
	}
}
