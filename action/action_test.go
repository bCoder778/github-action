package action

import (
	"github.com/bCoder778/log"
	"testing"
)

func TestA(t *testing.T) {
	if A(){
		log.Failf("failed")
	}
}