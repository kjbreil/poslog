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

func (p *POSLog) counts() {
	if len(p.Transaction) > 0 {
		for i := range p.Transaction {
			p.Transaction[i].counts()
			p.Transaction[i].id()
		}
	}
	return
}

func (tr *Transaction) counts() {
	if tr.RetailTransaction != nil {
		tc := tr.RetailTransaction.counts()
		tr.TransactionCounts = &tc
	}
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
func (tr *Transaction) end() (endstring string) {
	tm, err := time.Parse("2006-01-02T15:04:05", tr.EndDateTime)
	if err != nil {
		log.Fatalln(err)
	}
	format := "20060102150405"
	endstring = tm.Format(format)
	return
}

// id appends the transactionID to a transaction
func (tr *Transaction) id() {
	var tida []string
	// first is the enddateint, good for sorting
	tida = append(tida, string(tr.end()))
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

func (rt *RetailTransaction) counts() TransactionCounts {
	var tc TransactionCounts

	for _, li := range rt.LineItem {
		if li.AgeRestriction != nil {
			tc.AgeRestrictionCount++
		}
		if li.CRMCustomVariable != nil {
			tc.CRMCustomVariableCount++
		}
		if li.CardActivation != nil {
			tc.CardActivationCount++
		}
		if li.ElectronicSignature != nil {
			tc.ElectronicSignatureCount++
		}
		if li.ItemNotFound != nil {
			tc.ItemNotFoundCount++
		}
		if li.ItemRestriction != nil {
			tc.ItemRestrictionCount++
		}
		if li.LoyaltyMembership != nil {
			tc.LoyaltyMembershipCount++
		}
		if li.LoyaltyReward != nil {
			tc.LoyaltyRewardCount++
		}
		if li.Sale != nil {
			tc.SaleCount++
		}
		if li.Tax != nil {
			tc.TaxCount++
		}
		if li.Tender != nil {
			tc.TenderCount++
		}
	}

	return tc
}

// Type returns a string of the line item
func (li *LineItem) Type() string {
	if li.AgeRestriction != nil {
		return "AgeRestriction"
	}
	if li.CRMCustomVariable != nil {
		return "CRMCustomVariable"
	}
	if li.CardActivation != nil {
		return "CardActivation"
	}
	if li.ElectronicSignature != nil {
		return "ElectronicSignature"
	}
	if li.ItemNotFound != nil {
		return "ItemNotFound"
	}
	if li.ItemRestriction != nil {
		return "ItemRestriction"
	}
	if li.LoyaltyMembership != nil {
		return "LoyaltyMembership"
	}
	if li.LoyaltyReward != nil {
		return "LoyaltyReward"
	}
	if li.Sale != nil {
		return "Sale"
	}
	if li.Tax != nil {
		return "Tax"
	}
	if li.Tender != nil {
		return "Tender"
	}
	return "UNKOWN"
}
