CREATE TABLE IF NOT EXISTS konsumen_info(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nik VARCHAR(64) NOT NULL,
    full_name VARCHAR(128) NOT NULL,
    legal_name VARCHAR(64) NOT NULL,
    tempat_lahir VARCHAR(64) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    gaji INT NOT NULL,
    foto_ktp VARCHAR(255),
    foto_selfie VARCHAR(255)
);