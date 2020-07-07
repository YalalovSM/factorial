package server

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/yalalovsm/factorial/calculations"
)

const (
	factorialPath = "/factorial/"
)

// Server ...
type Server struct{}

// New ...
func New() *Server {
	return &Server{}
}

// Start ...
func (s *Server) Start() error {
	http.HandleFunc(factorialPath, handleFactorial)
	return http.ListenAndServe(":8080", nil)
}

func handleFactorial(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, factorialPath)
	if len(p) > 1 {
		n, err := strconv.Atoi(p[1])
		if err != nil {
			responseBadRequest(w)
			return
		}

		f := calculations.FactorialTree(n)
		responseWithFactorial(w, f)
	} else {
		responseBadRequest(w)
	}
}

func responseBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("You have to provide number!\n"))
}

func responseWithFactorial(w http.ResponseWriter, f *big.Int) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(f.String()))
}
