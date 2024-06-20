CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    proveedor_id INT NOT NULL,
    producto_id INT NOT NULL,
    fecha_pedido DATE NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (proveedor_id) REFERENCES proveedores(id),
    FOREIGN KEY (producto_id) REFERENCES productos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);