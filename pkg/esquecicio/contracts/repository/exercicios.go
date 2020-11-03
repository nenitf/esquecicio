package repository

import "github.com/nenitf/esquecicio/pkg/esquecicio/entity"

type Exercicios interface {
	FindPull() []entity.Exercicio
	FindPush() []entity.Exercicio
	FindLegs() []entity.Exercicio
	FindAbs() []entity.Exercicio
}
