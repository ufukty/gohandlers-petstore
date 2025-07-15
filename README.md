# Petstore with Gohandlers

This repository implements a petstore example with [Gohandlers](https://github.com/ufukty/gohandlers). Click the link to learn about Gohandlers.

## Understanding the files

There are user provided and Gohandlers generated files in this repository. By default Gohandlers suffixes the files it generates with `.gh.go`.

### User provided files

- [`main.go`](main.go) demonstrates the usage of `ListHandlers` method with `http.ServeMux`.

- [`handlers/pets`](handlers/pets) directory contain [`create.go`](handlers/pets/create.go), [`delete.go`](handlers/pets/delete.go), [`get.go`](handlers/pets/get.go) and [`list.go`](handlers/pets/list.go) shows the suggested format for handler files which contain one handler, its request binding type and optionally its response binding type.

- [`Makefile`](Makefile), [`handlers/pets/Makefile`](handlers/pets/Makefile) shows the gohandlers commands and exact args used in generating this example.

### Gohandlers generated files

- [`handlers/pets/gh.go`](handlers/pets/gh.go) contains the builder, parser and validator for request bindings; builder and writer for response bindings and path, method and reference listers for global handlers and method handlers.

- [`client/client.gh.go`](client/client.gh.go) contains the real `Client` implementation as well as the `Interface` declaration for using `Mock` implementation during tests.

- [`handlers/pets/gh.yml`](handlers/pets/gh.yml) contains information of handler names, paths and methods that are meant help outside clients.

## Runtime output

```sh
$ go run .

registering ListPets as GET /pets
registering CreatePet as POST /create-pet
registering DeletePet as DELETE /pets/{id}
registering GetPet as GET /pets/{id}
```

## Re-generating files

```sh
find . -name '*.gh.go' -delete
make -B all
git status
```
