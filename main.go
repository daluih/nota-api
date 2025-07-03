package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// itens da nota fiscal
type Item struct {
	Codigo     string  `json:"codigo"`
	Descricao  string  `json:"descricao"`
	Quantidade int     `json:"quantidade"`
	ValorUnit  float64 `json:"valor_unitario"`
}

// tipo
type Nota struct {
	Numero string `json:"numero_nota"`
	Itens  []Item `json:"itens"`
}

// que vai guardar as notas
var notas []Nota

func main() {
	// lê o arquivo JSON (substituição do ioutil.ReadFile)
	data, err := os.ReadFile("itens.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}
	err = json.Unmarshal(data, &notas)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	// Handler para raiz /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API rodando. Use /health para status ou /notas/{numero}/itens para consultar notas."))
	})

	// Handler para /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Redireciona /notas (sem barra) para /notas/ (com barra)
	http.HandleFunc("/notas", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/notas/", http.StatusMovedPermanently)
	})

	// Handler para /notas/ com a lógica principal
	http.HandleFunc("/notas/", func(w http.ResponseWriter, r *http.Request) {
		// só aceita método GET
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Método não permitido"))
			return
		}

		// pega número da nota da URL
		url := strings.TrimPrefix(r.URL.Path, "/notas/")
		partes := strings.Split(url, "/")
		if len(partes) != 2 || partes[1] != "itens" {
			w.WriteHeader(404)
			w.Write([]byte("Página não encontrada"))
			return
		}
		numero := partes[0]

		// procura nota
		for _, n := range notas {
			if n.Numero == numero {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(n)
				return
			}
		}

		// se não achou a nota
		w.WriteHeader(404)
		w.Write([]byte("Nota não encontrada"))
	})

	fmt.Println("API rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
