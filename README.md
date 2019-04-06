# CodeJam Solutions

A collection of my own solutions to past questions in [Google Code Jam](https://codingcompetitions.withgoogle.com/codejam/).

Just a fun, personal project.

# Running with input 
Structured as `package main` files with a `func main` that reads from stdin and outputs to stdout, and offloads to another function for the calculation.

Compile with `go build` and then pipe the input file to the binary.
`go build mything && ./mything < input`

# Tips
Don't forget to output `Case #X: ` at the start of your answers.
Note that this numbering starts at 1, *not* 0.

The version of Go used in CodeJam is `1.7.4` which is several versions behind. It does not include strings.Builder (introduced in 1.10)
