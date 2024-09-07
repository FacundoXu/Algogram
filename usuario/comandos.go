package usuario

// ➤ Módulos

import (
	"algogram/errores"
	"algogram/tdas_tp"
	"fmt"
	"tdas/hash_cerrado"
)

// ➤ Constantes

const (
	// Comandos
	LOGIN              = "login"
	LOGOUT             = "logout"
	PUBLICAR           = "publicar"
	VER_SIGUIENTE_FEED = "ver_siguiente_feed"
	LIKEAR_POST        = "likear_post"
	MOSTRAR_LIKES      = "mostrar_likes"

	// Mensajes
	HOLA           = "Hola"
	ADIOS          = "Adios"
	POST_PUBLICADO = "Post publicado"
	POST_LIKEADO   = "Post likeado"

	// Auxiliares
	NADIE_LOGGEADO = "💤"
)

func Comandos(cmd []string, postID *int, loggeado *string, publicaciones hash_cerrado.Diccionario[int, tdas_tp.Post], usuarios hash_cerrado.Diccionario[string, tdas_tp.Usuario]) {

	// Condiciones de programa dependiendo de la entrada ingresada
	switch cmd[0] {

	case LOGIN:
		Login(cmd, loggeado, usuarios)

	case LOGOUT:
		Logout(loggeado)

	case PUBLICAR:
		Publicar(cmd, postID, loggeado, publicaciones, usuarios)

	case VER_SIGUIENTE_FEED:
		VerSiguienteFeed(loggeado, usuarios)

	case LIKEAR_POST:
		LikearPost(cmd, loggeado, publicaciones)

	case MOSTRAR_LIKES:
		MostrarLikes(cmd, publicaciones)

	default:
		fmt.Println(errores.ErrorComando{}.Error()) // Error cuando ingresa un comando invalido
	}
}
