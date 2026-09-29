package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Struct untuk memetakan struktur JSON
type Profile struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Bio  string `json:"bio"`
}

type Skill struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Level    string `json:"level"`
}

type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	GithubURL   string   `json:"github_url"`
	DemoURL     string   `json:"demo_url"`
}

type PortfolioData struct {
	Profile  Profile   `json:"profile"`
	Skills   []Skill   `json:"skills"`
	Projects []Project `json:"projects"`
}

// Middleware untuk mengizinkan CORS (Cross-Origin Resource Sharing)
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mengizinkan semua origin untuk mengakses API ini
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Handler untuk membaca file JSON dan mengembalikan respons
func getPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	// Baca file data.json
	fileData, err := os.ReadFile("data/data.json")
	if err != nil {
		http.Error(w, "Gagal membaca data portfolio", http.StatusInternalServerError)
		log.Println("Error reading json file:", err)
		return
	}

	// Set header HTTP menjadi application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileData)
}

func main() {
	// Mendaftarkan route API
	http.HandleFunc("/api/portfolio", enableCORS(getPortfolioHandler))

	port := ":8080"
	fmt.Printf("Server Golang berjalan di http://localhost%s\n", port)

	// Jalankan server HTTP
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
