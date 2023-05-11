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

Os dois kernels estão em diferentes diretórios, um chamado EP e o outro IS.

Caso queira utilizar o EP, basta acessar a pasta do EP:
```sh
cd EP
```

Para limpar builds anteriores:
```sh
make clean
```

Para executar a versão do EP:
```sh
make EP CLASS=<class>
```

Onde a <class> pode ser "S", "W", "A", "B", "C", "D", "E"
 

Caso queira utilizar o IS, basta acessar a pasta do IS:
 
```sh
cd IS
```

Para limpar builds anteriores:
```sh
make clean
```

Para executar a versão do IS:
```sh
make IS CLASS=<class>
```

Onde a <class> pode ser "S", "W", "A", "B", "C", "D", "E"



   [NPB Benchmark em C++]: <https://github.com/GMAP/NPB-CPP>
   [GitHub]: <https://github.com/>
   [Go]: <https://go.dev/>
   [Makefile]: <https://www.cs.colby.edu/maxwell/courses/tutorials/maketutor/>
