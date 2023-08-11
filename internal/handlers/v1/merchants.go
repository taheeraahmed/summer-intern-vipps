package v1_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
	"github.com/vippsas/summerstudents-backend/internal/database/rizz"
)

func MerchantsGetAllHandler(api app.Api) summerstudents_backend_app.MerchantsGetAllV1HandlerFunc {
	return func(merchantsGetAllV1Params summerstudents_backend_app.MerchantsGetAllV1Params) middleware.Responder {
		log := api.Logger
		repo := api.Repository

		// Fetching all merchants' vippsnumbers
		merchants, err := repo.GetAllMerchants()
		if err != nil {
			log.WithFields(logrus.Fields{
				"operation": "MerchantsGetAllHandler",
				"error":     err,
			}).Error("Unable to get all merchants")
			return summerstudents_backend_app.NewMerchantsGetAllV1InternalServerError()
		}

		response := summerstudents_backend_app.NewMerchantsGetAllV1OK()
		response.WithPayload(merchants)
		log.WithFields(logrus.Fields{
			"operation": "MerchantsGetAllHandler",
			"merchants": merchants,
		}).Debug("Successfully fetched all merchants")
		return response
	}
}

func MerchantsGetHandler(api app.Api) summerstudents_backend_app.MerchantsGetV1HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsGetV1Params) middleware.Responder {
		log := api.Logger
		repo := api.Repository
		vippsnumber := params.Vippsnummer

		merchant, err := repo.GetMerchant(vippsnumber)
		if err != nil {
			log.WithFields(logrus.Fields{
				"operation":   "MerchantsGetHandler",
				"error":       err,
				"vippsnummer": vippsnumber,
			}).Error("Unable to get merchant")
			return summerstudents_backend_app.NewMerchantsGetV1InternalServerError()
		}

		if merchant == nil {
			log.WithFields(logrus.Fields{
				"operation":   "MerchantsGetHandler",
				"error":       err,
				"vippsnummer": vippsnumber,
			}).Info("Merchant not found")
			return summerstudents_backend_app.NewMerchantsGetV1NotFound()
		}

		response := summerstudents_backend_app.NewMerchantsGetV1OK()
		response.WithPayload(merchant)

		log.WithFields(logrus.Fields{
			"operation":   "MerchantsGetHandler",
			"vippsnummer": vippsnumber,
			"merchant":    merchant,
		}).Debug("Successfully fetched merchant")
		return response
	}
}

func MerchantsPostHandler(api app.Api) summerstudents_backend_app.MerchantsPostV1HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsPostV1Params) middleware.Responder {
		log := api.Logger
		repo := api.Repository

		err := repo.CreateMerchant(params.Merchant)
		if err != nil {
			// repository.DuplicateMerchantError
			if _, ok := err.(*rizz.DuplicateMerchantError); ok {
				log.WithFields(logrus.Fields{
					"operation": "MerchantsPostHandler",
					"error":     err,
				}).Info("Merchant already exists")
				return summerstudents_backend_app.NewMerchantsPostV1Conflict()
			}
			log.WithFields(logrus.Fields{
				"operation": "MerchantsPostHandler",
				"error":     err,
			}).Error("Unable to create merchant")
			return summerstudents_backend_app.NewMerchantsPostV1InternalServerError()
		}

		log.WithFields(logrus.Fields{
			"operation": "MerchantsPostHandler",
			"merchant":  params.Merchant,
		}).Debug("Successfully created merchant")
		return summerstudents_backend_app.NewMerchantsPostV1Created()
	}
}

func MerchantsPatchHandler(api app.Api) summerstudents_backend_app.MerchantsPatchV1HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsPatchV1Params) middleware.Responder {
		log := api.Logger
		repo := api.Repository
		vippsnumber := params.Vippsnummer
		merchantUpdates := params.Merchant

		merchant, err := repo.UpdateMerchant(vippsnumber, merchantUpdates)
		if err != nil {
			if _, ok := err.(*rizz.NoRowsUpdatedError); ok {
				log.WithFields(logrus.Fields{
					"operation":     "MerchantsPatchHandler",
					"updatedFields": merchantUpdates,
					"vippsnummer":   vippsnumber,
				}).Info("Merchant not found")
				return summerstudents_backend_app.NewMerchantsPatchV1NotFound()
			}
			log.WithFields(logrus.Fields{
				"operation":     "MerchantsPatchHandler",
				"updatedFields": merchantUpdates,
				"vippsnummer":   vippsnumber,
				"error":         err,
			}).Error("Unable to update merchant")
			return summerstudents_backend_app.NewMerchantsPatchV1InternalServerError()
		}

		log.WithFields(logrus.Fields{
			"operation":     "MerchantsPatchHandler",
			"updatedFields": merchantUpdates,
			"vippsnummer":   vippsnumber,
		}).Debug("Successfully updated merchant")

		response := summerstudents_backend_app.NewMerchantsPatchV1OK()
		response.WithPayload(merchant)

		return response
	}
}

func MerchantAgreementsGetAllHandler(api app.Api) summerstudents_backend_app.MerchantAgreementsGetAllV1HandlerFunc {
	return func(params summerstudents_backend_app.MerchantAgreementsGetAllV1Params) middleware.Responder {
		log := api.Logger
		repo := api.Repository
		agreements, err := repo.GetAgreementsByMerchant(params.Vippsnummer)

		if err != nil {
			log.WithFields(logrus.Fields{
				"operation":   "MerchantAgreementsGetAllHandler",
				"vippsnummer": params.Vippsnummer,
				"error":       err,
			}).Error(err)

			return summerstudents_backend_app.NewMerchantAgreementsGetAllV1NotFound()
		}

		response := summerstudents_backend_app.NewMerchantAgreementsGetAllV1OK()

		return response.WithPayload(agreements)
	}
}
