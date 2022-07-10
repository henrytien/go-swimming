package main

import (
	"expvar"
	"flag"
	"fmt"
//	"html/template"
    "io"
	"log"
	"net/http"
	"sync"
	"time"
)

// Command-line flags.
var (
	httpAddr   = flag.String("http", ":8080", "Listen address")
	pollPeriod = flag.Duration("poll", 1*time.Second, "Poll period")
	version    = flag.String("version", "1.8", "Go version")
)

const baseChangeURL = "https://www.google.com"

func main() {
	flag.Parse()
	changeURL := fmt.Sprintf("%s %s", baseChangeURL, *version)
    log.Printf(changeURL)
    //http.HandleFunc("/hello",getHello)
	http.Handle("/", NewServer(*version, baseChangeURL, *pollPeriod))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// Server implements the outyet server.
// It serves the user interface (it's an http.Handler)
// and polls the remote repository for changes.
type Server struct {
	version string
	url     string
	period  time.Duration

	mu  sync.RWMutex // protects the yes variable
	yes bool
}

// NewServer returns an initialized outyet server.
func NewServer(version, url string, period time.Duration) *Server {
	s := &Server{version: version, url: url, period: period}
	go s.poll()
	return s
}

// poll polls the change URL for the specified period until the tag exits.
// the it sets the Server's yes field true and exits.
func (s *Server) poll() {
	for !isTagged(baseChangeURL) {
        fmt.Println("poll")
		pollSleep(s.period)
	}
	s.mu.Lock()
	s.yes = true
	s.mu.Unlock()
	pollDone()
    log.Print("isTagged")
}

// Exported variable for monitoring the server.
// These are exported via HTTP as a JSON object at /debug/vars.
var (
	hitCount       = expvar.NewInt("hitCount")
	pollCount      = expvar.NewInt("pollCount")
	pollError      = expvar.NewString("pollErorr")
	pollErrorCount = expvar.NewInt("pollErrorCount")
)

// isTagged makes and HTTP HEAD request to the given URL and reports whether it
// returned a 200 OK response.
func isTagged(url string) bool {
	pollCount.Add(1)
    log.Print(pollCount)
    r, err := http.Head(url)
    if err != nil {
        log.Print(err)
        pollError.Set(err.Error())
        pollErrorCount.Add(1)
        return false
    }
    
    return r.StatusCode == http.StatusOK
}

// Hooks that may be overridden for integration tests.
var (
    pollSleep = time.Sleep
    pollDone = func() {
        log.Print("pollDone")
    }
)

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    hitCount.Add(1)
    log.Print(hitCount)
    s.mu.RLock()
    data := struct {
        URL string
        Version string
        Yes bool
    } {
        s.url,
        s.version,
        s.yes,
    }
    fmt.Printf("go / Request\n %s", data)
    io.WriteString(w, "This is called by Henry")
    s.mu.RUnlock()
    /*
    err := tmpl.Execute(w, data)
    if err != nil {
        log.Print(err)
    }
    */
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

/*
// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse('
<!DOCTYPE html><html><body><center>
	<h2>Is Go {{.Version}} out yet?</h2>
	<h1>
	{{if .Yes}}
		<a href="{{.URL}}">YES!</a>
	{{else}}
		No. :-(
	{{end}}
	</h1>
</center></body></html>
`))
*/
