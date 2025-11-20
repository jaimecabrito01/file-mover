# FileMover Version 1.0.0 !!!

Um movimentador de arquivos que organiza seus downloads.
- Pega qualquer arquivo novo na pasta de downloads e move para outra de acordo com seu tipo
  Ex: se for .mp4 -> /Videos


## Índice

- [Instalação](#instalação)
- [Uso](#uso)
- [Configuração](#configuração)
- [Contribuição](#contribuição)
- [Licença](#licença)
- [Autores](#autores)

## Instalação
Antes de fazer a instalação recomendo baixar o Go para testar a aplicação


[Golang](https://go.dev/doc/install)

Clone o repositório:

```bash
git clone https://github.com/jaimecabrito01/separador-arquivos-service.git
cd separador-arquivos-service
```

Instale as dependências:

```bash
go mod tidy
```

## Uso

Execute o programa principal:

```bash
go run main.go
```
Ou compile para binário:

```bash
go build -o meu-app main.go
./meu-app
```

Adapte as instruções acima conforme a estrutura do seu projeto.

## Configuração

 ainda nao tem.

## Contribuição

Contribuições são bem-vindas! Para contribuir, siga os passos:

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/NomeDaFeature`)
3. Faça commit das suas alterações (`git commit -am 'Adiciona uma nova feature'`)
4. Envie para o branch (`git push origin feature/NomeDaFeature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## Autores

- [jaimecabrito01](https://github.com/jaimecabrito01)

