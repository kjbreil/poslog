package poslog

import (
	"os"
	"testing"
)

func TestReadAllFromZip(t *testing.T) {
	ps := zipReadAllXML("./input/zip/POSLog-20161230.zip")

	os.Mkdir("output/test", 0777)
	for _, p := range ps.POSLogs {

		WriteJSON("output/test", p)
	}

}
