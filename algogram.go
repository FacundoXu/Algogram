package main

// ➤ Módulos

import (
	"algogram/auxiliares"
	"algogram/errores"
	"algogram/usuario"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ➤ Constantes

const (
	MAX_PARAMETROS_REQUERIDOS = 2
	USUARIOS                  = 1
	NADIE_LOGGEADO            = "💤"
)

func main() {

	// Variables a utilizar
	args := os.Args

	// Error si no hay suficientes parametros ingresados
	if len(args) < MAX_PARAMETROS_REQUERIDOS {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}

	// Cargamos la informacion de los usuarios registrados
	usuarios := auxiliares.CargarUsuarios(args[USUARIOS])

	// Error si no existen los archivos
	if usuarios == nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
	}

	// Publicaciones de los usuarios, usuario loggeado actual y ID del respectivo post publicado
	publicaciones := auxiliares.CrearPublicaciones()
	loggeado := NADIE_LOGGEADO
	postID := 0

	// Pedimos entradas al usuario
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		// Entrada al usuario
		entrada := s.Text()

		// Separamos la entrada
		cmd := strings.Split(entrada, " ")

		// Acciones dependiendo del comando ingresado por el usuario
		usuario.Comandos(cmd, &postID, &loggeado, publicaciones, usuarios)
	}
}
