package test

import (
	"strings"
	"testing"
)

func TestResultConstructor(t *testing.T) {
	m := make(map[string]int)
	m["oks_triggered"] = 1
	m["warns_triggered"] = 2
	m["crits_triggered"] = 0

	r := NewResult(m)

	if r.Ok != 1 {
		t.Error("Ok should be initialized with value 1")
	}
	if r.Warn != 2 {
		t.Error("Warn should be initialized with value 1")
	}
	if r.Crit != 0 {
		t.Error("Crit should be initialized with value 1")
	}
	if r.Error != false {
		t.Error("Error should be initialized with value false")
	}
}

func TestResultCompareOk(t *testing.T) {
	m1 := make(map[string]int)
	m1["oks_triggered"] = 1
	m1["warns_triggered"] = 2
	m1["crits_triggered"] = 0
	r1 := NewResult(m1)

	m2 := make(map[string]int)
	m2["oks_triggered"] = 1
	m2["warns_triggered"] = 2
	m2["crits_triggered"] = 0
	r2 := NewResult(m2)

	r1.Compare(r2)

	if r1.Message != "OK" {
		t.Error("Comparison message should be OK")
	}

	if r1.Passed != true {
		t.Error("Comparison result should be true")
	}

}

func TestResultCompareNOk(t *testing.T) {
	m1 := make(map[string]int)
	m1["oks_triggered"] = 2
	m1["warns_triggered"] = 2
	m1["crits_triggered"] = 0
	r1 := NewResult(m1)

	m2 := make(map[string]int)
	m2["oks_triggered"] = 1
	m2["warns_triggered"] = 2
	m2["crits_triggered"] = 0
	r2 := NewResult(m2)

	s := "FAIL\nShould have triggered 1 Ok alerts, triggered 2\nAlerts triggered (ok: 2, warn: 2, crit: 0)\n"

	r1.Compare(r2)

	if strings.Compare(r1.Message, s) != 0 {
		t.Error(r1.Message)
		t.Error(s)
	}

	if r1.Passed != false {
		t.Error("Comparison result should be false")
	}
}
