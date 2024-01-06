# Manual Técnico

<p>
Para la realización de este proyecto se llevaron a cabo el uso de diferentes Struct para el backend.
<p>

## TablaHash struct
<p>
Esta struct fue utilizada para definir la tabla hash, la cual fue utilizada para guaradar a los estudiantes. A este fue asociado diferentes métodos los cuales se detallan a continuación.

- calculoIndice: Este método recibe como parámetro el carnet del estudiante para luego dividirlo en digito por digito para guardar estos digitos en un arreglo, luego se toma cada digito guardado para obtener su valor en ASCII, al obtner este valor se ingresa en un nuevo arreglo para luego sumar cada uno de estos valores. La suma obtenida se divide dentro de la capacidad actual de la tabla hash para obtener el indice.

- capacidadTabla: Este método se usa para saber cuantos elementos hay en la tabla hash y ver que aun no se exceda la cantidad de elementos al el porcentaje de utilización de la tabla, pero de ser el caso que se haya excedido se empieza a hacer el procedimiento necesario para aumentar la capacidad de la tabla hash.

- nuevaCapacidad: Este método se utiliza para obtener la nueva capacidad para la tabla basandose en la serie de Fibonacci.

- reInsertar: Este método se utiliza para volver a ingresar todos los datos en la nueva tabla hash.

- reCalculoIndice: Este método se utiliza para calcular de nuevo el indice, debido a que se utiliza para el ingreso de datos a la tabla hash el direccionamiento abierto por ende se podrán generar colisiones las cuales son manejadas por medio de este método, para el manejo de estas colisiones se utiliza el salto al cuadrado.

- nuevoIndice: Este método se utiliza para verificar que el nuevo indice que se ha obtenido no exceda la capacidad de la tabla, sí se da este caso. El método realiza los procedimientos necesarios para obtener un índice el cual no exceda la capacidad.

- Insertar: Este método recibe como parámetros los valores de carnet, nombre, password y cursos que corresponden a un estudiante para luego realizar los procedimientos necesarios para ingresar los datos en la tabla hash.

- Buscar: Este método recibe como parámetro los valroes de carnet y password, usando estos valores se busca dentro de la tabla hash para ver si existen estos datos y devuelve un valor booleano para verificar si los datos si se encuentran en la tabla hash o no.


<p>
<p align="center">
  <img src="https://i.postimg.cc/tgVgdn5P/Captura-de-pantalla-2024-01-05-203141.png" alt="Descripción de la imagen">
</p>

### NodoMatriz struct

<p>
Este struct se utilizo para definir los atributos que tienen los nodos de la tabla hash, siendo estos los siguientes.

- LLave int 
- Persona *Persona
<p>

### Persona struct

<p>
Este struct se utilizo para definir atributos que hacen referencia a los datos de los estudiantes que guarda cada nodo de la tabla hash . Siendo estos datos el carnet, nombre, password y un arreglo que contiene los cursos que esta asignado el estudiante.
<p>

## Grafo

<p>
Este struct se utilizo para definir un grafo que contendra los cursos y sus post requisitos. A este struct se asociaron diferentes métodos los cuales se detallan a continuación:

- InsertarValores: Este método se utilizo para ingresar los diferentes cursos y sus post requisitos al grafo. Para relizacionar los nodos del grafo se utilizo una matriz disperasa por lo cual dentro del flujo de este método se llaman a los métodos insertarFila e insertarColumna.

- insertarFila: Este método recibe como parámetro el curso, usando este valor se crea un nodo que represente dicho curso. Se realiza un recorrido hacia abajo para verificar si dicho curso ya se encuentra entre las filas, si ya se encuentra solo se detiene la ejecución por medio de un return pero si no existe dicho curso dentro de las filas entonces crea uno nuevo que represente este.

- insertarColumna: Este método recibe como parámetro el curso y su postrequisito, con estos datos realiza primero una busqueda en las filas para encontrar el curso y luego se utiliza el apuntador siguiente del nodo del curso encontrado para llegar hasta el final y agregar ahí el nodo correspondiente a el curso post requisito.

- Reporte: Este método es utilizado para crear un reporte en graphivz el cual mostrará de manera grafica la realizacion de los cursos con sus cursos post requisito.

- retornarValoresMatriz: Este método se utiliza para recorrer el grafo y de este modo se forme una cadena de texto la cual mostrar la relacion de los nodos para poder crear el reporte en graphviz.
<p>
<p align="center">
  <img src="https://i.postimg.cc/nhyq8D3V/Captura-de-pantalla-2024-01-05-205325.png" alt="Descripción de la imagen">
</p>

### NodoListaAdyacencia struct

<p>
Este struct se utilizo para definir los atributos de los nodos del grafo, siendo estos los siguientes:

- Siguiente *NodoListaAdyacencia
-	Abajo     *NodoListaAdyacencia
-	Valor     string
<p>

