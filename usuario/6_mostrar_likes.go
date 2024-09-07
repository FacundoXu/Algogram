package usuario

// ➤ Módulos

import (
	"algogram/tdas_tp"
	"algogram/errores"
	"fmt"
	"strconv"
	"tdas/hash_cerrado"
)

func MostrarLikes(cmd []string, publicaciones hash_cerrado.Diccionario[int, tdas_tp.Post]) {

	// Convertimos el ID ingresado a un numero
	id, err := strconv.Atoi(cmd[1])

	// Error si el ID ingresado tiene mal formato
	if err != nil {
		fmt.Println(errores.ErrorFormato{}.Error())
		return
	}

	// Error si no existe tal publicacion con el respectivo ID o la publicacion no tiene likes
	if !publicaciones.Pertenece(id) || publicaciones.Obtener(id).ObtenerLikes() == 0 {
		fmt.Println(errores.ErrorPostSinLikes{}.Error())
		return
	}

	// Obtenemos los likes de la publicacion y los usuarios que le dieron like
	publicaciones.Obtener(id).ImprimirLikes()
}
