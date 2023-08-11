package v0_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
)

func MerchantsGetAllHandler(api app.Api) summerstudents_backend_app.MerchantsGetAllV0HandlerFunc {
	return func(merchantsGetAllParams summerstudents_backend_app.MerchantsGetAllV0Params) middleware.Responder {
		offers0 := []int64{100, 150, 200}
		offers1 := []int64{200, 250, 300}
		offers2 := []int64{50, 250, 300}

		merchant0 := models.Merchant{
			SalesUnit:       "DyrebeskyttelsenNO AS",
			CoverURL:        "https://vipps.no/media/images/Skjermbilde_2023-06-13_kl._1.max-1400x800.jpegquality-60.png",
			HasRecurring:    true,
			LogoURL:         "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			MerchantName:    "Dyrebeskyttelsen Norge",
			Offers:          offers0,
			TermsConditions: "www.dyrebeskyttelsen.no/terms-and-conditions",
			Vippsnummer:     123,
			Description:     "Vi beskytter dyr og sånn XD",
		}

		merchant1 := models.Merchant{
			SalesUnit:       "DyrebeskyttelsenNO AS",
			CoverURL:        "https://vipps.no/media/images/Skjermbilde_2023-06-13_kl._1.max-1400x800.jpegquality-60.png",
			HasRecurring:    true,
			LogoURL:         "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			MerchantName:    "Dyrebeskyttelsen Norge",
			Offers:          offers1,
			TermsConditions: "www.dyrebeskyttelsen.no/terms-and-conditions",
			Vippsnummer:     123,
			Description:     "Vi beskytter dyr og sånn XD",
		}

		merchant2 := models.Merchant{
			SalesUnit:       "DyrebeskyttelsenNO AS",
			CoverURL:        "https://vipps.no/media/images/Skjermbilde_2023-06-13_kl._1.max-1400x800.jpegquality-60.png",
			HasRecurring:    true,
			LogoURL:         "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			MerchantName:    "Dyrebeskyttelsen Norge",
			Offers:          offers2,
			TermsConditions: "www.dyrebeskyttelsen.no/terms-and-conditions",
			Vippsnummer:     123,
			Description:     "Vi beskytter dyr og sånn XD",
		}

		var merchants []*models.Merchant

		merchants = append(merchants, &merchant0, &merchant1, &merchant2)

		log := api.Logger

		response := summerstudents_backend_app.NewMerchantsGetAllV0OK()

		response.WithPayload(merchants)

		log.WithFields(logrus.Fields{
			"operation": "MerchantsGetAllHandler",
		}).Info("Fake merchants returned")
		return response
	}
}

func MerchantsPatchHandler(api app.Api) summerstudents_backend_app.MerchantsPatchV0HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsPatchV0Params) middleware.Responder {
		log := api.Logger

		log.WithFields(logrus.Fields{
			"operation": "MerchantsPatchHandler",
		}).Info("Fake merchant updated")

		response := summerstudents_backend_app.NewMerchantsPatchV0OK()

		response.WithPayload(&models.Merchant{
			SalesUnit:       "DyrebeskyttelsenNO AS",
			CoverURL:        "https://vipps.no/media/images/Skjermbilde_2023-06-13_kl._1.max-1400x800.jpegquality-60.png",
			HasRecurring:    true,
			LogoURL:         "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			MerchantName:    "Dyrebeskyttelsen Norge",
			Offers:          []int64{100, 150, 200},
			TermsConditions: "www.dyrebeskyttelsen.no/terms-and-conditions",
			Vippsnummer:     params.Vippsnummer,
			Description:     "Vi beskytter dyr og sånn XD",
		})

		return response
	}
}

func MerchantsGetHandler(api app.Api) summerstudents_backend_app.MerchantsGetV0HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsGetV0Params) middleware.Responder {
		vippsnumber := params.Vippsnummer
		offers := []int64{100, 150, 200}
		merchant := models.Merchant{
			SalesUnit:       "DyrebeskyttelsenNO AS",
			CoverURL:        "https://vipps.no/media/images/Skjermbilde_2023-06-13_kl._1.max-1400x800.jpegquality-60.png",
			HasRecurring:    true,
			LogoURL:         "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			MerchantName:    "Dyrebeskyttelsen Norge",
			Offers:          offers,
			TermsConditions: "www.dyrebeskyttelsen.no/terms-and-conditions",
			Vippsnummer:     vippsnumber,
			Description:     "Vi beskytter dyr og sånn XD",
		}

		log := api.Logger

		response := summerstudents_backend_app.NewMerchantsGetV0OK()

		response.WithPayload(&merchant)

		log.WithFields(logrus.Fields{
			"operation": "MerchantsGetHandler",
		}).Info("Fake merchant returned")
		return response
	}
}

func MerchantsPostHandler(api app.Api) summerstudents_backend_app.MerchantsPostV0HandlerFunc {
	return func(params summerstudents_backend_app.MerchantsPostV0Params) middleware.Responder {
		log := api.Logger

		log.WithFields(logrus.Fields{
			"operation": "MerchantsPostHandler",
		}).Info("Fake merchant created")
		return summerstudents_backend_app.NewMerchantsPostV0Created()
	}
}

func MerchantAgreementsGetAllHandler(api app.Api) summerstudents_backend_app.MerchantAgreementsGetAllV0HandlerFunc {
	return func(params summerstudents_backend_app.MerchantAgreementsGetAllV0Params) middleware.Responder {
		testAgreement := models.Agreement{
			AgreementURL:         "https://example.com",
			Amount:               300,
			CustomerID:           100304,
			IntervalCount:        1,
			IntervalUnit:         "MONTH",
			RecurringAgreementID: 93219123,
			StartDate:            strfmt.Date{},
			Status:               "ACTIVE",
			StatusChangeTime:     strfmt.DateTime{},
			Vippsnummer:          params.Vippsnummer,
		}

		var items = []*models.Agreement{
			&testAgreement,
			&testAgreement,
			&testAgreement,
		}

		response := summerstudents_backend_app.NewMerchantAgreementsGetAllV0OK()

		response.WithPayload(items)
		return response
	}
}
