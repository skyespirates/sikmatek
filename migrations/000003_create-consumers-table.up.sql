CREATE TABLE IF NOT EXISTS consumers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nik VARCHAR(16) UNIQUE,
    full_name VARCHAR(128),
    legal_name VARCHAR(64),
    tempat_lahir VARCHAR(128),
    tanggal_lahir DATE,
    gaji INT,
    foto_ktp VARCHAR(128),
    foto_selfie VARCHAR(128),
    is_verified BOOLEAN CHECK (is_verified IN (0,1)) DEFAULT 0,
    user_id INT NOT NULL UNIQUE,
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id)
);