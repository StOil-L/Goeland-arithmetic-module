# Goéland - Arithmetic module first version

Goéland is an automated theorem prover using the tableau method for first order logic.

All the work shown in this repository is derived from the work of the team working on Goéland.

Please note that Goéland is licensed under the CeCILL 2.1 License.

### You can find the Goéland prover on [its dedicated GitHub repository](https://github.com/GoelandProver/Goeland).

## Authors

Cédric BERTHET

Enzo GOULESQUE

[Lorenzo PUCCIO](https://github.com/StOil-L)

[Margaux RENOIR](https://www.linkedin.com/in/margaux-renoir-244479220/)

[Tom SIMULA](https://www.linkedin.com/in/tom-simula-5039b8193/)

## Context

This project was carried out during my studies at the University of Montpellier as a third year Computer Science Bachelors Degree student.

It constitutes as the sole assignment for the L3 Programming Project subject (subject code: HLIN405). Its preliminary study work was carried out as the sole assignment for the CMI Long Integration Project (subject code: HAI508I).

This repository provides the first version of Goéland's arithmetic module. Its theory rests in the use of the Simplex and Branch-&-Bound algorithms, as well as the use of Gomory Cuts (thus offering the Branch-&-Cut algorithm as an improvement to the Branch-&-Bound algorithm). These, like most of the Goéland prover, were implemented in Go.

## Features

Depending on which file the user executes, they can access different features.

The features accessed through the BranchAndCut folder executable are:
- run resolution tests of the Simplex algorithm with the odd number option of your choice
- run resolution tests of the Branch & Cut algorithm with the even number option of your choice

The ArithmTests folder executable allows the user to run a number of unit tests for various operations on various types (integers, fractions). These operations include:
- unary:
  - negative
  - type check: integer, fraction
  - normalization (eg. = becomes <= & >=)
  - approximation: floor, ceiling, truncate, round
- binary:
  - sum
  - difference
  - product
  - quotient_, remainder_:
    - E: euclidian division
    - T: truncated to integer
    - T: floored to integer
  - equality
  - comparison: less than, greater than, less or equal to, greater or equal to

These operations also apply to each step of the Simplex algorithm.

## Getting Started

In order to run this, please ensure your computer has version 1.17+ of Go installed.

Once you've made sure of that, navigate to each folder and run the `go build` command. This should build an executable file, no matter your operating system. You are then free to run any of the executable files from your terminal.
