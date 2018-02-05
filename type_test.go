package poslog

import (
	"testing"
)

func TestReadPOSLog(t *testing.T) {
	filename := "./input/xml/POSLog-201612301530-54.xml"
	p := ImportXML(filename)
}
