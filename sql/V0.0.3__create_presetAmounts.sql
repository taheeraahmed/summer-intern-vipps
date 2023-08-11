CREATE TABLE RIZZ.presetAmounts (
    vippsnummer INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    CONSTRAINT FK_vippsnummer FOREIGN KEY (vippsnummer) REFERENCES RIZZ.merchants (vippsnummer)
);