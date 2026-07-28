package dao

import (
	"database/sql"
	"deber4-docente_dao/model"
	"fmt"
	"log"
)

type DocenteDAO struct {
	db *sql.DB
}

func NewDocenteDAO(db *sql.DB) *DocenteDAO {
	return &DocenteDAO{db: db}
}

func (d *DocenteDAO) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS docentes (
		id TEXT PRIMARY KEY,
		nombre TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		departamento TEXT,
		cargo TEXT,
		anios_antiguedad INTEGER DEFAULT 0
	);`

	_, err := d.db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear tabla docentes: %w", err)
	}

	log.Println("Tabla docentes creada/verificada exitosamente")
	return nil
}

func (d *DocenteDAO) Insert(docente *model.Docente) error {
	query := `
	INSERT INTO docentes (id, nombre, email, departamento, cargo, anios_antiguedad)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := d.db.Exec(query,
		docente.ID,
		docente.Nombre,
		docente.Email,
		docente.Departamento,
		docente.Cargo,
		docente.AniosAntiguedad,
	)

	if err != nil {
		return fmt.Errorf("error al insertar docente: %w", err)
	}

	log.Printf("Docente %s insertado exitosamente", docente.ID)
	return nil
}

func (d *DocenteDAO) GetByID(id string) (*model.Docente, error) {
	query := `
	SELECT id, nombre, email, departamento, cargo, anios_antiguedad
	FROM docentes
	WHERE id = ?`

	row := d.db.QueryRow(query, id)

	var docente model.Docente
	err := row.Scan(
		&docente.ID,
		&docente.Nombre,
		&docente.Email,
		&docente.Departamento,
		&docente.Cargo,
		&docente.AniosAntiguedad,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("docente con ID %s no encontrado", id)
		}
		return nil, fmt.Errorf("error al buscar docente: %w", err)
	}

	return &docente, nil
}

func (d *DocenteDAO) GetByEmail(email string) (*model.Docente, error) {
	query := `
	SELECT id, nombre, email, departamento, cargo, anios_antiguedad
	FROM docentes
	WHERE email = ?`

	row := d.db.QueryRow(query, email)

	var docente model.Docente
	err := row.Scan(
		&docente.ID,
		&docente.Nombre,
		&docente.Email,
		&docente.Departamento,
		&docente.Cargo,
		&docente.AniosAntiguedad,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("docente con email %s no encontrado", email)
		}
		return nil, fmt.Errorf("error al buscar docente: %w", err)
	}

	return &docente, nil
}
