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

### Names

All names start with either a letter or an underscore `_`. By convention, names should be in lower camel case. If a name begins with an upper-case letter it is `exported` which means it is visible and accessable outside of its package.

### Short Variable Declaration

Delcare and initialize a local variable:

```go
name := expression
```

Short variable declarations should be used over a normal variable declaration in most cases.

### Packages and Files

Packages in Go are the same as modules or libraries in other languages.

```go
$GOPATH/src/gopl.io/ch1/helloworld
```

`helloworld` contains one or more .go files needed for the helloworld package. Packages act as namespaces. You can hide information in packages. Remember: indentifiers starting with an upper-case letter are `exported` or visible outside of the package.

## Basic Data Types

### Integers

Integers allow signed and unsigned integer math. They come in four signed sizes: `int8`, `int16`, `int32`, and `int64`; as well as four unsigned sizes: `uint8`, `uint16`, `uint32`, `uint64`. There is also a generic `int` and `uint` that use the most efficient size. Keep in mind that an `int` of size 8 is not the same type as a `int8`. There also exists a `uintptr` type that is big enough to hold the bits of a pointer value.

### Floats

Floats come in two sizes: `float32` and `float64`. `float32` provides 6 decimals digits of percision, where as `float64` provides 15 decimal digits of percision. Use `float64` wherever possible to help avoid rounding errors.

### Complex Numbers

Complex numbers come in two sizes: `complex64` and `complex128`. This type creates a complex number from real and imaginary components.

```go
var x complex128 = complex(1,2)  // 1+2i
var y complex128 = complex(3,4)  // 3+4i
fmt.Println(x * y)               // "(-5+10i)"
fmt.Println(real(x * y))         // "-5"
fmt.Println(imag(x * y))         // "10"
```

Note the built in `real()` and `imag()` functions extract the real and imaginary components.

## Composite Types

### Arrays

Arrays are fixed-length sequences of zero or more elements of a single type. Because of their fixed-length, they are rarely used.

```go
var a [3]int               // array of 3 ints
fmt.Println(a[0])          // print the first element
fmt.Println(a[len(a) - 1]) // print the last element
```

### Slices

Slices are connected to an underlying array. Slices gives access to a subsequence of the array. Slices have three compnents: `pointer` to the first element of the slice (not the array), `length` the number of slice elements (this number cannot exceed the cap), and `capacity` is usually the number of elements between the start of the slice and the end of the underlying array.

### Maps

Maps are references to a hash table. Written: `map[K]V` where _K_ and _V_ are the types of the map's _keys_ and _values_.

```go
ages := make(map[string]int) // empty map of ints with strings as keys

ages := map[string]int {     // map literal
 "alice": 31,
 "charlie": 34,
}
```

## Functions

### Multiple Return Values

Functions can return more than one result. Often times this is the result and an error value.

```go
func name(args)([]string, error) {
 // ...
 return nil, err
}
```

### Anonymous Function

Named functions must be declared at the package level, where as function literals can denote a function value within any expression.

```go
// squares() returns a function that returns
// the next square number each time it is called
func squares() func() int {
  var x int
  return func() int {
    x++
    return x * x
  }
}

func main() {
   f := squares()
   fmt.Println(f()) // "1"
   fmt.Println(f()) // "4"
   fmt.Println(f()) // "9"
   fmt.Println(f()) // "16"
```

### Deferred Function Calls

Deferred function calls call a function from inside another function. The deferred function and arguments are evaluated when the line is executed but the call does not get made until the containing function is finished. This is great for opening/closing files or locking/unlocking resources for example.

```go
func title(url string) error {
   resp, err := http.Get(url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   // ...
}
```

## Methods

### Method Declarations

Method declarations are like function declarations but with an extra parameter before the name. This parameter attaches the function to the type of that parameter. This parameter is called the `receiver`. The receiver is vommonly named as the first letter of the type it represents.

`p.Distance` is called a `selector` because it selects the appropriate Distance method. This allows us to use the name `Distance` for other methods, so long as they belong to different types.

### Methods with a Pointer Receiver

If a function needs to update a variable of if we wish to avoid copying an argument, we must pass the address of the variable using a pointer.

```go
func (p *Point) ScaleBy(factor float64) {
   p.X *= factor
   p.Y *= factor
}
```

The name of this method is `(*Point).ScaleBy`. The parentheses are necessary.

### Nil is a Valid Receiver Value

```go
// An IntList is a linked list of integers
// A nil *IntList represents an empty list
type IntList struct {
   Value int
   Tail *IntList
}

// sum returns the sum of the list elements
func(list *IntList) Sum() int {
   if list == nil {
      return 0
   }
   return list.Value + list.Tail.Sum()
}
```

## Interfaces

### Interfaces as Contracts

Interfaces are abstract types. Interfaces don't expose the representation or external structure of its values or the set of basic operations they support. They only reveal some of its methods.

### Interface Types

An interface type specifies a set of methods that a concrete type must possess to be considered an instance of that interface.

```go
package io

type Reader interface {
   Read(p[]byte) (n int, err error)
}

// Embedding an interface allows us to
// name another interface instead of writing
// out all of its methods

type ReaderWriter interface {
   Reader
   Write(p[]byte) (n int, err error)
}
```

### Interface Satisfaction

A type satisfies an interface only if it posesses all the methods the interface requires.

## Goroutines and Channels

### Goroutines

Every concurrently executed activity is called a goroutine. You can assume a goroutine is similar to a thread, though there are some differences.

New goroutines are created with the `go statement`.

```go
f()    // Call f(); wait for a return
go f() // Create a new goroutine that calls f(); don't wait for a return
```

When the main function returns all goroutines are terminated and the program exits. Only the `main goroutine` can tell other goroutines to quit. No other goroutine may stop another, but there are ways to request a goroutine to stop.

### Channels

Channels are connections between concurrent goroutines. Channels allow goroutines to send values to another goroutine. Channels are type specific, and a channel's type is called it's `element type`.

```go
chan int
```

Using the `make` function to create a new channel.

```go
ch := make(chan int) // ch has type `chan int`
```

Channels can be told to `send`, `receive`, and `close`.

```go
// Send

ch <- x // Send x through the channel

// Receive

x = <- ch // Receive expression in an assignment statement

<- ch     // Receive statement; result is discarded

// Close

close(ch) // Subsequent attempts to send will result in panic
```

### Buffered/Unbuffered Channels

```go
ch = make(chan int)    // Unbuffered channel
ch = make(chan int, 0) // Unbuffered channel
ch = make(chan int, 3) // Buffered channel with a capasity of 3
```

## Concurency with Shared Variables

## Packages and the Go Tool

## Testing

## Reflection

## Low-Level Programming
