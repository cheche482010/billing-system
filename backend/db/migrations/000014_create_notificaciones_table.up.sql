CREATE TABLE notificaciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    destinatario_id INT NOT NULL,
    mensaje TEXT NOT NULL,
    fecha_notificacion DATE NOT NULL,
    leido BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (destinatario_id) REFERENCES usuarios(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);