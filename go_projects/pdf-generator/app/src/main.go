package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"pdf-generator/src/pdf"
	"pdf-generator/src/repository"
	"strconv"
)

func main() {
	repo := repository.NewRepository()
	tmplPath := filepath.Join("templates", "report.html")
	pdfGen := pdf.NewGenerator(tmplPath)
	mux := http.NewServeMux()
	mux.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {

		var id int
		var action string

		_, err := fmt.Sscanf(r.URL.Path, "/order/%d/%s", &id, &action)

		if err != nil || action != "report" {
			http.NotFound(w, r)
			return
		}

		order, found := repo.GetOrderByID(id)
		if !found {
			http.Error(w, "Заказ не найден", http.StatusNotFound)
			return
		}

		pdfBytes, err := pdfGen.Generate(order)
		if err != nil {
			log.Printf("Error generating PDF: %v", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=order_%d.pdf", id))
		w.Header().Set("Content-Length", strconv.Itoa(len(pdfBytes)))

		if _, err := w.Write(pdfBytes); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
