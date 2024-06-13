-- Tabla Clientes
CREATE TABLE Clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    cedula VARCHAR(20) NOT NULL UNIQUE,
    telefono VARCHAR(15),
    correo VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- Tabla Productos
CREATE TABLE Productos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    precio DECIMAL(10, 2) NOT NULL,
    cantidad INT NOT NULL,
    codigo_barra VARCHAR(50) NOT NULL UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla Facturas
CREATE TABLE Facturas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    num_factura VARCHAR(6) NOT NULL UNIQUE,
    fecha DATE NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Detalle_Factura_Producto (
    id INT AUTO_INCREMENT PRIMARY KEY,
    factura_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,
    precio_unitario DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (factura_id) REFERENCES Facturas(id),
    FOREIGN KEY (producto_id) REFERENCES Productos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE Cliente_Facturas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_id INT NOT NULL,
    factura_id INT NOT NULL,
    fecha_compra DATE NOT NULL,
    FOREIGN KEY (cliente_id) REFERENCES Clientes(id),
    FOREIGN KEY (factura_id) REFERENCES Facturas(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- Tabla Usuarios
CREATE TABLE Usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    contraseña VARCHAR(255) NOT NULL,
    tipo ENUM('cajero', 'admin') NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla Proveedores
CREATE TABLE Proveedores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    contacto VARCHAR(100),
    direccion VARCHAR(255),
    telefono VARCHAR(15),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    proveedor_id INT NOT NULL,
    producto_id INT NOT NULL,
    fecha_pedido DATE NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (proveedor_id) REFERENCES Proveedores(id),
    FOREIGN KEY (producto_id) REFERENCES Productos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Pedido_Producto (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pedido_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,
    FOREIGN KEY (pedido_id) REFERENCES Pedidos(id),
    FOREIGN KEY (producto_id) REFERENCES Productos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla HistorialPrecios (Historial de precios de los productos)
CREATE TABLE HistorialPrecios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    producto_id INT NOT NULL,
    fecha DATE NOT NULL,
    precio DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (producto_id) REFERENCES Productos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla Pagos (Rastreo de pagos recibidos contra las facturas emitidas)
CREATE TABLE Pagos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_factura_id INT NOT NULL,
    fecha_pago DATE NOT NULL,
    pago_total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (cliente_factura_id) REFERENCES Cliente_Facturas(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE MetodoPago (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pago_id INT NOT NULL,
    metodo_pago VARCHAR(50) NOT NULL,
    monto DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (pago_id) REFERENCES Pagos(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla UsuariosRolesPermisos (Roles y permisos de los usuarios)
CREATE TABLE UsuariosRolesPermisos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT NOT NULL,
    rol VARCHAR(20) NOT NULL,
    permiso VARCHAR(255) NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES Usuarios(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabla Notificaciones (Sistema de notificaciones)
CREATE TABLE Notificaciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    destinatario_id INT NOT NULL,
    mensaje TEXT NOT NULL,
    fecha_notificacion DATE NOT NULL,
    leido BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (destinatario_id) REFERENCES Usuarios(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE Bitacora (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tabla_modificada VARCHAR(64) NOT NULL,
    accion VARCHAR(10) NOT NULL CHECK (accion IN ('INSERT', 'UPDATE', 'DELETE')),
    id_registro INT,
    id_usuario INT,
    descripcion_accion TEXT,
    fecha_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_usuario) REFERENCES Usuarios(id)
);
