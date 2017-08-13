package poslog

import (
	"testing"
)

func TestWriteXML(t *testing.T) {
	data := importXML("./input/xml/POSLog-201612301530-54.xml")
	writeXML(data, "./output/TestOut.xml")
}

func TestWriteJSON(t *testing.T) {
	data := importXML("./input/xml/POSLog-201612301530-54.xml")
	writeJSON(data, "./output/TestOut.json")
}
