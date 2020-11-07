package esquecicio

import (
	"errors"
	"math/rand"
	"time"

	"github.com/nenitf/esquecicio/pkg/esquecicio/contracts/repository"
	"github.com/nenitf/esquecicio/pkg/esquecicio/entity"
)

type TreinoCalisteniaDTO struct {
	Tipos []string
}

type TreinoCalistenia struct {
	repo repository.Exercicios
}

func NewTreinoCalistenia(r repository.Exercicios) *TreinoCalistenia {
	return &TreinoCalistenia{
		repo: r,
	}
}

func (s *TreinoCalistenia) Execute(dto TreinoCalisteniaDTO) (entity.Treino, error) {
	var treino entity.Treino
	var exercicios []entity.Exercicio
	var requisitos []string
	for _, tipo := range dto.Tipos {
		switch tipo {
		case "pull":
			exercicios = append(exercicios, s.repo.FindPull()...)
			requisitos = append(requisitos, "costas", "biceps")
			break
		case "push":
			exercicios = append(exercicios, s.repo.FindPush()...)
			requisitos = append(requisitos, "ombro", "peito", "triceps")
			break
		case "legs":
			exercicios = append(exercicios, s.repo.FindAbs()...)
			requisitos = append(requisitos, "abdomen")
			break
		case "abs":
			exercicios = append(exercicios, s.repo.FindAbs()...)
			requisitos = append(requisitos, "abdomen")
			break
		default:
			return treino, errors.New("Escolha pull, push, legs e/ou abs")
		}
	}

	if len(dto.Tipos) <= 0 {
		return treino, errors.New("Escolha pull, push, legs e/ou abs")
	}

	if len(exercicios) <= 0 {
		return treino, errors.New("Não foram encontrados exercicios")
	}

	for _, requisito := range requisitos {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(exercicios), func(i, j int) {
			exercicios[i], exercicios[j] = exercicios[j], exercicios[i]
		})
		for _, exercicio := range exercicios {
			for _, grupo := range exercicio.Grupo {
				if requisito == grupo {
					treino.Exercicios = append(treino.Exercicios, exercicio)
					goto Encontrou
				}
			}
		}
		return treino, errors.New("Não foi encontrado exercicio")
	Encontrou:
	}
	return treino, nil
}
