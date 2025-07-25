package converter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// PDFConverter é responsável pela conversão de HTML para PDF usando Gotenberg
type PDFConverter struct {
	gotenbergURL string
	client       *http.Client
}

// NewPDFConverter cria uma nova instância do conversor PDF
func NewPDFConverter() *PDFConverter {
	// Tentar carregar .env se existir
	_ = godotenv.Load()

	// Obter URL do Gotenberg da variável de ambiente ou usar padrão
	gotenbergURL := os.Getenv("GOTENBERG_URL")
	if gotenbergURL == "" {
		gotenbergURL = "https://demo.gotenberg.dev"
	}

	return &PDFConverter{
		gotenbergURL: gotenbergURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ConvertHTMLToPDF converte conteúdo HTML para PDF usando Gotenberg
func (c *PDFConverter) ConvertHTMLToPDF(htmlContent []byte) ([]byte, error) {
	// Preparar multipart/form-data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Adicionar arquivo index.html
	htmlPart, err := writer.CreateFormFile("files", "index.html")
	if err != nil {
		return nil, fmt.Errorf("erro ao criar form file: %w", err)
	}
	if _, err = htmlPart.Write(htmlContent); err != nil {
		return nil, fmt.Errorf("erro ao escrever HTML: %w", err)
	}

	// Configurar dimensões da página (80mm = 3.15 polegadas)
	if err = writer.WriteField("paperWidth", "3.15in"); err != nil {
		return nil, fmt.Errorf("erro ao definir paperWidth: %w", err)
	}

	// Usar singlePage=true para altura dinâmica baseada no conteúdo
	if err = writer.WriteField("singlePage", "true"); err != nil {
		return nil, fmt.Errorf("erro ao definir singlePage: %w", err)
	}

	// Configurar margens zeradas
	if err = writer.WriteField("marginTop", "0"); err != nil {
		return nil, fmt.Errorf("erro ao definir marginTop: %w", err)
	}
	if err = writer.WriteField("marginBottom", "0"); err != nil {
		return nil, fmt.Errorf("erro ao definir marginBottom: %w", err)
	}
	if err = writer.WriteField("marginLeft", "0"); err != nil {
		return nil, fmt.Errorf("erro ao definir marginLeft: %w", err)
	}
	if err = writer.WriteField("marginRight", "0"); err != nil {
		return nil, fmt.Errorf("erro ao definir marginRight: %w", err)
	}

	// Ativar printBackground para renderizar cores e imagens de fundo
	if err = writer.WriteField("printBackground", "true"); err != nil {
		return nil, fmt.Errorf("erro ao definir printBackground: %w", err)
	}

	// Configurar escala
	if err = writer.WriteField("scale", "1.0"); err != nil {
		return nil, fmt.Errorf("erro ao definir scale: %w", err)
	}

	// Fechar o writer
	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("erro ao fechar multipart writer: %w", err)
	}

	// Criar requisição HTTP
	var req *http.Request
	req, err = http.NewRequest("POST", c.gotenbergURL+"/forms/chromium/convert/html", &requestBody)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	// Definir Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Executar requisição
	var resp *http.Response
	resp, err = c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição: %w", err)
	}
	defer resp.Body.Close()

	// Verificar status da resposta
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("erro na API do Gotenberg (status %d): %s", resp.StatusCode, string(body))
	}

	// Ler o PDF gerado
	var pdfBytes []byte
	pdfBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler PDF: %w", err)
	}

	if len(pdfBytes) == 0 {
		return nil, fmt.Errorf("PDF gerado está vazio")
	}

	return pdfBytes, nil
}