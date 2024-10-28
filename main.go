package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Definición de la estructura Proyecto
type Proyecto struct {
	NumeroControl  string `json:"numero_control"`
	NombreAlumno   string `json:"nombre_alumno"`
	Carrera        string `json:"carrera"`
	NombreProyecto string `json:"nombre_proyecto"`
}

var proyectos []Proyecto

// Función principal
func main() {
	// Valores predeterminados
	proyectos = []Proyecto{
		{NumeroControl: "12345", NombreAlumno: "Juan Pérez", Carrera: "Ingeniería de Software", NombreProyecto: "Sistema de Gestión"},
		{NumeroControl: "67890", NombreAlumno: "María López", Carrera: "Administración", NombreProyecto: "Aplicación Móvil"},
	}

	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/proyectos", GetAllProyectos).Methods("GET")
	router.HandleFunc("/proyectos/{numero_control}", GetProyectoByID).Methods("GET")
	router.HandleFunc("/proyectos", CreateProyecto).Methods("POST")
	router.HandleFunc("/proyectos/{numero_control}", DeleteProyecto).Methods("DELETE")

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Obtener todos los proyectos
func GetAllProyectos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(proyectos)
}

// Obtener un proyecto por número de control
func GetProyectoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numeroControl := vars["numero_control"]

	for _, p := range proyectos {
		if p.NumeroControl == numeroControl {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Proyecto no encontrado", http.StatusNotFound)
}

// Crear un nuevo proyecto
func CreateProyecto(w http.ResponseWriter, r *http.Request) {
	var nuevoProyecto Proyecto
	if err := json.NewDecoder(r.Body).Decode(&nuevoProyecto); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	proyectos = append(proyectos, nuevoProyecto)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoProyecto)
}

// Eliminar un proyecto por número de control
func DeleteProyecto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numeroControl := vars["numero_control"]

	for i, p := range proyectos {
		if p.NumeroControl == numeroControl {
			proyectos = append(proyectos[:i], proyectos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Proyecto no encontrado", http.StatusNotFound)
}
