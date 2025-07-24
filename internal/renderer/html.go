package renderer

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"

	"github.com/marcelo-cunha/nfce-render/internal/xmlparser"
	"github.com/skip2/go-qrcode"
)

// HTMLRenderer é responsável pela renderização do DANFE em HTML
type HTMLRenderer struct {
	nfe *xmlparser.NFeProc
}

// NewHTMLRenderer cria uma nova instância do renderizador HTML
func NewHTMLRenderer(nfe *xmlparser.NFeProc) *HTMLRenderer {
	return &HTMLRenderer{
		nfe: nfe,
	}
}

// RenderToWriter renderiza o DANFE em HTML para um io.Writer
func (r *HTMLRenderer) RenderToWriter(writer io.Writer) error {
	// Criar template com funções auxiliares
	tmpl := template.New("danfe").Funcs(template.FuncMap{
		"formatCNPJ":     xmlparser.FormatCNPJ,
		"formatCPF":      xmlparser.FormatCPF,
		"formatCEP":      xmlparser.FormatCEP,
		"formatCurrency": xmlparser.FormatCurrency,
		"formatQuantity": xmlparser.FormatQuantity,
		"formatDate": func(t time.Time) string {
			return t.Format("02/01/2006 15:04:05")
		},
		"formatDateOnly": func(t time.Time) string {
			return t.Format("02/01/2006")
		},
		"getPaymentMethod": xmlparser.GetPaymentMethodDescription,
		"generateQRCode":   r.generateQRCodeHTML,
		"upper":            strings.ToUpper,
		"add": func(a, b int) int {
			return a + b
		},
		"formatKey": func(key string) string {
			if len(key) == 0 {
				return key
			}
			var result strings.Builder
			for i, char := range key {
				if i > 0 && i%4 == 0 {
					result.WriteString(" ")
				}
				result.WriteRune(char)
			}
			return result.String()
		},
	})

	// Parse do template
	var err error
	tmpl, err = tmpl.Parse(danfeTemplate)
	if err != nil {
		return fmt.Errorf("erro ao fazer parse do template: %w", err)
	}

	// Executar template
	data := struct {
		NFe *xmlparser.NFeProc
	}{
		NFe: r.nfe,
	}

	if err := tmpl.Execute(writer, data); err != nil {
		return fmt.Errorf("erro ao executar template: %w", err)
	}

	return nil
}

// generateQRCodeHTML gera um QR Code em formato HTML
func (r *HTMLRenderer) generateQRCodeHTML(content string) template.HTML {
	if content == "" {
		return template.HTML("")
	}

	// Gerar QR Code como PNG
	pngBytes, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return template.HTML("")
	}

	// Converter para base64 para embedding no HTML
	base64String := base64.StdEncoding.EncodeToString(pngBytes)
	return template.HTML(fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="QR Code" style="width: 25mm; height: 25mm;">`, base64String))
}

// Template HTML do DANFE
const danfeTemplate = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DANFE NFC-e</title>
    <style>
        @page {
            size: 80mm auto;
            margin: 0mm;
        }
        
        @media print {
            body { 
                margin: 0; 
                padding: 0;
                width: 80mm;
            }
            .danfe { 
                box-shadow: none; 
                border: none; 
                width: 80mm;
                max-width: 80mm;
            }
        }
        
        body {
            font-family: 'Arial', sans-serif;
            font-size: 8px;
            margin: 0;
            padding: 0;
            background-color: white;
            line-height: 1.2;
            width: 80mm;
            max-width: 80mm;
            min-height: auto;
        }
        
        .danfe {
            width: 80mm;
            max-width: 80mm;
            margin: 0;
            background-color: white;
            padding: 2mm;
            box-sizing: border-box;
            min-height: auto;
        }
        
        .header {
            text-align: center;
            margin-bottom: 4px;
        }
        
        .company-name {
            font-weight: bold;
            font-size: 14px;
            margin-bottom: 1px;
        }
        
        .cnpj {
            font-weight: bold;
            font-size: 12px;
            margin-bottom: 1px;
        }
        
        .address {
            font-size: 11px;
            margin-bottom: 1px;
            line-height: 1.0;
        }
        
        .document-title {
            font-weight: bold;
            font-size: 13px;
            margin: 4px 0 2px 0;
        }
        
        .document-subtitle {
            font-size: 10px;
            margin-bottom: 4px;
        }
        
        .section-title {
            font-weight: bold;
            font-size: 11px;
            text-align: center;
            margin: 4px 0 2px 0;
            border-top: 1px dashed #333;
            border-bottom: 1px dashed #333;
            padding: 1px 0;
        }
        
        .item {
            margin-bottom: 3px;
            font-size: 10px;
        }
        
        .item-line {
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-weight: bold;
            margin-bottom: 1px;
        }
        
        .item-code {
            font-weight: bold;
        }
        
        .item-desc {
            margin-bottom: 1px;
        }
        
        .item-values {
            text-align: right;
            font-weight: bold;
        }
        
        .totals {
            margin-bottom: 4px;
            font-size: 10px;
        }
        
        .total-line {
            display: flex;
            justify-content: space-between;
            margin-bottom: 0px;
        }
        
        .total-final {
            font-weight: bold;
            font-size: 11px;
            border-top: 1px solid #333;
            padding-top: 1px;
        }
        
        .payment {
            margin-bottom: 4px;
            font-size: 10px;
        }
        
        .payment-line {
            display: flex;
            justify-content: space-between;
            margin-bottom: 0px;
        }
        
        .consumer {
            margin-bottom: 4px;
            font-size: 10px;
        }
        
        .nfc-info {
            margin-bottom: 4px;
            font-size: 10px;
        }
        
        .footer {
            text-align: center;
            font-size: 10px;
            margin-bottom: 4px;
        }
        
        .key {
            word-break: break-all;
            font-size: 9px;
            margin: 2px 0;
            text-align: center;
        }
        
        .qr-code {
            text-align: center;
            padding: 2px;
            margin-top: 4px;
        }
        
        .qr-code img {
            width: 20mm;
            height: 20mm;
            border: 1px solid #ddd;
        }
        
        .qr-text {
            font-size: 9px;
            margin-top: 1px;
            color: #666;
        }
    </style>
</head>
<body>
    <div class="danfe">
        
        <div class="header">
            
            <div style="margin-bottom: 8px;">
                <!-- Cabeçalho sem logo -->
            </div>
            
            
            <div class="company-name">{{.NFe.NFe.InfNFe.Emit.XNome}}</div>
            <div class="cnpj">CNPJ: {{formatCNPJ .NFe.NFe.InfNFe.Emit.CNPJ}}</div>
            <div class="address">
                {{.NFe.NFe.InfNFe.Emit.EnderEmit.XLgr}}, {{.NFe.NFe.InfNFe.Emit.EnderEmit.Nro}}<br>
                {{.NFe.NFe.InfNFe.Emit.EnderEmit.XBairro}}, {{.NFe.NFe.InfNFe.Emit.EnderEmit.XMun}}-{{.NFe.NFe.InfNFe.Emit.EnderEmit.UF}}<br>
                CEP: {{formatCEP .NFe.NFe.InfNFe.Emit.EnderEmit.CEP}}
            </div>
            <div class="document-title">DANFE NFC-e</div>
            <div class="document-subtitle">Documento Auxiliar da Nota Fiscal de Consumidor Eletrônica</div>
        </div>

        
        <div class="section-title">ITENS</div>
        
        {{range $index, $item := .NFe.NFe.InfNFe.Det}}
        <div class="item">
            <div class="item-line">
                <span class="item-code">{{printf "%02d" (add $index 1)}} - {{$item.Prod.CProd}}</span>
                <span class="item-values">{{formatQuantity $item.Prod.QCom}}{{$item.Prod.UCom}} x {{formatCurrency $item.Prod.VUnCom}} = {{formatCurrency $item.Prod.VProd}}</span>
            </div>
            <div class="item-desc">{{$item.Prod.XProd}}</div>
        </div>
        {{end}}
        

        
        <div class="section-title">TOTAIS</div>
        <div class="totals">
            <div class="total-line">
                <span>Qtde itens:</span>
                <span>{{len .NFe.NFe.InfNFe.Det}}</span>
            </div>
            <div class="total-line">
                <span>Valor total:</span>
                <span>{{formatCurrency .NFe.NFe.InfNFe.Total.ICMSTot.VProd}}</span>
            </div>
            {{if gt .NFe.NFe.InfNFe.Total.ICMSTot.VDesc 0.0}}
            
            
            <div class="total-line">
                <span>Desconto:</span>
                <span>{{formatCurrency .NFe.NFe.InfNFe.Total.ICMSTot.VDesc}}</span>
            </div>
            {{end}}
            {{if gt .NFe.NFe.InfNFe.Total.ICMSTot.VOutro 0.0}}
            <div class="total-line">
                <span>Outros valores:</span>
                <span>{{formatCurrency .NFe.NFe.InfNFe.Total.ICMSTot.VOutro}}</span>
            </div>
            {{end}}
            
            <div class="total-line total-final">
                <span>TOTAL A PAGAR:</span>
                <span>{{formatCurrency .NFe.NFe.InfNFe.Total.ICMSTot.VNF}}</span>
            </div>
        </div>

        
        <div class="section-title">PAGAMENTO</div>
        <div class="payment">
            {{range .NFe.NFe.InfNFe.Pag.DetPag}}
            <div class="payment-line">
                <span>{{getPaymentMethod .TPag}}:</span>
                <span>{{formatCurrency .VPag}}</span>
            </div>
            {{end}}
            {{if gt .NFe.NFe.InfNFe.Pag.VTroco 0.0}}
            <div class="payment-line">
                <span>Troco:</span>
                <span>{{formatCurrency .NFe.NFe.InfNFe.Pag.VTroco}}</span>
            </div>
            {{end}}
        </div>

        
        <div class="section-title">CONSUMIDOR</div>
        <div class="consumer">
            {{if .NFe.NFe.InfNFe.Dest}}
                {{.NFe.NFe.InfNFe.Dest.XNome}}
            {{else}}
                CONSUMIDOR NÃO IDENTIFICADO
            {{end}}
        </div>

        
        <div class="nfc-info">
            NFC-e Nº {{.NFe.NFe.InfNFe.Ide.NNF}} Série {{.NFe.NFe.InfNFe.Ide.Serie}}<br>
            Emissão: {{formatDate .NFe.NFe.InfNFe.Ide.DHEmi}}
        </div>

        
        <div class="section-title">DADOS DA NFC-e</div>
        <div class="footer">
            Protocolo: {{.NFe.ProtNFe.InfProt.NProt}}<br>
            Autorização: {{formatDate .NFe.ProtNFe.InfProt.DhRecbto}}<br>
            <div class="key">{{formatKey .NFe.GetChaveAcesso}}</div>
        

        {{if .NFe.NFe.InfNFe.InfAdic}}
            {{if .NFe.NFe.InfNFe.InfAdic.InfCpl}}
                <div style="margin-top: 8px; font-size: 10px;">
                    <strong>Informações de interesse do contribuinte:</strong><br>
                    {{.NFe.NFe.InfNFe.InfAdic.InfCpl}}
                </div>
                
            {{end}}
        {{end}}
        </div>
        
        {{if .NFe.GetQRCode}}
            
            <div class="qr-code">
                {{generateQRCode .NFe.GetQRCode}}
                <div class="qr-text">Consulta via Leitor QR Code</div>
            </div>
            
        {{end}}
    </div>
</body>
</html>
`
