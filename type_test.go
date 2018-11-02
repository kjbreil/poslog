package poslog

import "testing"

// func TestReadPOSLog(t *testing.T) {
// 	filename := "./input/xml/POSLog-201612301530-54.xml"
// 	p := ReadXML(filename)
// }

func TestPOSLog(t *testing.T) {
	// data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	data, _ := Read("./sample/POSLog.xml")
	data.WriteXML("./output/TestOut.xml")
}
