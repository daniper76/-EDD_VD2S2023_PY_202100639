package Cola

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: 0}

	if nota >= 90 && nota <= 100 {
		nuevoNodo.Prioridad = 1
	} else if nota >= 75 && nota <= 89 {
		nuevoNodo.Prioridad = 2
	} else if nota >= 65 && nota <= 74 {
		nuevoNodo.Prioridad = 3
	} else if nota >= 61 && nota <= 64 {
		nuevoNodo.Prioridad = 4
	} else {
		return
	}

	if c.Longitud == 0 {
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		aux := c.Primero
		for aux.Siguiente != nil {
			if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad == nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else if aux.Siguiente.Prioridad > nuevoNodo.Prioridad && aux.Prioridad < nuevoNodo.Prioridad {
				nuevoNodo.Siguiente = aux.Siguiente
				aux.Siguiente = nuevoNodo
				c.Longitud++
				return
			} else {
				aux = aux.Siguiente
			}
		}
		if aux.Prioridad > nuevoNodo.Prioridad {
			nuevoNodo.Siguiente = aux
			c.Primero = nuevoNodo
			c.Longitud++
		} else if aux.Prioridad <= nuevoNodo.Prioridad {
			aux.Siguiente = nuevoNodo
			c.Longitud++
		}
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay tutores en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}

func (c *Cola) LeerCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}

		carnet, _ := strconv.Atoi(linea[0])
		nota, _ := strconv.Atoi(linea[3])
		c.Encolar(carnet, linea[1], "0"+linea[2], nota)
	}
}

func (c *Cola) Primero_Cola() {
	if c.Longitud == 0 {
		fmt.Println("No hay mas Tutores")
	} else {
		if c.Primero != nil {
			fmt.Println("Actual: ", c.Primero.Tutor.Carnet)
			fmt.Println("Nombre: ", c.Primero.Tutor.Nombre)
			fmt.Println("Curso: ", c.Primero.Tutor.Curso)
			fmt.Println("Nota: ", c.Primero.Tutor.Nota)
			fmt.Println("Prioridad: ", c.Primero.Prioridad)
			if c.Primero.Siguiente != nil {
				fmt.Println("Siguiente: ", c.Primero.Siguiente.Tutor.Carnet)
			} else {
				fmt.Print("Siguiente: No hay mas tutores por evaluar")
			}
		}
	}
}
