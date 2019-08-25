package goal

import (
	"log"
	"testing"
)

type XT struct {
	*testing.T
}

func (t *XT) assertNotNull(who interface{}, err error) {
	if err != nil {
		log.Println(err)
	}

	if who == nil {
		t.FailNow()
	}
}


func (t *XT) assertNull(who interface{}, err error) {
	if err != nil {
		log.Println(err)
	}

	if who != nil {
		t.FailNow()
	}
}