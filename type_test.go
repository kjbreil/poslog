package poslog

import "testing"

// func TestReadPOSLog(t *testing.T) {
// 	filename := "./input/xml/POSLog-201612301530-54.xml"
// 	p := ReadXML(filename)
// }

func TestPOSLog(t *testing.T) {
	// data := ImportXML("./input/xml/POSLog-201612301530-54.xml")
	data, err := Read("./sample/POSLog.xml")
	if err != nil {
		t.Fatal(err)
	}
	data.WriteXML("./output/TestOut.xml")
	// data, _ := ZipReadAllXML("./sample/matching1.zip")
	// t.Log(len(data))
	// t.Fail()
}
