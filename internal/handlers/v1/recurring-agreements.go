package v1_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
	"github.com/vippsas/summerstudents-backend/internal/database/rizz"
)

func RecurringAgreementsGetHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsGetV1HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsGetV1Params) middleware.Responder {
		repo := api.Repository
		log := api.Logger

		agreement, err := repo.GetRecurringAgreement(params.AgreementID)

		if err != nil {
			log.WithFields(logrus.Fields{
				"operation":   "RecurringAgreementsGetV1Handler",
				"agreementID": params.AgreementID,
				"error":       err,
			}).Error("Unable to get agreement")
			return summerstudents_backend_app.NewRecurringAgreementsGetV1InternalServerError()
		}

		if agreement == nil {
			log.WithFields(logrus.Fields{
				"operation":   "RecurringAgreementsGetV1Handler",
				"agreementID": params.AgreementID,
			}).Info("Agreement not found")
			return summerstudents_backend_app.NewRecurringAgreementsGetV1NotFound()
		}

		response := summerstudents_backend_app.NewRecurringAgreementsGetV1OK()
		response.WithPayload(agreement)

		log.WithFields(logrus.Fields{
			"operation":   "RecurringAgreementsGetV1Handler",
			"agreementID": params.AgreementID,
		}).Debug("Agreement returned")
		return response
	}
}

func RecurringAgreementsPatchHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsPatchV1HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsPatchV1Params) middleware.Responder {
		repo := api.Repository
		log := api.Logger
		recurringAgreementUpdates := params.Agreement

		agreement, err := repo.UpdateRecurringAgreement(params.AgreementID, recurringAgreementUpdates)

		if err != nil {
			if _, ok := err.(*rizz.NoRowsUpdatedError); ok {
				log.WithFields(logrus.Fields{
					"operation":     "RecurringAgreementsUpdateV1Handler",
					"updatedFields": recurringAgreementUpdates,
					"agreementID":   params.AgreementID,
				}).Info("Agreement not found")
				return summerstudents_backend_app.NewRecurringAgreementsPatchV1NotFound()
			}

			log.WithFields(logrus.Fields{
				"operation":     "RecurringAgreementsUpdateV1Handler",
				"updatedFields": recurringAgreementUpdates,
				"agreementID":   params.AgreementID,
				"error":         err,
			}).Error("Unable to update agreement")
			return summerstudents_backend_app.NewRecurringAgreementsPatchV1InternalServerError()
		}

		log.WithFields(logrus.Fields{
			"operation":     "RecurringAgreementsUpdateV1Handler",
			"updatedFields": recurringAgreementUpdates,
			"agreementID":   params.AgreementID,
		}).Info("Agreement updated")

		response := summerstudents_backend_app.NewRecurringAgreementsPatchV1OK()
		response.WithPayload(agreement)

		return response
	}
}

func RecurringAgreementsPostHandler(api app.Api) summerstudents_backend_app.RecurringAgreementsPostV1HandlerFunc {
	return func(params summerstudents_backend_app.RecurringAgreementsPostV1Params) middleware.Responder {
		repo := api.Repository
		log := api.Logger

		agreementId, err := repo.CreateRecurringAgreement(params.Agreement)
		if err != nil {
			log.Errorf("Failed to create agreement: %v", err)
			merchant, invalidMerchantError := repo.GetMerchant(params.Agreement.Vippsnummer)

			if merchant == nil {
				log.WithFields(logrus.Fields{
					"operation":  "RecurringAgreementsV1Handler",
					"agreement":  params.Agreement,
					"error":      err,
					"merchantID": params.Agreement.Vippsnummer,
				}).Error("Merchant not found", invalidMerchantError)
				return summerstudents_backend_app.NewRecurringAgreementsPostV1NotFound()
			}

			log.WithFields(logrus.Fields{
				"operation": "RecurringAgreementsV1Handler",
				"agreement": params.Agreement,
				"error":     err,
			}).Error("Unable to create agreement")
			return summerstudents_backend_app.NewRecurringAgreementsPostV1InternalServerError()
		}

		recurringAgreement := models.AgreementReturn{
			AgreementID:          agreementId,
			ChargeID:             "chr_5kSeqz",
			UUID:                 "9c2ca95c-245f-4a2e-aab2-4a08eb78e6fb",
			VippsConfirmationURL: "https://api.vipps.no/v2/register/U6JUjQXq8HQmmV",
		}

		response := summerstudents_backend_app.NewRecurringAgreementsPostV1OK()
		response.WithPayload(&recurringAgreement)

		log.WithFields(logrus.Fields{
			"operation":   "RecurringAgreementsV1Handler",
			"agreement":   recurringAgreement,
			"agreementId": agreementId,
		}).Info("Agreement created")
		return response
	}
}
