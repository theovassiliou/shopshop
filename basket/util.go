package shopshop

import (
	log "github.com/sirupsen/logrus"
)

// AssertNoErrorFatal fails on error
func AssertNoErrorFatal(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// AssertNoError returns false on error, else true
func AssertNoError(err error) bool {
	if err != nil {
		return false
	}
	return true
}
