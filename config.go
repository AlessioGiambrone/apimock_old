package main

import "os"

// Gets the variable from the environment. `def` is the default value
// that gets used if no env is found with that name.
func getenv(varName, def string) string {
	if newVar := os.Getenv(varName); newVar != "" {
		return newVar
	}
	return def
}
