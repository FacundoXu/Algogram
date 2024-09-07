package tdas_tp

// ➤ Módulos

import (
	"algogram/errores"
	"tdas/heap"
)

// ➤ Estructuras

type usuario struct {
	id     int
	nombre string
	feed   heap.ColaPrioridad[*Post]
}

// ➤ Funciones auxiliares

func distancia(a usuario, b Usuario) int {
	distancia := a.VerID() - b.VerID()

	if distancia > 0 {
		return distancia
	}
	return distancia * (-1)
}

// ➤ Funciones de comparación auxiliares

func afinidadPosts(a Post, b Post) int {
	if a.VerID()-b.VerID() > 0 {
		return -1
	}
	if a.VerID()-b.VerID() < 0 {
		return 1
	}
	return 0
}

func (usuario usuario) afinidadUsuarios(a *Post, b *Post) int {
	distancia1 := distancia(usuario, (*a).VerAutor())
	distancia2 := distancia(usuario, (*b).VerAutor())

	if (distancia1 - distancia2) > 0 {
		return -1
	}
	if (distancia1 - distancia2) < 0 {
		return 1
	}
	return afinidadPosts((*a), (*b))
}

// ➤ Primitivas de un Usuario

func CrearUsuario(id int, nombre string) Usuario {
	usuario := new(usuario)
	usuario.id = id
	usuario.nombre = nombre
	usuario.feed = heap.CrearHeap[*Post](usuario.afinidadUsuarios)
	return usuario
}

func (usuario *usuario) VerID() int {
	return usuario.id
}

func (usuario *usuario) VerNombre() string {
	return usuario.nombre
}

func (usuario *usuario) GuardarPost(post *Post) {
	usuario.feed.Encolar(post)
}

func (usuario *usuario) VerSiguienteFeed() (*Post, error) {
	if usuario.feed.EstaVacia() {
		return nil, errores.ErrorFeedVacio{}
	}
	return usuario.feed.Desencolar(), nil
}
