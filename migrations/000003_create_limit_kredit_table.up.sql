CREATE TABLE IF NOT EXISTS limit_kredit(
    id INT AUTO_INCREMENT PRIMARY KEY,
    limit_nominal INT NOT NULL,
    limit_terpakai INT NOT NULL,
    id_konsumen INT NOT NULL,
    id_tenor INT NOT NULL,
    CONSTRAINT fk_konsumen_info
        FOREIGN KEY (id_konsumen)
        REFERENCES konsumen_info(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_tenor
        FOREIGN KEY (id_tenor)
        REFERENCES tenors(id)
        ON DELETE CASCADE
);