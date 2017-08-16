package poslog

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	p.appendFilename(filepath.Base(filename))
	// p.appendDayID()

	return
}

func importReaderXML(f io.Reader, filename string) (p POSLog) {
	byteXML, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("FUCK", err)
	}
	xml.Unmarshal(byteXML, &p)
	p.appendFilename(filepath.Base(filename))
	// p.appendDayID()
	return
}

// WriteXML writes a POSLog object to an XML file given as first argument
func WriteXML(filename string, p POSLog) {
	// Drop any other extension and stick a xml on there
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	filename = filename + ".xml"

	ioutil.WriteFile(filename, createXML(p), 0666)
}

func createXML(p POSLog) []byte {
	xs, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return xs
}

// WriteJSON writes a POSLog object to an json file given as first argument
func WriteJSON(filename string, p POSLog) {
	// Drop any other extension and stick a json on there
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	filename = filename + ".json"
	ioutil.WriteFile(filename, createJSON(p), 0666)
}

func createJSON(p POSLog) []byte {
	js, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return js
}

func WriteJSONs(folder string, ps POSLogs) {
	os.Mkdir(folder, 0777)
	for _, p := range ps.POSLogs {
		op := filepath.Join(folder, p.filename)
		WriteJSON(op, p)
	}
}
