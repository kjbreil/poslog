package poslog

import (
	"testing"
)

func TestWriteXML(t *testing.T) {
	// data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	data := ReadXML("./sample/POSLog.xml")
	data.WriteXML("./output/TestOut.xml")
}
