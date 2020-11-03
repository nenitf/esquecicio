# Esquecicio

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](https://gist.github.com/nenitf/1cf5182bff009974bf436f978eea1996#emojicom)

Gerador de circuito de treino

## Utilização

1. Acesse a pasta do [executável](https://github.com/nenitf/esquecicio/releases) ou `cmd/main.go`
2. Crie arquivo `exercicios.json` com base em *exercicios_????.json*
3. Execute com `cmd.exe` ou `go run .`
    - `.\cmd.exe -tipo=pull` ou `go run . -tipo=pushlegs`
    - `.\cmd.exe -help` ou `go run . -help`

## Desenvolvimento

### Testes

```sh
go test ./pkg/esquecicio/usecase
```

## Referências

- [Como Começar na Calistenia - Guia Prático](https://bookdown.org/kaiquegalois/guia_pratico/como-organizar-seus-treinos.html#intermedi%C3%A1rio---push-pull)
