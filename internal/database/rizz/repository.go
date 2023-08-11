package rizz

import (
	"database/sql"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) DeleteMerchant(vippsnumber int64) (*models.Merchant, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	query := "DELETE FROM RIZZ.presetAmounts OUTPUT DELETED.amount WHERE vippsnummer = @number"

	var presetAmounts []int64
	err = tx.Select(&presetAmounts, query, sql.Named("number", vippsnumber))
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}

	query = "DELETE RIZZ.merchants OUTPUT DELETED.* FROM RIZZ.Merchants WHERE vippsnummer = @number"
	var deletedDBMerchant dbMerchant
	err = tx.Get(&deletedDBMerchant, query, sql.Named("number", vippsnumber))
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}

	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return nil, rErr
		}
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}

	return convertDBMerchantToAPIMerchant(&deletedDBMerchant, presetAmounts), nil
}

type MerchantNotFoundError struct{}

func (e *MerchantNotFoundError) Error() string {
	return "merchant not found"
}

func offersAreEqual(offers1 []int64, offers2 []int64) bool {
	if len(offers1) != len(offers2) {
		return false
	}
	for i, offer := range offers1 {
		if offer != offers2[i] {
			return false
		}
	}
	return true
}

func (r *Repository) UpdateMerchant(vippsnummer int64, updates *models.MerchantUpdateBody) (*models.Merchant, error) {
	merchant, err := r.getMerchant(vippsnummer)
	if err != nil {
		return nil, err
	}
	if merchant == nil {
		return nil, &MerchantNotFoundError{}
	}

	presetAmounts, err := r.getMerchantPresetAmounts(vippsnummer)
	if err != nil {
		return nil, err
	}

	if updates.MerchantName != nil {
		merchant.OrganizationName = *updates.MerchantName
	}
	if updates.SalesUnit != nil {
		merchant.SalesUnit = *updates.SalesUnit
	}
	if updates.Description != nil {
		merchant.Description = *updates.Description
	}
	if updates.LogoURL != nil {
		merchant.LogoURL = *updates.LogoURL
	}
	if updates.CoverURL != nil {
		merchant.CoverURL = *updates.CoverURL
	}
	if updates.TermsConditions != nil {
		merchant.TermsAndConditionsURL = *updates.TermsConditions
	}
	if updates.HasRecurring != nil {
		merchant.RecurringOption = *updates.HasRecurring
	}
	if updates.MinimumAmount != nil {
		merchant.MinimumAmount = *updates.MinimumAmount
	}

	if updates.Greeting != nil {
		merchant.Greeting = *updates.Greeting
	}

	tsx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	if updates.Offers != nil {
		if !offersAreEqual(presetAmounts, updates.Offers) {
			if err := r.deleteMerchantPresetAmounts(tsx, vippsnummer); err != nil {
				if rErr := tsx.Rollback(); rErr != nil {
					return nil, rErr
				}
				return nil, err
			}
			if err := r.createPresetAmounts(tsx, updates.Offers, vippsnummer); err != nil {
				if rErr := tsx.Rollback(); rErr != nil {
					return nil, rErr
				}
				return nil, err
			}
		}
	}

	query := `
		UPDATE RIZZ.merchants
		SET
		    organizationName=:organizationName,
		    salesUnit=:salesUnit,
		    description=:description,
		    logoURL=:logoURL,
		    coverURL=:coverURL,
		    termsAndConditionsURL=:termsAndConditionsURL,
		    recurringOption=:recurringOption,
		    greeting=:greeting,
		    minimumAmount=:minimumAmount
		
		WHERE vippsnummer=:vippsnummer
		
	`
	result, err := tsx.NamedExec(query, merchant)
	if err != nil {
		if rErr := tsx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		if rErr := tsx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}

	if rowsAffected == 0 {
		if rErr := tsx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, &MerchantNotFoundError{}
	}

	if err := tsx.Commit(); err != nil {
		if rErr := tsx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}

	return convertDBMerchantToAPIMerchant(merchant, updates.Offers), nil
}

func (r *Repository) GetMerchant(vippsnumber int64) (*models.Merchant, error) {
	merchant, err := r.getMerchant(vippsnumber)
	if err != nil {
		return nil, err
	}

	offers, err := r.getMerchantPresetAmounts(vippsnumber)
	if err != nil {
		return nil, err
	}

	if merchant == nil {
		return nil, nil
	}

	return convertDBMerchantToAPIMerchant(merchant, offers), nil
}

func (r *Repository) GetAllMerchants() ([]*models.Merchant, error) {
	// TODO: This is an N+1 query and will slow down as the number of merchants increases

	const query = "SELECT * FROM RIZZ.Merchants"

	var merchants []*dbMerchant
	err := r.db.Select(&merchants, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return []*models.Merchant{}, nil
		} else {
			return nil, fmt.Errorf("error getting all merchants: %w", err)
		}
	}

	var merchantsAPI []*models.Merchant
	for _, merchant := range merchants {
		offers, err := r.getMerchantPresetAmounts(merchant.Vippsnummer)
		if err != nil {
			return nil, fmt.Errorf("error getting preset amounts for merchant %d: %w", merchant.Vippsnummer, err)
		}
		merchantsAPI = append(merchantsAPI, convertDBMerchantToAPIMerchant(merchant, offers))
	}

	return merchantsAPI, nil
}

func (r *Repository) getMerchant(vippsnumber int64) (*dbMerchant, error) {
	const query = "SELECT * FROM RIZZ.Merchants WHERE vippsnummer = @number"

	var merchant dbMerchant
	err := r.db.Get(&merchant, query, sql.Named("number", vippsnumber))

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned - handle this as needed
			return nil, nil
		} else {
			// A different error occurred
			return nil, err
		}
	}
	return &merchant, nil
}

func (r *Repository) getMerchantPresetAmounts(vippsnumber int64) ([]int64, error) {
	const query = "SELECT amount FROM RIZZ.presetAmounts WHERE vippsnummer = @number"

	var offers []int64
	err := r.db.Select(&offers, query, sql.Named("number", vippsnumber))

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned - handle this as needed
			return nil, nil
		} else {
			// A different error occurred
			return nil, err
		}
	}

	return offers, nil
}

func (r *Repository) createPresetAmounts(tx *sqlx.Tx, offers []int64, vippsnummer int64) error {
	const query = `INSERT INTO RIZZ.presetAmounts (amount, vippsnummer) VALUES (:amount, :vippsnummer)`
	for _, amount := range offers {
		_, err := tx.NamedExec(
			query,
			&dbPresetAmount{
				Amount:      amount,
				Vippsnummer: vippsnummer,
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) createMerchant(tx *sqlx.Tx, merchant dbMerchant) error {
	const query = `INSERT INTO RIZZ.merchants
				(organizationName, vippsnummer, salesUnit, description, coverURL, logoURL, termsAndConditionsURL, recurringOption, minimumAmount)
				VALUES (:organizationName, :vippsnummer, :salesUnit, :description, :coverURL, :logoURL, :termsAndConditionsURL, :recurringOption, :minimumAmount)`
	_, err := tx.NamedExec(
		query,
		&merchant)
	return err
}

// DuplicateMerchantError is returned when a merchant with the same vippsnummer already exists
type DuplicateMerchantError struct{}

func (e *DuplicateMerchantError) Error() string {
	return "duplicate merchant"
}

func (r *Repository) CreateMerchant(apiMerchant *models.Merchant) error {
	merchant, offers, err := convertAPIMerchantToDBMerchant(apiMerchant)
	if err != nil {
		return err
	}
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	if err := r.createMerchant(tx, *merchant); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return rErr
		}
		return &DuplicateMerchantError{}
	}
	if err := r.createPresetAmounts(tx, offers, merchant.Vippsnummer); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return rErr
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return rErr
		}
		return err
	}
	return nil
}

func (r *Repository) CreateRecurringAgreement(agreement *models.AgreementBody) (int64, error) {
	dbAgreement := createDBAgreement(agreement)
	rows, err := r.db.NamedQuery(
		`INSERT INTO RIZZ.recurringAgreements (vippsnummer, customerID, amount, startDate, intervalUnit, intervalCount, agreementURL, status, statusChangeTime, paymentDay)
		OUTPUT INSERTED.recurringAgreementID
		VALUES (:vippsnummer, :customerID, :amount, :startDate, :intervalUnit, :intervalCount, :agreementURL, :status, GETDATE(), :paymentDay)`,
		dbAgreement)
	if err != nil {
		return -1, err
	}
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return -1, err
		}
		return -1, fmt.Errorf("no rows returned")
	}
	var id int64
	if err := rows.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

type NoRowsUpdatedError struct{}

func (e *NoRowsUpdatedError) Error() string {
	return "no rows were updated"
}

func (r *Repository) GetRecurringAgreement(id int64) (*models.Agreement, error) {
	query := "SELECT * FROM RIZZ.recurringAgreements WHERE recurringAgreementID = @id"
	var DBAgreement dbRecurringAgreement
	err := r.db.Get(&DBAgreement, query, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return convertDBAgreementToAPIAgreement(&DBAgreement), nil
}

func (r *Repository) DeleteRecurringAgreement(id int64) error {
	query := "UPDATE RIZZ.recurringAgreements SET status = 'STOPPED' WHERE recurringAgreementID = :id"

	namedParams := map[string]interface{}{
		"id": id,
	}

	result, err := r.db.NamedExec(query, namedParams)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return &NoRowsUpdatedError{}
	}
	return nil
}

func (r *Repository) UpdateRecurringAgreement(id int64, recurringAgreementUpdates *models.AgreementUpdateBody) (*models.Agreement, error) {
	query := "SELECT * FROM RIZZ.recurringAgreements WHERE recurringAgreementID = @id"
	var existingRecurringAgreement dbRecurringAgreement
	err := r.db.Get(&existingRecurringAgreement, query, sql.Named("id", id))
	if err != nil {
		// TODO: This doesn't account for the query failing for other reasons than no rows being returned
		return nil, &NoRowsUpdatedError{}
	}
	if recurringAgreementUpdates.Amount != nil {
		existingRecurringAgreement.Amount = *recurringAgreementUpdates.Amount
	}
	if recurringAgreementUpdates.IntervalCount != nil {
		existingRecurringAgreement.IntervalCount = *recurringAgreementUpdates.IntervalCount
	}
	if recurringAgreementUpdates.IntervalUnit != nil {
		existingRecurringAgreement.IntervalUnit = *recurringAgreementUpdates.IntervalUnit
	}
	if recurringAgreementUpdates.Status != nil {
		if existingRecurringAgreement.Status != *recurringAgreementUpdates.Status {
			existingRecurringAgreement.StatusChangeTime = time.Now()
		}
		existingRecurringAgreement.Status = *recurringAgreementUpdates.Status
	}
	if existingRecurringAgreement.IntervalUnit == "MONTH" && recurringAgreementUpdates.PaymentDay != nil {
		existingRecurringAgreement.PaymentDay = *recurringAgreementUpdates.PaymentDay
	}
	query = `
		UPDATE RIZZ.recurringAgreements
		SET
		    amount=:amount,
		    intervalCount=:intervalCount,
		    intervalUnit=:intervalUnit,
		    status=:status,
		    statusChangeTime=:statusChangeTime,
		    paymentDay=:paymentDay
		
		WHERE recurringAgreementID=:recurringAgreementID
	`
	result, err := r.db.NamedExec(query, existingRecurringAgreement)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, err
	} else if rowsAffected > 1 {
		return nil, fmt.Errorf("more than one row was updated")
	}
	return convertDBAgreementToAPIAgreement(&existingRecurringAgreement), nil
}

func (r *Repository) GetUserRecurringAgreements(customerID int64) ([]*dbRecurringAgreement, error) {
	const query = "SELECT * FROM RIZZ.recurringAgreements WHERE customerID = @id"
	var agreements []*dbRecurringAgreement
	err := r.db.Select(&agreements, query, sql.Named("id", customerID))

	if err != nil {
		return nil, err
	}

	return agreements, nil
}

func (r *Repository) deleteMerchantPresetAmounts(tsx *sqlx.Tx, vippsnummer int64) error {
	const query = "DELETE FROM RIZZ.presetAmounts OUTPUT DELETED.amount WHERE vippsnummer = @number"

	var _presetAmounts []int64
	err := tsx.Select(&_presetAmounts, query, sql.Named("number", vippsnummer))
	if err != nil {
		return err
	}

	return nil
}

func createDBAgreement(agreement *models.AgreementBody) *dbRecurringAgreement {
	var paymentDay int64
	if agreement.IntervalUnit == "MONTH" {
		paymentDay = int64(time.Now().Day())
	} else {
		paymentDay = 0
	}

	return &dbRecurringAgreement{
		Vippsnummer:   agreement.Vippsnummer,
		CustomerID:    agreement.CustomerID,
		Amount:        agreement.Amount,
		StartDate:     time.Now(),
		IntervalUnit:  agreement.IntervalUnit,
		IntervalCount: agreement.IntervalCount,
		AgreementURL:  agreement.AgreementURL,
		Status:        agreement.Status,
		PaymentDay:    paymentDay,
	}
}

func convertDBAgreementToAPIAgreement(agreement *dbRecurringAgreement) *models.Agreement {
	return &models.Agreement{
		AgreementURL:         agreement.AgreementURL,
		Amount:               agreement.Amount,
		CustomerID:           agreement.CustomerID,
		IntervalCount:        agreement.IntervalCount,
		IntervalUnit:         agreement.IntervalUnit,
		StartDate:            strfmt.Date(agreement.StartDate),
		Status:               agreement.Status,
		Vippsnummer:          agreement.Vippsnummer,
		RecurringAgreementID: agreement.RecurringAgreementID,
		StatusChangeTime:     strfmt.DateTime(agreement.StatusChangeTime),
		PaymentDay:           agreement.PaymentDay,
	}
}

func (r *Repository) GetAgreementsByMerchant(vippsnummer int64) ([]*models.Agreement, error) {
	query := "SELECT * FROM RIZZ.recurringAgreements WHERE vippsnummer = @vippsnummer"

	var agreements []*dbRecurringAgreement

	err := r.db.Select(&agreements, query, sql.Named("vippsnummer", vippsnummer))

	if err != nil {
		return nil, err
	}

	var result []*models.Agreement

	for _, agreement := range agreements {
		result = append(result, convertDBAgreementToAPIAgreement(agreement))
	}
	return result, nil
}
