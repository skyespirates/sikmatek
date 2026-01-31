CREATE TABLE IF NOT EXISTS transactions(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nomor_kontrak VARCHAR(128) NOT NULL,
    otr INT NOT NULL,
    admin_fee INT NOT NULL,
    jumlah_cicilan INT NOT NULL,
    jumlah_bunga INT NOT NULL,
    nama_asset VARCHAR(128) NOT NULL,
    id_konsumen INT NOT NULL,
    id_tenor INT NOT NULL,
    id_limit INT NOT NULL,
    CONSTRAINT fk_konsumen_info
        FOREIGN KEY (id_konsumen)
        REFERENCES konsumen_info(id),
    CONSTRAINT fk_tenor
        FOREIGN KEY (id_tenor)
        REFERENCES tenors(id),
    CONSTRAINT fk_limit
        FOREIGN KEY (id_limit)
        REFERENCES limit_kredit(id)
);