package weather

import "errors"

var errorNoCityFound = errors.New("no match found")
var invalidAPIKey = errors.New("Invalid API key.")
