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

### 1.1 Hello, World

Go code is organized into packages similar to libraries or modules from other languages. A package is one or more `.go` source files in a single directory that define what a package does.

A source file starts with a package declaration. In the `helloworld` example, the package is `main` stating which package the file belongs to, followed by a list of other packages that it imports, and then declarations of the program that are stored in that file.

The `fmt package` contains functions for printing formatted output and scanning input.

`Package main` is special. It defines a standalone executable, not a library. The `main function` is also special. It marks the beginning of the program's execution.

> A program will not compile if there are missing imports or if there are unnecessary ones.

### 1.2 Command-Line Arguments

#### Comments

```Go
// In Go, comments look like this

/* 
 Or like
 this
 */
```

#### String Concatenation
Strings can be concatentated with the `+` operator

```Go
s += "string" + "concatenation"
```

#### Increment and Decrement
Increment and Decrement are postfix only; allowing `i++` but not `++i` for example.

#### Loops

* for loops are the only loop statements in Go.

```Go
for initialization; condition; post {
    // Zero or more statements
}
```

* If initialization and post are omitted, you get a traditional while loop.

```Go
for condition {
    // Break when condition is false
}
```

* Infinite Loops

```Go
for {
    // Broken with break or return statements
}
```

#### Temporary Values

Go does not allow unused variables, so when you need one, use the blank identifier `_`.

#### Variable Declaration

Each of the following declarations are equivalent:

* s:= " "
* var s string
* var s = " "
* var s string = " "

Most of the time, you'll want to use one of the first two, with `s := " "` being preferable when the initial value is _not important_, and with `var s string` being preferable with the inital value _is important_.

### 1.3 Finding Duplicate Lines

#### bufio Package

`bufio` is used for input and output in Go. This package gives us access to the `Scanner` type which reads input and breaks it into lines or words.

```Go
input := bufio.NewScanner(os.Stdin)
```

`input.Scan()` reads the next line and removes the newline, returning true when there is input, and false if there is no more.

`input.Text()` retrieves the result of `input.Scan()`.

### 1.6 Fetching URLs Concurrently

#### Goroutines

A `gorutine` is a concurrent function execution, and a `channel` is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.

### 1.7 A Web Server

Go allows a simple statement such as a local variable declaration to precede the if condition. This reduces the scope of the variable `err` below:

```Go
// Good
err := r.ParseForm()
if err != nil {
  log.print(err)
}

// Better
if err := r.ParseForm() {
  log.print(err)
}
```

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
