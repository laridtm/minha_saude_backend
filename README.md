# Minha saúde - Backend

Backend escrito em Go para o projeto mobile Minha saúde

## Getting Started

Primeiro precisamos baixar todas as dependências do projeto que estão descritas no arquivo [go.mod](go.mod). Para isso basta executar o comando

```sh
$ go mod
```

Feito isso, precisamos subir as dependências de infraestrutura do projeto, que no momento é apenas o MongoDB:

```sh
$ make env
```

Esse comando acima, irá subir um MongoDB na versão 5.0 e irá incluir algumas informações já de início para fins de teste. Essas informações podem ser encontradas [aqui](build/mongo-init.js).

Para startar o projeto, você pode executar o comando abaixo. E para parar, simplemente pressione `Ctrl + C` no seu teclado.

```sh
$ make start
```

Assim como temos um comando para subir as dependências de infraestrutura do projeto, também existe um comando para remove-las:

```sh
$ make env-stop
```

