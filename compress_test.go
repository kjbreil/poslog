package poslog

import (
	"os"
	"testing"
)

func TestReadAllFromZip(t *testing.T) {
	ps := zipReadAllXML("./input/zip/POSLog-20161230.zip")

	// for _, p := range ps.POSLogs {
	// 	fmt.Println(p)
	// }
	os.Mkdir("output/test", 0777)
	WriteJSONs("output/test", ps)

}
