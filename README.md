# Hexagonal Architecture Golang with Framework Gin


## Description

This repository is to reference the hexagonal architecture folder structure


## Structure

> This is structure to folder in the project.

    ├── CMD
    │   └── httpServer
    ├─── internal         
    │        ├── core
    │        │    ├─── domain
    │        │    ├─── ports
    │        │    └─── services    
    │        ├── handlers    
    │        └── repositories    
    └─── pkg

## Getting Started

You only need the following steps:

* Install dependencies with `go get -u dependencies`
* Run the software with `cd cmd/httpServer && go run .`