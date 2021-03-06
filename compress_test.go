package poslog

import (
	"io"
	"os"
	"testing"
)

// func TestReadAllFromZip(t *testing.T) {
// 	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")

// 	// for _, p := range ps.POSLogs {
// 	// 	fmt.Println(p)
// 	// }
// 	os.Mkdir("output/test", 0777)
// 	WriteJSONs("output/test", ps)

// }

func TestTarJSONs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/json_test.tar")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	tarJSONs(writer, ps)

}
func TestGzipJSONs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/json_test.tar.gz")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	GzipJSONs(writer, ps)

}

func TestTarXMLs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/xml_test.tar")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	tarXMLs(writer, ps)

}

func TestGzipXMLs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/xml_test.tar.gz")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	gzipXMLs(writer, ps)

}

func TestXzXMLs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/xml_test.tar.xz")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	xzXMLs(writer, ps)

}

func TestSnappyXMLs(t *testing.T) {
	ps := ZipReadAllXML("./input/zip/POSLog-20161230.zip")
	file, err := os.Create("./output/xml_test.snappy")
	if err != nil {
		t.Fail()
	}
	writer := io.Writer(file)

	snappyXMLs(writer, ps)

}
