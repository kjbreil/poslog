package poslog

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// func (p *POSLog) transactionCounts() {
// 	for _, tr := range p.Transaction {
// 		if tr.RetailTransaction != nil {
// 			tr.RetailTransaction.counts()
// 		}
// 	}
// }

func NewPOSLog(trs []*Transaction) (p POSLog) {
	for _, eachTransaction := range trs {
		p.Transaction = append(p.Transaction, eachTransaction)
	}
	return
}

// each returns all transactions with information pre processed and appended to type
func (p *POSLog) each() (trs []*Transaction) {
	if p.Transaction != nil && len(p.Transaction) != 0 {
		for _, eachTransaction := range p.Transaction {
			if eachTransaction.RetailTransaction != nil {
				eachTransaction.RetailTransaction.counts()
				trs = append(trs, eachTransaction)
			}
		}
	}
	return
}

// filename is to append the filename to the poslog object
func (p *POSLog) filename(filename string) {
	p.Filename = &filename
}

// storeID returns a string of the store id for the poslog. This is not only informational but assures there is no major problem
// with the poslog xml file
func (p *POSLog) storeID() (storeID int) {
	if len(p.Transaction) == 0 {
		log.Println("No store id in ", &p.Filename)
		storeID = 0
		return
	}

	storeID = p.Transaction[0].RetailStoreID

	for _, t := range p.Transaction {
		if t.RetailStoreID != storeID {
			log.Fatalln("Multiple Store ID's in single POSLog, I cannot handle that yet", t.RetailStoreID, storeID, p.Filename)
		}
	}
	p.RetailStoreID = &storeID
	return
}

// buisnessDayDate returns the buisness date of all transactions in a poslog file, all transactions in a poslog file should have SAME buisness date
func (p *POSLog) buisnessDayDate() (buisnessDayDate *string) {
	if len(p.Transaction) == 0 {
		log.Println("No bid in ", p.Filename)
		*buisnessDayDate = "0"
		return
	}

	buisnessDayDate = &p.Transaction[0].BusinessDayDate

	for _, tr := range p.Transaction {
		if *buisnessDayDate != tr.BusinessDayDate {
			log.Fatalln("Multiple Buisness days in single poslog file, I cannot handle that yet", tr.BusinessDayDate, *buisnessDayDate, *p.Filename)
		}
	}

	p.BusinessDayDate = buisnessDayDate
	return
}

// dayid returns an int in dayid format YYYYMMDD
// Should actuall Validate this as go time 20060102 like EndInt
func (tr *Transaction) dayid() (dayid *string) {

	inputFormat := "2006-01-02"

	time, err := time.Parse(inputFormat, tr.BusinessDayDate)
	if err != nil {
		log.Fatalln(err)
	}

	format := "20060102"
	mdy := time.Format(format)

	// did, err := strconv.Atoi(mdy)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	return &mdy
}

// End returns an int of the transaction end datetime
func (tr *Transaction) end() (endstring *string) {
	time, err := time.Parse(time.RFC3339, tr.BusinessDayDate)
	if err != nil {
		log.Fatalln(err)
	}
	format := "20060102150405"
	*endstring = time.Format(format)
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

func (rt *RetailTransaction) counts() {
	var lineItemCounts LineItemCounts
	// rt.LineItemCounts.AgeRestrictionCount = 0
	// rt.LineItemCounts.CRMCustomVariableCount = 0
	// rt.LineItemCounts.CardActivationCount = 0
	// rt.LineItemCounts.ElectronicSignatureCount = 0
	// rt.LineItemCounts.ItemNotFoundCount = 0
	// rt.LineItemCounts.ItemRestrictionCount = 0
	// rt.LineItemCounts.LoyaltyMembershipCount = 0
	// rt.LineItemCounts.LoyaltyRewardCount = 0
	// rt.LineItemCounts.SaleCount = 0
	// rt.LineItemCounts.TaxCount = 0
	// rt.LineItemCounts.TenderCount = 0
	for _, li := range rt.LineItem {
		if li.AgeRestriction != nil {
			lineItemCounts.AgeRestrictionCount++
		}
		if li.CRMCustomVariable != nil {
			lineItemCounts.CRMCustomVariableCount++
		}
		if li.CardActivation != nil {
			lineItemCounts.CardActivationCount++
		}
		if li.ElectronicSignature != nil {
			lineItemCounts.ElectronicSignatureCount++
		}
		if li.ItemNotFound != nil {
			lineItemCounts.ItemNotFoundCount++
		}
		if li.ItemRestriction != nil {
			lineItemCounts.ItemRestrictionCount++
		}
		if li.LoyaltyMembership != nil {
			lineItemCounts.LoyaltyMembershipCount++
		}
		if li.LoyaltyReward != nil {
			lineItemCounts.LoyaltyRewardCount++
		}
		if li.Sale != nil {
			lineItemCounts.SaleCount++
		}
		if li.Tax != nil {
			lineItemCounts.TaxCount++
		}
		if li.Tender != nil {
			lineItemCounts.TenderCount++
		}
	}
	rt.LineItemCounts = &lineItemCounts
}
