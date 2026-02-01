CREATE TABLE IF NOT EXISTS limit_usages(
    id INT AUTO_INCREMENT PRIMARY KEY,
    used_amount INT NOT NULL,
    nomor_kontrak VARCHAR(128) NOT NULL,
    limit_id INT NOT NULL,
    CONSTRAINT fk_contract FOREIGN KEY (nomor_kontrak) REFERENCES contracts(nomor_kontrak),
    CONSTRAINT fk_credit_limits FOREIGN KEY (limit_id) REFERENCES credit_limits(id)
);