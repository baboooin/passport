package main

import (
	"encoding/json"
	"runtime"
	"io"
	"log"
	"net/http"
	// "strings"

	"github.com/gorilla/mux"
	// "github.com/gorilla/rpc"
)

func Static(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("static"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	//log 
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func postBody(r *http.Request) ( result interface{}, err error) {
	bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &result)
		return result, err
}

func main() {
	runtime.GOMAXPROCS(2)
	r := mux.NewRouter()
	
	//r.HandleFunc("/", Static).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/about").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/login").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/register").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/SignIn").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/SignUp").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	r.PathPrefix("/ForgotPassword").Handler(http.FileServer(http.Dir("./static/frontend/build/")))
	
	
	r.HandleFunc("/", login).Methods("POST")
	//r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/login", login).Methods("POST") // login & password auth
	r.HandleFunc("/getCode", getCode).Methods("POST") // get Code call
	r.HandleFunc("/getToken", getToken).Methods("POST") // get Code call
	r.HandleFunc("/register", register).Methods("POST")


	// r.Use(loggingMiddleware)

	srv := &http.Server{Addr: ":7000", Handler: r}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("HTTP Server started")
	}

	// 	go func() {

	// 	}()

	// 	var wait time.Duration

	// 	c := make(chan os.Signal, 1)
	// 	signal.Notify(c, os.Interrupt)
	// 	<-c
	// 	ctx, cancel := context.WithTimeout(context.Background(), wait)
	// 	defer cancel()
	// 	srv.Shutdown(ctx)
	// 	log.Println("shutting down")
	// 	os.Exit(0)
}
