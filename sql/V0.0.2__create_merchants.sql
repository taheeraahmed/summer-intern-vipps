CREATE TABLE RIZZ.Merchants (
    vippsnummer INTEGER NOT NULL,
    organizationName VARCHAR(255) NOT NULL,
    salesUnit VARCHAR(255) NOT NULL,
    description VARCHAR(4096) NOT NULL,
    logoURL VARCHAR(4096),
    coverURL VARCHAR(4096),
    recurringOption BIT NOT NULL,
    termsAndConditionsURL VARCHAR(4096),
    CONSTRAINT PK_vippsnummer PRIMARY KEY (vippsnummer)
);
