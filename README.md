
# Sistema de Facturación

Este proyecto es un sistema de facturación completo, construido con tecnologías modernas: un backend en Go y un frontend en React. El sistema permite la gestión de clientes, facturas, y transacciones, ofreciendo una interfaz intuitiva para administrar y realizar seguimiento de las facturas emitidas.

## Estructura del Proyecto

El proyecto está dividido en dos principales partes: el backend en Go y el frontend en React. A continuación, se describe la estructura de carpetas y archivos clave de cada parte:

- **Controllers**
  - `clientes_controller.go`

- **Db**
  - `DB.sql`
  - `config.go`
  - `crud.go`
  - `db.go`

- **Models**
  - `clientes_model.go`

- **Repositories**
  - `clientes_repository.go`

- **Routes**
  - `route.go`

- **Services**
  - `clientes_service.go`

- **Utils**
  - `errors.go`
  - `pattern.go`
  - `validators.go`


#### Descripcion:

- **Controllers**: Contiene los controladores que manejan las solicitudes HTTP.
- **Db**: Gestiona la conexión a la base de datos y operaciones CRUD.
- **Models**: Define los modelos de datos.
- **Repositories**: Implementa la lógica de acceso a datos.
- **Routes**: Define las rutas de la API.
- **Services**: Contiene la lógica de negocio.
- **Utils**: Utilidades compartidas como manejo de errores y validaciones.

### Frontend (React)

- **Frontend**

  - **Public**
    - `favicon.ico`
    -... (otros archivos públicos)

  - **Src**
    - `App.js`
    - **Components**
      - **Navbar**
        - `Navbar.js`
        - `Navbar.scss`
    - `index.css`
    - `index.js`
    - **Pages**
      - **Login**
        - `Login.js`
        - `Login.scss`
    - `reportWebVitals.js`
    - **Routes**
      - `routes.js`
    -... (otros archivos dentro de src)

#### Descripcion:

- **Public**: Contiene recursos estáticos como imágenes y archivos de configuración web.
- **Src**: La carpeta principal donde se encuentra todo el código JavaScript y React.
  - **Components**: Componentes reutilizables de UI.
  - **Pages**: Páginas individuales de la aplicación.
  - **Routes**: Configuración de rutas de la aplicación.