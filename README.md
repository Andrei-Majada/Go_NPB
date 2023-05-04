# The NAS Parallel Benchmarks for evaluating Go parallel programming

Desenvolvedores: Caroline Evangelista e Andrei Majada

## Proposta

A proposta desse projeto é a de reproduzir o NPB utilizando a linguagem Go e realizar análises comparativas entre o código desenvolvido e o [NPB Benchmark em C++].

## Tech

Dillinger uses a number of open source projects to work properly:

- [Go] - Linguagem de programação Go.
- [Makefile] - Makefiles são uma maneira simples de organizar a compilação do código.


E o projeto está disponível em um repositório publico no [GitHub].

## Installation

O projeto requer que tenha [Go] instalado na máquina, de preferência na versão  go 1.18.

Para limpar builds anteriores:
```sh
make clean
```

Para executar uma versão:
```sh
make <benchmark-name> CLASS=<class>
```
Onde o <benchmark-name> pode ser "ep" ou "is". 

E a <class> pode ser "S", "W", "A", "B", "C", "D", "E"



   [NPB Benchmark em C++]: <https://github.com/GMAP/NPB-CPP>
   [GitHub]: <https://github.com/>
   [Go]: <https://go.dev/>
   [Makefile]: <https://www.cs.colby.edu/maxwell/courses/tutorials/maketutor/>
