package tdas_tp

type Usuario interface {

	// VerID devuelve un numero entero que representa el ID asignado cuando el usuario fue registrado
	VerID() int

	// VerNombre devuelve una cadena que representa el nombre del usuario
	VerNombre() string

	// GuardarPost recibe una publicacion y lo agrega al feed del usuario
	GuardarPost(*Post)

	// VerSiguienteFeed devuelve la proxima publicacion habil que haya en el feed del usuario, error si
	// el feed del usuario no tiene mas publicaciones para ver
	VerSiguienteFeed() (*Post, error)
}
