package tdas_tp

// ➤ Módulos

import (
	"fmt"
	"strings"
	"tdas/abb"
)

// ➤ Estructuras

type post struct {
	id    int
	autor Usuario
	texto string
	likes abb.DiccionarioOrdenado[string, string]
}

// ➤ Primitivas de un Post

func CrearPost(id int, autor Usuario, texto string) Post {
	post := new(post)
	post.id = id
	post.autor = autor
	post.texto = texto
	post.likes = abb.CrearABB[string, string](strings.Compare)
	return post
}

func (post *post) VerID() int {
	return post.id
}

func (post *post) VerAutor() Usuario {
	return post.autor
}

func (post *post) VerTexto() string {
	return post.texto
}

func (post *post) ObtenerLikes() int {
	return post.likes.Cantidad()
}

func (post *post) GuardarLike(nombre string) {
	post.likes.Guardar(nombre, "👍")
}

func (post *post) ImprimirPost() string {
	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", post.VerID(), post.autor.VerNombre(), post.VerTexto(), post.ObtenerLikes())
}

func (post *post) ImprimirLikes() {
	fmt.Println("El post tiene", post.ObtenerLikes(), "likes:")

	for iter := post.likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		fmt.Printf("\t%s\n", nombre)
	}
}
