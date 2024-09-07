package tdas_tp

type Post interface {

	// VerID devuelve un numero entero que representa el ID de la publicacion
	VerID() int

	// VerAutor devuelve el usuario que realizó dicha publicacion
	VerAutor() Usuario

	// VerTexto devuelve una cadena que representa el mensaje que contiene dicha publicacion
	VerTexto() string

	// ObtenerLikes devuelve un numero entero que representa la cantidad de likes que tiene dicha publicacion
	ObtenerLikes() int

	// ImprimirPost devuelve una cadena con el siguiente formato:
	// Post ID <id de la publicacion>
	// <nombre del usuario que realizo la publicacion> dijo: <texto de la publicacion>
	// Likes: <cantidad de likes que tiene la publicacion>
	ImprimirPost() string

	// GuardarLike recibe una cadena que representa el nombre del usuario que dio like a dicha publicacion
	// y lo almacena como uno de los usuarios que dio like a esta publicacion
	GuardarLike(string)

	// ImprimirLikes muestra la cantidad de likes que tiene la publicacion y los nombres de los usuarios
	// que dieron like a dicha publicacion en orden alfabetico con el siguiente formato:
	// El post tiene <cantidad de likes> likes:
	//     <nombre de usuario que dio like>
	//     <nombre de usuario que dio like>
	//     <nombre de usuario que dio like>
	//     ...
	ImprimirLikes()
}
