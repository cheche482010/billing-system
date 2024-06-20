CREATE TABLE facturas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    num_factura VARCHAR(6) NOT NULL UNIQUE,
    fecha DATE NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);