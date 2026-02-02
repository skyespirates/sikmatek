CREATE TABLE IF NOT EXISTS credit_limits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    requested_limit INT NOT NULL,
    status ENUM('PENDING', 'APPROVED', 'REJECTED') DEFAULT 'PENDING',
    approved_by INT,
    approved_at TIMESTAMP,
    consumer_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (approved_by) REFERENCES users(id),
    CONSTRAINT fk_consumer FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);