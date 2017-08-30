package test

import (
	"github.com/gpestana/kapacitor-unit/io"
	"testing"
)

func TestValidateRecAndData(t *testing.T) {
	k := io.NewK("localhost")
	r := Result{}
	d := []string{"data1", "data2"}
	tst := Test{}

	tst.Data = d
	tst.RecId = "e24db07d-1646-4bb3-a445-828f5049bea0"
	tst.Result = r

	tst.validate(k)

	if tst.Result.Error != true {
		t.Error("Test initialized with recording_id and test must be invalid")
	}
}

func TestValidateRecOk(t *testing.T) {
	k := io.NewK("localhost")
	r := Result{}
	tst := Test{}

	tst.RecId = "e24db07d-1646-4bb3-a445-828f5049bea0"
	tst.Result = r

	tst.validate(k)

	if tst.Result.Error != false {
		t.Error("Test initialized only with recording_id must be valid")
	}
}

func TestValidateDataOk(t *testing.T) {
	k := io.NewK("localhost")
	r := Result{}
	tst := Test{}
	d := []string{"data1", "data2"}

	tst.Data = d
	tst.Result = r

	tst.validate(k)

	if tst.Result.Error != false {
		t.Error("Test initialized only with data must be valid")
	}
}
