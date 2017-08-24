package test

type Result struct {
	Ok   string //`json:"ok,string,omitempty"`
	Warn string //`json:"warn,string,omitempty"`
	Crit string //`json:"crit,string,omitempty"`
}
