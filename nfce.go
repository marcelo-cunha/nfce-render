package nfce

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/marcelo-cunha/nfce-render/converter"
	"github.com/marcelo-cunha/nfce-render/renderer"
	"github.com/marcelo-cunha/nfce-render/xmlparser"
)

// Format representa os formatos de saída suportados
type Format string

const (
	FormatHTML Format = "html"
	FormatPDF  Format = "pdf"
)

// GenerateOptions contém as opções para geração do DANFE
type GenerateOptions struct {
	Format Format
}

// Generator é responsável pela geração de DANFEs
type Generator struct {
	nfe *xmlparser.NFeProc
}

// NewGenerator cria uma nova instância do gerador
func NewGenerator(xmlContent []byte) (*Generator, error) {
	nfe, err := xmlparser.ParseXML(xmlContent)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do XML: %w", err)
	}

	return &Generator{
		nfe: nfe,
	}, nil
}

// NewGeneratorFromFile cria uma nova instância do gerador a partir de um arquivo
func NewGeneratorFromFile(xmlPath string) (*Generator, error) {
	xmlContent, err := os.ReadFile(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo XML: %w", err)
	}

	return NewGenerator(xmlContent)
}

// GenerateToWriter gera o DANFE e escreve no writer fornecido
func (g *Generator) GenerateToWriter(writer io.Writer, options GenerateOptions) error {
	switch options.Format {
	case FormatHTML:
		return g.generateHTML(writer)
	case FormatPDF:
		return g.generatePDF(writer)
	default:
		return fmt.Errorf("formato não suportado: %s", options.Format)
	}
}

// GenerateToFile gera o DANFE e salva em um arquivo
func (g *Generator) GenerateToFile(outputPath string, options GenerateOptions) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %w", err)
	}
	defer file.Close()

	return g.GenerateToWriter(file, options)
}

// generateHTML gera o DANFE em formato HTML
func (g *Generator) generateHTML(writer io.Writer) error {
	htmlRenderer := renderer.NewHTMLRenderer(g.nfe)
	return htmlRenderer.RenderToWriter(writer)
}

// generatePDF gera o DANFE em formato PDF
func (g *Generator) generatePDF(writer io.Writer) error {
	// Primeiro gerar HTML em memória
	htmlRenderer := renderer.NewHTMLRenderer(g.nfe)
	
	// Renderizar HTML para buffer
	var htmlBuffer bytes.Buffer
	if err := htmlRenderer.RenderToWriter(&htmlBuffer); err != nil {
		return fmt.Errorf("erro ao renderizar HTML: %w", err)
	}
	
	// Usar o conversor PDF
	pdfConverter := converter.NewPDFConverter()
	pdfBytes, err := pdfConverter.ConvertHTMLToPDF(htmlBuffer.Bytes())
	if err != nil {
		return fmt.Errorf("erro ao converter para PDF: %w", err)
	}
	
	// Escrever PDF no writer
	_, err = writer.Write(pdfBytes)
	return err
}

// GetNFe retorna a estrutura NFeProc parseada
func (g *Generator) GetNFe() *xmlparser.NFeProc {
	return g.nfe
}

// IsNFCe verifica se é uma NFC-e
func (g *Generator) IsNFCe() bool {
	return g.nfe.IsNFCe()
}

// Version da biblioteca
const Version = "1.0.0"

// GetVersion retorna a versão da biblioteca
func GetVersion() string {
	return Version
}