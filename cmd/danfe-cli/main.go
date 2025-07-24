package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcelo-cunha/nfce-render/internal/danfe"
)

func main() {
	// Definir flags
	var (
		inputFile  = flag.String("input", "", "Caminho para o arquivo XML da NF-e (obrigatório)")
		outputFile = flag.String("output", "", "Caminho para o arquivo de saída (obrigatório)")
		format     = flag.String("format", "html", "Formato de saída: html ou pdf")
		version    = flag.Bool("version", false, "Exibir versão")
		help       = flag.Bool("help", false, "Exibir ajuda")
	)

	// Parse das flags
	flag.Parse()

	// Exibir versão
	if *version {
		fmt.Printf("danfe-cli v%s\n", danfe.GetVersion())
		os.Exit(0)
	}

	// Exibir ajuda
	if *help {
		printHelp()
		os.Exit(0)
	}

	// Validar argumentos obrigatórios
	if *inputFile == "" {
		fmt.Fprintf(os.Stderr, "Erro: arquivo de entrada é obrigatório\n\n")
		printUsage()
		os.Exit(1)
	}

	if *outputFile == "" {
		fmt.Fprintf(os.Stderr, "Erro: arquivo de saída é obrigatório\n\n")
		printUsage()
		os.Exit(1)
	}

	// Validar se o arquivo de entrada existe
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Erro: arquivo de entrada não encontrado: %s\n", *inputFile)
		os.Exit(1)
	}

	// Validar formato
	*format = strings.ToLower(*format)
	if *format != "html" && *format != "pdf" {
		fmt.Fprintf(os.Stderr, "Erro: formato não suportado: %s. Use 'html' ou 'pdf'\n", *format)
		os.Exit(1)
	}

	// Criar diretório de saída se não existir
	outputDir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar diretório de saída: %v\n", err)
		os.Exit(1)
	}

	// Configurar opções
	options := &danfe.GenerateOptions{
		Format: *format,
	}

	// Gerar DANFE
	fmt.Printf("Gerando DANFE...\n")
	fmt.Printf("  Entrada: %s\n", *inputFile)
	fmt.Printf("  Saída: %s\n", *outputFile)
	fmt.Printf("  Formato: %s\n", *format)

	if err := danfe.SaveDANFEFromFile(*inputFile, *outputFile, options); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao gerar DANFE: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("DANFE gerado com sucesso!\n")
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Uso: %s --input <arquivo.xml> --output <arquivo.html> [opções]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Opções:\n")
	flag.PrintDefaults()
}

func printHelp() {
	fmt.Printf("DANFE CLI - Gerador de DANFE NFC-e v%s\n\n", danfe.GetVersion())
	fmt.Printf("Descrição:\n")
	fmt.Printf("  Gera DANFE (Documento Auxiliar da Nota Fiscal Eletrônica) para NFC-e\n")
	fmt.Printf("  a partir de um arquivo XML da Nota Fiscal Eletrônica.\n\n")

	printUsage()

	fmt.Printf("\nExemplos:\n")
	fmt.Printf("  # Gerar DANFE em HTML\n")
	fmt.Printf("  %s --input ./examples/example.xml --output ./out/nota.html\n\n", os.Args[0])



	fmt.Printf("  # Gerar DANFE em PDF\n")
	fmt.Printf("  %s --input ./examples/example.xml --output ./out/nota.pdf --format pdf\n\n", os.Args[0])

	fmt.Printf("Formatos suportados:\n")
	fmt.Printf("  html - Gera DANFE em formato HTML (padrão)\n")
	fmt.Printf("  pdf  - Gera DANFE em formato PDF\n\n")

	fmt.Printf("Notas:\n")
	fmt.Printf("  - Apenas NFC-e (modelo 65) é suportada atualmente\n")

	fmt.Printf("  - O diretório de saída será criado automaticamente se não existir\n")
}
