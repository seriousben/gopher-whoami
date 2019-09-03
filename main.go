package main

import (
	"log"
	"os"
	"strconv"
	"time"

	gopher "github.com/prologic/go-gopher"
)

var started string

func gopherHost() string {
	return os.Getenv("GOPHER_HOST")
}

func gopherPort() int {
	i, _ := strconv.Atoi(os.Getenv("GOPHER_PORT"))
	return i
}

func index(w gopher.ResponseWriter, r *gopher.Request) {
	w.WriteInfo("started: " + started)
	w.WriteInfo("now:     " + time.Now().UTC().Format(time.RFC3339))
	w.WriteInfo("addr:    " + gopherHost() + ":" + os.Getenv("GOPHER_PORT"))
	w.WriteInfo("")
	w.WriteInfo("++++++++++++++++++++++++++++++")
	w.WriteInfo("+        gopher woami        +")
	w.WriteInfo("++++++++++++++++++++++++++++++")
	w.WriteInfo("")
	w.WriteInfo("")
	w.WriteItem(&gopher.Item{
		Type:        gopher.DIRECTORY,
		Selector:    "/hello",
		Description: "hello",
		Host:        gopherHost(),
		Port:        gopherPort(),
	})
	w.WriteItem(&gopher.Item{
		Type:        gopher.FILE,
		Selector:    "/foo",
		Description: "foo",
		Host:        gopherHost(),
		Port:        gopherPort(),
	})
	w.WriteItem(&gopher.Item{
		Type:        gopher.DIRECTORY,
		Selector:    "/",
		Description: "(Remote) Floodgap",
		Host:        "gopher.floodgap.com",
		Port:        70,
	})
}

func hello(w gopher.ResponseWriter, r *gopher.Request) {
	w.WriteInfo("Hello World!")
}

func foo(w gopher.ResponseWriter, r *gopher.Request) {
	w.Write([]byte("Foo!"))
}

func main() {
	started = time.Now().UTC().Format(time.RFC3339)
	mux := gopher.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/foo", foo)
	mux.HandleFunc("/hello", hello)

	log.Fatal(gopher.ListenAndServe("localhost:7000", mux))
}
