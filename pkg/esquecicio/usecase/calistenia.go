package esquecicio

import (
	"errors"
	"math/rand"
	"time"

	"github.com/nenitf/esquecicio/pkg/esquecicio/contracts/repository"
	"github.com/nenitf/esquecicio/pkg/esquecicio/entity"
)

type TreinoCalisteniaDTO struct {
	Tipo string
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
	switch dto.Tipo {
	case "pull":
		exercicios = s.repo.FindPull()
		exercicios = append(exercicios, s.repo.FindAbs()...)
		requisitos = []string{"costas", "biceps", "abdomen"}
		break
	case "push":
		exercicios = s.repo.FindPush()
		exercicios = append(exercicios, s.repo.FindAbs()...)
		requisitos = []string{"ombro", "peito", "triceps", "abdomen"}
		break
	case "legs":
		exercicios = s.repo.FindLegs()
		exercicios = append(exercicios, s.repo.FindAbs()...)
		requisitos = []string{"perna", "abdomen"}
		break
	case "pulllegs":
		exercicios = s.repo.FindPull()
		exercicios = append(exercicios, s.repo.FindLegs()...)
		exercicios = append(exercicios, s.repo.FindAbs()...)
		requisitos = []string{"costas", "biceps", "perna", "abdomen"}
		break
	case "pushlegs":
		exercicios = s.repo.FindPush()
		exercicios = append(exercicios, s.repo.FindLegs()...)
		exercicios = append(exercicios, s.repo.FindAbs()...)
		requisitos = []string{"ombro", "peito", "triceps", "perna", "abdomen"}
		break
	default:
		return treino, errors.New("Escolha pull, push ou legs")
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
