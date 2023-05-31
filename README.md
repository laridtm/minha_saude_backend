# Minha saúde - Backend

Backend written in Go for the Minha Saúde mobile project.

It is a personal health management application for the `Trabalho de Conclusão de Curso Sistemas de Informação - Universidade do Sul de Santa Catarina/ UNISUL`.

## Minha saúde - Frontend Mobile iOS

Frontend written in Swift for the project in this [link](https://github.com/laridtm/minha_saude)

### Architecture

The main responsibility of the project is to serve a RESTful interface so that the mobile service can communicate with the server. This responsibility is divided between four main layers:

**Model**: In this layer we have the definition of the structures of entities that are used in the project.

**Controller**: The layer that receives HTTP requests and converts them to a known structure within the project (previously defined in the Model layer).

**Service**: Layer that contains the business logic of the actions received in the control layer, such as validating whether the model received by the HTTP request is coherent.

**Repository**: Layer responsible for assembling all queries in the Mongo DB database.

![Captura de Tela 2023-05-31 às 17 19 31](https://github.com/laridtm/minha_saude_backend/assets/55598696/79b737e7-ef1e-4cee-92b9-88070cc0e013)

### Getting Started

First we need to download all the project dependencies that are described in the file[go.mod](go.mod). To do this, just run the command

```sh
$ go mod
```

Once that's done, we need to upload the project's infrastructure dependencies, which at the moment is just MongoDB:

```sh
$ make env
```

This command above will upload a MongoDB in version 5.0 and will include some initial information for testing purposes. This information can be found [here](build/mongo-init.js).

To start the project, you can run the command below. And to stop it, simply press `Ctrl + C` on your keyboard.

```sh
$ make start
```

Just as we have a command to upload the project's infrastructure dependencies, there is also a command to remove them:

```sh
$ make env-stop
```

### Author

| [![Larissa](https://avatars.githubusercontent.com/u/55598696?v=4&s=80)](https://github.com/laridtm/) | [@laridtm](https://github.com/laridtm/) |
| ------ | ------ |
