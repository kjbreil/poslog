// Package poslog works with POSLog XML files.
// Currently writing types for NCR ACS POSLog
// But can be extended for other POS vendores and all
// NCR specific tags are (will be) set as options
package poslog

import (
	"encoding/xml"
)

// POSLog the main type of a POSLog XMl file. The type contains
// the marshaling information to marshal and unmarshal to json and xml
// currently it is not complete and does not handle all fields in
// source XML
type POSLog struct {
	Filename         *string        `xml:"Filename,omitempty" json:"Filename,omitempty" db:"filename,omitempty"`
	RetailStoreID    *int           `xml:"RetailStoreID,omitempty" json:"RetailStoreID,omitempty" db:"retail_store_id,omitempty"`
	BusinessDayDate  *string        `xml:"BusinessDayDate,omitempty" json:"BusinessDayDate,omitempty" db:"buisness_day_date,omitempty"`
	TransactionCount *int           `xml:"TransactionCount,omitempty" json:"TransactionCount,omitempty"  db:"transaction_count,omitempty"`
	XmlnsAcs         *string        `xml:"xmlns acs,attr,omitempty"  json:",omitempty"`
	XmlnsAcssm       *string        `xml:"xmlns acssm,attr,omitempty"  json:",omitempty"`
	XmlnsAs          *string        `xml:"xmlns as,attr,omitempty"  json:",omitempty"`
	XmlnsMsxsl       *string        `xml:"xmlns msxsl,attr,omitempty"  json:",omitempty"`
	XmlnsPoslog      *string        `xml:"xmlns poslog,attr,omitempty"  json:",omitempty"`
	XmlnsRaw         *string        `xml:"xmlns raw,attr,omitempty"  json:",omitempty"`
	Xmlns            *string        `xml:"xmlns,attr,omitempty"  json:",omitempty"`
	XmlnsXsi         *string        `xml:"xmlns xsi,attr,omitempty"  json:",omitempty"`
	Transaction      []*Transaction `xml:"http://www.nrf-arts.org/IXRetail/namespace/ Transaction,omitempty" json:"Transaction,omitempty" db:"http://www.nrf-arts.org/IXRetail/namespace/ Transaction,omitempty"`
	XMLName          xml.Name       `xml:"http://www.nrf-arts.org/IXRetail/namespace/ POSLog,omitempty" json:"POSLog,omitempty"`
}

// Transaction is the body of POSLog, each action at the POS is a transaction
type Transaction struct {
	TransactionID      string              `json:"TransactionID" db:"transaction_id"`
	TransactionCounts  *TransactionCounts  `xml:"TransactionCounts,omitempty" json:"TransactionCounts,omitempty"`
	BusinessDayDate    string              `xml:"BusinessDayDate" json:"BusinessDayDate" db:"buisness_day_date"`
	ControlTransaction *ControlTransaction `xml:"ControlTransaction,omitempty" json:"ControlTransaction,omitempty" db:"control_transaction,omitempty"`
	CurrencyCode       *string             `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" db:"currency_code,omitempty"`
	EndDateTime        string              `xml:"EndDateTime" json:"EndDateTime" db:"end_date_time"`
	OperatorID         *OperatorID
	RetailStoreID      int                `xml:"RetailStoreID" json:"RetailStoreID" db:"retail_store_id"`
	RetailTransaction  *RetailTransaction `xml:"RetailTransaction,omitempty" json:"RetailTransaction,omitempty"`
	SequenceNumber     int                `xml:"SequenceNumber" json:"SequenceNumber" db:"sequence_number"`
	WorkstationID      int                `xml:"WorkstationID" json:"WorkstationID" db:"workstation_id"`
	XMLName            xml.Name           `xml:"Transaction,omitempty" json:"Transaction,omitempty"`
}

// OperatorID is the name and into of operator
type OperatorID struct {
	OperatorID   int    `xml:",chardata" json:"OperatorID,omitempty" db:"operator_id,omitempty"`
	OperatorName string `xml:" OperatorName,attr,omitempty"  json:"operator_name,omitempty"`
}

// RetailTransaction is any "sale" transaction
type RetailTransaction struct {
	AttrVersion      *string          `xml:" Version,attr,omitempty"  json:",omitempty" db:"version,omitempty"`
	LineItem         []*LineItem      `xml:"LineItem,omitempty" json:"LineItem,omitempty"`
	ReceiptDateTime  string           `xml:"ReceiptDateTime,omitempty" json:"ReceiptDateTime,omitempty" db:"receipt_date_time,omitempty"`
	Total            []*Total         `xml:"Total,omitempty" json:"Total,omitempty"`
	TransactionCount *string          `xml:"TransactionCount,omitempty" json:"TransactionCount,omitempty" db:"transaction_count,omitempty"`
	TransactionLink  *TransactionLink `xml:"TransactionLink,omitempty" json:"TransactionLink,omitempty"`
	XMLName          xml.Name         `xml:"RetailTransaction,omitempty" json:"RetailTransaction,omitempty"`
	// namespace ACSIR
	PerformanceMetrics *PerformanceMetrics `xml:"PerformanceMetrics,omitempty" json:"PerformanceMetrics,omitempty"`
	ItemCount          *int                `xml:"ItemCount,omitempty" json:"ItemCount,omitempty" db:"item_count,omitempty"`
}

// LineItem is each line at the register, in order of squence
// Besides the attributes and sequance number each type is a different type of lineitem
// to confirm a whole xml has already been inserted we need to summarize each line item
type LineItem struct {
	SequenceNumber  int            `xml:"SequenceNumber" json:"SequenceNumber" db:"SequenceNumber"`
	AttrEntryMethod *string        `xml:" EntryMethod,attr,omitempty"  json:",omitempty"`
	AttrVoidFlag    *string        `xml:" VoidFlag,attr,omitempty"  json:",omitempty"`
	Tax             *Tax           `xml:"Tax,omitempty" json:"Tax,omitempty" db:"Tax,omitempty"`
	LoyaltyReward   *LoyaltyReward `xml:"LoyaltyReward,omitempty" json:"LoyaltyReward,omitempty" db:"LoyaltyReward,omitempty"`
	Tender          *Tender        `xml:"Tender,omitempty" json:"Tender,omitempty" db:"Tender,omitempty"`
	Sale            *Sale          `xml:"Sale,omitempty" json:"Sale,omitempty" db:"Sale,omitempty"`
	// namespace ACSIR
	AttrAcsSpaceKeyedPrice       *string              `xml:"keyedPrice,attr,omitempty"  json:",omitempty"`
	AttrAcsSpaceOperatorOverride *string              `xml:"OperatorOverride,attr,omitempty"  json:",omitempty"`
	AttrAcsSpacePriceRequired    *string              `xml:"priceRequired,attr,omitempty"  json:",omitempty"`
	AttrAcsSpaceWeightItem       *string              `xml:"weightItem,attr,omitempty"  json:",omitempty"`
	AgeRestriction               *AgeRestriction      `xml:"AgeRestriction,omitempty" json:"AgeRestriction,omitempty" db:"AgeRestriction,omitempty"`
	CRMCustomVariable            *CRMCustomVariable   `xml:"CRMCustomVariable,omitempty" json:"CRMCustomVariable,omitempty" db:"CRMCustomVariable,omitempty"`
	CardActivation               *CardActivation      `xml:"CardActivation,omitempty" json:"CardActivation,omitempty" db:"CardActivation,omitempty"`
	ElectronicSignature          *ElectronicSignature `xml:"ElectronicSignature,omitempty" json:"ElectronicSignature,omitempty" db:"ElectronicSignature,omitempty"`
	ItemNotFound                 *ItemNotFound        `xml:"ItemNotFound,omitempty" json:"ItemNotFound,omitempty" db:"ItemNotFound,omitempty"`
	ItemRestriction              *ItemRestriction     `xml:"ItemRestriction,omitempty" json:"ItemRestriction,omitempty" db:"ItemRestriction,omitempty"`
	LoyaltyMembership            *LoyaltyMembership   `xml:"LoyaltyMembership,omitempty" json:"LoyaltyMembership,omitempty" db:"LoyaltyMembership,omitempty"`
	XMLName                      xml.Name             `xml:"LineItem,omitempty" json:"LineItem,omitempty"`
}

// TransactionCounts is the count of each type of lineitem in a POSLog file, used for verfication of data
type TransactionCounts struct {
	AgeRestrictionCount      int `xml:"AgeRestrictionCount,omitempty" json:"AgeRestrictionCount,omitempty" db:"AgeRestrictionCount,omitempty"`
	CRMCustomVariableCount   int `xml:"CRMCustomVariableCount,omitempty" json:"CRMCustomVariableCount,omitempty" db:"CRMCustomVariableCount,omitempty"`
	CardActivationCount      int `xml:"CardActivationCount,omitempty" json:"CardActivationCount,omitempty" db:"CardActivationCount,omitempty"`
	ElectronicSignatureCount int `xml:"ElectronicSignatureCount,omitempty" json:"ElectronicSignatureCount,omitempty" db:"ElectronicSignatureCount,omitempty"`
	ItemNotFoundCount        int `xml:"ItemNotFoundCount,omitempty" json:"ItemNotFoundCount,omitempty" db:"ItemNotFoundCount,omitempty"`
	ItemRestrictionCount     int `xml:"ItemRestrictionCount,omitempty" json:"ItemRestrictionCount,omitempty" db:"ItemRestrictionCount,omitempty"`
	LoyaltyMembershipCount   int `xml:"LoyaltyMembershipCount,omitempty" json:"LoyaltyMembershipCount,omitempty" db:"LoyaltyMembershipCount,omitempty"`
	LoyaltyRewardCount       int `xml:"LoyaltyRewardCount,omitempty" json:"LoyaltyRewardCount,omitempty" db:"LoyaltyRewardCount,omitempty"`
	SaleCount                int `xml:"SaleCount,omitempty" json:"SaleCount,omitempty" db:"SaleCount,omitempty"`
	TaxCount                 int `xml:"TaxCount,omitempty" json:"TaxCount,omitempty" db:"TaxCount,omitempty"`
	TenderCount              int `xml:"TenderCount,omitempty" json:"TenderCount,omitempty" db:"TenderCount,omitempty"`
}

// Sale is a line item "sold" item
type Sale struct {
	AttrItemType           string                `xml:" ItemType,attr,omitempty"  json:",omitempty" db:"item_type,omitempty"`
	Description            *string               `xml:"Description,omitempty" json:"Description,omitempty" db:"description,omitempty"`
	DiscountAmount         *string               `xml:"DiscountAmount,omitempty" json:"DiscountAmount,omitempty" db:"discount_amount,omitempty"`
	ExtendedAmount         *string               `xml:"ExtendedAmount,omitempty" json:"ExtendedAmount,omitempty" db:"extended_amount,omitempty"`
	ExtendedDiscountAmount *string               `xml:"ExtendedDiscountAmount,omitempty" json:"ExtendedDiscountAmount,omitempty" db:"extended_discount_amount,omitempty"`
	ItemID                 string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"item_id,omitempty"`
	MerchandiseHierarchy   *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty"`
	POSIdentity            *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty"`
	Quantity               *string               `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"quantity,omitempty"`
	RegularSalesUnitPrice  *string               `xml:"RegularSalesUnitPrice,omitempty" json:"RegularSalesUnitPrice,omitempty" db:"regular_sale_unit_price,omitempty"`
	XMLName                xml.Name              `xml:"Sale,omitempty" json:"Sale,omitempty"`
	// namespace ACS-IR
	OperatorSequence *string    `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	ReportCode       *string    `xml:"ReportCode,omitempty" json:"ReportCode,omitempty" db:"ReportCode,omitempty"`
	SaleableMediaID  *int       `xml:"SaleableMediaID,omitempty" json:"SaleableMediaID,omitempty" db:"SaleableMediaID,omitempty"`
	Itemizers        *Itemizers `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
}

// POSIdentity contains basic item information, UPC and department as the itemID and Qualifier
type POSIdentity struct {
	AttrPOSIDType string   `xml:"POSIDType,attr"  json:",omitempty"`
	POSItemID     *int     `xml:"POSItemID,omitempty" json:"POSItemID,omitempty" db:"POSItemID,omitempty"`
	Qualifier     *string  `xml:"Qualifier,omitempty" json:"Qualifier,omitempty" db:"Qualifier,omitempty"`
	XMLName       xml.Name `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty"`
}

// MerchandiseHierarchy is the merchandising department of an item
type MerchandiseHierarchy struct {
	AttrLevel string   `xml:" Level,attr"  json:",omitempty" db:"level_name,omitempty"`
	Text      string   `xml:",chardata" json:",omitempty" db:"level,omitempty"`
	XMLName   xml.Name `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty"`
	// namespace ACSIR
	AttrAcsSpaceDepartmentDescription string `xml:"DepartmentDescription,attr"  json:",omitempty"`
}

// Tender is information about the tender amount(s) and type(s)
type Tender struct {
	AttrTenderType string         `xml:" TenderType,attr"  json:",omitempty" db:"tender_type,omitempty"`
	AttrTypeCode   string         `xml:" TypeCode,attr"  json:",omitempty" db:"type_code,omitempty"`
	Amount         *string        `xml:"Amount,omitempty" json:"Amount,omitempty" db:"amount,omitempty"`
	Authorization  *Authorization `xml:"Authorization,omitempty" json:"authorization,omitempty"`
	Cashback       *string        `xml:"Cashback,omitempty" json:"Cashback,omitempty" db:"cashback,omitempty"`
	Coupon         *Coupon        `xml:"Coupon,omitempty" json:"Coupon,omitempty"`
	CreditDebit    *CreditDebit   `xml:"CreditDebit,omitempty" json:"CreditDebit,omitempty"`
	TenderChange   *TenderChange  `xml:"TenderChange,omitempty" json:"TenderChange,omitempty"`
	TenderID       *int           `xml:"TenderID,omitempty" json:"TenderID,omitempty" db:"TenderID,omitempty"`
	XMLName        xml.Name       `xml:"Tender,omitempty" json:"Tender,omitempty"`
	// namespace ACS-IR
	OperatorSequence              *int   `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	AttrAcsSpaceTenderDescription string `xml:"TenderDescription,attr"  json:",omitempty"`
}

// Authorization is the data returned after credit/debit authorization
type Authorization struct {
	AttrElectronicSignature string   `xml:" ElectronicSignature,attr"  json:",omitempty"`
	AttrHostAuthorized      string   `xml:" HostAuthorized,attr"  json:",omitempty"`
	AuthorizationCode       *string  `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" db:"AuthorizationCode,omitempty"`
	AuthorizationDateTime   *string  `xml:"AuthorizationDateTime,omitempty" json:"AuthorizationDateTime,omitempty" db:"AuthorizationDateTime,omitempty"`
	AuthorizedChangeAmount  *string  `xml:"AuthorizedChangeAmount,omitempty" json:"AuthorizedChangeAmount,omitempty" db:"AuthorizedChangeAmount,omitempty"`
	ReferenceNumber         *string  `xml:"ReferenceNumber,omitempty" json:"ReferenceNumber,omitempty" db:"ReferenceNumber,omitempty"`
	RequestedAmount         *string  `xml:"RequestedAmount,omitempty" json:"RequestedAmount,omitempty" db:"RequestedAmount,omitempty"`
	XMLName                 xml.Name `xml:"Authorization,omitempty" json:"Authorization,omitempty"`
}

// Credit Debit is just a flag of if the transaction is a credit or debit transaciton
type CreditDebit struct {
	AttrCardType string `xml:" CardType,attr"  json:",omitempty"`
	// namespace ACSIR
	AttrAcsSpaceCreditDescription string   `xml:"CreditDescription,attr"  json:",omitempty"`
	XMLName                       xml.Name `xml:"CreditDebit,omitempty" json:"CreditDebit,omitempty"`
}

type LoyaltyReward struct {
	PromotionID *int     `xml:"PromotionID,omitempty" json:"PromotionID,omitempty" db:"PromotionID,omitempty"`
	EventID     *int     `xml:"EventID,omitempty" json:"EventID,omitempty" db:"EventID,omitempty"`
	ReasonCode  *string  `xml:"ReasonCode,omitempty" json:"ReasonCode,omitempty" db:"ReasonCode,omitempty"`
	XMLName     xml.Name `xml:"LoyaltyReward,omitempty" json:"LoyaltyReward,omitempty"`
	// namespace ACS-IR
	ExtendedRewardAmount           *string                    `xml:"ExtendedRewardAmount,omitempty" json:"ExtendedRewardAmount,omitempty" db:"ExtendedRewardAmount,omitempty"`
	Itemizers                      *Itemizers                 `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
	OperatorSequenceReference      *OperatorSequenceReference `xml:"OperatorSequenceReference,omitempty" json:"OperatorSequenceReference,omitempty" db:"OperatorSequenceReference,omitempty"`
	RewardBasis                    *RewardBasis               `xml:"RewardBasis,omitempty" json:"RewardBasis,omitempty" db:"RewardBasis,omitempty"`
	RewardCategory                 *string                    `xml:"RewardCategory,omitempty" json:"RewardCategory,omitempty" db:"RewardCategory,omitempty"`
	RewardType                     *string                    `xml:"RewardType,omitempty" json:"RewardType,omitempty" db:"RewardType,omitempty"`
	RewardLevel                    *string                    `xml:"RewardLevel,omitempty" json:"RewardLevel,omitempty" db:"RewardLevel,omitempty"`
	AttrAcsSpaceDetailedData       string                     `xml:"detailedData,attr"  json:",omitempty"`
	AttrAcsSpaceMembershipRequired string                     `xml:"membershipRequired,attr"  json:",omitempty"`
	AttrAcsSpaceSummarizedActivity string                     `xml:"summarizedActivity,attr"  json:",omitempty"`
	BaseRewardAmount               *string                    `xml:"BaseRewardAmount,omitempty" json:"BaseRewardAmount,omitempty" db:"BaseRewardAmount,omitempty"`
	CustomOfferID                  *int                       `xml:"CustomOfferID,omitempty" json:"CustomOfferID,omitempty" db:"CustomOfferID,omitempty"`
}

type RewardBasis struct {
	AmountUsed           *string               `xml:"AmountUsed,omitempty" json:"AmountUsed,omitempty" db:"AmountUsed,omitempty"`
	ItemDescription      *string               `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	QuantityUsed         *string               `xml:"QuantityUsed,omitempty" json:"QuantityUsed,omitempty" db:"QuantityUsed,omitempty"`
	WeightUsed           *string               `xml:"WeightUsed,omitempty" json:"WeightUsed,omitempty" db:"WeightUsed,omitempty"`
	XMLName              xml.Name              `xml:"RewardBasis,omitempty" json:"RewardBasis,omitempty"`
}

type OperatorSequenceReference string

type Total struct {
	AttrTotalType string   `xml:" TotalType,attr"  json:",omitempty" db:"total_type,omitempty"`
	Text          string   `xml:",chardata" json:",omitempty" db:"text,omitempty"`
	XMLName       xml.Name `xml:"Total,omitempty" json:"Total,omitempty"`
}

type Tax struct {
	Amount        *string  `xml:"Amount,omitempty" json:"Amount,omitempty" db:"Amount,omitempty"`
	Percent       *string  `xml:"Percent,omitempty" json:"Percent,omitempty" db:"Percent,omitempty"`
	Reason        *string  `xml:"Reason,omitempty" json:"Reason,omitempty" db:"Reason,omitempty"`
	TaxableAmount *string  `xml:"TaxableAmount,omitempty" json:"TaxableAmount,omitempty" db:"TaxableAmount,omitempty"`
	XMLName       xml.Name `xml:"Tax,omitempty" json:"Tax,omitempty"`
	// namespace ACS-IR
	AttrAcsSpaceTaxDescription string `xml:"TaxDescription,attr"  json:",omitempty"`
	AttrAcsSpaceTaxID          string `xml:"TaxID,attr"  json:",omitempty"`
}

type ElectronicSignature struct {
	Svg     *Svg     `xml:"svg,omitempty" json:"svg,omitempty" db:"svg,omitempty"`
	XMLName xml.Name `xml:"ElectronicSignature,omitempty" json:"ElectronicSignature,omitempty"`
}

type Path struct {
	AttrD   string   `xml:" d,attr"  json:",omitempty"`
	XMLName xml.Name `xml:"path,omitempty" json:"path,omitempty"`
}

type LoyaltyMembership struct {
	HouseholdID     *int       `xml:"HouseholdID,omitempty" json:"HouseholdID,omitempty" db:"HouseholdID,omitempty"`
	LoyaltyID       *LoyaltyID `xml:"LoyaltyID,omitempty" json:"LoyaltyID,omitempty" db:"LoyaltyID,omitempty"`
	MembershipID    *int       `xml:"MembershipID,omitempty" json:"MembershipID,omitempty" db:"MembershipID,omitempty"`
	MembershipLevel *string    `xml:"MembershipLevel,omitempty" json:"MembershipLevel,omitempty" db:"MembershipLevel,omitempty"`
	XMLName         xml.Name   `xml:"LoyaltyMembership,omitempty" json:"LoyaltyMembership,omitempty"`
}

type LoyaltyID struct {
	AttrType string   `xml:" Type,attr"  json:",omitempty"`
	Text     string   `xml:",chardata" json:",omitempty"`
	XMLName  xml.Name `xml:"LoyaltyID,omitempty" json:"LoyaltyID,omitempty"`
}

type CRMCustomVariable struct {
	ID      *int     `xml:"ID,omitempty" json:"ID,omitempty" db:"ID,omitempty"`
	Type    *string  `xml:"Type,omitempty" json:"Type,omitempty" db:"Type,omitempty"`
	Value   *string  `xml:"Value,omitempty" json:"Value,omitempty" db:"Value,omitempty"`
	XMLName xml.Name `xml:"CRMCustomVariable,omitempty" json:"CRMCustomVariable,omitempty"`
}

// Coupon is a MFG Coupon
type Coupon struct {
	AttrCouponType string        `xml:" CouponType,attr"  json:",omitempty"`
	ExpirationDate *string       `xml:"ExpirationDate,omitempty" json:"ExpirationDate,omitempty" db:"ExpirationDate,omitempty"`
	PrimaryLabel   *PrimaryLabel `xml:"PrimaryLabel,omitempty" json:"PrimaryLabel,omitempty" db:"PrimaryLabel,omitempty"`
	Quantity       *string       `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"Quantity,omitempty"`
	ScanCode       *string       `xml:"ScanCode,omitempty" json:"ScanCode,omitempty" db:"ScanCode,omitempty"`
	XMLName        xml.Name      `xml:"Coupon,omitempty" json:"Coupon,omitempty"`
	// namespace ACSIR
	Item *Item `xml:"Item,omitempty" json:"Item,omitempty" db:"Item,omitempty"`
}

type PrimaryLabel struct {
	XMLName xml.Name `xml:"PrimaryLabel,omitempty" json:"PrimaryLabel,omitempty"`
}

// Item is a PLU item
type Item struct {
	AttrItemType           string                `xml:" ItemType,attr"  json:",omitempty"`
	Description            *string               `xml:"Description,omitempty" json:"Description,omitempty" db:"Description,omitempty"`
	DiscountAmount         *string               `xml:"DiscountAmount,omitempty" json:"DiscountAmount,omitempty" db:"DiscountAmount,omitempty"`
	ExtendedAmount         *string               `xml:"ExtendedAmount,omitempty" json:"ExtendedAmount,omitempty" db:"ExtendedAmount,omitempty"`
	ExtendedDiscountAmount *string               `xml:"ExtendedDiscountAmount,omitempty" json:"ExtendedDiscountAmount,omitempty" db:"ExtendedDiscountAmount,omitempty"`
	ItemID                 string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	Itemizers              *Itemizers            `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
	MerchandiseHierarchy   *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	OperatorSequence       *int                  `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	POSIdentity            *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	Quantity               *string               `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"Quantity,omitempty"`
	RegularSalesUnitPrice  *string               `xml:"RegularSalesUnitPrice,omitempty" json:"RegularSalesUnitPrice,omitempty" db:"RegularSalesUnitPrice,omitempty"`
	ReportCode             *string               `xml:"ReportCode,omitempty" json:"ReportCode,omitempty" db:"ReportCode,omitempty"`
	SaleableMediaID        *int                  `xml:"SaleableMediaID,omitempty" json:"SaleableMediaID,omitempty" db:"SaleableMediaID,omitempty"`
	XMLName                xml.Name              `xml:"Item,omitempty" json:"Item,omitempty"`
}

type AgeRestriction struct {
	AttrBirthdate string   `xml:" Birthdate,attr"  json:",omitempty"`
	AttrVerified  string   `xml:" Verified,attr"  json:",omitempty"`
	Text          string   `xml:",chardata" json:",omitempty"`
	XMLName       xml.Name `xml:"AgeRestriction,omitempty" json:"AgeRestriction,omitempty"`
}

type CardActivation struct {
	AccountNumber  *string  `xml:"AccountNumber,omitempty" json:"AccountNumber,omitempty" db:"AccountNumber,omitempty"`
	CardType       *string  `xml:"CardType,omitempty" json:"CardType,omitempty" db:"CardType,omitempty"`
	EntryMode      *string  `xml:"EntryMode,omitempty" json:"EntryMode,omitempty" db:"EntryMode,omitempty"`
	ItemID         string   `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	PurchaseAmount *string  `xml:"PurchaseAmount,omitempty" json:"PurchaseAmount,omitempty" db:"PurchaseAmount,omitempty"`
	Track1         *string  `xml:"Track1,omitempty" json:"Track1,omitempty" db:"Track1,omitempty"`
	Track2         *string  `xml:"Track2,omitempty" json:"Track2,omitempty" db:"Track2,omitempty"`
	XMLName        xml.Name `xml:"CardActivation,omitempty" json:"CardActivation,omitempty"`
}

// TenderChange is amount given back on tenders that allow change
type TenderChange struct {
	Amount  *string  `xml:"Amount,omitempty" json:"Amount,omitempty" db:"Amount,omitempty"`
	XMLName xml.Name `xml:"TenderChange,omitempty" json:"TenderChange,omitempty"`
}

type ItemRestriction struct {
	EndDay               *string               `xml:"EndDay,omitempty" json:"EndDay,omitempty" db:"EndDay,omitempty"`
	ItemDescription      *string               `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	XMLName              xml.Name              `xml:"ItemRestriction,omitempty" json:"ItemRestriction,omitempty"`
}

type ControlTransaction struct {
	AttrVersion     string           `xml:" Version,attr"  json:",omitempty"`
	NoSale          *string          `xml:"NoSale,omitempty" json:"NoSale,omitempty" db:"NoSale,omitempty"`
	OperatorSignOff *OperatorSignOff `xml:"OperatorSignOff,omitempty" json:"OperatorSignOff,omitempty" db:"OperatorSignOff,omitempty"`
	OperatorSignOn  *OperatorSignOn  `xml:"OperatorSignOn,omitempty" json:"OperatorSignOn,omitempty" db:"OperatorSignOn,omitempty"`
	PriceLookup     *PriceLookup     `xml:"PriceLookup,omitempty" json:"PriceLookup,omitempty" db:"PriceLookup,omitempty"`
	ReasonCode      *string          `xml:"ReasonCode,omitempty" json:"ReasonCode,omitempty" db:"ReasonCode,omitempty"`
	XMLName         xml.Name         `xml:"ControlTransaction,omitempty" json:"ControlTransaction,omitempty"`
}

type OperatorSignOn struct {
	CloseBusinessDayDate           *string  `xml:"CloseBusinessDayDate,omitempty" json:"CloseBusinessDayDate,omitempty" db:"CloseBusinessDayDate,omitempty"`
	CloseTransactionSequenceNumber *int     `xml:"CloseTransactionSequenceNumber,omitempty" json:"CloseTransactionSequenceNumber,omitempty" db:"CloseTransactionSequenceNumber,omitempty"`
	EndDateTimestamp               *string  `xml:"EndDateTimestamp,omitempty" json:"EndDateTimestamp,omitempty" db:"EndDateTimestamp,omitempty"`
	OpenBusinessDayDate            *string  `xml:"OpenBusinessDayDate,omitempty" json:"OpenBusinessDayDate,omitempty" db:"OpenBusinessDayDate,omitempty"`
	OpenTransactionSequenceNumber  *int     `xml:"OpenTransactionSequenceNumber,omitempty" json:"OpenTransactionSequenceNumber,omitempty" db:"OpenTransactionSequenceNumber,omitempty"`
	StartDateTimestamp             *string  `xml:"StartDateTimestamp,omitempty" json:"StartDateTimestamp,omitempty" db:"StartDateTimestamp,omitempty"`
	XMLName                        xml.Name `xml:"OperatorSignOn,omitempty" json:"OperatorSignOn,omitempty"`
}

type OperatorSignOff struct {
	CloseBusinessDayDate           *string  `xml:"CloseBusinessDayDate,omitempty" json:"CloseBusinessDayDate,omitempty" db:"CloseBusinessDayDate,omitempty"`
	CloseTransactionSequenceNumber *int     `xml:"CloseTransactionSequenceNumber,omitempty" json:"CloseTransactionSequenceNumber,omitempty" db:"CloseTransactionSequenceNumber,omitempty"`
	EndDateTimestamp               *string  `xml:"EndDateTimestamp,omitempty" json:"EndDateTimestamp,omitempty" db:"EndDateTimestamp,omitempty"`
	OpenBusinessDayDate            *string  `xml:"OpenBusinessDayDate,omitempty" json:"OpenBusinessDayDate,omitempty" db:"OpenBusinessDayDate,omitempty"`
	OpenTransactionSequenceNumber  *int     `xml:"OpenTransactionSequenceNumber,omitempty" json:"OpenTransactionSequenceNumber,omitempty" db:"OpenTransactionSequenceNumber,omitempty"`
	StartDateTimestamp             *string  `xml:"StartDateTimestamp,omitempty" json:"StartDateTimestamp,omitempty" db:"StartDateTimestamp,omitempty"`
	XMLName                        xml.Name `xml:"OperatorSignOff,omitempty" json:"OperatorSignOff,omitempty"`
}

type PriceLookup struct {
	ItemCount *int     `xml:"ItemCount,omitempty" json:"ItemCount,omitempty" db:"ItemCount,omitempty"`
	Items     *Items   `xml:"Items,omitempty" json:"Items,omitempty" db:"Items,omitempty"`
	XMLName   xml.Name `xml:"PriceLookup,omitempty" json:"PriceLookup,omitempty"`
}

type Items struct {
	ItemID  string   `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	XMLName xml.Name `xml:"Items,omitempty" json:"Items,omitempty"`
}

type TransactionLink struct {
	AttrEntryMethod string   `xml:" EntryMethod,attr"  json:",omitempty"`
	AttrReasonCode  string   `xml:" ReasonCode,attr"  json:",omitempty"`
	BusinessDayDate *string  `xml:"BusinessDayDate,omitempty" json:"BusinessDayDate,omitempty" db:"BusinessDayDate,omitempty"`
	RetailStoreID   *int     `xml:"RetailStoreID,omitempty" json:"RetailStoreID,omitempty" db:"RetailStoreID,omitempty"`
	SequenceNumber  int      `xml:"SequenceNumber,omitempty" json:"SequenceNumber,omitempty" db:"SequenceNumber,omitempty"`
	WorkstationID   *int     `xml:"WorkstationID,omitempty" json:"WorkstationID,omitempty" db:"WorkstationID,omitempty"`
	XMLName         xml.Name `xml:"TransactionLink,omitempty" json:"TransactionLink,omitempty"`
}

type ItemNotFound struct {
	Disposition          *string               `xml:"Disposition,omitempty" json:"Disposition,omitempty" db:"Disposition,omitempty"`
	ItemDescription      *string               `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	XMLName              xml.Name              `xml:"ItemNotFound,omitempty" json:"ItemNotFound,omitempty"`
}
