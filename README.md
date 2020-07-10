# NADC - Intro to RPC
The second workshop of Night Login App Development Community (NADC) - Intro to REST RPC using Go. This is an example of basic implementation of traditional RPC. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

1. Clone repository: `git clone git@github.com:ridwanakf/nadc-intro-to-rpc.git`
2. Run dep on both server and client directories: `dep ensure -v`

### Run Project

Running server:

```$xslt
make run-server
```

Running client:

```$xslt
make run-client
```

## Directory Structure

This repository is organized in the following directory structure.

```
nadc-intro-to-rpc
|-- client                                # Directory of RPC Client code
|   |-- internal                          # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app.go                            # main package, and entry point of the app
|   |-- vendor                            # Dependencies folder that's maintained by dep tool https://golang.github.io/dep/
|   |-- Gopkg.lock                        # https://golang.github.io/dep/docs/Gopkg.lock.html
|   |-- Gopkg.toml                        # https://golang.github.io/dep/docs/Gopkg.toml.html
|-- server                                # Directory of RPC Server code
|   |-- internal                          # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app.go                            # main package, and entry point of the app
|   |-- vendor                            # Dependencies folder that's maintained by dep tool https://golang.github.io/dep/
|   |-- Gopkg.lock                        # https://golang.github.io/dep/docs/Gopkg.lock.html
|   |-- Gopkg.toml                        # https://golang.github.io/dep/docs/Gopkg.toml.html
|-- Makefile                              # Makefile

```

## Tech Stacks

- Golang
