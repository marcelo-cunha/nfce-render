package nfce

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Options contém as opções para geração do DANFE
type Options struct {
	Format string // Formato de saída: "html" ou "pdf" (padrão: "html")
}

// GenerateDANFE gera um DANFE a partir do XML da NF-e
func GenerateDANFE(xmlContent []byte, options *Options) ([]byte, error) {
	// Criar gerador
	generator, err := NewGenerator(xmlContent)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar gerador: %w", err)
	}

	// Validar se é NFC-e
	if !generator.IsNFCe() {
		return nil, fmt.Errorf("apenas NFC-e (modelo 65) é suportada atualmente")
	}

	// Configurar opções padrão
	if options == nil {
		options = &Options{
			Format: "html",
		}
	}
	if options.Format == "" {
		options.Format = "html"
	}

	// Converter para formato do novo pacote
	var format Format
	switch options.Format {
	case "html":
		format = FormatHTML
	case "pdf":
		format = FormatPDF
	default:
		return nil, fmt.Errorf("formato não suportado: %s. Use 'html' ou 'pdf'", options.Format)
	}

	// Gerar usando o novo gerador
	var buf bytes.Buffer
	if err := generator.GenerateToWriter(&buf, GenerateOptions{Format: format}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GenerateDANFEFromFile gera um DANFE a partir de um arquivo XML
func GenerateDANFEFromFile(xmlFilePath string, options *Options) ([]byte, error) {
	// Ler arquivo XML
	xmlContent, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo XML: %w", err)
	}

	return GenerateDANFE(xmlContent, options)
}

// SaveDANFE gera e salva um DANFE em arquivo
func SaveDANFE(xmlContent []byte, outputPath string, options *Options) error {
	// Gerar DANFE
	content, err := GenerateDANFE(xmlContent, options)
	if err != nil {
		return err
	}

	// Salvar arquivo
	if err := os.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	return nil
}

// SaveDANFEFromFile gera e salva um DANFE a partir de um arquivo XML
func SaveDANFEFromFile(xmlFilePath, outputPath string, options *Options) error {
	// Ler arquivo XML
	xmlContent, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo XML: %w", err)
	}

	return SaveDANFE(xmlContent, outputPath, options)
}

// WriteDANFE gera um DANFE e escreve para um io.Writer
func WriteDANFE(xmlContent []byte, writer io.Writer, options *Options) error {
	// Criar gerador
	generator, err := NewGenerator(xmlContent)
	if err != nil {
		return fmt.Errorf("erro ao criar gerador: %w", err)
	}

	// Validar se é NFC-e
	if !generator.IsNFCe() {
		return fmt.Errorf("apenas NFC-e (modelo 65) é suportada atualmente")
	}

	// Configurar opções padrão
	if options == nil {
		options = &Options{
			Format: "html",
		}
	}
	if options.Format == "" {
		options.Format = "html"
	}

	// Converter para formato do novo pacote
	var format Format
	switch options.Format {
	case "html":
		format = FormatHTML
	case "pdf":
		format = FormatPDF
	default:
		return fmt.Errorf("formato não suportado: %s. Use 'html' ou 'pdf'", options.Format)
	}

	return generator.GenerateToWriter(writer, GenerateOptions{Format: format})
}