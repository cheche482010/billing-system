CREATE TABLE cliente_facturas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_id INT NOT NULL,
    factura_id INT NOT NULL,
    fecha_compra DATE NOT NULL,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id),
    FOREIGN KEY (factura_id) REFERENCES facturas(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);