package main

import (
	"testing"
)

func Test_fillBlob(t *testing.T) {
	rnd, err := fillBlob(1)
	if err != nil {
		t.Errorf("got error %v",err)
	}

	if len(rnd) > 4 || len (rnd) < 1 {
		t.Errorf("wrong array size %v",len(rnd))
	}

}
