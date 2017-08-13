package poslog

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// ImportXML takes a POSLog XML file as the argument and returns
// a POSLog
func ImportXML(filename string) (p POSLog) {
	byteXML, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("FUCK", err)
	}

	xml.Unmarshal(byteXML, &p)

	return
}

func importReaderXML(f io.ReadCloser) (p POSLog) {
	byteXML, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("FUCK", err)
	}
	xml.Unmarshal(byteXML, &p)
	return
}

// WriteXML writes a POSLog object to an XML file given as first argument
func WriteXML(fn string, p POSLog) {
	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
	fn = fn + ".xml"
	XMLString, err := xml.MarshalIndent(p, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fn, XMLString, 0666)
}

// WriteJSON writes a POSLog object to an json file given as first argument
func WriteJSON(fn string, p POSLog) {
	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
	fn = fn + ".json"
	XMLString, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fn, XMLString, 0666)
}
