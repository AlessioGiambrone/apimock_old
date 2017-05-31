package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/pierreprinetti/apimock/store"
)

type Item interface {
	Value() []byte
	ContentType() string
	Read(p []byte) (n int, err error)
}

type Store interface {
	Get(string) Item
	Set(Item)
}

type StoreHandler struct {
	Store
}

func ItemFromRequest(r *http.Request) Item {
	body, _ := ioutil.ReadAll(r.Body)
	return store.NewItem(r.URL.Path, r.Header.Get("Content-Type"), body)
}

func (s *StoreHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, s.Get(r.URL.Path))
}

func (s *StoreHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	item := ItemFromRequest(r)
	s.Set(item)
}

func (s *StoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.HandleGet(w, r)
		return
	case http.MethodPut:
		s.HandlePut(w, r)
		return
	}
}

func set(key string, value []byte, contentType string) {
	if overrideContentType != "" {
		contentType = overrideContentType
	} else {
		if contentType == "" {
			contentType = defaultContentType
		}
	}
	store[key] = entry{Value: value, ContentType: contentType}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	log.Println("Not implemented", path)

	w.WriteHeader(http.StatusNotImplemented)
	_, err := w.Write([]byte("Not implemented"))
	check(err)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	e, ok := store[path]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", e.ContentType)
	_, err := w.Write(e.Value)
	check(err)
}

func headHandler(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	value, err := ioutil.ReadAll(r.Body)
	check(err)

	set(path, value, r.Header.Get("Content-Type"))

	getHandler(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]

	if _, ok := store[path]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(store, path)
	w.WriteHeader(http.StatusNoContent)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	localStore := store.New()

	router := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	router.HandleFunc("/{path:.*}", getHandler).Methods("GET")
	router.HandleFunc("/{path:.*}", headHandler).Methods("HEAD")
	router.HandleFunc("/{path:.*}", putHandler).Methods("PUT")
	router.HandleFunc("/{path:.*}", postHandler).Methods("POST")
	router.HandleFunc("/{path:.*}", deleteHandler).Methods("DELETE")
	router.HandleFunc("/{path:.*}", optionsHandler).Methods("OPTIONS")

	n := negroni.New(negroni.NewRecovery(), newLogger(), newCors())
	n.UseHandler(router)
	n.Run(host)
}
