CREATE TABLE bitacora (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tabla_modificada VARCHAR(64) NOT NULL,
    accion VARCHAR(10) NOT NULL CHECK (accion IN ('INSERT', 'UPDATE', 'DELETE')),
    id_registro INT,
    id_usuario INT,
    descripcion_accion TEXT,
    fecha_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_usuario) REFERENCES usuarios(id)
);