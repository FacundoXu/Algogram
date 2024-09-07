package usuario

// ➤ Módulos

import (
	"algogram/errores"
	"algogram/tdas_tp"
	"fmt"
	"strings"
	"tdas/hash_cerrado"
)

func Login(cmd []string, loggeado *string, usuarios hash_cerrado.Diccionario[string, tdas_tp.Usuario]) {

	// Convertimos el nombre ingresado a una cadena
	nombre := strings.Join(cmd[1:], " ")

	// Error si ya habia alguien loggeado
	if *loggeado != NADIE_LOGGEADO {
		fmt.Println(errores.ErrorUsuarioYaLoggeado{}.Error())
		return
	}

	// Error si no existe tal usuario registrado
	if !usuarios.Pertenece(nombre) {
		fmt.Println(errores.ErrorUsuarioNoExistente{}.Error())
		return
	}

	// El usuario es correcto, ahora pasa a ser el usuario loggeado actual
	*loggeado = nombre
	fmt.Println(HOLA, nombre)
}
