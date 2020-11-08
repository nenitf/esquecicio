package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/nenitf/esquecicio/pkg/esquecicio/entity"
	esquecicio "github.com/nenitf/esquecicio/pkg/esquecicio/usecase"
)

type Exercicios struct {
	Pull []entity.Exercicio
	Push []entity.Exercicio
	Legs []entity.Exercicio
	Abs  []entity.Exercicio
}

type Semana struct {
	Segunda []string
	Terca   []string
	Quarta  []string
	Quinta  []string
	Sexta   []string
	Sabado  []string
	Domingo []string
}

type Config struct {
	Exercicios Exercicios
	Semana     Semana
}

type ExerciciosRepo struct {
	Config Config
}

func (r ExerciciosRepo) FindPull() []entity.Exercicio {
	return r.Config.Exercicios.Pull
}

func (r ExerciciosRepo) FindPush() []entity.Exercicio {
	return r.Config.Exercicios.Push
}

func (r ExerciciosRepo) FindLegs() []entity.Exercicio {
	return r.Config.Exercicios.Legs
}

func (r ExerciciosRepo) FindAbs() []entity.Exercicio {
	return r.Config.Exercicios.Abs
}

func main() {
	cwd, err := os.Getwd()
	configfilename := "exercicios.json"
	configpath := filepath.FromSlash(cwd + "/" + configfilename)

	data, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Config não foi encontrado: ", err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Config não foi parseado: ", err)
	}

	r := ExerciciosRepo{
		Config: config,
	}
	s := esquecicio.NewTreinoCalistenia(r)

	var tipo []string

	switch time.Now().Weekday() {
	case time.Sunday:
		tipo = config.Semana.Domingo
		break
	case time.Monday:
		tipo = config.Semana.Segunda
		break
	case time.Tuesday:
		tipo = config.Semana.Terca
		break
	case time.Wednesday:
		tipo = config.Semana.Quarta
		break
	case time.Thursday:
		tipo = config.Semana.Quinta
		break
	case time.Friday:
		tipo = config.Semana.Sexta
		break
	case time.Saturday:
		tipo = config.Semana.Sabado
		break
	}

	dto := esquecicio.TreinoCalisteniaDTO{
		Tipos: tipo,
	}

	treino, err := s.Execute(dto)
	if err != nil {
		log.Fatal("Treino não foi criado, erro: ", err)
	}

	for _, exercicio := range treino.Exercicios {
		fmt.Println(exercicio.Nome)
	}
	fmt.Scanln()
}
