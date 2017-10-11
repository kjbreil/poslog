package poslog

import (
	"log"
	"testing"
)

func TestXML(t *testing.T) {
	var x XTT
	log.Println("I Made a Type")
	x.load("./sample/SMALL.xml")
	x.toType()

}
