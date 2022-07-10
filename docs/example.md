# Go example projects

This repository contains a collection of Go programs and libraries that demonstrate the language, standard libraries, and tools.

## Clone the project

```shell
$ git clone https://go.googlesource.com/example
$ cd example
```

## [hello](https://github.com/henrytien/go-swimming/blob/master/src/example/hello/hello.go)

What is learning from this demo?

```
$ cd hello
$ go build
```

A trivial "Hello, world" program that uses a stringutil package.

Command [hello](https://github.com/golang/example/blob/master/hello) covers:

- The basic form of an executable command
- Importing packages (from the standard library and the local repository)
- Printing strings ([fmt](https://golang.org/pkg/fmt/))

Library [stringutil](https://github.com/golang/example/blob/master/stringutil) covers:

- The basic form of a library
- Conversion between string and []rune
- Table-driven unit tests ([testing](https://golang.org/pkg/testing/))

## outyet

### packages

**expvar**

> Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.

**flag**

> Package flag implements command-line flag parsing.

**fmt**

> Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
>
> General:
>
> ```
> %v	the value in a default format
> 	when printing structs, the plus flag (%+v) adds field names
> %#v	a Go-syntax representation of the value
> %T	a Go-syntax representation of the type of the value
> %%	a literal percent sign; consumes no value
> ```

**html/template**

> Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.
>
> The documentation here focuses on the security features of the package. For information about how to program the templates themselves, see the documentation for text/template.

**log**

> Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error and prints the date and time of each logged message. Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.

**http**

> Package http provides HTTP client and server implementations.
>
> Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:

**sync**

> Package sync provides basic synchronization primitives such as **mutual exclusion locks**. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.
>
> Values containing the types defined in this package should not be copied.

**time**

> Package time provides functionality for measuring and displaying time.
>
> The calendrical calculations always assume a Gregorian calendar, with no leap seconds.

### golang syntax

**Goroutines**

```go
go s.poll()
```

> A *goroutine* is a lightweight thread managed by the Go runtime.
>
> ```go
> func main() {
> 	message := "Hello Go routine"
> 	go func() {
> 		fmt.Println(message)
> 	}()
> 	message = "Message Changed"
> 	time.Sleep(time.Millisecond * 10)
> }
> ```

**RLock**

> // RLock locks rw for reading.
>
> //
>
> // It should not be used for recursive read locking; a blocked Lock
>
> // call excludes new readers from acquiring the lock. See the
>
> // documentation on the RWMutex type.



### build error

```
./main.go:7:2: imported and not used: "html/template"
./main.go:26:19: cannot use NewServer(*version, changeURL, *pollPeriod) (value of type *Server) as type http.Handler in argument to http.Handle:
	*Server does not implement http.Handler (missing ServeHTTP method)
```

Because of the `ServerHTTP` interface does not implement.

