package poslog

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// POSLogs is an array of poslog grouped by store(s) and dayid(s)
type POSLogs struct {
	DayIDs  []string
	Stores  []int
	POSLogs []POSLog
}

func (ps *POSLogs) appendDayID(d string) {
	for _, c := range ps.DayIDs {
		if c == d {
			continue
		} else {
			ps.DayIDs = append(ps.DayIDs, d)
		}
		return
	}
}

func (ps *POSLogs) appendStore(s int) {
	for _, c := range ps.Stores {
		if c == s {
			continue
		} else {
			ps.Stores = append(ps.Stores, s)
		}
		return
	}
}

func (ps *POSLogs) appendPOSLog(p POSLog) {
	ps.POSLogs = append(ps.POSLogs, p)
	return
}

func zipReadAllXML(archive string) (ps POSLogs) {

	an := strings.TrimSuffix(filepath.Base(archive), filepath.Ext(filepath.Base(archive)))
	reader, err := zip.OpenReader(archive)
	if err != nil {
		fmt.Println("FUCK THE ZIP DIDN'T WORK")
	}

	for _, file := range reader.File {
		ext := filepath.Ext(file.Name)
		noext := strings.TrimSuffix(file.Name, ext)
		ofp := filepath.Join("output", an)
		os.MkdirAll(ofp, 0777)

		if ext == ".xml" {

			ofn := filepath.Join(ofp, noext)
			ofn = "./" + ofn
			fileReader, err := file.Open()
			if err != nil {
				fmt.Println("ERRUR")
			}
			defer fileReader.Close()

			p := importReaderXML(fileReader, file.Name)

			ps.appendPOSLog(p)
		}
	}

	return
}
