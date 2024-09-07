package auxiliares

import (
	"algogram/errores"
	"algogram/tdas_tp"
	"bufio"
	"fmt"
	"os"
	"tdas/hash_cerrado"
)

// CargarUsuarios() recibe por parametro la ruta que contiene los usuarios registrados.
// Devuelve un diccionario con sus respectivos usuarios.
func CargarUsuarios(ruta string) hash_cerrado.Diccionario[string, tdas_tp.Usuario] {
	usuarios := hash_cerrado.CrearHash[string, tdas_tp.Usuario]()

	// Abrimos el archivo
	archivo, err := os.Open(ruta)

	// Devuelve un error en caso de que no exista el archivo
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		os.Exit(0)
	}

	// Cierra el archivo al final de la ejecucion de la funcion
	defer archivo.Close()

	// Leemos cada linea del archivo y guardamos los usuarios
	linea := bufio.NewScanner(archivo)
	id := 0

	for linea.Scan() {
		// Por cada usuario del archivo creamos un usuario y lo guardamos en el diccionario
		nombre := linea.Text()
		usuario := tdas_tp.CrearUsuario(id, nombre)
		usuarios.Guardar(nombre, usuario)
		id++
	}
	return usuarios
}

// CrearPublicaciones() devuelve un diccionario de publicaciones habil para guardar posts
func CrearPublicaciones() hash_cerrado.Diccionario[int, tdas_tp.Post] {
	return hash_cerrado.CrearHash[int, tdas_tp.Post]()
}
