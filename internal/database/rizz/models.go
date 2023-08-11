package rizz

import (
	"fmt"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"time"
)

type dbRecurringAgreement struct {
	RecurringAgreementID int64     `json:"recurringAgreementID" db:"recurringAgreementID"`
	Vippsnummer          int64     `json:"vippsnummer" db:"vippsnummer"`
	CustomerID           int64     `json:"customerID" db:"customerID"`
	Amount               int64     `json:"amount" db:"amount"`
	StartDate            time.Time `json:"startDate" db:"startDate"`
	PaymentDay           int64     `json:"paymentDay" db:"paymentDay"`
	IntervalUnit         string    `json:"intervalUnit" db:"intervalUnit"`
	IntervalCount        int64     `json:"intervalCount" db:"intervalCount"`
	AgreementURL         string    `json:"agreementURL" db:"agreementURL"`
	Status               string    `json:"status" db:"status"`
	StatusChangeTime     time.Time `json:"statusChangeTime" db:"statusChangeTime"`
}

type dbMerchant struct {
	Vippsnummer           int64  `json:"vippsnummer" db:"vippsnummer"`
	OrganizationName      string `json:"organizationName" db:"organizationName"`
	SalesUnit             string `json:"salesUnit" db:"salesUnit"`
	Description           string `json:"description" db:"description"`
	LogoURL               string `json:"logoURL" db:"logoURL"`
	CoverURL              string `json:"coverURL" db:"coverURL"`
	RecurringOption       bool   `json:"recurringOption" db:"recurringOption"`
	TermsAndConditionsURL string `json:"termsAndConditionsURL" db:"termsAndConditionsURL"`
	MinimumAmount         int64  `json:"minimumAmount" db:"minimumAmount"`
	Greeting              string `json:"greeting" db:"greeting"`
}

type dbPresetAmount struct {
	Vippsnummer int64 `json:"vippsnummer"`
	Amount      int64 `json:"amount"`
}

func convertAPIMerchantToDBMerchant(merchant *models.Merchant) (*dbMerchant, []int64, error) {
	if merchant == nil || merchant.MerchantName == "" || merchant.Vippsnummer == 0 || merchant.SalesUnit == "" ||
		merchant.Description == "" || merchant.CoverURL == "" || merchant.LogoURL == "" ||
		merchant.TermsConditions == "" || merchant.Offers == nil {

		return nil, nil, fmt.Errorf("one or more fields are missing")
	}
	return &dbMerchant{
		OrganizationName:      merchant.MerchantName,
		Vippsnummer:           merchant.Vippsnummer,
		SalesUnit:             merchant.SalesUnit,
		Description:           merchant.Description,
		CoverURL:              merchant.CoverURL,
		LogoURL:               merchant.LogoURL,
		TermsAndConditionsURL: merchant.TermsConditions,
		RecurringOption:       merchant.HasRecurring,
		MinimumAmount:         merchant.MinimumAmount,
		Greeting:              merchant.Greeting,
	}, merchant.Offers, nil
}

func convertDBMerchantToAPIMerchant(merchant *dbMerchant, offers []int64) *models.Merchant {
	return &models.Merchant{
		CoverURL:        merchant.CoverURL,
		Description:     merchant.Description,
		HasRecurring:    merchant.RecurringOption,
		LogoURL:         merchant.LogoURL,
		MerchantName:    merchant.OrganizationName,
		Offers:          offers,
		SalesUnit:       merchant.SalesUnit,
		TermsConditions: merchant.TermsAndConditionsURL,
		Vippsnummer:     merchant.Vippsnummer,
		MinimumAmount:   merchant.MinimumAmount,
		Greeting:        merchant.Greeting,
	}
}
