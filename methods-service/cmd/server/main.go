package main

import (
    "os"
    "context"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "methods-service/internal/client"
    "methods-service/internal/service"

    "github.com/gorilla/mux"
)

func main() {
    // Create gRPC client to user-service
    userServiceHost := os.Getenv("USER_SERVICE_HOST")
    if userServiceHost == "" {
        userServiceHost = "user-service"
    }
    userServicePort := os.Getenv("USER_SERVICE_PORT") 
    if userServicePort == "" {
        userServicePort = "50051"
    }
    userClient, err := client.NewUserClient(userServiceHost + ":" + userServicePort)
    if err != nil {
        log.Printf("Warning: Failed to create user client: %v", err)
        log.Printf("Starting with mock data mode")
        userClient = nil
    }
    defer func() {
        if userClient != nil {
            userClient.Close()
        }
    }()

    methodsService := service.NewMethodsService(userClient)

    r := mux.NewRouter()
    r.HandleFunc("/method1", func(w http.ResponseWriter, r *http.Request) {
        handleMethod(methodsService.ProcessMethod1, w, r)
    }).Methods("GET")
    r.HandleFunc("/method2", func(w http.ResponseWriter, r *http.Request) {
        handleMethod(methodsService.ProcessMethod2, w, r)
    }).Methods("GET")

    log.Println("Methods service starting on :8080")
    serverPort := os.Getenv("SERVER_PORT")
    if serverPort == "" {
        serverPort = ":8080"
    }
    log.Println("Methods service starting on " + serverPort)
    log.Fatal(http.ListenAndServe(serverPort, r))
}

func handleMethod(methodFunc func(ctx context.Context, waitTime int) ([]string, error), w http.ResponseWriter, r *http.Request) {
    waitTimeStr := r.URL.Query().Get("wait")
    waitTime := 5 // default
    if waitTimeStr != "" {
        if parsed, err := strconv.Atoi(waitTimeStr); err == nil {
            waitTime = parsed
        }
    }

    names, err := methodFunc(r.Context(), waitTime)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "user_names": names,
        "wait_time":  waitTime,
    })
}

