package poslog

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (p *POSLog) pAppend(filename string) {

	p.filename(filepath.Base(filename))
	p.buisnessDayDate()
	p.storeID()
	p.counts()
}

// Read takes a POSLog XML file as the argument and returns
// a POSLog object
func Read(filename string) (p POSLog, err error) {
	byteXML, err := ioutil.ReadFile(filename)
	if err != nil {
		return p, fmt.Errorf("problem reading XML File with error: %v", err)
	}
	err = xml.Unmarshal(byteXML, &p)

	if err != nil {
		return p, fmt.Errorf("XML unmarshal error: %v", err)
	}

	p.pAppend(filename)

	return
}

func importReaderXML(f io.Reader, filename string) (p POSLog) {
	byteXML, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Problem reading XML File ")
		log.Println(err)
	}
	xml.Unmarshal(byteXML, &p)

	p.pAppend(filename)

	return
}

// ZipReadAllXML Reads all XML from a passed archive
func ZipReadAllXML(archive string) (ps []POSLog, err error) {

	if _, err := os.Stat(archive); os.IsNotExist(err) {
		return nil, fmt.Errorf("Archive does not exist: %s - %v", archive, err)
	}

	reader, err := zip.OpenReader(archive)
	if err != nil {
		return nil, fmt.Errorf("Could not open the zip archive for some reason: %s - %v", archive, err)
	}

	for _, file := range reader.File {
		ext := filepath.Ext(file.Name)

		if ext == ".xml" {
			fileReader, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer fileReader.Close()

			p := importReaderXML(fileReader, file.Name)

			ps = append(ps, p)
		}
	}

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
