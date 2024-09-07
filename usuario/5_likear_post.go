package usuario

// ➤ Módulos

import (
	"algogram/tdas_tp"
	"algogram/errores"
	"fmt"
	"strconv"
	"tdas/hash_cerrado"
)

func LikearPost(cmd []string, loggeado *string, publicaciones hash_cerrado.Diccionario[int, tdas_tp.Post]) {

	// Convertimos el ID ingresado a un numero
	id, err := strconv.Atoi(cmd[1])

	// Error si el ID ingresado tiene mal formato
	if err != nil {
		fmt.Println(errores.ErrorFormato{}.Error())
		return
	}

	// Error si no habia nadie loggeado o si no existe tal publicacion con el respectivo ID
	if *loggeado == NADIE_LOGGEADO || !publicaciones.Pertenece(id) {
		fmt.Println(errores.ErrorPostInexistente{}.Error())
		return
	}

	// El usuario loggeado likea la publicacion del correspondiente ID
	publicaciones.Obtener(id).GuardarLike(*loggeado)
	fmt.Println(POST_LIKEADO)
}
