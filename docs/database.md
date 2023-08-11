# RIZZ Database Schema

In the directories `sql` and `testdb` the SQL-scripts to generate the database and populate it with test data are located.

The database schema for the Rizz backend is shown below. The database is hosted on Azure SQL Server.

![db-schema](https://github.com/vippsas/summerstudents-backend/raw/main/docs/images/db-schema.png)

## 1. Merchants Table (`RIZZ.Merchants`)

This table stores information about the merchants.

| Column               | Data Type     | Description                             | Constraints              |
|----------------------|---------------|-----------------------------------------|--------------------------|
| vippsnummer          | INTEGER       | Unique identifier for the merchant      | PRIMARY KEY              |
| organizationName     | VARCHAR(255)  | Name of the organization                | NOT NULL                 |
| salesUnit            | VARCHAR(255)  | Sales unit of the organization          | NOT NULL                 |
| description          | VARCHAR(4096) | Description of the organization         | NOT NULL                 |
| logoURL              | VARCHAR(4096) | URL to the organization's logo          |                          |
| coverURL             | VARCHAR(4096) | URL to the organization's cover image   |                          |
| recurringOption      | BIT           | Option for recurring agreements         | NOT NULL                 |
| termsAndConditionsURL| VARCHAR(4096) | URL to terms and conditions             |                          |
| minimumAmount        | INT           | Minimum amount for transactions         | NOT NULL, DEFAULT 0      |
| greeting             | VARCHAR(200)  | Greeting for the merchant               | NOT NULL, DEFAULT ''    |


## 2. Preset Amounts Table (`RIZZ.presetAmounts`)

This table stores the preset amounts for each merchant.

| Column       | Data Type | Description                                   | Constraints                                      |
|--------------|-----------|-----------------------------------------------|--------------------------------------------------|
| vippsnummer  | INTEGER   | Unique number for the merchant                | FOREIGN KEY REFERENCES `RIZZ.merchants(vippsnummer)`  |
| amount       | INTEGER   | Preset amount                                 | NOT NULL                                         |

## 3. Recurring Agreements Table (`RIZZ.recurringAgreements`)

This table stores information about the recurring agreements.

| Column              | Data Type     | Description                              | Constraints              |
|---------------------|---------------|------------------------------------------|--------------------------|
| recurringAgreementID| INTEGER       | Unique identifier for the agreement      | PRIMARY KEY              |
| vippsnummer         | INTEGER       | Merchant's unique identifier             | FOREIGN KEY              |
| customerID          | INTEGER       | Customer's unique identifier             | NOT NULL                 |
| amount              | INTEGER       | Amount of the agreement                  | NOT NULL                 |
| startDate           | DATE          | Start date of the agreement              | NOT NULL                 |
| intervalUnit        | VARCHAR(255)  | Unit of time for payment intervals       | NOT NULL                 |
| intervalCount       | INTEGER       | Number of intervals between payments     | NOT NULL                 |
| agreementURL        | VARCHAR(4096) | URL related to the agreement             | NOT NULL                 |
| status              | VARCHAR(10)   | Current status of the agreement          | NOT NULL                 |
| statusChangeTime    | DATETIME      | Date and time of status change           | NOT NULL                 |
| paymentDay          | INTEGER       | Day of the month for payments            | Conditional (interval)   |

