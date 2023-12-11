package main

import (
	"ProyectoFase1/Cola"
	"ProyectoFase1/Listas"
	"fmt"
)

var colaPrioridad *Cola.Cola = &Cola.Cola{Primero: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}

func main() {
	opcion := 0
	bandera := true
	for bandera {
		fmt.Println("1-. Iniciar Sesión")
		fmt.Println("2-. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			MenuLogin()
		case 2:
			bandera = false
		}
	}
}

func MenuLogin() {
	usuario := ""
	password := ""
	fmt.Println("----------------------------------")
	fmt.Print("Ingrese su Usuario: ")
	fmt.Scanln(&usuario)
	fmt.Print("Ingrese su Contraseña: ")
	fmt.Scanln(&password)
	fmt.Println("----------------------------------")

	if usuario == "ADMIN_202100639" && password == "admin" {
		fmt.Println("--------Bienvenido ADMIN_202100639----------")
		MenuAdmin()
	} else if listaDoble.Buscar(usuario, password) {
		fmt.Println("--------Bienvenido alumno ", usuario, "-----------")
		MenuAlumno()
	} else {
		fmt.Println("¡ERROR EN CREDENCIALES!")
	}
}

func MenuAdmin() {
	opcion := 0
	bandera := true
	for bandera {
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar Cursos al Sistema")
		fmt.Println("4. Control de Estudiantes Tutores")
		fmt.Println("5. Reportes")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			fmt.Println("Se cargo los cursos")
		case 4:
			ControlEstudiantes()
		case 5:
			listaDoble.Reporte()
			listaDobleCircular.Reporte()
		case 6:
			bandera = false
		}

	}
}

func MenuAlumno() {
	opcion := 0
	bandera := true
	for bandera {
		fmt.Println("1-. Ver tutores disponibles")
		fmt.Println("2-. Asignarse a Tutores")
		fmt.Println("3-. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			CargaTutores()
		case 2:
			CargaEstudiantes()
		case 3:
			bandera = false
		}

	}
}

func CargaTutores() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	colaPrioridad.LeerCSV(ruta)
	fmt.Println("Se cargo a la Cola los tutores")
}

func CargaEstudiantes() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	listaDoble.LeerCSV(ruta)
	fmt.Println("Se cargo los estudiantes")
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("═══════════════════")
		colaPrioridad.Primero_Cola()
		fmt.Println("═══════════════════")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			listaDobleCircular.Agregar(colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nombre, colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Nota)
			colaPrioridad.Descolar()
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}
