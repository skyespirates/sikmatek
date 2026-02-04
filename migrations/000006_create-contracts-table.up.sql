CREATE TABLE IF NOT EXISTS contracts (
    nomor_kontrak VARCHAR(128) PRIMARY KEY,
    otr INT,
    admin_fee INT,
    jumlah_bunga INT,
    tenor INT CHECK(tenor IN (1, 2, 3, 6)),
    total_pembiayaan INT GENERATED ALWAYS AS (otr+admin_fee+jumlah_bunga) STORED,
    status ENUM('DRAFT', 'QUOTED', 'CANCELLED', 'CONFIRMED', 'ACTIVE', 'COMPLETED', 'DEFAULT') DEFAULT 'DRAFT',
    consumer_id INT NOT NULL,
    product_id INT NOT NULL,
    limit_id INT NOT NULL,
    CONSTRAINT fk_consumers FOREIGN KEY (consumer_id) REFERENCES consumers(id),
    CONSTRAINT fk_products FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_limits FOREIGN KEY (limit_id) REFERENCES credit_limits(id)
);