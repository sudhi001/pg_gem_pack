package models

// Branche model for handling payment details

type Subscribable struct {
	BaseModel
	UID                               string  `gorm:"index"` // Unique identifier, indexed
	Name                              string  `gorm:"index"` // Name of the branch, indexed
	Code                              string  `gorm:"index"` // Branch code, indexed
	Duration                          int     `json:"duration"`
	AddMonths                         int     `json:"addMonths"`
	SchemeType                        int     `json:"schemeType"`
	AllowVariableInst                 bool    `json:"allowVariableInst"`
	DiscountType                      int     `json:"discountType"`
	DiscountValue                     float64 `json:"discountValue"`
	Purity                            float64 `json:"purity"`
	MinAmount                         float64 `json:"minAmount"`
	MultipleOf                        float64 `json:"multipleOf"`
	SchemeCode                        int     `json:"schemeCode"`
	KycRequired                       bool    `json:"kycRequired"`
	BankAccountDetailsRequired        bool    `json:"bankAccountDetailsRequired"`
	NoOfIntallementsAllowedWithoutKyc int     `json:"noOfIntallementsAllowedWithoutKyc"`
	AllowAfter360DaysEntry            bool    `json:"allowAfter360DaysEntry"`
	DisToBeGivenAfterHowManyMonths    int     `json:"disToBeGivenAfterHowManyMonths"`
	GraceDaysOnLbrChargesDis          int     `json:"graceDaysOnLbrChargesDis"`
	MaturityDate                      string  `json:"maturityDate"`
	SchemeTypeDesc                    string  `json:"schemeTypeDesc"`
	DiscountTypeDesc                  string  `json:"discountTypeDesc"`
	DurationDesc                      string  `json:"durationDesc"`
	SaleRate                          float64 `json:"saleRate"`
	AllowVariableInstallmentUpto      int     `json:"allowVariableInstallmentUpto"`
	MinimumVariableAmountPer          float64 `json:"minimumVariableAmountPer"`
	MaximumVariableAmountPer          float64 `json:"maximumVariableAmountPer"`
	DoNotAllowInstMoreThanDuration    string  `json:"doNotAllowInstMoreThanDurationLookUp"`
	BranchCode                        string  `gorm:"index"` // Branch code this rate applies to, indexed
	Description                       *string `gorm:"default:null"`
	ShortDescription                  *string `gorm:"default:null"`
	TermsAndCondition                *string `gorm:"default:null"`
	Thumbnail                         *string `gorm:"default:null"`
	PrimaryImage                      *string `gorm:"default:null"`
}
