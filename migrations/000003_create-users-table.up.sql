CREATE TABLE IF NOT EXISTS users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    role_id INT NOT NULL,
    consumer_id INT,
    CONSTRAINT pk_roles FOREIGN KEY (role_id) REFERENCES roles(id),
    CONSTRAINT pk_consumers FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);