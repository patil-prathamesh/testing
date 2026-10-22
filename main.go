package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Store struct {
	store  map[int]Student
	nextId int
	mu     sync.Mutex
}

func (s *Store) Get(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("[recovery] caught a panic")
	// 	}
	// }()
	fmt.Println(s.store)
	sid := r.PathValue("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		WriteJSON(w, 400, "id is non-numeric")
		return
	}

	panic("panic in get ")

	s.mu.Lock()
	student, ok := s.store[int(id)]
	s.mu.Unlock()

	if !ok {
		WriteJSON(w, 404, "not found")
		return
	}

	WriteJSON(w, http.StatusOK, student)
}

func (s *Store) Create(w http.ResponseWriter, r *http.Request) {
	var student Student
	if err := ReadJSON(w, r, &student); err != nil {
		var mbe *http.MaxBytesError
		if errors.As(err, &mbe) {
			writeError(w, http.StatusRequestEntityTooLarge, "body too large")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
	}

	student.Name = strings.TrimSpace(student.Name)
	if student.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is empty")
	}

	s.mu.Lock()
	student.ID = s.nextId
	s.store[s.nextId] = student
	s.nextId++
	s.mu.Unlock()

	fmt.Println(student)

	WriteJSON(w, http.StatusCreated, "created")
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("body must contain a single json object")
	}
	return nil
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("content/type", "application/json")
	w.WriteHeader(status)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("inside panic")
				fmt.Println(string(debug.Stack()))
				writeError(w, 500, "panic")
			}
		}()
		next.ServeHTTP(w, r)
		fmt.Println("after recover func")
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sr := &statusRecorder{w, 200}
		fmt.Println("before")
		next.ServeHTTP(sr, r) // call the wrapped handler
		fmt.Println("after")
		log.Printf("%s %d %s %v", r.Method, sr.status, r.URL.Path, time.Since(start))
	})
}

func main() {
	store := Store{map[int]Student{}, 1, sync.Mutex{}}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /students", store.Create)
	mux.Handle("GET /students/{id}", Logging(Recover(http.HandlerFunc(store.Get))))

	http.ListenAndServe(":8080", mux)
}
