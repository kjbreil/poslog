package poslog

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
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

func tarJSONs(wr io.Writer, ps POSLogs) {
	tw := tar.NewWriter(wr)
	defer tw.Close()
	for _, file := range ps.POSLogs {
		posLogString := createJSON(file)
		filename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + ".json"
		hdr := &tar.Header{
			Name: filename,
			Mode: 0666,
			Size: int64(len(posLogString)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(posLogString)); err != nil {
			log.Fatalln(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

}

func tarXMLs(wr io.Writer, ps POSLogs) {
	tw := tar.NewWriter(wr)
	defer tw.Close()
	for _, file := range ps.POSLogs {
		posLogString := createXML(file)
		filename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + ".xml"
		hdr := &tar.Header{
			Name: filename,
			Mode: 0666,
			Size: int64(len(posLogString)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(posLogString)); err != nil {
			log.Fatalln(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

}

func gzipJSONs(file io.Writer, ps POSLogs) {

	gz := gzip.NewWriter(file)
	defer gz.Close()
	tarJSONs(gz, ps)

	if err := gz.Close(); err != nil {
		log.Fatalln(err)
	}

}

func gzipXMLs(file io.Writer, ps POSLogs) {
	gz := gzip.NewWriter(file)
	defer gz.Close()
	tarXMLs(gz, ps)

	if err := gz.Close(); err != nil {
		log.Fatalln(err)
	}

}
