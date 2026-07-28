package main

import (
	"deber4-docente_dao/dao"
	"deber4-docente_dao/dataaccess"
	"deber4-docente_dao/model"
	"log"
)

func main() {

	db := dataaccess.InitDB()
	defer db.Close()

	log.Println("Base de datos inicializada correctamente")

	docenteDAO := dao.NewDocenteDAO(db)

	if err := docenteDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}

	docente1 := &model.Docente{
		ID:              "D001",
		Nombre:          "Ana García",
		Email:           "ana.garcia@email.com",
		Departamento:    "Informática",
		Cargo:           "Profesora",
		AniosAntiguedad: 5,
	}
	if err := docenteDAO.Insert(docente1); err != nil {
		log.Printf("Error al insertar: %v", err)
	}

	docente2 := &model.Docente{
		ID:              "D002",
		Nombre:          "Carlos Ruiz",
		Email:           "carlos.ruiz@email.com",
		Departamento:    "Matemáticas",
		Cargo:           "Profesor",
		AniosAntiguedad: 3,
	}
	if err := docenteDAO.Insert(docente2); err != nil {
		log.Printf("Error al insertar: %v", err)
	}

	docente, err := docenteDAO.GetByID("D001")
	if err != nil {
		log.Printf("Error al buscar: %v", err)
	} else {
		log.Printf("Docente encontrado: %+v", docente)
	}

	docentePorEmail, err := docenteDAO.GetByEmail("carlos.ruiz@email.com")
	if err != nil {
		log.Printf("Error al buscar por email: %v", err)
	} else {
		log.Printf("Docente encontrado por email: %+v", docentePorEmail)
	}
}
