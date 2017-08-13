package poslog

import (
	"testing"
)

func TestWriteXML(t *testing.T) {
	data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	WriteXML("./output/TestOut.xml", data)
}

func TestWriteJSON(t *testing.T) {
	data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	WriteJSON("./output/TestOut.kadjasljdla", data)
}
