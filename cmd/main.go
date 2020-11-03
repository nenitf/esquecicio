package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/nenitf/esquecicio/pkg/esquecicio/entity"
	esquecicio "github.com/nenitf/esquecicio/pkg/esquecicio/usecase"
)

type Config struct {
	Pull []entity.Exercicio
	Push []entity.Exercicio
	Legs []entity.Exercicio
	Abs  []entity.Exercicio
}

type ExerciciosRepo struct {
	config Config
}

func (r ExerciciosRepo) FindPull() []entity.Exercicio {
	return r.config.Pull
}

func (r ExerciciosRepo) FindPush() []entity.Exercicio {
	return r.config.Push
}

func (r ExerciciosRepo) FindLegs() []entity.Exercicio {
	return r.config.Legs
}

func (r ExerciciosRepo) FindAbs() []entity.Exercicio {
	return r.config.Abs
}

func main() {
	data, err := ioutil.ReadFile("exercicios.json")
	if err != nil {
		log.Fatal("Config não foi encontrado: ", err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Config não foi parseado: ", err)
	}

	r := ExerciciosRepo{
		config: config,
	}
	s := esquecicio.NewTreinoCalistenia(r)

	var padrao string

	switch time.Now().Weekday() {
	case time.Sunday:
		padrao = "push"
		break
	case time.Monday:
		padrao = "pulllegs"
		break
	case time.Tuesday:
		padrao = "pushlegs"
		break
	case time.Wednesday:
		padrao = "legs"
		break
	case time.Thursday:
		padrao = "pulllegs"
		break
	case time.Friday:
		padrao = "pushlegs"
		break
	case time.Saturday:
		padrao = "pull"
		break
	}

	flagTipo := flag.String("tipo", padrao, "pull, push ou legs")
	flag.Parse()

	dto := esquecicio.TreinoCalisteniaDTO{
		Tipo: *flagTipo,
	}

	treino, err := s.Execute(dto)
	if err != nil {
		log.Fatal("Treino não foi criado, erro: ", err)
	}

	for _, exercicio := range treino.Exercicios {
		fmt.Println(exercicio.Nome)
	}
}
