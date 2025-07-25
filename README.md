# NFCE Render

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.21)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

Biblioteca Go para geração de DANFE (Documento Auxiliar da Nota Fiscal Eletrônica) para NFC-e (Nota Fiscal de Consumidor Eletrônica).

## Características

- Suporte para NFC-e (modelo 65)
- Geração em formato HTML e PDF
- API simples e intuitiva
- Módulo Go reutilizável

## Instalação

```bash
go get github.com/marcelo-cunha/nfce-render
```

## Uso Básico

```go
package main

import (
    "fmt"
    "os"
    nfce "github.com/marcelo-cunha/nfce-render"
)

func main() {
    // Ler XML da NFC-e
    xmlContent, err := os.ReadFile("nfce.xml")
    if err != nil {
        panic(err)
    }

    // Gerar DANFE em HTML
    options := &nfce.Options{Format: "html"}
    danfeBytes, err := nfce.GenerateDANFE(xmlContent, options)
    if err != nil {
        panic(err)
    }

    // Salvar arquivo
    err = os.WriteFile("danfe.html", danfeBytes, 0644)
    if err != nil {
        panic(err)
    }

    fmt.Println("DANFE gerado com sucesso!")
}
```

## API

### Funções Principais

```go
// Gerar DANFE a partir de conteúdo XML
danfeBytes, err := nfce.GenerateDANFE(xmlContent, options)

// Gerar DANFE a partir de arquivo XML
danfeBytes, err := nfce.GenerateDANFEFromFile("nfce.xml", options)

// Salvar DANFE diretamente em arquivo
err := nfce.SaveDANFE(xmlContent, "danfe.html", options)

// Salvar de arquivo XML para arquivo de saída
err := nfce.SaveDANFEFromFile("nfce.xml", "danfe.html", options)

// Escrever para io.Writer
err := nfce.WriteDANFE(xmlContent, writer, options)
```

### Opções

```go
type Options struct {
    Format string // "html" ou "pdf"
}
```

### API Avançada

```go
// Criar gerador
generator, err := nfce.NewGenerator(xmlContent)

// Verificar se é NFC-e
if generator.IsNFCe() {
    // Gerar para writer específico
    options := nfce.GenerateOptions{Format: nfce.FormatHTML}
    err = generator.GenerateToWriter(writer, options)
}
```

## Formatos Suportados

- **HTML**: Formato padrão, ideal para visualização web
- **PDF**: Requer serviço Gotenberg para conversão

## Configuração PDF

Para geração de PDF, configure a variável de ambiente:

```bash
export GOTENBERG_URL="http://localhost:3000"
```

Ou use o arquivo `.env`:

```
GOTENBERG_URL=http://localhost:3000
```

## Dependências

- Go 1.21+
- github.com/skip2/go-qrcode (geração de QR Code)
- github.com/joho/godotenv (variáveis de ambiente)
- Gotenberg (apenas para PDF)

## Limitações

- Suporta apenas NFC-e (modelo 65)
- PDF requer serviço Gotenberg externo

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Versão

Versão estável: **v1.0.0**