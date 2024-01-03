# Manual Técnico

<p>
Para la realización de este proyecto se llevaron a cabo el uso de diferentes Struct, a las cuales fueron asociadas diferentes métodos.
<p>

### Matriz struct
<p>
Esta struct fue utilizada para definir la matriz dispersa, la cual fue utilizada para guaradar las asignaciones de los estudiantes a los diferentes cursos. A este fue asociado diferentes métodos los cuales se detallan a continuación.

- Insertar_Elemento: Este método recibe como parámetros el carnet del tutor, el carnet del estudiante y el codigo del curso. Fue utilizada para crear nuevos nodos los cuales relacionan al estudiante con el tutor que imparte el curso. Debido a que los tutores se almacenan en las cabeceras y los alumnos asignados en las filas, durante la ejecución del método se van haciendo diferentes validaciones para saber si el alumno asinado ya se encuentra en las filas o si el tutor del curso ya se encuentra en la cabecera y dependiendo del caso que ocurra se iran crando nuevos nodos en la cabecera o en las filas de ser necesario para realizar las conecciones del nuevo nodo.

- nuevaFila: Este método recibe como parámetro la posición y, el carnet del estudiante y el curso al cual se desea asignar. Se utiliza este método para crear un nuevo nodo en las filas el cual representara al estudiante asignado y devuelve dicho nodo creado en las filas.

- insertarFila: Este método recibe como parámetro nuevoNodo *NodoMatriz que es el nodo que se desea conectar y nodoRaiz *NodoMatriz que puede ser la raiz de la matriz donde inicia todo o puede ser un nodo de la cabecera. Este método se utiliza para hacer las conecciones del nuevo nodos con su respectivo nodo de la cabecera.

- nuevaColumna: Este método recibe como parámetro la posición x, el carnet del tutor y el curso el cual imparte. Se utiliza este método para crear un nuevo nodo en la cabecera el cual representara al tutor del curso y devuelve dicho nodo creado en la cabecera.

- insertarColumna: Este método recibe como parámetro nuevoNodo *NodoMatriz que es el nodo que se desea conectar y nodoRaiz *NodoMatriz que puede ser la raiz de la matriz donde inicia todo o puede ser un nodo de las filas. Este método se utiliza para hacer las conecciones del nodo con su respectivo nodo de la fila.

- buscarFila: Este método recib como parámetro el carnet del estudiante y usando dicho carnet revisa en las filas si existe un nodo con dicho carnet. Si existe devulve ese nodo de la fila sino devulve nil.

- buscarColumna: Este métdo recibe como parámetro el carnet del tutor y el curso, usando estos parametros realiza una busqueda en los nodos de la cabecera para ver si algun nodo tiene valores iguales y sí existe devuelve un nodo de la cabecera, de lo contrario devuleve nil. 

- Reporte: Este método se utilizo para realizar un reporte en graphviz, donde se pueda visualizar en los nodos de la cabecera el carnet de cada tutor de determinado curso y en los nodos de la fila el carnet de los alumnos que se asignan a determinado curso. Estos nodos de cabercera y los de fila estaran conectados a un nodo que muestra el código del curso al cual el estudiante se asigna y el cual es impartido por el tutor correspondiente.
<p>
<p align="center">
  <img src="https://i.postimg.cc/8cwRpFD3/Captura-de-pantalla-2023-12-19-110008.png" alt="Descripción de la imagen">
</p>

### NodoMatriz struct

<p>
Este struct se utilizo para definir los atributos que tienen los nodos de la matriz, siendo estos los siguientes.

- Siguiente *NodoMatriz 
- Anterior  *NodoMatriz
- Abajo     *NodoMatriz
- Arriba    *NodoMatriz
- PosX      int
- PosY      int
- Dato      *Dato
<p>

### Dato struct

<p>
Este struct se utilizo para definir atributos que hacen referencia a los datos que guarda cada nodo de la matriz dispersa. Siendo estos datos el carnet del estudiante que se va a asignar, el carnet del tutor que impartira el curso y el código del curso.
<p>

## Lista Doble

<p>
Este struct se utilizo para definir la lista doble que contendra los estudiantes cargados al sistema. A este struct se asociaron diferentes métodos los cuales se detallan a continuación:

- Agregar: Este método recibe como parámetro el nombre y el carnet del estudiante, usando estos datos se crea un nuevo nodo el cual se ingresa en la lista doble.
- Buscar: Este método recibe como parámetro el carnet del estudiante y su contraseña siendo esta el mismo que su carnet. Utilizando estos datos realiza una busqueda dentro de la lista doble para encontrar algun nodo que contenga un carnet igual al que se envia en el parámetro.
- LeerCSV: Este método se utiliza para leer un archivo csv el cual contiene los datos de carnet y nombre de los estudiantes que se desean cargar al sistema y aquí mismo se llama al método Agregar para ingresar todos los datos en la lista doble.
- Reporte: Este método se utiliza para crear un reporte usando graphviz en donde se muestre la lista doble y cada uno de sus nodos muestren el nombre y carnet de cada estudiante guardado en el sistema.
<p>
<p align="center">
  <img src="https://i.postimg.cc/13wTXLr0/Captura-de-pantalla-2023-12-19-130035.png" alt="Descripción de la imagen">
</p>

### NodoListaDoble struct

<p>
Este struct se utilizo para definir los atributos de los nodos de la lista doble, siendo estos los siguientes:

- Alumno    *Alumno
- Siguiente *NodoListaDoble
- Anterior  *NodoListaDoble
<p>

### Alumno struct

<p>
Este struct se utilizo para definir atributos que hacen referencia a los datos que guarda cada nodo de la lista doble. Siendo estos datos el carnet y el nombre del estudiante.
<p>

## Cola struct

<p>
Este struct se utilizo para definir una cola de prioridad la cual contendra a los posibles tutores para cada curso. A este struct se asociaron diferents métodos los cuales se detallan a continaución:

- Encolar: Este método recibe como parámetro el nombre del tutor, el carnet del tutor, el curso que va a impartir y la nota. Utilizando estos datos se crea un nodo al cual se le asigna una prioridad dependiendo de la nota que tenga el tutor y dependiendo de la prioridad que tenga cada nodo se ira agregando a la cola de prioridad.
- Desencolar: Este método se utilizo para eliminar al primer nodo en la cola de prioridad.
- LeerCSV: Este método se utilizo para leer el achivo CSV donde vienen los datos de los tutores y usando estos datos se crearon nuevos nodos los cuales fueron agregados a la cola de prioridad por medio de la función Encolar.
- Primero_Cola: Este método fue utilizado para mostrar los datos que guarda el primer nodo en la cola de prioridad, además de mostrar el carnet que se encuentra en el siguiente nodo.
<p>
<p align="center">
  <img src="https://i.postimg.cc/zfJ83dxx/Captura-de-pantalla-2023-12-20-162038.png" alt="Descripción de la imagen">
</p>

### NodoCola struct
<p>
Este struct se utilizo para definir los atributos que tienen los nodos de la cola de prioridad, siendo estos los siguientes:

- Tutor     *Tutores
- Siguiente *NodoCola
- Prioridad int
<p>

### Tutores struct
<p>
Este struct se utilizo para definir atributos que hacen referencia a los datos que guarda cada nodo de la cola de prioridad. Siendo estos datos: el carnet del tutor, el nombre del tutor, el curso que busca impartir y la nota.
<p>

## ListaDobleCircular struct
<p>
Este struct se utilizo para definir la lista circula doblemente enlazada que contendra a los tutores que fueron aceptados por el administrador para impartir determinado curso. A esta struct fueron asociados varios métodos, los cuales se detallan a continuación.

- Agregar: Este método recibe como parámetros los datos del tutor que se va a aceptar, los cuales son:  Carnet, nombre, curso y nota. Usando estos datos se crea un nuevo nodo el cual se introducira en la lista circular doblemente enlazada.
- Reporte: Este método se utiliza para generar un reporte en graphviz que muestra de forma gráfica la lista circular doblemente enlzada y en cada uno de sus nodos se muestra el nombre y el carnet del tutor aceptado.
- Mostrar: Este método se utiliza para recorrer la lista circular doblemente enlazada e ir mostrando el curso y el nombre del tutor que lo imparte.
- Buscar: Este método recibe como parametro el código de algún curso y se utiliza para recorrer la lista circular doblemente enlazada y verificar si hay algun nodo que contenga el curso que se envió por parámetro y de esta manera retornar un valor booleano el cual determinara si el curso tiene un tutor asignado en el sistema.
- BuscarTutor: Este método recibe como parámetro el código de algún curso y se utiliza para recorrer la lista circular doblemente enlazada y verificar si algún nodo contiene dicho curso que se envió como parámetro. De encontrar algún nodo con este curso se devuleve dicho nodo encontrado.
- CambiarTutor: Este método recibe como parámetro el carnet, nombre, curso y nota del nuevo tutor. Usando el valor de curso obtenido por parámetro se relizada una busqueda en la lista circular doblemente enlazada y al encontra el nodo deseado se sustituyen los valores de nombre, carnet y nota del nodo por los valores obtenidos por parámetro.
<p>

<p align="center">
  <img src="https://i.postimg.cc/gJDvfysw/Captura-de-pantalla-2023-12-20-170340.png" alt="Descripción de la imagen">
</p>

### NodoListaCircular struct
<p>
Este struct se utilizo para definir los atributos que tienen los nodos de la cola de prioridad, siendo estos los siguientes:

- Tutor     *Tutores
- Siguiente *NodoListaCircular
- Anterior  *NodoListaCircular
<p>

## ArbolAVL

<p>
  Este struct se utilizo para definir el arbol AVL, a este struct fueron asociados diferentes métodos los cuales se detallan a continuación:

- InsertarElemento: Este método recibe como parámetro el nombre y el codigo del curso, usando estos valores se crea un nuevo nodo y llama a el método de InsertarNodo el cual recibe como parámetro el nodo creado y el nodo raiz del arbol. Por medio de esto se agrega el nuevo nodo a el arbol AVL.
- InsertarNodo: Este método recibe como parámetro el nodo creado y el nodo raiz del arbol, es utilizado para agregar el nuevo nodo a el arbol y luego de haberlo agregado se realiza una verificación si esta balanceado el arbol, si este no es el caso se lleva a cabo procedimientos para balancear el arbol.
- RotaciónI: Este método recibe el nodo que necesita ser balanceado y durante la ejecución de este método se realiza una rotación hacia la izquierda.
- RotaciónD: Este método recibe el nodo que necesita ser balanceado y durante la ejecución de este método se realiza una rotación hacia la derecha.
- altura: Este método recibe un nodo del arbol AVL y devuelve el valor de la altura que tiene.
- equilibrio: Este método recibe un nodo del arbol AVL y devuelve el valor del factor de equilibrio que tiene.
- Busqueda: Este métdo recibe como parámetro el código del curso y realiza una busqueda en el arbol AVL para determinar si algun nodo tiene el valor del código del curso que se recibio por parámetro, para esto devuleve un valor booleano para determinar si el curso se encuentra o no en el sistema.
- busqueda_arbol: Este métdo recibe como parámetro el código del curso y la raíz del arbol para realizar una busqueda en el arbol AVL para determinar si algun nodo tiene el valor del código del curso que se recibio por parámetro, sí lo encuentra devuelve dicho nodo.
- LeerJson: Este método se utiliza para leer el archivo Json que contiene los datos del nombre y código de los cursos, y luego llama al método InsertarElemento para ingresar esos datos al árbol AVL.
- Graficar: Este método se utiliza para realizar un reporte en graphviz donde se muestra en forma gráfica el arbol AVL, mostrando en sus nodos los códigos de los cursos ingresados en el sistema.
<p>

<p align="center">
  <img src="https://i.postimg.cc/439d7pSY/Captura-de-pantalla-2023-12-20-174847.png" alt="Descripción de la imagen">
</p>

### nodoArbol struct
Este struct se utilizo para definir los atributos que tienen los nodos del arbol AVL, siendo estos los siguientes:

-	Izquierdo         *NodoArbol
-	Derecho           *NodoArbol
-	Valor             string
-	Nombre            string
-	Altura            int
-	Factor_Equilibrio int

