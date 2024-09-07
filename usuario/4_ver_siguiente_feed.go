package usuario

// ➤ Módulos

import (
	"algogram/tdas_tp"
	"algogram/errores"
	"fmt"
	"tdas/hash_cerrado"
)

func VerSiguienteFeed(loggeado *string, usuarios hash_cerrado.Diccionario[string, tdas_tp.Usuario]) {

	// Error si no habia nadie loggeado
	if *loggeado == NADIE_LOGGEADO {
		fmt.Println(errores.ErrorFeedVacio{}.Error())
		return
	}

	// Obtenemos el usuario correspondiente
	usuario := usuarios.Obtener(*loggeado)

	// Obtenemos la siguiente publicacion del feed del usuario
	post, err := usuario.VerSiguienteFeed()

	// Error si no hay mas publicaciones para ver
	if err != nil {
		fmt.Println(errores.ErrorFeedVacio{}.Error())
		return
	}

	// Obtenemos los resultados de la publicacion
	fmt.Println((*post).ImprimirPost())
}
