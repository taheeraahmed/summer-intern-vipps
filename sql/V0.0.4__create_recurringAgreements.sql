CREATE TABLE RIZZ.recurringAgreements(
    recurringAgreementID INTEGER IDENTITY(1,1) NOT NULL,
    vippsnummer INTEGER NOT NULL,
    customerID INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    startDate DATE NOT NULL,
    intervalUnit VARCHAR(255) NOT NULL,
    intervalCount INTEGER NOT NULL,
    agreementURL VARCHAR(4096) NOT NULL,
    status VARCHAR(10) NOT NULL,
    CONSTRAINT PK_recurringAgreementID_recurringAgreements PRIMARY KEY (recurringAgreementID),
    CONSTRAINT FK_vippsnummer_recurringAgreements FOREIGN KEY (vippsnummer) REFERENCES RIZZ.merchants (vippsnummer)
);
