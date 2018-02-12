package poslog

import (
	"encoding/xml"
)

// Itemizers is item iteizers
type Itemizers struct {
	AttrFoodStampable string `xml:" FoodStampable,attr"  json:",omitempty" db:"IdleTime,omitempty"`
	AttrItemizer1     string `xml:" Itemizer1,attr"  json:",omitempty" db:"itemizer_1,omitempty"`
	AttrItemizer2     string `xml:" Itemizer2,attr"  json:",omitempty" db:"itemizer_2,omitempty"`
	AttrItemizer3     string `xml:" Itemizer3,attr"  json:",omitempty" db:"itemizer_3,omitempty"`
	AttrItemizer4     string `xml:" Itemizer4,attr"  json:",omitempty" db:"itemizer_4,omitempty"`
	AttrTax1          string `xml:" Tax1,attr"  json:",omitempty" db:"tax_1,omitempty"`
	AttrTax2          string `xml:" Tax2,attr"  json:",omitempty" db:"tax_2,omitempty"`
	AttrTax3          string `xml:" Tax3,attr"  json:",omitempty" db:"tax_3,omitempty"`
	AttrTax4          string `xml:" Tax4,attr"  json:",omitempty" db:"tax_4,omitempty"`
	Itemizers         string `xml:",chardata" json:",omitempty" db:"itemizers,omitempty"`
}

// PerformanceMetrics contains transaction part timings
type PerformanceMetrics struct {
	IdleTime   *string  `xml:"IdleTime,omitempty" json:"IdleTime,omitempty" db:"IdleTime,omitempty"`
	RingTime   *string  `xml:"RingTime,omitempty" json:"RingTime,omitempty" db:"RingTime,omitempty"`
	TenderTime *string  `xml:"TenderTime,omitempty" json:"TenderTime,omitempty" db:"TenderTime,omitempty"`
	XMLName    xml.Name `xml:"PerformanceMetrics,omitempty" json:"PerformanceMetrics,omitempty"`
}

// Svg holds the vector graphics of a signature
type Svg struct {
	AttrHeight string   `xml:" height,attr"  json:",omitempty"`
	AttrStyle  string   `xml:" style,attr"  json:",omitempty"`
	AttrWidth  string   `xml:" width,attr"  json:",omitempty"`
	AttrXmlns  string   `xml:" xmlns,attr"  json:",omitempty"`
	Path       []*Path  `xml:"path,omitempty" json:"path,omitempty" db:"path,omitempty"`
	XMLName    xml.Name `xml:"svg,omitempty" json:"svg,omitempty"`
}
