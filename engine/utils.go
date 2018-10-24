package engine

import "os"

// GetKey - returns the value from the env var APP_KEY
func GetKey() string {
	return os.Getenv("APP_KEY")
}

// GetName - returns the value from the env var APP_NAME
func GetName() string {
	return os.Getenv("APP_NAME")
}

// hasKey - checks that the application has the key env var
func hasKey() bool {
	return GetKey() != ""
}

// hasName - checks that the application has the name env var
func hasName() bool {
	return GetName() != ""
}

// IsDebug - checks that the application is in debug mode
func IsDebug() bool {
	return os.Getenv("APP_DEBUG") != ""
}
