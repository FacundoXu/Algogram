package usuario

// ➤ Módulos

import (
	"algogram/errores"
	"fmt"
)

func Logout(loggeado *string) {

	// Error si no habia nadie loggeado
	if *loggeado == NADIE_LOGGEADO {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{}.Error())
		return
	}

	// Hay un usuario loggeado y lo desloggeamos
	*loggeado = NADIE_LOGGEADO
	fmt.Println(ADIOS)
}
