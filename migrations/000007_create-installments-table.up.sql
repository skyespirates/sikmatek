CREATE TABLE IF NOT EXISTS installments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nomor_kontrak VARCHAR(128) NOT NULL,
    bulan_ke INT NOT NULL,
    jumlah_tagihan INT,
    due_date DATE,
    status ENUM('UNPAID', 'PAID', 'LATE') DEFAULT 'UNPAID',
    paid_at TIMESTAMP,
    CONSTRAINT fk_contracts FOREIGN KEY (nomor_kontrak) REFERENCES contracts(nomor_kontrak)
);