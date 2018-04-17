package commands

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	//"testing"
	//"github.com/stretchr/testify/assert"
)

// Test proper UI serve
func TestCmd_Serve(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)
	/*
		if assert.Error(t, result) {
			assert.Contains(t, result.Error(), "is invalid against swagger specification 2.0")
			assert.Contains(t, result.Error(), "definitions.RRSets in body must be of type array")
		}
	*/
}
