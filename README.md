# Sistema de Facturación

Este proyecto es un sistema de facturación completo, construido con tecnologías modernas: un backend en Go y un frontend en React. El sistema permite la gestión de clientes, facturas, y transacciones, ofreciendo una interfaz intuitiva para administrar y realizar seguimiento de las facturas emitidas.

## Estructura del Proyecto

El proyecto está dividido en dos principales partes: el backend en Go y el frontend en React. A continuación, se describe la estructura de carpetas y archivos clave de cada parte:

### Backend (Go)

/backend 
├── controllers 
│ └── clientes_controller.go 
├── db 
│ ├── DB.sql 
│ ├── config.go 
│ ├── crud.go 
│ └── db.go 
├── models 
│ └── clientes_model.go 
├── repositories 
│ └── clientes_repository.go 
├── routes 
│ └── route.go 
├── services │ 
└── clientes_service.go 
└── utils 
├── errors.go 
├── pattern.go 
└── validators.go


- **Controllers**: Contiene los controladores que manejan las solicitudes HTTP.
- **Db**: Gestiona la conexión a la base de datos y operaciones CRUD.
- **Models**: Define los modelos de datos.
- **Repositories**: Implementa la lógica de acceso a datos.
- **Routes**: Define las rutas de la API.
- **Services**: Contiene la lógica de negocio.
- **Utils**: Utilidades compartidas como manejo de errores y validaciones.

### Frontend (React)

/frontend 
├── public 
│ ├── favicon.ico 
│ └──...
├── src 
│ ├── App.js 
│ ├── components 
│ │ └── navbar 
│ │ ├── Navbar.js 
│ │ └── Navbar.scss 
│ ├── index.css 
│ ├── index.js 
│ ├── pages 
│ │ └── login 
│ │ ├── Login.js 
│ │ └── Login.scss 
│ ├── reportWebVitals.js 
│ ├── routes 
│ │ └── routes.js 
│ └──...


- **Public**: Contiene recursos estáticos como imágenes y archivos de configuración web.
- **Src**: La carpeta principal donde se encuentra todo el código JavaScript y React.
  - **Components**: Componentes reutilizables de UI.
  - **Pages**: Páginas individuales de la aplicación.
  - **Routes**: Configuración de rutas de la aplicación.
