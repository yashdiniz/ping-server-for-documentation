package main

import (
    "fmt"
    "log"
    "net/http"
    "errors"
    "github.com/joho/godotenv"
    "os"
)

func Load() error {
    if _, err := os.Stat("./.env"); errors.Is(err, os.ErrNotExist) {
        return godotenv.Load(".env")    // read through environment variables.
    }
    return nil
}

func main() {

    if err := Load(); err != nil {
        log.Fatal("Error while loading environment variables! ", err)
    }

    mux := http.NewServeMux()    // a new Server mux
    mux.HandleFunc("/a/health", func (w http.ResponseWriter, r *http.Request) {    // handling /health
        switch r.Method {
        case "GET":    // on GET /health
            w.WriteHeader(http.StatusOK)
            fmt.Fprint(w, "Server is healthy!")    // reply with "Server is healthy!"
        default:
            http.NotFound(w, r)    // on any other request method, throw a 404
        }
    })
    mux.HandleFunc("/a/svcName", func (w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":    // on GET /svcName
            w.WriteHeader(http.StatusOK)
            fmt.Fprint(w, os.Getenv("SERVICE_NAME"))    // take service name from the environment.
        default:
            http.NotFound(w, r)    // throw a 404
        }
    })

    server := http.Server{
        Addr: ":8080",    // server listens at :8080
        Handler: mux,
    }

    log.Print("Listening at ", server.Addr)
    log.Fatal(server.ListenAndServe())
}
