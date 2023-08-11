UPDATE [RIZZ].[recurringAgreements]
SET paymentDay = DATEPART(DAY, [startDate])
WHERE intervalUnit = 'MONTH'