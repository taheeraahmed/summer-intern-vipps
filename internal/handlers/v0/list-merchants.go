package v0_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations/summerstudents_backend_app"
	"github.com/vippsas/summerstudents-backend/internal/app"
)

func CustomerAgreementsGetAllHandler(api app.Api) summerstudents_backend_app.CustomerAgreementsGetAllV0HandlerFunc {
	return func(params summerstudents_backend_app.CustomerAgreementsGetAllV0Params) middleware.Responder {
		active_1 := models.AgreementDetails{
			AgreementURL:         "https://api.vipps.no/v2/agreements/123/charges/chr_5kSeqz",
			Amount:               100,
			CustomerID:           1,
			IntervalCount:        0,
			IntervalUnit:         "MONTH",
			LogoURL:              "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			RecurringAgreementID: 1,
			SalesUnit:            "DyrebeskyttelsenNO AS",
			StartDate:            strfmt.Date{},
			Status:               "ACTIVE",
			Vippsnummer:          123,
			StatusChangeTime:     strfmt.DateTime{},
		}

		active_2 := models.AgreementDetails{
			AgreementURL:         "https://api.vipps.no/v2/agreements/123/charges/chr_5kSeqz",
			Amount:               5000,
			CustomerID:           1,
			IntervalCount:        1,
			IntervalUnit:         "MONTH",
			LogoURL:              "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			RecurringAgreementID: 1,
			SalesUnit:            "DyrebeskyttelsenNO AS",
			StartDate:            strfmt.Date{},
			Status:               "ACTIVE",
			Vippsnummer:          123,
			StatusChangeTime:     strfmt.DateTime{},
		}

		paused_1 := models.AgreementDetails{
			AgreementURL:         "https://api.vipps.no/v2/agreements/123/charges/chr_5kSeqz",
			Amount:               10000,
			CustomerID:           1,
			IntervalCount:        1,
			IntervalUnit:         "MONTH",
			LogoURL:              "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			RecurringAgreementID: 1,
			SalesUnit:            "DyrebeskyttelsenNO AS",
			StartDate:            strfmt.Date{},
			Status:               "PAUSED",
			Vippsnummer:          123,
			StatusChangeTime:     strfmt.DateTime{},
		}

		stopped_1 := models.AgreementDetails{
			AgreementURL:         "https://api.vipps.no/v2/agreements/123/charges/chr_5kSeqz",
			Amount:               2500,
			CustomerID:           1,
			IntervalCount:        1,
			IntervalUnit:         "MONTH",
			LogoURL:              "https://vipps.no/media/images/vipps-rgb-orange-neg.width-400.jpegquality-60.png",
			RecurringAgreementID: 1,
			SalesUnit:            "DyrebeskyttelsenNO AS",
			StartDate:            strfmt.Date{},
			Status:               "STOPPED",
			Vippsnummer:          123,
			StatusChangeTime:     strfmt.DateTime{},
		}

		agreement := models.AgreementGroups{
			Active:  []*models.AgreementDetails{&active_1, &active_2},
			Paused:  []*models.AgreementDetails{&paused_1},
			Stopped: []*models.AgreementDetails{&stopped_1},
		}

		response := summerstudents_backend_app.NewCustomerAgreementsGetAllV0OK()

		response.WithPayload(&agreement)

		log := api.Logger
		log.WithFields(logrus.Fields{
			"operation": "CustomerAgreementsGetAllHandler",
		}).Debug("Fake agreement returned")
		return response
	}
}
