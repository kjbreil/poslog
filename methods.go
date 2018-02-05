package poslog

import (
	"log"
	"strconv"
	"strings"
)

// filename is to append the filename to the poslog object
func (p *POSLog) filename(filename string) {
	p.Filename = &filename
}

// dayid returns an int in dayid format YYYYMMDD
// Should actuall Validate this as go time 20060102 like EndInt
func (tr *Transaction) dayid() (dayid *int) {
	format := "20061002"
	mdy := tr.BusinessDayDate.Format(format)

	did, err := strconv.Atoi(mdy)
	if err != nil {
		log.Fatalln(err)
	}
	return &did
}

// storeID returns a string of the store id for the poslog. This is not only informational but assures there is no major problem
// with the poslog xml file
func (p *POSLog) storeID() (storeID int) {
	if len(p.Transaction) == 0 {
		log.Println("No store id in ", p.Filename)
		storeID = 0
		return
	}

	storeID = p.Transaction[0].RetailStoreID

	for _, t := range p.Transaction {
		if t.RetailStoreID != storeID {
			log.Fatalln("Multiple Store ID's in single POSLog, I cannot handle that yet", t.RetailStoreID, storeID, p.Filename)
		}
	}
	return
}

// End returns an int of the transaction end datetime
func (tr *Transaction) end() (endstring *string) {
	format := "20060102150405"
	*endstring = tr.EndDateTime.Format(format)
	return
}

// id appends the transactionID to a transaction
func (tr *Transaction) id() {
	var tida []string
	// first is the enddateint, good for sorting
	tida = append(tida, string(*tr.end()))
	// next the dayid, this is buisness date
	tida = append(tida, string(*tr.dayid()))
	// next store number
	tida = append(tida, strconv.Itoa(tr.RetailStoreID))
	// Terminal Number
	tida = append(tida, strconv.Itoa(tr.WorkstationID))
	// SequenceNumber
	tida = append(tida, strconv.Itoa(tr.SequenceNumber))

	tr.TransactionID = strings.Join(tida, "-")

	return
}
