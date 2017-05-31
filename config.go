package main

import "os"

var port string
var defaultContentType string
var overrideContentType string

// Gets the variable from the environment. `def` is the default value
// that gets used if no env is found with that name.
func getenv(varName, def string) string {
	if newVar := os.Getenv(varName); newVar != "" {
		return newVar
	}
	return def
}

func init() {
	port = getenv("PORT", "80")
	defaultContentType = getenv("DEFAULT_CONTENT_TYPE", "text/plain")
	overrideContentType = getenv("FORCED_CONTENT_TYPE", "")
}
