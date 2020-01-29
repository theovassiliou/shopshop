package shopshop

import (
	log "github.com/sirupsen/logrus"
)

// AssertNoError fails on error
func AssertNoError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
