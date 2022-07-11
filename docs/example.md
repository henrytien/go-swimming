# Go example projects

This repository contains a collection of Go programs and libraries that demonstrate the language, standard libraries, and tools.



## Getting Started

[Golang documentation](https://go.dev/doc/)

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

```
$ cd outyet
$ go build
```

A web server that answers the question: "Is Go 1.x out yet?"

Topics covered:

- Command-line flags ([flag](https://golang.org/pkg/flag/))
- Web servers ([net/http](https://golang.org/pkg/net/http/))
- HTML Templates ([html/template](https://golang.org/pkg/html/template/))
- Logging ([log](https://golang.org/pkg/log/))
- Long-running background processes
- Synchronizing data access between goroutines ([sync](https://golang.org/pkg/sync/))
- Exporting server state for monitoring ([expvar](https://golang.org/pkg/expvar/))
- Unit and integration tests ([testing](https://golang.org/pkg/testing/))
- Dependency injection
- Time ([time](https://golang.org/pkg/time/))

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

**testing**

> Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form.

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

**Channels**

```go
// Replace the pollSleep with a closure that we can block and unblock.
	sleep := make(chan bool)
	pollSleep = func(time.Duration) {
		sleep <- true
		sleep <- true
	}
```

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

### build error

```
./main.go:7:2: imported and not used: "html/template"
./main.go:26:19: cannot use NewServer(*version, changeURL, *pollPeriod) (value of type *Server) as type http.Handler in argument to http.Handle:
	*Server does not implement http.Handler (missing ServeHTTP method)
```

Because of the `ServerHTTP` interface does not implement.



## [appengine-hello](https://github.com/golang/example/blob/master/appengine-hello)

A trivial "Hello, world" App Engine application intended to be used as the starting point for your own code. Please see [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go) and [Quickstart for Go in the App Engine Standard Environment](https://cloud.google.com/appengine/docs/standard/go/quickstart).

## App Engine

[app.yaml](https://cloud.google.com/appengine/docs/standard/go/config/appref)

### Building a Go App on App Engine

To set up your `main.go` file:

1. In your `go-app/` folder, create a `main.go` file.

2. Add the `package main` statement to treat your code as an executable program:

   ```
   package main
   ```

   To successfully deploy a service in the Go 1.12+ runtimes, a `package main` statement must be defined at the beginning of at least one of your Go source files.

   

   **Learn more**: You do not have to put the `app.yaml` file and the `main.go` file in the same directory. You can specify the directory containing your main package by using `main:` in your project's `app.yaml` file. For more information, see the [`app.yaml` reference](https://cloud.google.com/appengine/docs/standard/go/config/appref).

   

3. Import the following packages:

   [appengine/go11x/helloworld/helloworld.go](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   [View on GitHub](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   ```golang
   import (
           "fmt"
           "log"
           "net/http"
           "os"
   )
   ```

4. Define your HTTP handler:

   [appengine/go11x/helloworld/helloworld.go](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   [View on GitHub](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   ```golang
   // indexHandler responds to requests with our greeting.
   func indexHandler(w http.ResponseWriter, r *http.Request) {
           if r.URL.Path != "/" {
                   http.NotFound(w, r)
                   return
           }
           fmt.Fprint(w, "Hello, World!")
   }
   ```

   The `http.ResponseWriter` object assembles the HTTP server response; by writing to it, you send data to the browser. The `http.Request` object is a data structure that represents the incoming HTTP request.

5. Register your HTTP handler:

   [appengine/go11x/helloworld/helloworld.go](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   [View on GitHub](https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine/go11x/helloworld/helloworld.go)

   ```golang
   func main() {
           http.HandleFunc("/", indexHandler)
   
           port := os.Getenv("PORT")
           if port == "" {
                   port = "8080"
                   log.Printf("Defaulting to port %s", port)
           }
   
           log.Printf("Listening on port %s", port)
           if err := http.ListenAndServe(":"+port, nil); err != nil {
                   log.Fatal(err)
           }
   }
   ```

   The `main` function is the entry point of your executable program, so it starts the application. It begins with a call to the `http.HandleFunc` function which tells the `http` package to handle all requests to the web root (`"/"`) with the `indexHandler` function.

   If the `PORT` environment variable is not set, port `8080` is used as a default. When your app is running on App Engine, the `PORT` environment variable is set for you, but when testing your app locally, you can set `PORT` to any preferred value.

   

## FAQ

- Why the C function can't return multiply value, but the Golang function could be do that?

  You can read these articles and getting the answer.

  - [The Function Stack](https://www.tenouk.com/Bufferoverflowc/Bufferoverflow2a.html)
  - [Why do byte spills occur and what do they achieve?](https://stackoverflow.com/questions/16453314/why-do-byte-spills-occur-and-what-do-they-achieve)
  - [Friday Q&A 2011-12-16: Disassembling the Assembly, Part 1](https://mikeash.com/pyblog/friday-qa-2011-12-16-disassembling-the-assembly-part-1.html)
  - [x86 calling conventions](https://en.wikipedia.org/wiki/X86_calling_conventions)
  - [Call Stack](https://en.wikipedia.org/wiki/Call_stack)
  - [Chapter I: A Primer on Go Assembly](https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md)

- Use of defer in Go?

  https://stackoverflow.com/questions/47607955/use-of-defer-in-go

- How does ServeHTTP work?

  https://stackoverflow.com/questions/49668070/how-does-servehttp-work

## Reference

- [Go Modules with Private GIT Repository](https://medium.com/swlh/go-modules-with-private-git-repository-3940b6835727)

  



