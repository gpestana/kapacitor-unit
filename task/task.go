package task

import ()

//FS configurations, namely path where TICKscripts are located
type Task struct {
	Path string
}

func (task Task) Load() []string {
	return make([]string, 10)
}
