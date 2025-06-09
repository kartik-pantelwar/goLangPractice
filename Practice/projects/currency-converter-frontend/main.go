package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type CurrencyResponse struct {
    Currencies map[string]string `json:"currencies"`
}

func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getCurrencies(w http.ResponseWriter, r *http.Request) {
    enableCORS(w)
    
    if r.Method == "OPTIONS" {
        return
    }

    // Fetch all available currencies from the API
    resp, err := http.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json")
    if err != nil {
        http.Error(w, "Failed to fetch currencies", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Failed to read response", http.StatusInternalServerError)
        return
    }

    // Parse the response
    var currencies map[string]string
    if err := json.Unmarshal(body, &currencies); err != nil {
        http.Error(w, "Failed to parse currencies", http.StatusInternalServerError)
        return
    }

    // Return as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(CurrencyResponse{Currencies: currencies})
}

func main() {
    http.HandleFunc("/api/currencies", getCurrencies)
    
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}