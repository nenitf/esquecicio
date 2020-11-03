package esquecicio_test

import (
	"sort"
	"testing"

	"github.com/nenitf/esquecicio/pkg/esquecicio/mock"
	esquecicio "github.com/nenitf/esquecicio/pkg/esquecicio/usecase"
)

/*
func TestDeveTer3Exercicios(t *testing.T) {
	sut := esquecicio.NewTreinoCalistenia(mock.ExerciciosRepoPadrao())

	var tests = []struct {
		in string
	}{
		{"pull"},
		{"push"},
		{"legs"},
	}
	for _, test := range tests {
		dto := esquecicio.TreinoCalisteniaDTO{
			Tipo: test.in,
		}

		treino, err := sut.Execute(dto)
		if err != nil {
			t.Fatal(err)
		}

		got := len(treino.Exercicios)
		want := 3
		if got != want {
			t.Error("For", test.in, "got:", got, "want:", want)
		}
	}
}
*/

func TestDeveTerOmbroPeitoTricepsAbdomenEmPush(t *testing.T) {
	sut := esquecicio.NewTreinoCalistenia(mock.ExerciciosRepoPadrao())
	dto := esquecicio.TreinoCalisteniaDTO{
		Tipo: "push",
	}

	treino, err := sut.Execute(dto)
	if err != nil {
		t.Fatal(err)
	}

	var got []string
	for _, e := range treino.Exercicios {
		got = append(got, e.Grupo...)
	}
	want := []string{"ombro", "peito", "triceps", "abdomen"}
	sort.Strings(got)
	sort.Strings(want)
	for i := 0; i < len(want); i++ {
		encontrou := false
		for j := 0; j < len(got); j++ {
			if want[i] == got[j] {
				encontrou = true
				break
			}
		}
		if !encontrou {
			t.Error("got:", got, "want:", want[i])
		}
	}
}
