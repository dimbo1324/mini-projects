package pdf

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"pdf-generator/src/repository"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Generator struct {
	templatePath string
}

func NewGenerator(templatePath string) *Generator {
	return &Generator{templatePath: templatePath}
}

func (g *Generator) Generate(order repository.Order) ([]byte, error) {
	tmpl, err := template.ParseFiles(g.templatePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга шаблона: %w", err)
	}
	var htmlBody bytes.Buffer
	if err := tmpl.Execute(&htmlBody, order); err != nil {
		return nil, fmt.Errorf("ошибка выполнения шаблона: %w", err)
	}
	tmpFile, err := os.CreateTemp("", "order-*.html")
	if err != nil {
		return nil, fmt.Errorf("не удалось создать temp файл: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.Write(htmlBody.Bytes()); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации конвертера: %w", err)
	}
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)
	page := wkhtmltopdf.NewPage(tmpFile.Name())
	page.Encoding.Set("utf-8")
	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		return nil, fmt.Errorf("ошибка создания PDF: %w", err)
	}
	return pdfg.Bytes(), nil
}
