package usuario

// ➤ Módulos

import (
	"algogram/errores"
	"algogram/tdas_tp"
	"fmt"
	"strings"
	"tdas/hash_cerrado"
)

func Publicar(cmd []string, postID *int, loggeado *string, publicaciones hash_cerrado.Diccionario[int, tdas_tp.Post], usuarios hash_cerrado.Diccionario[string, tdas_tp.Usuario]) {

	// Error si no habia nadie loggeado
	if *loggeado == NADIE_LOGGEADO {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{}.Error())
		return
	}

	// Convertimos el texto ingresado a una cadena
	texto := strings.Join(cmd[1:], " ")

	// Creamos un post con su respectivo ID, nombre del usuario que realizo la publicacion y el texto
	post := tdas_tp.CrearPost(*postID, usuarios.Obtener(*loggeado), texto)

	// Una vez creado lo guardamos en las publicaciones
	publicaciones.Guardar(*postID, post)

	// Para cada usuario registrado guardamos en sus feed la publicacion realizada
	for iter := usuarios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {

		nombre, usuario := iter.VerActual()

		if nombre != *loggeado {
			usuario.GuardarPost(&post)
		}
	}

	// Aumentamos el ID para el proximo post
	*postID++
	fmt.Println(POST_PUBLICADO)
}
