import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import "./App.css";
import Principal_Admin from "./pages/Administrador/Principal_Admin";
import Carga_Archivos from "./pages/Administrador/Carga_Archivos";
import Libros_Control from "./pages/Administrador/Libros_Control";
import Tabla_Alumnos from "./pages/Administrador/Tabla_Alumnos";
import Cargar_Libro from "./pages/Tutor/Cargar_Libro";
import Carga_Publicacion from "./pages/Tutor/Carga_Publicacion";
import Principal_Estudiante from "./pages/Estudiante/Principal_Estudiante";
import Reporte from "./pages/Administrador/Reporte";
import Cursos_Estudiante from "./pages/Estudiante/Cursos_Estudiante";
import Publicacion_Estudiante from "./pages/Estudiante/Publicacion_Estudiante";

function App() {
  const [count, setCount] = useState(0);

  return (
    
    <>
      <Routes>
        <Route path="/" element={<Login />} />

        <Route path="/principal/admin" element={<Principal_Admin />} />
        <Route path="/principal/admin/archivo" element={<Carga_Archivos />} />
        <Route path="/principal/admin/libros" element={<Libros_Control />} />
        <Route path="/principal/admin/alumnos" element={<Tabla_Alumnos />} />
        <Route path="/principal/admin/reporte" element={<Reporte />} />

        <Route path="/principal/tutor" element={<Cargar_Libro />} />
        <Route
          path="/principal/tutor/publicacion"
          element={<Carga_Publicacion />}
        />

        <Route
          path="/principal/estudiante"
          element={<Principal_Estudiante />}
        />
        <Route
          path="/principal/estudiante/libro"
          element={<Cursos_Estudiante />}
        />
        <Route
          path="/principal/estudiante/publicacion"
          element={<Publicacion_Estudiante />}
        />
      </Routes>
    </>
  );
}

export default App;