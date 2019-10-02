# Notes for _The Go Programming Language_

## Preface

### The Language

On the surface Go is similar to C. Go is known for its new and efficient facilities for concurentcy, as well as its unusually felxible take on data abstraction and object-oriented programming.

Go is a general-purpose language, but it is especially good for building infrastructure, tools, and systems for programmers. Go's balance between expressiveness and safety makes it a populare replacement for untyped scripting languages. Programs written in Go are generally faster and experience far fewer unexpected type errors than you might expect from programs written in dynamic languages.

Go's compiler, tools, and libraries are all open-source with contributions coming from a worldwide community.

### The Go Project

The Go project was born out of necessity due to several software systems at Google suffering from an explosion of complexity.

> Only through simplicity of design can a system remain stable, secure, and coherent as it grows.

The Go project is more than just the language. It is also comprised of tools, standard libraries, and a culture of radical simplicity. Go comes with very few features, it focuses on doing the basics well and is unlikely to add more. The language is mature and stable and guarantees backwards compatability.

## Tutorial

### Hello, World

Go code is organized into packages similar to libraries or modules from other languages. A package is one or more `.go` source files in a single directory that define what a package does.

A source file starts with a package declaration. In the `helloworld` example, the package is `main` stating which package the file belongs to, followed by a list of other packages that it imports, and then declarations of the program that are stored in that file.

The `fmt package` contains functions for printing formatted output and scanning input.

`Package main` is special. It defines a standalone executable, not a library. The `main function` is also special. It marks the beginning of the program's execution.

> A program will not compile if there are missing imports or if there are unnecessary ones.

### Command-Line Arguments



## Program Structure

## Basic Data Types

## Composite Types

## Functions

## Methods

## Interfaces

## Goroutines and Channels

## Concurency with Shared Variables

## Packages and the Go Tool

## Testing

## Reflection

## Low-Level Programming
