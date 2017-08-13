package poslog

import (
	"path/filepath"
	"testing"
)

func TestAppendPOSLog(t *testing.T) {
	filename := "./input/xml/POSLog-201612301530-54.xml"
	p := ImportXML(filename)
	if len(p.DayID.DayID) != 8 {
		t.Fail()
	}
	if filepath.Base(filename) != p.Filename {
		t.Fail()
	}
}
