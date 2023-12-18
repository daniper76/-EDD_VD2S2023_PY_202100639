package main

import (
	"ProyectoFase1/ArbolAVL"
	"ProyectoFase1/Cola"
	"ProyectoFase1/Listas"
	"ProyectoFase1/MatrizDispersa"
	"fmt"
	"strconv"
)

var colaPrioridad *Cola.Cola = &Cola.Cola{Primero: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.Matriz = &MatrizDispersa.Matriz{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolCursos *ArbolAVL.ArbolAVL = &ArbolAVL.ArbolAVL{Raiz: nil}
var loggeado_estudiante string = ""

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
		loggeado_estudiante = usuario
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
			CargarCursos()
		case 4:
			ControlEstudiantes()
		case 5:
			listaDoble.Reporte()
			listaDobleCircular.Reporte()
			matrizDispersa.Reporte()
			arbolCursos.Graficar()
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
			listaDobleCircular.Mostrar()
		case 2:
			AsignarCurso()
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

func CargarCursos() {
	ruta := ""
	fmt.Print("Nombre de Archivo: ")
	fmt.Scanln(&ruta)
	arbolCursos.LeerJson(ruta)
	fmt.Println("Se cargo los cursos")
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
			existe_tutor := listaDobleCircular.BuscarTutor(colaPrioridad.Primero.Tutor.Curso)
			if existe_tutor != nil {
				if existe_tutor.Tutor.Nota < colaPrioridad.Primero.Tutor.Nota {
					listaDobleCircular.CambiarTutor(colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nota, colaPrioridad.Primero.Tutor.Nombre)
					colaPrioridad.Descolar()
					fmt.Println("Se sustituyó tutor de curso actual")
				} else {
					fmt.Println("Tutor actual tiene mejor nota")
					colaPrioridad.Descolar()
				}
			} else {
				listaDobleCircular.Agregar(colaPrioridad.Primero.Tutor.Carnet, colaPrioridad.Primero.Tutor.Nombre, colaPrioridad.Primero.Tutor.Curso, colaPrioridad.Primero.Tutor.Nota)
				colaPrioridad.Descolar()
			}
		} else if opcion == 2 {
			colaPrioridad.Descolar()
		} else if opcion == 3 {
			salir = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}

func AsignarCurso() {
	opcion := ""
	salir := false
	for !salir {
		fmt.Println("Teclee el codigo del curso: ")
		fmt.Scanln(&opcion)
		//Iria el primer If del Arbol (pendiente)
		if arbolCursos.Busqueda(opcion) {
			if listaDobleCircular.Buscar(opcion) {
				TutorBuscado := listaDobleCircular.BuscarTutor(opcion)
				estudiante, err := strconv.Atoi(loggeado_estudiante)
				if err != nil {
					break
				}
				matrizDispersa.Insertar_Elemento(estudiante, TutorBuscado.Tutor.Carnet, opcion)
				fmt.Println("Se asigno Correctamente....")
				break
			} else {
				fmt.Println("No hay tutores para ese curso....")
				break
			}
		} else {
			fmt.Println("El curso no existe en el sistema")
			break
		}

	}
}
