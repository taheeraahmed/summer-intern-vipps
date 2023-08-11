package v1_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
)

type RizzRepository interface {
	GetMerchant(vippsnumber int64) (*models.Merchant, error)
	GetAllMerchants() ([]*models.Merchant, error)
	CreateMerchant(merchant *models.Merchant) error
	DeleteMerchant(vippsnumber int64) (*models.Merchant, error)
	GetRecurringAgreement(id int64) (*models.Agreement, error)
	CreateAgreement(agreement *models.AgreementBody) (int64, error)
	DeleteRecurringAgreement(id int64) error
	UpdateRecurringAgreement(id int64, recurringAgreementUpdates *models.AgreementUpdateBody) error
	GetUserRecurringAgreements(customerID int64) ([]*models.Agreement, error)
}

func CustomerAgreementsGetAllHandler(api app.Api) summerstudents_backend_app.CustomerAgreementsGetAllV1HandlerFunc {
	return func(params summerstudents_backend_app.CustomerAgreementsGetAllV1Params) middleware.Responder {
		logger := api.Logger
		logger.WithFields(logrus.Fields{
			"operation":  "CustomerAgreementsGetAllHandler",
			"customerID": params.CustomerID,
		}).Debug("Getting agreements for customer")

		agreements, err := api.Repository.GetUserRecurringAgreements(params.CustomerID)

		if err != nil {
			logger.WithFields(logrus.Fields{
				"operation":  "CustomerAgreementsGetAllHandler",
				"error":      err,
				"customerID": params.CustomerID,
			}).Error("Unable to get agreements for this user")
			return summerstudents_backend_app.NewCustomerAgreementsGetAllV1InternalServerError()
		}

		active := []*models.AgreementDetails{}
		paused := []*models.AgreementDetails{}
		stopped := []*models.AgreementDetails{}

		for _, agreement := range agreements {
			merchant, err := api.Repository.GetMerchant(agreement.Vippsnummer) // replace this
			if err != nil {
				logger.WithFields(logrus.Fields{
					"operation": "CustomerAgreementsGetAllHandler",
					"error":     err,
				}).Error("Unable to get the merchant attached to the agreement")
				return summerstudents_backend_app.NewCustomerAgreementsGetAllV1InternalServerError()
			}

			agreementDetail := models.AgreementDetails{
				RecurringAgreementID: agreement.RecurringAgreementID,
				Vippsnummer:          agreement.Vippsnummer,
				CustomerID:           agreement.CustomerID,
				Amount:               agreement.Amount,
				StartDate:            strfmt.Date(agreement.StartDate),
				IntervalUnit:         agreement.IntervalUnit,
				IntervalCount:        agreement.IntervalCount,
				AgreementURL:         agreement.AgreementURL,
				Status:               agreement.Status,
				SalesUnit:            merchant.SalesUnit,
				LogoURL:              merchant.LogoURL,
				StatusChangeTime:     strfmt.DateTime(agreement.StatusChangeTime),
				PaymentDay:           agreement.PaymentDay,
			}
			switch agreementDetail.Status {
			case "ACTIVE":
				active = append(active, &agreementDetail)
			case "PAUSED":
				paused = append(paused, &agreementDetail)
			case "STOPPED":
				stopped = append(stopped, &agreementDetail)
			default:
				logger.WithFields(logrus.Fields{
					"operation": "CustomerAgreementsGetAllHandler",
					"msg":       "Invalid status code in agreement",
					"agreement": agreement,
				}).Error("Invalid status code in agreement: " + agreement.Status)
				return summerstudents_backend_app.NewCustomerAgreementsGetAllV1InternalServerError()
			}
		}

		response := summerstudents_backend_app.NewCustomerAgreementsGetAllV1OK()

		body := models.AgreementGroups{
			Active:  active,
			Paused:  paused,
			Stopped: stopped,
		}

		response.WithPayload(&body)

		logger.WithFields(logrus.Fields{
			"operation":  "CustomerAgreementsGetAllHandler",
			"agreements": body,
			"customer":   params.CustomerID,
		}).Debug("Successfully fetched all agreements for customer")
		return response
	}
}
