package mock

import (
	"github.com/nenitf/esquecicio/pkg/esquecicio/contracts/repository"
	"github.com/nenitf/esquecicio/pkg/esquecicio/entity"
)

type ExerciciosRepo struct {
	MFindPull *[]entity.Exercicio
	MFindPush *[]entity.Exercicio
	MFindLegs *[]entity.Exercicio
	MFindAbs  *[]entity.Exercicio
}

func (r ExerciciosRepo) genericFind(prop *[]entity.Exercicio) []entity.Exercicio {
	if prop == nil {
		panic("Sem mock")
	}

	return *prop
}

func (r ExerciciosRepo) FindPull() []entity.Exercicio {
	return r.genericFind(r.MFindPull)
}

func (r ExerciciosRepo) FindPush() []entity.Exercicio {
	return r.genericFind(r.MFindPush)
}

func (r ExerciciosRepo) FindLegs() []entity.Exercicio {
	return r.genericFind(r.MFindLegs)
}

func (r ExerciciosRepo) FindAbs() []entity.Exercicio {
	return r.genericFind(r.MFindAbs)
}

func ExerciciosRepoPadrao() repository.Exercicios {
	pull := []entity.Exercicio{
		entity.Exercicio{
			Nome:        "biceps curl",
			Dificuldade: 2,
			Grupo:       []string{"biceps"},
		},
		entity.Exercicio{
			Nome:        "pull up arqueiro",
			Dificuldade: 3,
			Grupo:       []string{"biceps"},
		},
		entity.Exercicio{
			Nome:        "pull up australiano",
			Dificuldade: 3,
			Grupo:       []string{"costas"},
		},
		entity.Exercicio{
			Nome:        "chin up",
			Dificuldade: 3,
			Grupo:       []string{"costas", "biceps"},
		},
		entity.Exercicio{
			Nome:        "pull up",
			Dificuldade: 3,
			Grupo:       []string{"costas"},
		},
	}
	push := []entity.Exercicio{
		entity.Exercicio{
			Nome:        "dips",
			Dificuldade: 2,
			Grupo:       []string{"triceps"},
		},
		entity.Exercicio{
			Nome:        "push up",
			Dificuldade: 2,
			Grupo:       []string{"peito"},
		},
		entity.Exercicio{
			Nome:        "push up inclinado",
			Dificuldade: 2,
			Grupo:       []string{"ombro"},
		},
		entity.Exercicio{
			Nome:        "rotação ombro argolas",
			Dificuldade: 2,
			Grupo:       []string{"ombro"},
		},
		entity.Exercicio{
			Nome:        "extensão triceps testa",
			Dificuldade: 2,
			Grupo:       []string{"triceps"},
		},
		entity.Exercicio{
			Nome:        "extensão triceps sentado",
			Dificuldade: 2,
			Grupo:       []string{"triceps"},
		},
	}
	legs := []entity.Exercicio{
		entity.Exercicio{
			Nome:        "agachamento",
			Dificuldade: 1,
			Grupo:       []string{"perna"},
		},
		entity.Exercicio{
			Nome:        "agachamento bulgaro",
			Dificuldade: 3,
			Grupo:       []string{"perna"},
		},
		entity.Exercicio{
			Nome:        "extensão panturrilha",
			Dificuldade: 1,
			Grupo:       []string{"panturrilha"},
		},
	}
	abs := []entity.Exercicio{
		entity.Exercicio{
			Nome:        "montanha",
			Dificuldade: 1,
			Grupo:       []string{"abdomen"},
		},
		entity.Exercicio{
			Nome:        "l sit",
			Dificuldade: 1,
			Grupo:       []string{"abdomen"},
		},
	}
	repo := ExerciciosRepo{
		MFindPull: &pull,
		MFindPush: &push,
		MFindLegs: &legs,
		MFindAbs:  &abs,
	}
	return repo
}
