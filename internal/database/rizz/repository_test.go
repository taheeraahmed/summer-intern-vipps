package rizz

import (
	"database/sql"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/vippsas/summerstudents-backend/generated/models"
	"testing"
	"time"
)

func TestGetMerchant(t *testing.T) {
	// Set up the test database and cleanup function
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)

	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	// Insert test data into the database
	vippsnumber := int64(123456)

	err = createTestMerchant(repo)

	assert.NoError(t, err)

	// Test getting the merchant from the database
	apiMerchant, err := repo.GetMerchant(vippsnumber)
	assert.NoError(t, err)

	fmt.Println("API MERCHANT:")
	fmt.Println(apiMerchant)
	fmt.Println(err)

	// Assert the results
	assert.NotNil(t, apiMerchant)
	assert.Equal(t, vippsnumber, apiMerchant.Vippsnummer)
	assert.Equal(t, "Merchant AS", apiMerchant.MerchantName)
	assert.Equal(t, "Test Sales Unit", apiMerchant.SalesUnit)
	assert.Equal(t, "Test Merchant Description", apiMerchant.Description)
	assert.Equal(t, "https://example.com/logo", apiMerchant.LogoURL)
	assert.Equal(t, "https://example.com/cover", apiMerchant.CoverURL)
	assert.True(t, apiMerchant.HasRecurring)
	assert.Equal(t, "https://example.com/terms", apiMerchant.TermsConditions)
	assert.Equal(t, []int64{100, 200, 300}, apiMerchant.Offers)

	// Test getting a merchant that does not exist
	apiMerchant, err = repo.GetMerchant(789012)
	assert.NoError(t, err)
	assert.Nil(t, apiMerchant)

	// Test getting presetAmounts for a merchant that does not exist
	presetAmounts, err := repo.getMerchantPresetAmounts(789012)
	assert.NoError(t, err)
	assert.Nil(t, presetAmounts)

	// Test getting presetAmounts for a merchant that does exist
	presetAmounts, err = repo.getMerchantPresetAmounts(vippsnumber)
	assert.NoError(t, err)
	assert.NotNil(t, presetAmounts)
	assert.Equal(t, []int64{100, 200, 300}, presetAmounts)
}

func TestGetAllMerchants(t *testing.T) {
	// Set up the test database and cleanup function
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)
	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	// Insert test data into the database
	merchants := []*models.Merchant{
		{
			Vippsnummer:     123456,
			MerchantName:    "Merchant 1",
			SalesUnit:       "Sales Unit 1",
			Description:     "Description 1",
			LogoURL:         "https://example.com/logo1",
			CoverURL:        "https://example.com/cover1",
			HasRecurring:    true,
			Offers:          []int64{100, 200, 300},
			TermsConditions: "https://example.com/terms1",
		},
		{
			Vippsnummer:     789012,
			MerchantName:    "Merchant 2",
			SalesUnit:       "Sales Unit 2",
			Description:     "Description 2",
			LogoURL:         "https://example.com/logo2",
			CoverURL:        "https://example.com/cover2",
			HasRecurring:    false,
			Offers:          []int64{100, 200, 300},
			TermsConditions: "https://example.com/terms2",
		},
	}

	for _, merchant := range merchants {
		err = repo.CreateMerchant(merchant)
		assert.NoError(t, err)
	}

	// Call the function with the test database
	apiMerchants, err := repo.GetAllMerchants()
	assert.NoError(t, err)

	// Assert the results
	assert.Equal(t, 2, len(apiMerchants))

	// Check that the returned merchants match the expected results
	assert.Equal(t, merchants[0].Vippsnummer, apiMerchants[0].Vippsnummer)
	assert.Equal(t, merchants[0].MerchantName, apiMerchants[0].MerchantName)
	assert.Equal(t, merchants[0].SalesUnit, apiMerchants[0].SalesUnit)
	assert.Equal(t, merchants[0].Description, apiMerchants[0].Description)
	assert.Equal(t, merchants[0].LogoURL, apiMerchants[0].LogoURL)
	assert.Equal(t, merchants[0].CoverURL, apiMerchants[0].CoverURL)
	assert.Equal(t, merchants[0].HasRecurring, apiMerchants[0].HasRecurring)
	assert.Equal(t, merchants[0].TermsConditions, apiMerchants[0].TermsConditions)

	assert.Equal(t, merchants[1].Vippsnummer, apiMerchants[1].Vippsnummer)
	assert.Equal(t, merchants[1].MerchantName, apiMerchants[1].MerchantName)
	assert.Equal(t, merchants[1].SalesUnit, apiMerchants[1].SalesUnit)
	assert.Equal(t, merchants[1].Description, apiMerchants[1].Description)
	assert.Equal(t, merchants[1].LogoURL, apiMerchants[1].LogoURL)
	assert.Equal(t, merchants[1].CoverURL, apiMerchants[1].CoverURL)
	assert.Equal(t, merchants[1].HasRecurring, apiMerchants[1].HasRecurring)
	assert.Equal(t, merchants[1].TermsConditions, apiMerchants[1].TermsConditions)
}

func TestCreateMerchant(t *testing.T) {
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)
	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	vippsnummer := int64(123456)

	// check that merchant does not exist
	merchant, err := repo.GetMerchant(vippsnummer)
	assert.Nil(t, merchant)
	assert.NoError(t, err)

	// add new merchant
	testMerchant := &models.Merchant{
		Vippsnummer:     123456,
		MerchantName:    "Merchant AS",
		SalesUnit:       "Test Sales Unit",
		Description:     "Test Merchant Description",
		LogoURL:         "https://example.com/logo",
		CoverURL:        "https://example.com/cover",
		HasRecurring:    true,
		Offers:          []int64{100, 200, 300},
		TermsConditions: "https://example.com/terms",
	}

	err = repo.CreateMerchant(testMerchant)
	assert.NoError(t, err)

	// check that merchant exists
	merchant, err = repo.GetMerchant(vippsnummer)
	assert.NotNil(t, merchant)
	assert.NoError(t, err)

	// check for correct values

	assert.Equal(t, testMerchant.Vippsnummer, merchant.Vippsnummer)
	assert.Equal(t, testMerchant.MerchantName, merchant.MerchantName)
	assert.Equal(t, testMerchant.SalesUnit, merchant.SalesUnit)
	assert.Equal(t, testMerchant.Description, merchant.Description)
	assert.Equal(t, testMerchant.LogoURL, merchant.LogoURL)
	assert.Equal(t, testMerchant.CoverURL, merchant.CoverURL)
	assert.Equal(t, testMerchant.HasRecurring, merchant.HasRecurring)
	assert.Equal(t, testMerchant.TermsConditions, merchant.TermsConditions)
}

func TestCreateMerchant_FailedCreation(t *testing.T) {
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)
	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	vippsnummer := int64(123456)

	// Check that merchant does not exist
	merchant, err := repo.GetMerchant(vippsnummer)
	assert.Nil(t, merchant)
	assert.NoError(t, err)

	// Attempt to add a new merchant with missing fields (failed creation)
	testMerchant := &models.Merchant{
		// Missing required fields, e.g., MerchantName, SalesUnit, Description, etc.
	}

	err = repo.CreateMerchant(testMerchant)
	assert.Error(t, err)
	assert.Nil(t, merchant)

	// Check that merchant still does not exist after the failed creation
	merchant, err = repo.GetMerchant(vippsnummer)
	assert.Nil(t, merchant)
	assert.NoError(t, err)
}

func TestCreateMerchant_DuplicateID(t *testing.T) {
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)
	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	vippsnummer := int64(123456)

	// Create a merchant with the specified Vippsnummer (duplicate ID)
	testMerchant := &models.Merchant{
		Vippsnummer:     vippsnummer,
		MerchantName:    "Merchant 1",
		SalesUnit:       "Test Sales Unit 1",
		Description:     "Test Merchant Description 1",
		LogoURL:         "https://example.com/logo1",
		CoverURL:        "https://example.com/cover1",
		HasRecurring:    true,
		Offers:          []int64{100, 200, 300},
		TermsConditions: "https://example.com/terms1",
	}

	err = repo.CreateMerchant(testMerchant)
	assert.NoError(t, err)
	fmt.Println("added first time")

	err = repo.CreateMerchant(testMerchant)
	fmt.Println(err)
	assert.Error(t, err)
	fmt.Println("added second time")

	// Check that merchant with Vippsnummer (duplicate ID) still exists
	merchant, err := repo.GetMerchant(vippsnummer)
	fmt.Println("get merchant")
	assert.NotNil(t, merchant)
	assert.NoError(t, err)

	// Check for correct values of the existing merchant
	assert.Equal(t, testMerchant.Vippsnummer, merchant.Vippsnummer)
	assert.Equal(t, testMerchant.MerchantName, merchant.MerchantName)
	assert.Equal(t, testMerchant.SalesUnit, merchant.SalesUnit)
	assert.Equal(t, testMerchant.Description, merchant.Description)
	assert.Equal(t, testMerchant.LogoURL, merchant.LogoURL)
	assert.Equal(t, testMerchant.CoverURL, merchant.CoverURL)
	assert.Equal(t, testMerchant.HasRecurring, merchant.HasRecurring)
	assert.Equal(t, testMerchant.TermsConditions, merchant.TermsConditions)
}

func TestConvertAPIMerchantToDBMerchant(t *testing.T) {
	merchant := &models.Merchant{
		MerchantName:    "Test Merchant",
		Vippsnummer:     123456,
		SalesUnit:       "Test Sales Unit",
		Description:     "Test Description",
		CoverURL:        "https://example.com/cover",
		LogoURL:         "https://example.com/logo",
		TermsConditions: "https://example.com/terms",
		HasRecurring:    true,
		Offers:          []int64{100, 200, 300},
	}

	dbMerchant, offers, err := convertAPIMerchantToDBMerchant(merchant)

	assert.NoError(t, err)
	assert.NotNil(t, dbMerchant)
	assert.Equal(t, merchant.MerchantName, dbMerchant.OrganizationName)
	assert.Equal(t, merchant.Vippsnummer, dbMerchant.Vippsnummer)
	assert.Equal(t, merchant.SalesUnit, dbMerchant.SalesUnit)
	assert.Equal(t, merchant.Description, dbMerchant.Description)
	assert.Equal(t, merchant.CoverURL, dbMerchant.CoverURL)
	assert.Equal(t, merchant.LogoURL, dbMerchant.LogoURL)
	assert.Equal(t, merchant.TermsConditions, dbMerchant.TermsAndConditionsURL)
	assert.Equal(t, merchant.HasRecurring, dbMerchant.RecurringOption)
	assert.Equal(t, merchant.Offers, offers)
}

func TestConvertAPIMerchantToDBMerchant_MissingFields(t *testing.T) {
	merchant := &models.Merchant{
		// Missing required fields
	}

	dbMerchant, offers, err := convertAPIMerchantToDBMerchant(merchant)

	assert.Error(t, err)
	assert.Nil(t, dbMerchant)
	assert.Nil(t, offers)
	assert.Equal(t, "one or more fields are missing", err.Error())
}

func TestConvertDBMerchantToAPIMerchant(t *testing.T) {
	dbMerchant := &dbMerchant{
		OrganizationName:      "Test Merchant",
		Vippsnummer:           123456,
		SalesUnit:             "Test Sales Unit",
		Description:           "Test Description",
		CoverURL:              "https://example.com/cover",
		LogoURL:               "https://example.com/logo",
		TermsAndConditionsURL: "https://example.com/terms",
		RecurringOption:       true,
	}

	offers := []int64{100, 200, 300}

	merchant := convertDBMerchantToAPIMerchant(dbMerchant, offers)

	assert.NotNil(t, merchant)
	assert.Equal(t, dbMerchant.OrganizationName, merchant.MerchantName)
	assert.Equal(t, dbMerchant.Vippsnummer, merchant.Vippsnummer)
	assert.Equal(t, dbMerchant.SalesUnit, merchant.SalesUnit)
	assert.Equal(t, dbMerchant.Description, merchant.Description)
	assert.Equal(t, dbMerchant.CoverURL, merchant.CoverURL)
	assert.Equal(t, dbMerchant.LogoURL, merchant.LogoURL)
	assert.Equal(t, dbMerchant.TermsAndConditionsURL, merchant.TermsConditions)
	assert.Equal(t, dbMerchant.RecurringOption, merchant.HasRecurring)
	assert.Equal(t, offers, merchant.Offers)
}

func createTestMerchant(r *Repository) error {
	testMerchant := &models.Merchant{
		Vippsnummer:     123456,
		MerchantName:    "Merchant AS",
		SalesUnit:       "Test Sales Unit",
		Description:     "Test Merchant Description",
		LogoURL:         "https://example.com/logo",
		CoverURL:        "https://example.com/cover",
		HasRecurring:    true,
		Offers:          []int64{100, 200, 300},
		TermsConditions: "https://example.com/terms",
	}

	err := r.CreateMerchant(testMerchant)

	return err
}

func TestCreateAgreement(t *testing.T) {
	// Set up the test database using setupTestRepo.
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)
	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	// Agreement requires a merchant
	err = createTestMerchant(repo)

	assert.NoError(t, err)

	// Create test agreement data.
	agreement := &models.AgreementBody{
		Vippsnummer:   123456,
		CustomerID:    505050,
		Amount:        1000,
		IntervalUnit:  "MONTH",
		IntervalCount: 1,
		AgreementURL:  "https://example.com/agreement",
		Status:        "ACTIVE",
	}

	// Call the CreateRecurringAgreement function.
	id, err := repo.CreateRecurringAgreement(agreement)

	// Check the result.
	assert.NoError(t, err)
	assert.NotEqual(t, int64(-1), id) // Ensure the returned ID is valid.

	// Optional: Query the database to verify the data was inserted correctly.
	// Example: Replace "yourTableName" with the actual table name in your database.
	var dbAgreement dbRecurringAgreement
	query := "SELECT * FROM RIZZ.recurringAgreements WHERE recurringAgreementID = @id"
	err = repo.db.Get(&dbAgreement, query, sql.Named("id", id))
	assert.NoError(t, err)

	// Compare the data in dbAgreement with the test data.
	assert.Equal(t, agreement.Vippsnummer, dbAgreement.Vippsnummer)
	assert.Equal(t, agreement.CustomerID, dbAgreement.CustomerID)
	assert.Equal(t, agreement.Amount, dbAgreement.Amount)
	// ... and so on for other fields.

	// Note: The above verification step is optional and depends on your specific use case.
}

func TestGetRecurringAgreement(t *testing.T) {
	// Set up the test database using setupTestRepo.
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)

	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	// agreement requires a merchant
	err = createTestMerchant(repo)
	assert.NoError(t, err)

	// Insert a test agreement into the database (using your CreateRecurringAgreement function).
	testAgreement := &models.AgreementBody{
		Vippsnummer:   123456,
		CustomerID:    505050,
		Amount:        1000,
		IntervalUnit:  "MONTH",
		IntervalCount: 1,
		AgreementURL:  "https://example.com/agreement",
		Status:        "ACTIVE",
	}

	id, err := repo.CreateRecurringAgreement(testAgreement)

	assert.NoError(t, err)

	// Call the GetRecurringAgreement function.
	agreement, err := repo.GetRecurringAgreement(id)
	// Check the result.
	assert.NoError(t, err)
	assert.NotNil(t, agreement)
	assert.Equal(t, id, agreement.RecurringAgreementID)
	assert.Equal(t, testAgreement.Vippsnummer, agreement.Vippsnummer)
	assert.Equal(t, testAgreement.CustomerID, agreement.CustomerID)
	assert.Equal(t, testAgreement.Amount, agreement.Amount)
	assert.Equal(t, testAgreement.IntervalUnit, agreement.IntervalUnit)
	assert.Equal(t, testAgreement.IntervalCount, agreement.IntervalCount)
	assert.Equal(t, testAgreement.AgreementURL, agreement.AgreementURL)
	assert.Equal(t, testAgreement.Status, agreement.Status)
	//... and so on for other fields.
}

func TestUpdateRecurringAgreement(t *testing.T) {
	// Set up the test repository and obtain the test database
	repo, cleanup, err := SetupTestRepo()
	assert.NoError(t, err)

	defer func() {
		// Clean up the test data after the test.
		err := cleanup()
		assert.NoError(t, err)
	}()

	err = createTestMerchant(repo)
	assert.NoError(t, err)

	// Create a new agreement in the database for testing
	testAgreement := &models.AgreementBody{
		Vippsnummer:   123456,
		CustomerID:    505050,
		Amount:        1000,
		IntervalUnit:  "MONTH",
		IntervalCount: 1,
		AgreementURL:  "https://example.com/agreement",
		Status:        "ACTIVE",
	}
	id, err := repo.CreateRecurringAgreement(testAgreement)
	assert.NoError(t, err)

	// Create an example agreement update body with some changes
	amount := int64(2000)
	status := "CANCELLED"
	recurringAgreementUpdates := &models.AgreementUpdateBody{
		Amount: &amount,
		Status: &status,
	}

	// Call the UpdateRecurringAgreement function
	updatedAgreement, err := repo.UpdateRecurringAgreement(id, recurringAgreementUpdates)
	assert.NoError(t, err)

	fmt.Println("UPDATE!!!")
	fmt.Println(updatedAgreement)
	fmt.Println(testAgreement)
	assert.NoError(t, err)
	assert.NotNil(t, updatedAgreement)

	// Assert that the updates were applied correctly
	assert.Equal(t, *recurringAgreementUpdates.Amount, updatedAgreement.Amount)
	assert.Equal(t, *recurringAgreementUpdates.Status, updatedAgreement.Status)
	assert.Equal(t, testAgreement.Vippsnummer, updatedAgreement.Vippsnummer)
	assert.Equal(t, testAgreement.CustomerID, updatedAgreement.CustomerID)
	// Add more assertions for other fields if needed
}

func TestGetUserRecurringAgreements(t *testing.T) {
	// Set up the test database and cleanup function
	repo, cleanup, err := SetupTestRepo()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer cleanup()

	err = createTestMerchant(repo)
	assert.NoError(t, err)

	// Insert test data into the database
	testAgreement := &models.AgreementBody{
		AgreementURL:  "https://example.com/agreement",
		Amount:        1000,
		CustomerID:    505050,
		IntervalCount: 1,
		IntervalUnit:  "MONTH",
		Status:        "ACTIVE",
		Vippsnummer:   123456,
	}
	id, err := repo.CreateRecurringAgreement(testAgreement)

	assert.NoError(t, err)

	// Call the function with the test database
	agreements, err := repo.GetUserRecurringAgreements(505050)
	assert.NoError(t, err)

	// Assert the results
	assert.Equal(t, 1, len(agreements))

	agreement := agreements[0]

	// Check the individual agreement details
	assert.Equal(t, int64(id), agreement.RecurringAgreementID)
	assert.Equal(t, int64(123456), agreement.Vippsnummer)
	assert.Equal(t, int64(505050), agreement.CustomerID)
	assert.Equal(t, int64(1000), agreement.Amount)
	assert.Equal(t, "MONTH", agreement.IntervalUnit)
	assert.Equal(t, int64(1), agreement.IntervalCount)
	assert.Equal(t, "https://example.com/agreement", agreement.AgreementURL)
	assert.Equal(t, "ACTIVE", agreement.Status)

}

func TestConvertAPIAgreementToDBAgreement(t *testing.T) {
	// Create a sample AgreementBody with test data.
	agreement := &models.AgreementBody{
		Vippsnummer:   123456,
		CustomerID:    505050,
		Amount:        1000,
		IntervalUnit:  "MONTH",
		IntervalCount: 1,
		AgreementURL:  "https://example.com/agreement",
		Status:        "ACTIVE",
	}

	// Call the function being tested.
	result := createDBAgreement(agreement)

	// Assert the converted dbRecurringAgreement.
	expected := &dbRecurringAgreement{
		Vippsnummer:   123456,
		CustomerID:    505050,
		Amount:        1000,
		StartDate:     result.StartDate, // The StartDate will be set dynamically, so just check the other fields.
		IntervalUnit:  "MONTH",
		IntervalCount: 1,
		AgreementURL:  "https://example.com/agreement",
		Status:        "ACTIVE",
	}

	assert.Equal(t, expected.Vippsnummer, result.Vippsnummer)
	assert.Equal(t, expected.CustomerID, result.CustomerID)
	assert.Equal(t, expected.Amount, result.Amount)
	assert.Equal(t, expected.IntervalUnit, result.IntervalUnit)
	assert.Equal(t, expected.IntervalCount, result.IntervalCount)
	assert.Equal(t, expected.AgreementURL, result.AgreementURL)
	assert.Equal(t, expected.Status, result.Status)
}

func TestConvertDBAgreementToAPIAgreement(t *testing.T) {
	// Create a sample dbRecurringAgreement for testing
	now := time.Now()
	dbAgreement := &dbRecurringAgreement{
		AgreementURL:         "https://example.com/agreement",
		Amount:               1000,
		CustomerID:           505050,
		IntervalCount:        1,
		IntervalUnit:         "MONTH",
		StartDate:            now,
		Status:               "ACTIVE",
		Vippsnummer:          123456,
		RecurringAgreementID: 123,
	}

	// Convert the sample dbRecurringAgreement to Agreement
	apiAgreement := convertDBAgreementToAPIAgreement(dbAgreement)

	// Create the expected Agreement object
	expectedAgreement := &models.Agreement{
		AgreementURL:         "https://example.com/agreement",
		Amount:               1000,
		CustomerID:           505050,
		IntervalCount:        1,
		IntervalUnit:         "MONTH",
		StartDate:            strfmt.Date(now),
		Status:               "ACTIVE",
		Vippsnummer:          123456,
		RecurringAgreementID: 123,
	}

	// Assert that the converted Agreement matches the expected Agreement
	assert.Equal(t, expectedAgreement, apiAgreement)
}
