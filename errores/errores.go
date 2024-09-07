package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "Error: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Faltan parámetros"
}

type ErrorFormato struct{}

func (e ErrorFormato) Error() string {
	return "Error: Formato de ID invalido"
}

type ErrorComando struct{}

func (e ErrorComando) Error() string {
	return "Error: Comando incorrecto"
}

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExistente struct{}

func (e ErrorUsuarioNoExistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorFeedVacio struct{}

func (e ErrorFeedVacio) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostSinLikes struct{}

func (e ErrorPostSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
