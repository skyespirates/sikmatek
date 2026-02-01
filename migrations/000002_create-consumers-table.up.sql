CREATE TABLE IF NOT EXISTS consumers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nik VARCHAR(16) NOT NULL UNIQUE,
    full_name VARCHAR(128) NOT NULL,
    legal_name VARCHAR(64),
    tempat_lahir VARCHAR(128),
    tanggal_lahir DATE,
    gaji INT,
    foto_ktp VARCHAR(128),
    foto_selfie VARCHAR(128)
);