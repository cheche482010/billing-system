CREATE TABLE pagos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_factura_id INT NOT NULL,
    fecha_pago DATE NOT NULL,
    pago_total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (cliente_factura_id) REFERENCES cliente_facturas(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);