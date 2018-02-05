package poslog

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// ReadXML takes a POSLog XML file as the argument and returns
// a POSLog object
func ReadXML(filename string) (p POSLog) {
	byteXML, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Problem reading XML File ")
		log.Println(err)
	}

	xml.Unmarshal(byteXML, &p)
	p.filename(filepath.Base(filename))

	return
}

func importReaderXML(f io.Reader, filename string) (p POSLog) {
	byteXML, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Problem reading XML File ")
		log.Println(err)
	}
	xml.Unmarshal(byteXML, &p)
	p.filename(filepath.Base(filename))

	return
}

// WriteXML writes a POSLog object to an XML file given as first argument
func (p *POSLog) WriteXML(filename string) {
	// Drop any other extension and stick a xml on there
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	filename = filename + ".xml"

	ioutil.WriteFile(filename, createXML(*p), 0666)
}

func createXML(p POSLog) []byte {
	xs, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return xs
}

// ZipReadAllXML Reads all XML from a passed archive
func ZipReadAllXML(archive string) (ps []POSLog) {

	reader, err := zip.OpenReader(archive)
	if err != nil {
		log.Println("Could not open the zip archive for some reason")
		log.Println(archive)
	}

	for _, file := range reader.File {
		ext := filepath.Ext(file.Name)

		if ext == ".xml" {
			fileReader, err := file.Open()
			if err != nil {
				fmt.Println("ERRUR")
			}
			defer fileReader.Close()

			p := importReaderXML(fileReader, file.Name)

			ps = append(ps, p)
		}
	}

	return
}
