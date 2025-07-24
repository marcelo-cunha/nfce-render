# nfce-render

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.21)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

Uma biblioteca Go para geração de DANFE (Documento Auxiliar da Nota Fiscal Eletrônica) para NFC-e (Nota Fiscal de Consumidor Eletrônica - modelo 65) em formatos HTML e PDF, utilizando o serviço Gotenberg para conversão de HTML para PDF.

## Características

- Geração de DANFE para NFC-e (modelo 65) a partir de XML
- Suporte para saída em HTML e PDF
- Integração com Gotenberg para conversão HTML para PDF
- Configuração flexível da URL do Gotenberg via variável de ambiente
- Configuração de largura fixa (80mm) e altura dinâmica para o PDF
- Geração de QR Code para consulta pública
- Arquitetura modular com separação de responsabilidades

## Requisitos

- Go 1.21 ou superior
- Dependências:
  - github.com/skip2/go-qrcode (para geração de QR Code)
  - github.com/joho/godotenv (para carregar variáveis de ambiente)
- Para geração de PDF: acesso a um servidor Gotenberg (por padrão usa https://demo.gotenberg.dev)

## Instalação

```bash
go get github.com/marcelo-cunha/nfce-render
```

## Exemplo de Uso

```go
package main

import (
	"fmt"
	"os"

	"github.com/marcelo-cunha/nfce-render/internal/danfe"
)

func main() {
	// Ler o arquivo XML da NFC-e
	xmlFilePath := "caminho/para/nfce.xml"
	
	// Configurar opções
	options := &danfe.GenerateOptions{
		Format: "pdf", // "html" ou "pdf"
	}
	
	// Gerar e salvar o DANFE
	err := danfe.SaveDANFEFromFile(xmlFilePath, "saida.pdf", options)
	if err != nil {
		fmt.Printf("Erro ao gerar DANFE: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("DANFE gerado com sucesso!")
}
```

## API

### Funções Principais

#### `GenerateDANFE(xmlContent []byte, options *GenerateOptions) ([]byte, error)`

Gera um DANFE a partir do conteúdo XML da NF-e.

- **Parâmetros**:
  - `xmlContent []byte`: Conteúdo do XML da NF-e
  - `options *GenerateOptions`: Opções de geração (pode ser nil para usar valores padrão)
- **Retorno**:
  - `[]byte`: Conteúdo do DANFE gerado (HTML ou PDF)
  - `error`: Erro, se houver

#### `GenerateDANFEFromFile(xmlFilePath string, options *GenerateOptions) ([]byte, error)`

Gera um DANFE a partir de um arquivo XML.

- **Parâmetros**:
  - `xmlFilePath string`: Caminho para o arquivo XML da NF-e
  - `options *GenerateOptions`: Opções de geração (pode ser nil para usar valores padrão)
- **Retorno**:
  - `[]byte`: Conteúdo do DANFE gerado (HTML ou PDF)
  - `error`: Erro, se houver

#### `SaveDANFE(xmlContent []byte, outputPath string, options *GenerateOptions) error`

Gera e salva um DANFE em arquivo.

- **Parâmetros**:
  - `xmlContent []byte`: Conteúdo do XML da NF-e
  - `outputPath string`: Caminho para o arquivo de saída
  - `options *GenerateOptions`: Opções de geração (pode ser nil para usar valores padrão)
- **Retorno**:
  - `error`: Erro, se houver

#### `SaveDANFEFromFile(xmlFilePath, outputPath string, options *GenerateOptions) error`

Gera e salva um DANFE a partir de um arquivo XML.

- **Parâmetros**:
  - `xmlFilePath string`: Caminho para o arquivo XML da NF-e
  - `outputPath string`: Caminho para o arquivo de saída
  - `options *GenerateOptions`: Opções de geração (pode ser nil para usar valores padrão)
- **Retorno**:
  - `error`: Erro, se houver

#### `WriteDANFE(xmlContent []byte, writer io.Writer, options *GenerateOptions) error`

Gera um DANFE e escreve para um io.Writer.

- **Parâmetros**:
  - `xmlContent []byte`: Conteúdo do XML da NF-e
  - `writer io.Writer`: Writer para escrever o conteúdo gerado
  - `options *GenerateOptions`: Opções de geração (pode ser nil para usar valores padrão)
- **Retorno**:
  - `error`: Erro, se houver

### Estruturas

#### `GenerateOptions`

```go
type GenerateOptions struct {
	Format string // Formato de saída: "html" ou "pdf" (padrão: "html")
}
```

## Configuração

### Variáveis de Ambiente

O projeto suporta configuração via arquivo `.env`. Copie o arquivo `.env.example` para `.env` e configure as variáveis:

```bash
cp .env.example .env
```

#### Variáveis Disponíveis

- `GOTENBERG_URL`: URL do servidor Gotenberg (padrão: https://demo.gotenberg.dev)

### Configurações do PDF

A geração de PDF utiliza o serviço Gotenberg com as seguintes configurações padrão:

- **URL do Gotenberg**: Configurável via `GOTENBERG_URL` (padrão: https://demo.gotenberg.dev/forms/chromium/convert/html)
- **Largura do papel**: 80mm (3.15 polegadas)
- **Altura do papel**: Dinâmica, baseada no conteúdo (usando `singlePage=true`)
- **Margens**: 0 (zero) em todos os lados
- **Escala**: 1.0
- **Impressão de fundo**: Ativada (`printBackground=true`)

## Limitações

- Atualmente, apenas NFC-e (modelo 65) é suportada
- Para uso em produção, recomenda-se configurar um servidor Gotenberg próprio em vez de usar o demo

## Uso via CLI

O projeto inclui uma ferramenta de linha de comando para geração de DANFE:

```bash
# Gerar DANFE em HTML
danfe-cli --input ./examples/example.xml --output ./out/nota.html

# Gerar DANFE em PDF
danfe-cli --input ./examples/example.xml --output ./out/nota.pdf --format pdf
```

### Opções disponíveis

- `--input`: Caminho para o arquivo XML da NF-e (obrigatório)
- `--output`: Caminho para o arquivo de saída (obrigatório)
- `--format`: Formato de saída: html ou pdf (padrão: html)
- `--version`: Exibir versão
- `--help`: Exibir ajuda

## Licença

Este projeto está licenciado sob a licença MIT.