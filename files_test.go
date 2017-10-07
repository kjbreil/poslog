package poslog

import (
	"testing"
)

func TestWriteXML(t *testing.T) {
	// data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	data := ImportXML("./input/xml/POSLog.xml")
	data.WriteXML("./output/TestOut.xml")
}

func TestWriteJSON(t *testing.T) {
	data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	WriteJSON("./output/TestOut", data)
}
