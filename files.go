package poslog

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func importXML(filename string) (p POSLog) {
	byteXML, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("FUCK", err)
	}

	xml.Unmarshal(byteXML, &p)

	return
}

func writeXML(p POSLog, fn string) {
	XMLString, err := xml.MarshalIndent(p, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fn, XMLString, 0666)
}

func writeJSON(p POSLog, fn string) {
	XMLString, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fn, XMLString, 0666)
}
