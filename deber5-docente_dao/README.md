# Deber 5 - DocenteDAO

## Descripción general

Continuación del Deber 4: se implementa nuevamente el patrón **DAO (Data
Access Object)** para la entidad `Docente`, usando `database/sql` y el
driver `github.com/glebarez/sqlite` sobre una base de datos SQLite.

La diferencia principal frente al Deber 4 es el flujo de `main.go`, que
en este caso:

1. Crea la tabla `docentes`.
2. Inserta **un solo docente** (Ana García).
3. Busca ese docente por su `ID`.
4. Intenta buscar por `Email` a un docente que **no fue insertado**
   (Carlos Ruiz), con el fin de comprobar el manejo del error
   `sql.ErrNoRows` en el método `GetByEmail`.

## Estructura del proyecto

\`\`\`
deber5-docente_dao/
├── dao/
│   └── docente_dao.go     # Lógica CRUD sobre la tabla "docentes"
├── dataaccess/
│   └── dataaccess.go      # Inicializa la conexión a la base de datos
├── model/
│   └── docente.go         # Struct Docente
├── main.go                # Punto de entrada
├── go.mod
└── README.md
\`\`\`

## Cómo funciona

### `dataaccess.InitDB()`
Abre la conexión SQLite con `sql.Open("sqlite", "competenciasdocentes.db")`
y verifica que funcione con `db.Ping()`. Devuelve un `*sql.DB` reutilizable
en toda la aplicación.

### `dao.DocenteDAO`
Encapsula el `*sql.DB` y expone las operaciones:

- `CreateTable()`: crea la tabla `docentes` si no existe, usando
  `db.Exec()`.
- `Insert(docente *model.Docente)`: inserta un registro usando consultas
  **parametrizadas** (`?`) para evitar inyección SQL.
- `GetByID(id string)`: busca un docente con `db.QueryRow()` y `Scan()`.
  Si no existe, devuelve un error descriptivo basado en `sql.ErrNoRows`.
- `GetByEmail(email string)`: misma lógica que `GetByID`, pero filtrando
  por la columna `email`.

### `main.go`
Orquesta el flujo: inicializa la base de datos, crea el DAO, crea la
tabla, inserta a Ana García, la busca por ID (éxito esperado) y luego
intenta buscar a Carlos Ruiz por email (error esperado, ya que no fue
insertado).

## Cómo ejecutar

\`\`\`bash
go mod tidy
go run main.go
\`\`\`

## Conceptos clave aplicados

- Patrón DAO para separar la lógica de acceso a datos.
- Uso de `database/sql` junto con el driver `glebarez/sqlite`.
- Consultas parametrizadas para prevenir inyección SQL.
- Manejo idiomático de errores en Go (`fmt.Errorf` + `%w`, `sql.ErrNoRows`).
- Uso de `defer` para el cierre seguro de la conexión.