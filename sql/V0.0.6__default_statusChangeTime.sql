UPDATE RIZZ.recurringAgreements
SET statusChangeTime = CURRENT_TIMESTAMP
WHERE statusChangeTime IS NULL;