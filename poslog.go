// Package poslog works with POSLog XML files.
// Currently writing types for NCR ACS POSLog
// But can be extended for other POS vendores and all
// NCR specific tags are (will be) set as options
package poslog

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"
	"time"
)

// POSLog the main type of a POSLog XMl file. The type contains
// the marshaling information to marshal and unmarshal to json and xml
// currently it is not complete and does not handle all fields in
// source XML
type POSLog struct {
	filename    string
	dayID       *string
	XmlnsAcs    *string        `xml:"xmlns acs,attr,omitempty"  json:",omitempty"`
	XmlnsAcssm  *string        `xml:"xmlns acssm,attr,omitempty"  json:",omitempty"`
	XmlnsAs     *string        `xml:"xmlns as,attr,omitempty"  json:",omitempty"`
	XmlnsMsxsl  *string        `xml:"xmlns msxsl,attr,omitempty"  json:",omitempty"`
	XmlnsPoslog *string        `xml:"xmlns poslog,attr,omitempty"  json:",omitempty"`
	XmlnsRaw    *string        `xml:"xmlns raw,attr,omitempty"  json:",omitempty"`
	Xmlns       *string        `xml:"xmlns,attr,omitempty"  json:",omitempty"`
	XmlnsXsi    *string        `xml:"xmlns xsi,attr,omitempty"  json:",omitempty"`
	Transaction []*Transaction `xml:"http://www.nrf-arts.org/IXRetail/namespace/ Transaction,omitempty" json:"Transaction,omitempty" db:"http://www.nrf-arts.org/IXRetail/namespace/ Transaction,omitempty"`
	XMLName     xml.Name       `xml:"http://www.nrf-arts.org/IXRetail/namespace/ POSLog,omitempty" json:"POSLog,omitempty"`
}

// Transaction is the body of POSLog, each action at the POS is a transaction
type Transaction struct {
	TransactionID      string              `json:"TransactionID" db:"TransactionID"`
	BusinessDayDate    string              `xml:"BusinessDayDate" json:"BusinessDayDate" db:"BusinessDayDate"`
	ControlTransaction *ControlTransaction `xml:"ControlTransaction,omitempty" json:"ControlTransaction,omitempty" db:"ControlTransaction,omitempty"`
	CurrencyCode       *string             `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" db:"CurrencyCode,omitempty"`
	EndDateTime        string              `xml:"EndDateTime" json:"EndDateTime" db:"EndDateTime"`
	OperatorID         *OperatorID
	RetailStoreID      int                `xml:"RetailStoreID" json:"RetailStoreID" db:"RetailStoreID"`
	RetailTransaction  *RetailTransaction `xml:"RetailTransaction,omitempty" json:"RetailTransaction,omitempty" db:"RetailTransaction,omitempty"`
	SequenceNumber     int                `xml:"SequenceNumber" json:"SequenceNumber" db:"SequenceNumber"`
	WorkstationID      int                `xml:"WorkstationID" json:"WorkstationID" db:"WorkstationID"`
	XMLName            xml.Name           `xml:"Transaction,omitempty" json:"Transaction,omitempty"`
}

// OperatorID is the name and into of operator
type OperatorID struct {
	OperatorID   int    `xml:",chardata" json:"OperatorID,omitempty" db:"OperatorID,omitempty"`
	OperatorName string `xml:" OperatorName,attr,omitempty"  json:",omitempty"`
}

// RetailTransaction is any "sale" transaction
type RetailTransaction struct {
	AttrVersion        *string             `xml:" Version,attr,omitempty"  json:",omitempty"`
	ItemCount          *ItemCount          `xml:"ItemCount,omitempty" json:"ItemCount,omitempty" db:"ItemCount,omitempty"`
	LineItem           []*LineItem         `xml:"LineItem,omitempty" json:"LineItem,omitempty" db:"LineItem,omitempty"`
	PerformanceMetrics *PerformanceMetrics `xml:"PerformanceMetrics,omitempty" json:"PerformanceMetrics,omitempty" db:"PerformanceMetrics,omitempty"`
	ReceiptDateTime    string              `xml:"ReceiptDateTime,omitempty" json:"ReceiptDateTime,omitempty" db:"ReceiptDateTime,omitempty"`
	Total              []*Total            `xml:"Total,omitempty" json:"Total,omitempty" db:"Total,omitempty"`
	TransactionCount   *string             `xml:"TransactionCount,omitempty" json:"TransactionCount,omitempty" db:"TransactionCount,omitempty"`
	TransactionLink    *TransactionLink    `xml:"TransactionLink,omitempty" json:"TransactionLink,omitempty" db:"TransactionLink,omitempty"`
	XMLName            xml.Name            `xml:"RetailTransaction,omitempty" json:"RetailTransaction,omitempty"`
}

// LineItem is each line at the register, in order of squence
type LineItem struct {
	AttrEntryMethod              *string              `xml:" EntryMethod,attr,omitempty"  json:",omitempty"`
	AttrAcsSpaceKeyedPrice       *string              `xml:"keyedPrice,attr,omitempty"  json:",omitempty"`
	AttrAcsSpaceOperatorOverride *string              `xml:"OperatorOverride,attr,omitempty"  json:",omitempty"`
	AttrAcsSpacePriceRequired    *string              `xml:"priceRequired,attr,omitempty"  json:",omitempty"`
	AttrVoidFlag                 *string              `xml:" VoidFlag,attr,omitempty"  json:",omitempty"`
	AttrAcsSpaceWeightItem       *string              `xml:"weightItem,attr,omitempty"  json:",omitempty"`
	AgeRestriction               *AgeRestriction      `xml:"AgeRestriction,omitempty" json:"AgeRestriction,omitempty" db:"AgeRestriction,omitempty"`
	CRMCustomVariable            *CRMCustomVariable   `xml:"CRMCustomVariable,omitempty" json:"CRMCustomVariable,omitempty" db:"CRMCustomVariable,omitempty"`
	CardActivation               *CardActivation      `xml:"CardActivation,omitempty" json:"CardActivation,omitempty" db:"CardActivation,omitempty"`
	ElectronicSignature          *ElectronicSignature `xml:"ElectronicSignature,omitempty" json:"ElectronicSignature,omitempty" db:"ElectronicSignature,omitempty"`
	ItemNotFound                 *ItemNotFound        `xml:"ItemNotFound,omitempty" json:"ItemNotFound,omitempty" db:"ItemNotFound,omitempty"`
	ItemRestriction              *ItemRestriction     `xml:"ItemRestriction,omitempty" json:"ItemRestriction,omitempty" db:"ItemRestriction,omitempty"`
	LoyaltyMembership            *LoyaltyMembership   `xml:"LoyaltyMembership,omitempty" json:"LoyaltyMembership,omitempty" db:"LoyaltyMembership,omitempty"`
	LoyaltyReward                *LoyaltyReward       `xml:"LoyaltyReward,omitempty" json:"LoyaltyReward,omitempty" db:"LoyaltyReward,omitempty"`
	Sale                         *Sale                `xml:"Sale,omitempty" json:"Sale,omitempty" db:"Sale,omitempty"`
	SequenceNumber               int                  `xml:"SequenceNumber" json:"SequenceNumber" db:"SequenceNumber"`
	Tax                          *Tax                 `xml:"Tax,omitempty" json:"Tax,omitempty" db:"Tax,omitempty"`
	Tender                       *Tender              `xml:"Tender,omitempty" json:"Tender,omitempty" db:"Tender,omitempty"`
	XMLName                      xml.Name             `xml:"LineItem,omitempty" json:"LineItem,omitempty"`
}

type Sale struct {
	AttrItemType           string                  `xml:" ItemType,attr,omitempty"  json:",omitempty"`
	Description            *Description            `xml:"Description,omitempty" json:"Description,omitempty" db:"Description,omitempty"`
	DiscountAmount         *DiscountAmount         `xml:"DiscountAmount,omitempty" json:"DiscountAmount,omitempty" db:"DiscountAmount,omitempty"`
	ExtendedAmount         *ExtendedAmount         `xml:"ExtendedAmount,omitempty" json:"ExtendedAmount,omitempty" db:"ExtendedAmount,omitempty"`
	ExtendedDiscountAmount *ExtendedDiscountAmount `xml:"ExtendedDiscountAmount,omitempty" json:"ExtendedDiscountAmount,omitempty" db:"ExtendedDiscountAmount,omitempty"`
	ItemID                 string                  `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	Itemizers              *Itemizers              `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
	MerchandiseHierarchy   *MerchandiseHierarchy   `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	OperatorSequence       *OperatorSequence       `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	POSIdentity            *POSIdentity            `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	Quantity               *Quantity               `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"Quantity,omitempty"`
	RegularSalesUnitPrice  *RegularSalesUnitPrice  `xml:"RegularSalesUnitPrice,omitempty" json:"RegularSalesUnitPrice,omitempty" db:"RegularSalesUnitPrice,omitempty"`
	ReportCode             *ReportCode             `xml:"ReportCode,omitempty" json:"ReportCode,omitempty" db:"ReportCode,omitempty"`
	SaleableMediaID        *SaleableMediaID        `xml:"SaleableMediaID,omitempty" json:"SaleableMediaID,omitempty" db:"SaleableMediaID,omitempty"`
	XMLName                xml.Name                `xml:"Sale,omitempty" json:"Sale,omitempty"`
}

type POSIdentity struct {
	AttrPOSIDType string   `xml:"POSIDType,attr"  json:",omitempty"`
	POSItemID     *string  `xml:"POSItemID,omitempty" json:"POSItemID,omitempty" db:"POSItemID,omitempty"`
	Qualifier     *string  `xml:"Qualifier,omitempty" json:"Qualifier,omitempty" db:"Qualifier,omitempty"`
	XMLName       xml.Name `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty"`
}

type MerchandiseHierarchy struct {
	AttrAcsSpaceDepartmentDescription string   `xml:"DepartmentDescription,attr"  json:",omitempty"`
	AttrLevel                         string   `xml:" Level,attr"  json:",omitempty"`
	Text                              string   `xml:",chardata" json:",omitempty"`
	XMLName                           xml.Name `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty"`
}

type Description string

type RegularSalesUnitPrice string

type ExtendedAmount string

type DiscountAmount string

type ExtendedDiscountAmount string

type Quantity string

type OperatorSequence string

type ReportCode string

type SaleableMediaID string

type Itemizers struct {
	AttrFoodStampable string `xml:" FoodStampable,attr"  json:",omitempty"`
	AttrItemizer1     string `xml:" Itemizer1,attr"  json:",omitempty"`
	AttrItemizer2     string `xml:" Itemizer2,attr"  json:",omitempty"`
	AttrTax1          string `xml:" Tax1,attr"  json:",omitempty"`
	Itemizers         string `xml:",chardata" json:",omitempty"`
}

type Tender struct {
	AttrAcsSpaceTenderDescription string            `xml:"TenderDescription,attr"  json:",omitempty"`
	AttrTenderType                string            `xml:" TenderType,attr"  json:",omitempty"`
	AttrTypeCode                  string            `xml:" TypeCode,attr"  json:",omitempty"`
	Amount                        *Amount           `xml:"Amount,omitempty" json:"Amount,omitempty" db:"Amount,omitempty"`
	Authorization                 *Authorization    `xml:"Authorization,omitempty" json:"Authorization,omitempty" db:"Authorization,omitempty"`
	Cashback                      *Cashback         `xml:"Cashback,omitempty" json:"Cashback,omitempty" db:"Cashback,omitempty"`
	Coupon                        *Coupon           `xml:"Coupon,omitempty" json:"Coupon,omitempty" db:"Coupon,omitempty"`
	CreditDebit                   *CreditDebit      `xml:"CreditDebit,omitempty" json:"CreditDebit,omitempty" db:"CreditDebit,omitempty"`
	OperatorSequence              *OperatorSequence `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	TenderChange                  *TenderChange     `xml:"TenderChange,omitempty" json:"TenderChange,omitempty" db:"TenderChange,omitempty"`
	TenderID                      *TenderID         `xml:"TenderID,omitempty" json:"TenderID,omitempty" db:"TenderID,omitempty"`
	XMLName                       xml.Name          `xml:"Tender,omitempty" json:"Tender,omitempty"`
}

type TenderID string

type Amount string

type Authorization struct {
	AttrElectronicSignature string                  `xml:" ElectronicSignature,attr"  json:",omitempty"`
	AttrHostAuthorized      string                  `xml:" HostAuthorized,attr"  json:",omitempty"`
	AuthorizationCode       *AuthorizationCode      `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" db:"AuthorizationCode,omitempty"`
	AuthorizationDateTime   *AuthorizationDateTime  `xml:"AuthorizationDateTime,omitempty" json:"AuthorizationDateTime,omitempty" db:"AuthorizationDateTime,omitempty"`
	AuthorizedChangeAmount  *AuthorizedChangeAmount `xml:"AuthorizedChangeAmount,omitempty" json:"AuthorizedChangeAmount,omitempty" db:"AuthorizedChangeAmount,omitempty"`
	ReferenceNumber         *ReferenceNumber        `xml:"ReferenceNumber,omitempty" json:"ReferenceNumber,omitempty" db:"ReferenceNumber,omitempty"`
	RequestedAmount         *RequestedAmount        `xml:"RequestedAmount,omitempty" json:"RequestedAmount,omitempty" db:"RequestedAmount,omitempty"`
	XMLName                 xml.Name                `xml:"Authorization,omitempty" json:"Authorization,omitempty"`
}

type RequestedAmount string

type AuthorizationCode string

type ReferenceNumber string

type AuthorizationDateTime string

type CreditDebit struct {
	AttrCardType                  string   `xml:" CardType,attr"  json:",omitempty"`
	AttrAcsSpaceCreditDescription string   `xml:"CreditDescription,attr"  json:",omitempty"`
	XMLName                       xml.Name `xml:"CreditDebit,omitempty" json:"CreditDebit,omitempty"`
}

type LoyaltyReward struct {
	AttrAcsSpaceDetailedData       string                     `xml:"detailedData,attr"  json:",omitempty"`
	AttrAcsSpaceMembershipRequired string                     `xml:"membershipRequired,attr"  json:",omitempty"`
	AttrAcsSpaceSummarizedActivity string                     `xml:"summarizedActivity,attr"  json:",omitempty"`
	BaseRewardAmount               *BaseRewardAmount          `xml:"BaseRewardAmount,omitempty" json:"BaseRewardAmount,omitempty" db:"BaseRewardAmount,omitempty"`
	CustomOfferID                  *CustomOfferID             `xml:"CustomOfferID,omitempty" json:"CustomOfferID,omitempty" db:"CustomOfferID,omitempty"`
	EventID                        *EventID                   `xml:"EventID,omitempty" json:"EventID,omitempty" db:"EventID,omitempty"`
	ExtendedRewardAmount           *ExtendedRewardAmount      `xml:"ExtendedRewardAmount,omitempty" json:"ExtendedRewardAmount,omitempty" db:"ExtendedRewardAmount,omitempty"`
	Itemizers                      *Itemizers                 `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
	OperatorSequenceReference      *OperatorSequenceReference `xml:"OperatorSequenceReference,omitempty" json:"OperatorSequenceReference,omitempty" db:"OperatorSequenceReference,omitempty"`
	PromotionID                    *PromotionID               `xml:"PromotionID,omitempty" json:"PromotionID,omitempty" db:"PromotionID,omitempty"`
	ReasonCode                     *ReasonCode                `xml:"ReasonCode,omitempty" json:"ReasonCode,omitempty" db:"ReasonCode,omitempty"`
	RewardBasis                    *RewardBasis               `xml:"RewardBasis,omitempty" json:"RewardBasis,omitempty" db:"RewardBasis,omitempty"`
	RewardCategory                 *RewardCategory            `xml:"RewardCategory,omitempty" json:"RewardCategory,omitempty" db:"RewardCategory,omitempty"`
	RewardLevel                    *RewardLevel               `xml:"RewardLevel,omitempty" json:"RewardLevel,omitempty" db:"RewardLevel,omitempty"`
	RewardType                     *RewardType                `xml:"RewardType,omitempty" json:"RewardType,omitempty" db:"RewardType,omitempty"`
	XMLName                        xml.Name                   `xml:"LoyaltyReward,omitempty" json:"LoyaltyReward,omitempty"`
}

type PromotionID string

type EventID string

type ReasonCode string

type RewardLevel string

type RewardCategory string

type RewardType string

type CustomOfferID string

type BaseRewardAmount string

type ExtendedRewardAmount string

type RewardBasis struct {
	AmountUsed           *AmountUsed           `xml:"AmountUsed,omitempty" json:"AmountUsed,omitempty" db:"AmountUsed,omitempty"`
	ItemDescription      *ItemDescription      `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	QuantityUsed         *QuantityUsed         `xml:"QuantityUsed,omitempty" json:"QuantityUsed,omitempty" db:"QuantityUsed,omitempty"`
	WeightUsed           *WeightUsed           `xml:"WeightUsed,omitempty" json:"WeightUsed,omitempty" db:"WeightUsed,omitempty"`
	XMLName              xml.Name              `xml:"RewardBasis,omitempty" json:"RewardBasis,omitempty"`
}

type QuantityUsed string

type WeightUsed string

type AmountUsed string

type OperatorSequenceReference string

type Total struct {
	AttrTotalType string   `xml:" TotalType,attr"  json:",omitempty"`
	Text          string   `xml:",chardata" json:",omitempty"`
	XMLName       xml.Name `xml:"Total,omitempty" json:"Total,omitempty"`
}

type ItemCount int

type PerformanceMetrics struct {
	IdleTime   *IdleTime   `xml:"IdleTime,omitempty" json:"IdleTime,omitempty" db:"IdleTime,omitempty"`
	RingTime   *RingTime   `xml:"RingTime,omitempty" json:"RingTime,omitempty" db:"RingTime,omitempty"`
	TenderTime *TenderTime `xml:"TenderTime,omitempty" json:"TenderTime,omitempty" db:"TenderTime,omitempty"`
	XMLName    xml.Name    `xml:"PerformanceMetrics,omitempty" json:"PerformanceMetrics,omitempty"`
}

type RingTime string

type IdleTime string

type TenderTime string

type Tax struct {
	AttrAcsSpaceTaxDescription string         `xml:"TaxDescription,attr"  json:",omitempty"`
	AttrAcsSpaceTaxID          string         `xml:"TaxID,attr"  json:",omitempty"`
	Amount                     *Amount        `xml:"Amount,omitempty" json:"Amount,omitempty" db:"Amount,omitempty"`
	Percent                    *Percent       `xml:"Percent,omitempty" json:"Percent,omitempty" db:"Percent,omitempty"`
	Reason                     *Reason        `xml:"Reason,omitempty" json:"Reason,omitempty" db:"Reason,omitempty"`
	TaxableAmount              *TaxableAmount `xml:"TaxableAmount,omitempty" json:"TaxableAmount,omitempty" db:"TaxableAmount,omitempty"`
	XMLName                    xml.Name       `xml:"Tax,omitempty" json:"Tax,omitempty"`
}

type TaxableAmount string

type Percent string

type Reason string

type ElectronicSignature struct {
	Svg     *Svg     `xml:"svg,omitempty" json:"svg,omitempty" db:"svg,omitempty"`
	XMLName xml.Name `xml:"ElectronicSignature,omitempty" json:"ElectronicSignature,omitempty"`
}

type Svg struct {
	AttrHeight string   `xml:" height,attr"  json:",omitempty"`
	AttrStyle  string   `xml:" style,attr"  json:",omitempty"`
	AttrWidth  string   `xml:" width,attr"  json:",omitempty"`
	AttrXmlns  string   `xml:" xmlns,attr"  json:",omitempty"`
	Path       []*Path  `xml:"path,omitempty" json:"path,omitempty" db:"path,omitempty"`
	XMLName    xml.Name `xml:"svg,omitempty" json:"svg,omitempty"`
}

type Path struct {
	AttrD   string   `xml:" d,attr"  json:",omitempty"`
	XMLName xml.Name `xml:"path,omitempty" json:"path,omitempty"`
}

type LoyaltyMembership struct {
	HouseholdID     *HouseholdID     `xml:"HouseholdID,omitempty" json:"HouseholdID,omitempty" db:"HouseholdID,omitempty"`
	LoyaltyID       *LoyaltyID       `xml:"LoyaltyID,omitempty" json:"LoyaltyID,omitempty" db:"LoyaltyID,omitempty"`
	MembershipID    *MembershipID    `xml:"MembershipID,omitempty" json:"MembershipID,omitempty" db:"MembershipID,omitempty"`
	MembershipLevel *MembershipLevel `xml:"MembershipLevel,omitempty" json:"MembershipLevel,omitempty" db:"MembershipLevel,omitempty"`
	XMLName         xml.Name         `xml:"LoyaltyMembership,omitempty" json:"LoyaltyMembership,omitempty"`
}

type LoyaltyID struct {
	AttrType string   `xml:" Type,attr"  json:",omitempty"`
	Text     string   `xml:",chardata" json:",omitempty"`
	XMLName  xml.Name `xml:"LoyaltyID,omitempty" json:"LoyaltyID,omitempty"`
}

type MembershipID string

type HouseholdID string

type MembershipLevel string

type ItemDescription string

type CRMCustomVariable struct {
	ID      *ID      `xml:"ID,omitempty" json:"ID,omitempty" db:"ID,omitempty"`
	Type    *Type    `xml:"Type,omitempty" json:"Type,omitempty" db:"Type,omitempty"`
	Value   *Value   `xml:"Value,omitempty" json:"Value,omitempty" db:"Value,omitempty"`
	XMLName xml.Name `xml:"CRMCustomVariable,omitempty" json:"CRMCustomVariable,omitempty"`
}

type Type string

type ID string

type Value string

type Coupon struct {
	AttrCouponType string          `xml:" CouponType,attr"  json:",omitempty"`
	ExpirationDate *ExpirationDate `xml:"ExpirationDate,omitempty" json:"ExpirationDate,omitempty" db:"ExpirationDate,omitempty"`
	Item           *Item           `xml:"Item,omitempty" json:"Item,omitempty" db:"Item,omitempty"`
	PrimaryLabel   *PrimaryLabel   `xml:"PrimaryLabel,omitempty" json:"PrimaryLabel,omitempty" db:"PrimaryLabel,omitempty"`
	Quantity       *Quantity       `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"Quantity,omitempty"`
	ScanCode       *ScanCode       `xml:"ScanCode,omitempty" json:"ScanCode,omitempty" db:"ScanCode,omitempty"`
	XMLName        xml.Name        `xml:"Coupon,omitempty" json:"Coupon,omitempty"`
}

type PrimaryLabel struct {
	XMLName xml.Name `xml:"PrimaryLabel,omitempty" json:"PrimaryLabel,omitempty"`
}

type ExpirationDate string

type ScanCode string

type Item struct {
	AttrItemType           string                  `xml:" ItemType,attr"  json:",omitempty"`
	Description            *Description            `xml:"Description,omitempty" json:"Description,omitempty" db:"Description,omitempty"`
	DiscountAmount         *DiscountAmount         `xml:"DiscountAmount,omitempty" json:"DiscountAmount,omitempty" db:"DiscountAmount,omitempty"`
	ExtendedAmount         *ExtendedAmount         `xml:"ExtendedAmount,omitempty" json:"ExtendedAmount,omitempty" db:"ExtendedAmount,omitempty"`
	ExtendedDiscountAmount *ExtendedDiscountAmount `xml:"ExtendedDiscountAmount,omitempty" json:"ExtendedDiscountAmount,omitempty" db:"ExtendedDiscountAmount,omitempty"`
	ItemID                 string                  `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	Itemizers              *Itemizers              `xml:"Itemizers,omitempty" json:"Itemizers,omitempty" db:"Itemizers,omitempty"`
	MerchandiseHierarchy   *MerchandiseHierarchy   `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	OperatorSequence       *OperatorSequence       `xml:"OperatorSequence,omitempty" json:"OperatorSequence,omitempty" db:"OperatorSequence,omitempty"`
	POSIdentity            *POSIdentity            `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	Quantity               *Quantity               `xml:"Quantity,omitempty" json:"Quantity,omitempty" db:"Quantity,omitempty"`
	RegularSalesUnitPrice  *RegularSalesUnitPrice  `xml:"RegularSalesUnitPrice,omitempty" json:"RegularSalesUnitPrice,omitempty" db:"RegularSalesUnitPrice,omitempty"`
	ReportCode             *ReportCode             `xml:"ReportCode,omitempty" json:"ReportCode,omitempty" db:"ReportCode,omitempty"`
	SaleableMediaID        *SaleableMediaID        `xml:"SaleableMediaID,omitempty" json:"SaleableMediaID,omitempty" db:"SaleableMediaID,omitempty"`
	XMLName                xml.Name                `xml:"Item,omitempty" json:"Item,omitempty"`
}

type AgeRestriction struct {
	AttrBirthdate string   `xml:" Birthdate,attr"  json:",omitempty"`
	AttrVerified  string   `xml:" Verified,attr"  json:",omitempty"`
	Text          string   `xml:",chardata" json:",omitempty"`
	XMLName       xml.Name `xml:"AgeRestriction,omitempty" json:"AgeRestriction,omitempty"`
}

type CardActivation struct {
	AccountNumber  *AccountNumber  `xml:"AccountNumber,omitempty" json:"AccountNumber,omitempty" db:"AccountNumber,omitempty"`
	CardType       *CardType       `xml:"CardType,omitempty" json:"CardType,omitempty" db:"CardType,omitempty"`
	EntryMode      *EntryMode      `xml:"EntryMode,omitempty" json:"EntryMode,omitempty" db:"EntryMode,omitempty"`
	ItemID         string          `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	PurchaseAmount *PurchaseAmount `xml:"PurchaseAmount,omitempty" json:"PurchaseAmount,omitempty" db:"PurchaseAmount,omitempty"`
	Track1         *Track1         `xml:"Track1,omitempty" json:"Track1,omitempty" db:"Track1,omitempty"`
	Track2         *Track2         `xml:"Track2,omitempty" json:"Track2,omitempty" db:"Track2,omitempty"`
	XMLName        xml.Name        `xml:"CardActivation,omitempty" json:"CardActivation,omitempty"`
}

type PurchaseAmount string

type EntryMode string

type Track1 string

type Track2 string

type AccountNumber string

type CardType string

type Cashback string

type AuthorizedChangeAmount string

type TenderChange struct {
	Amount  *Amount  `xml:"Amount,omitempty" json:"Amount,omitempty" db:"Amount,omitempty"`
	XMLName xml.Name `xml:"TenderChange,omitempty" json:"TenderChange,omitempty"`
}

type ItemRestriction struct {
	EndDay               *EndDay               `xml:"EndDay,omitempty" json:"EndDay,omitempty" db:"EndDay,omitempty"`
	ItemDescription      *ItemDescription      `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	XMLName              xml.Name              `xml:"ItemRestriction,omitempty" json:"ItemRestriction,omitempty"`
}

type EndDay string

type ControlTransaction struct {
	AttrVersion     string           `xml:" Version,attr"  json:",omitempty"`
	NoSale          *NoSale          `xml:"NoSale,omitempty" json:"NoSale,omitempty" db:"NoSale,omitempty"`
	OperatorSignOff *OperatorSignOff `xml:"OperatorSignOff,omitempty" json:"OperatorSignOff,omitempty" db:"OperatorSignOff,omitempty"`
	OperatorSignOn  *OperatorSignOn  `xml:"OperatorSignOn,omitempty" json:"OperatorSignOn,omitempty" db:"OperatorSignOn,omitempty"`
	PriceLookup     *PriceLookup     `xml:"PriceLookup,omitempty" json:"PriceLookup,omitempty" db:"PriceLookup,omitempty"`
	ReasonCode      *ReasonCode      `xml:"ReasonCode,omitempty" json:"ReasonCode,omitempty" db:"ReasonCode,omitempty"`
	XMLName         xml.Name         `xml:"ControlTransaction,omitempty" json:"ControlTransaction,omitempty"`
}

type OperatorSignOn struct {
	CloseBusinessDayDate           *CloseBusinessDayDate           `xml:"CloseBusinessDayDate,omitempty" json:"CloseBusinessDayDate,omitempty" db:"CloseBusinessDayDate,omitempty"`
	CloseTransactionSequenceNumber *CloseTransactionSequenceNumber `xml:"CloseTransactionSequenceNumber,omitempty" json:"CloseTransactionSequenceNumber,omitempty" db:"CloseTransactionSequenceNumber,omitempty"`
	EndDateTimestamp               *EndDateTimestamp               `xml:"EndDateTimestamp,omitempty" json:"EndDateTimestamp,omitempty" db:"EndDateTimestamp,omitempty"`
	OpenBusinessDayDate            *OpenBusinessDayDate            `xml:"OpenBusinessDayDate,omitempty" json:"OpenBusinessDayDate,omitempty" db:"OpenBusinessDayDate,omitempty"`
	OpenTransactionSequenceNumber  *OpenTransactionSequenceNumber  `xml:"OpenTransactionSequenceNumber,omitempty" json:"OpenTransactionSequenceNumber,omitempty" db:"OpenTransactionSequenceNumber,omitempty"`
	StartDateTimestamp             *StartDateTimestamp             `xml:"StartDateTimestamp,omitempty" json:"StartDateTimestamp,omitempty" db:"StartDateTimestamp,omitempty"`
	XMLName                        xml.Name                        `xml:"OperatorSignOn,omitempty" json:"OperatorSignOn,omitempty"`
}

type StartDateTimestamp string

type EndDateTimestamp string

type OpenBusinessDayDate string

type CloseBusinessDayDate string

type OpenTransactionSequenceNumber string

type CloseTransactionSequenceNumber string

type NoSale string

type OperatorSignOff struct {
	CloseBusinessDayDate           *CloseBusinessDayDate           `xml:"CloseBusinessDayDate,omitempty" json:"CloseBusinessDayDate,omitempty" db:"CloseBusinessDayDate,omitempty"`
	CloseTransactionSequenceNumber *CloseTransactionSequenceNumber `xml:"CloseTransactionSequenceNumber,omitempty" json:"CloseTransactionSequenceNumber,omitempty" db:"CloseTransactionSequenceNumber,omitempty"`
	EndDateTimestamp               *EndDateTimestamp               `xml:"EndDateTimestamp,omitempty" json:"EndDateTimestamp,omitempty" db:"EndDateTimestamp,omitempty"`
	OpenBusinessDayDate            *OpenBusinessDayDate            `xml:"OpenBusinessDayDate,omitempty" json:"OpenBusinessDayDate,omitempty" db:"OpenBusinessDayDate,omitempty"`
	OpenTransactionSequenceNumber  *OpenTransactionSequenceNumber  `xml:"OpenTransactionSequenceNumber,omitempty" json:"OpenTransactionSequenceNumber,omitempty" db:"OpenTransactionSequenceNumber,omitempty"`
	StartDateTimestamp             *StartDateTimestamp             `xml:"StartDateTimestamp,omitempty" json:"StartDateTimestamp,omitempty" db:"StartDateTimestamp,omitempty"`
	XMLName                        xml.Name                        `xml:"OperatorSignOff,omitempty" json:"OperatorSignOff,omitempty"`
}

type PriceLookup struct {
	ItemCount *ItemCount `xml:"ItemCount,omitempty" json:"ItemCount,omitempty" db:"ItemCount,omitempty"`
	Items     *Items     `xml:"Items,omitempty" json:"Items,omitempty" db:"Items,omitempty"`
	XMLName   xml.Name   `xml:"PriceLookup,omitempty" json:"PriceLookup,omitempty"`
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
	Disposition          *Disposition          `xml:"Disposition,omitempty" json:"Disposition,omitempty" db:"Disposition,omitempty"`
	ItemDescription      *ItemDescription      `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" db:"ItemDescription,omitempty"`
	ItemID               string                `xml:"ItemID,omitempty" json:"ItemID,omitempty" db:"ItemID,omitempty"`
	MerchandiseHierarchy *MerchandiseHierarchy `xml:"MerchandiseHierarchy,omitempty" json:"MerchandiseHierarchy,omitempty" db:"MerchandiseHierarchy,omitempty"`
	POSIdentity          *POSIdentity          `xml:"POSIdentity,omitempty" json:"POSIdentity,omitempty" db:"POSIdentity,omitempty"`
	XMLName              xml.Name              `xml:"ItemNotFound,omitempty" json:"ItemNotFound,omitempty"`
}

type Disposition string

func (p *POSLog) appendFilename(filename string) {
	p.filename = filename
}

//End return the end DateTime for the transaction
func (tr *Transaction) End() time.Time {
	// tf := "2006-01-02T15:04:05"
	eds := strings.Join([]string{tr.EndDateTime, "07:00"}, "-")
	edd, err := time.Parse(time.RFC3339, eds)
	if err != nil {
		log.Fatalln(err)
	}
	return edd
}

func (tr *Transaction) EndInt() (daytimeid *string) {
	format := "20060102150405"
	dtis := tr.End().Format(format)
	_, err := strconv.Atoi(dtis)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &dtis
}

// BusinessDay returns an int in dayid format YYYYMMDD
// Should actuall Validate this as go time 20060102 like EndInt
func (tr *Transaction) Dayid() (dayid *string) {
	// format := "20061002"
	mdya := strings.Split(string(tr.BusinessDayDate), "-")
	if len(mdya) == 3 {
		mdy := strings.Join(mdya, "")

		_, err := strconv.Atoi(mdy)
		if err != nil {
			log.Println(err)
			return nil
		}
		return &mdy
	}
	return nil
}

func (tr *Transaction) ID() {
	var tida []string
	// first is the enddateint, good for sorting
	tida = append(tida, string(*tr.EndInt()))
	// next the dayid, this is buisness date
	tida = append(tida, string(*tr.Dayid()))
	// next store number
	tida = append(tida, strconv.Itoa(tr.RetailStoreID))
	// Terminal Number
	tida = append(tida, strconv.Itoa(tr.WorkstationID))
	// SequenceNumber
	tida = append(tida, strconv.Itoa(tr.SequenceNumber))

	tr.TransactionID = strings.Join(tida, "-")

	return
}

/*
RetailStoreID
WorkstationID
SequenceNumber

*/
