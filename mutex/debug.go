package mutex

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type debug struct {
	timing   int64
	dir      string
	filename string
	line     int
}

func (d *debug) set() {
	d.timing = time.Now().UnixNano()
	if _, file, line, ok := runtime.Caller(2); ok {
		dir, filename := path.Split(file)
		dir = dir[strings.LastIndex(dir[:strings.LastIndex(dir, "/")], "/")+1:]
		d.dir = dir
		d.filename = filename
		d.line = line
	} else {
		d.dir = "???"
		d.filename = "???"
		d.line = 0
	}
}

func (d *debug) reset() {
	d.timing = 0
	d.dir = ""
	d.filename = ""
	d.line = -1
}

func (d *debug) Duration() int64 {
	if d.timing != 0 {
		return time.Now().UnixNano() - d.timing
	}
	return 0
}

func (d *debug) DurationSeconds() int64 {
	if d.timing != 0 {
		return (time.Now().UnixNano() - d.timing) / 1e9
	}
	return 0
}

func (d *debug) Line() int {
	return d.line
}

func (d *debug) Path() string {
	return d.dir + d.filename
}

func (d *debug) debugString() string {
	return fmt.Sprintf("Line:%v Path:%v Duration:%v", d.Line(), d.Path(), d.Duration())
}
