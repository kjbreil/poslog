// Package poslog works with POSLog XML files.
// Currently writing types for NCR ACS POSLog
// But can be extended for other POS vendores and all
// NCR specific tags are (will be) set as options
package poslog

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// POSLog is the super type of a POSLog XML file. The type contains
// the marshaling information to marshal and unmarshal to json and xml
// currently it is not complete and does not handle all fields in
// source XML
type POSLog struct {
	POSLog      string `xml:"POSLog" json:"POSLog"`
	XMLNSPOSLog string `xml:"xmlns poslog,attr"`
	XMLNS       string `xml:"xmlns,attr"`
	XMLNSACS    string `xml:"xmlns acs,attr"`
	XMLNSRaw    string `xml:"xmlns raw,attr"`
	XMLNSXSI    string `xml:"xmlns xsi,attr"`
	XMLNSMSXSL  string `xml:"xmlns msxsl,attr"`
	XMLNSAS     string `xml:"xmlns as,attr"`
	XMLNSACSSM  string `xml:"xmlns acssm,attr"`
	// Transaction does not have to exist in the transaction, thus a pointer
	// and we we have a POSLog type with nothing in it (unless there is
	// another top level array type)
	Transaction []*Transaction `xml:"Transaction" json:"Transaction"`
}

// Transaction is the main body of a POSLog XML. Each Transaction is not
// a literal transaction at the register but each complete actions
// at the register so a sign in to register would be a transactions
type Transaction struct {
	RetailStoreID     *int    `xml:"RetailStoreID" json:"RetailStoreID"`
	WorkstationID     *int    `xml:"WorkstationID" json:"WorkstationID"`
	SequenceNumber    *int    `xml:"SequenceNumber" json:"SequenceNumber"`
	BusinessDayDate   *string `xml:"BusinessDayDate" json:"BusinessDayDate"`
	EndDateTime       *string `xml:"EndDateTime" json:"EndDateTime"`
	OperatorID        *OperatorID
	CurrencyCode      *string `xml:"CurrencyCode" json:"CurrencyCode"`
	RetailTransaction *RetailTransaction
}

// OperatorID is the operator associeated with the transacions, some
// are done with system operators
type OperatorID struct {
	OperatorID   *int    `xml:",chardata" json:"OperatorID"`
	OperatorName *string `xml:"OperatorName,attr" json:"OperatorName"`
}

type RetailTransaction struct {
	Version            *string `xml:"Version,attr" json:"Version"`
	ReceiptDateTime    *string `xml:"ReceiptDateTime" json:"ReceiptDateTime"`
	TransactionCount   *int    `xml:"TransactionCount" json:"TransactionCount"`
	LineItem           []*LineItem
	Total              []*Total
	ItemCode           *int `xml:"ItemCount" json:"ItemCount"`
	PerformanceMetrics *PerformanceMetrics
}

// LineItem is a each action at the register, SequenceNumber will show actual sequence at register
type LineItem struct {
	EntryMethod    string  `xml:"EntryMethod,attr" json:"EntryMethod"`
	SequenceNumber int     `xml:"SequenceNumber" json:"SequenceNumber"`
	Sale           *Sale   `xml:"Sale,omitempty" json:"Sale,omitempty"`
	Tender         *Tender `xml:"Tender,omitempty" json:"Tender,omitempty"`
}

// Sale is a subtype of LineItem, the literal sale of an item
type Sale struct {
	ItemType               *string `xml:"ItemType,attr,omitempty" json:"ItemType,omitempty"`
	POSIdentity            *POSIdentity
	ItemID                 *string `xml:"ItemID" json:"ItemID"`
	MerchandiseHierarchy   *MerchandiseHierarchy
	Description            *string `Description:"ItemID"`
	RegularSalesUnitPrice  *string `xml:"RegularSalesUnitPrice" json:"RegularSalesUnitPrice"`
	ExtendedAmount         *string `xml:"ExtendedAmount" json:"ExtendedAmount"`
	DiscountAmount         *string `xml:"IteDiscountAmountmID" json:"IteDiscountAmountmID"`
	ExtendedDiscountAmount *string `ExtendedDiscountAmount:"ItemID"`
	Quantity               *string `xml:"Quantity" json:"Quantity"`
	OperatorSequence       *string `xml:"OperatorSequence" json:"OperatorSequence"`
	ReportCode             *string `xml:"ReportCode" json:"ReportCode"`
	SaleableMediaID        *string `xml:"SaleableMediaID" json:"SaleableMediaID"`
	Itemizers              *Itemizers
}

type POSIdentity struct {
	POSIDType *string `xml:"POSIDType,attr" json:"POSIDType"`
	POSItemID *string `xml:"POSItemID" json:"POSItemID"`
	Qualifier *int    `xml:"Qualifier" json:"Qualifier"`
}

type MerchandiseHierarchy struct {
	MerchandiseHierarchy     *int    `xml:",chardata" json:"MerchandiseHierarchy"`
	Level                    *string `xml:"Level,attr" json:"Level"`
	ACSDepartmentDescription *string `xml:"DepartmentDescription,attr" json:"DepartmentDescription"`
}

type Itemizers struct {
	Itemizers     int    `xml:",chardata" json:"Itemizers"`
	FoodStampable string `xml:"FoodStampable,attr,omitempty" json:"FoodStampable,omitempty"`
	Itemizer1     string `xml:"Itemizer1,attr,omitempty" json:"Itemizer1,omitempty"`
	Itemizer2     string `xml:"Itemizer2,attr,omitempty" json:"Itemizer2,omitempty"`
}

type Tender struct {
	TenderType        *string `xml:"TenderType,attr,omitempty" json:"TenderType,omitempty"`
	TypeCode          *string `xml:"TypeCode,attr,omitempty" json:"TypeCode,omitempty"`
	TenderDescription *string `xml:"TenderDescription,attr,omitempty" json:"TenderDescription,omitempty"`
	TenderID          *int    `xml:"TenderID,omitempty" json:"TenderID,omitempty"`
	Amount            *string `xml:"Amount,omitempty" json:"Amount,omitempty"`
}

type Total struct {
	Total     string `xml:",chardata" json:"Total"`
	TotalType string `xml:"TotalType,attr" json:"TotalType"`
}

type PerformanceMetrics struct {
	RingTime   int `xml:"RingTime" json:"RingTime"`
	IdleTime   int `xml:"IdleTime" json:"IdleTime"`
	TenderTime int `xml:"TenderTime" json:"TenderTime"`
}

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
