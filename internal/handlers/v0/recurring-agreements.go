package v0_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
	"time"
)

func RecurringAgreementsPatchHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsPatchV0HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsPatchV0Params) middleware.Responder {
		response := summerstudents_backend_app.NewRecurringAgreementsPatchV0OK()

		response.WithPayload(&models.Agreement{
			AgreementURL:     "https://api.vipps.no/v2/agreements/1001",
			Vippsnummer:      1001,
			CustomerID:       123456789,
			Amount:           100,
			StartDate:        strfmt.Date{},
			PaymentDay:       int64(time.Now().Day()),
			IntervalUnit:     "MONTH",
			IntervalCount:    1,
			Status:           "ACTIVE",
			StatusChangeTime: strfmt.DateTime{},
		})

		log := api.Logger

		log.WithFields(logrus.Fields{
			"operation": "RecurringAgreementsPatchHandler",
			"agreement": params.AgreementID,
		}).Info("Fake agreement updated")
		return response
	}
}
func RecurringAgreementsPostHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsPostV0HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsPostV0Params) middleware.Responder {

		recurringAgreement := models.AgreementReturn{
			AgreementID:          1001,
			ChargeID:             "chr_5kSeqz",
			UUID:                 "9c2ca95c-245f-4a2e-aab2-4a08eb78e6fb",
			VippsConfirmationURL: "https://api.vipps.no/v2/register/U6JUjQXq8HQmmV",
		}

		log := api.Logger

		response := summerstudents_backend_app.NewRecurringAgreementsPostV0OK()

		response.WithPayload(&recurringAgreement)

		log.WithFields(logrus.Fields{
			"operation": "RecurringAgreementsPostHandler",
		}).Info("Fake agreement created")
		return response
	}
}

func RecurringAgreementsGetHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsGetV0HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsGetV0Params) middleware.Responder {
		log := api.Logger

		recurringAgreement := models.AgreementReturn{
			AgreementID:          params.AgreementID,
			ChargeID:             "chr_5kSeqz",
			UUID:                 "9c2ca95c-245f-4a2e-aab2-4a08eb78e6fb",
			VippsConfirmationURL: "https://api.vipps.no/v2/register/U6JUjQXq8HQmmV",
		}

		response := summerstudents_backend_app.NewRecurringAgreementsGetV0OK()

		response.WithPayload(&recurringAgreement)

		log.WithFields(logrus.Fields{
			"operation": "RecurringAgreementsGetHandler",
		}).Debug("Fake agreement returned")
		return response
	}
}
