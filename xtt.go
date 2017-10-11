package poslog

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type XTT struct {
	Data     []byte  `json:"Data"`
	Type     *[]byte `json:"Type,omitempty"`
	Filename string  `json:"Filename"`
	XTag     *XTag   `json:"XTag,omitempty"`
}

type XTag struct {
	Name      string    `json:"Name"`
	Attrs     *[]string `json:"Attrs,omitempty"`
	Datatypes *string   `json:"Datatypes,omitempty"`
	Members   *[]XTag   `json:"Members,omitempty"`
}

func (x *XTT) load(f string) {
	xb, err := ioutil.ReadFile(f)
	x.Filename, x.Data = f, xb
	if err != nil {
		fmt.Println(err)
	}
}

func (x *XTT) toType() {
	var open bool
	var name []string
loop:
	for xc := range x.Data {
		if open && string(x.Data[xc]) == ">" {
			fmt.Println(strings.Join(name, ""))
			open = false
			continue loop
		} else if open {
			name = append(name, string(x.Data[xc]))
			continue loop
			// } else if string(x.Data[xc]) == "\n" || string(x.Data[xc]) == " " {
			// 	continue loop
		} else if string(x.Data[xc]) == "<" {
			open = true
			name = name[:0]
			continue loop
		} else {
			// log.Fatalln("Should be start tag first WTF: ", string(x.Data[xc]))
			continue loop
		}
	}
}
