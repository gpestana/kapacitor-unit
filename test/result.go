package test

import (
	"fmt"
	//"github.com/golang/glog"
	"strings"
)

type Result struct {
	Ok      int
	Warn    int
	Crit    int
	Message string
	Passed  bool
}

func NewResult(r map[string]int) Result {
	rf := new(Result)
	rf.Ok = r["oks_triggered"]
	rf.Warn = r["warns_triggered"]
	rf.Crit = r["crits_triggered"]
	return *rf
}

func (r *Result) Compare(r2 Result) {
	if *r == r2 {
		r.Passed = true
		r.Message = "Test passed"
	} else {
		r.Passed = false
		m := errorMessage(r2, *r)
		r.Message = m
	}
}

func errorMessage(rexp Result, r Result) string {
	s := []string{"Test Failed:"}
	if rexp.Ok != r.Ok {
		s = append(s, fmt.Sprintf("Should have triggered %v Ok alerts, triggered %v", rexp.Ok, r.Ok))
	}
	if rexp.Warn != r.Warn {
		s = append(s, fmt.Sprintf("Should have triggered %v Warning alerts, triggered %v", rexp.Warn, r.Warn))
	}
	if rexp.Crit != r.Crit {
		s = append(s, fmt.Sprintf("Should have triggered %v Critical alerts, triggered %v", rexp.Crit, r.Crit))
	}

	return strings.Join(s, " ")
}
