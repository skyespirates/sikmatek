CREATE TABLE IF NOT EXISTS limit_usages(
    id INT AUTO_INCREMENT PRIMARY KEY,
    used_amount INT NOT NULL,
    installment_id INT NOT NULL,
    limit_id INT NOT NULL,
    CONSTRAINT fk_installment FOREIGN KEY (installment_id) REFERENCES installments(id),
    CONSTRAINT fk_credit_limits FOREIGN KEY (limit_id) REFERENCES credit_limits(id)
);