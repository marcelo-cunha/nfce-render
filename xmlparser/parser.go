package xmlparser

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

// NFeProc representa a estrutura principal do XML da NF-e processada
type NFeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Versao  string   `xml:"versao,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	NFe     NFe      `xml:"NFe"`
	ProtNFe ProtNFe  `xml:"protNFe"`
}

// NFe representa a estrutura da Nota Fiscal Eletrônica
type NFe struct {
	XMLName    xml.Name    `xml:"NFe"`
	InfNFe     InfNFe      `xml:"infNFe"`
	InfNFeSupl *InfNFeSupl `xml:"infNFeSupl,omitempty"`
}

// InfNFe contém as informações da NF-e
type InfNFe struct {
	ID      string   `xml:"Id,attr"`
	Versao  string   `xml:"versao,attr"`
	Ide     Ide      `xml:"ide"`
	Emit    Emit     `xml:"emit"`
	Dest    *Dest    `xml:"dest,omitempty"`
	Det     []Det    `xml:"det"`
	Total   Total    `xml:"total"`
	Transp  *Transp  `xml:"transp,omitempty"`
	Cobr    *Cobr    `xml:"cobr,omitempty"`
	Pag     Pag      `xml:"pag"`
	InfAdic *InfAdic `xml:"infAdic,omitempty"`
}

// Ide contém as informações de identificação da NF-e
type Ide struct {
	CUF      string     `xml:"cUF"`
	CNF      string     `xml:"cNF"`
	NatOp    string     `xml:"natOp"`
	Mod      string     `xml:"mod"`
	Serie    string     `xml:"serie"`
	NNF      string     `xml:"nNF"`
	DHEmi    time.Time  `xml:"dhEmi"`
	DHSaiEnt *time.Time `xml:"dhSaiEnt,omitempty"`
	TpNF     string     `xml:"tpNF"`
	IDDest   string     `xml:"idDest"`
	CMunFG   string     `xml:"cMunFG"`
	TpImp    string     `xml:"tpImp"`
	TpEmis   string     `xml:"tpEmis"`
	CDV      string     `xml:"cDV"`
	TpAmb    string     `xml:"tpAmb"`
	FinNFe   string     `xml:"finNFe"`
	IndFinal string     `xml:"indFinal"`
	IndPres  string     `xml:"indPres"`
	ProcEmi  string     `xml:"procEmi"`
	VerProc  string     `xml:"verProc"`
}

// Emit contém as informações do emitente
type Emit struct {
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	XFant     string    `xml:"xFant,omitempty"`
	EnderEmit EnderEmit `xml:"enderEmit"`
	IE        string    `xml:"IE"`
	CRT       string    `xml:"CRT"`
}

// EnderEmit contém o endereço do emitente
type EnderEmit struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl,omitempty"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
	Fone    string `xml:"fone,omitempty"`
}

// Dest contém as informações do destinatário
type Dest struct {
	CNPJ      string     `xml:"CNPJ,omitempty"`
	CPF       string     `xml:"CPF,omitempty"`
	XNome     string     `xml:"xNome"`
	EnderDest *EnderDest `xml:"enderDest,omitempty"`
}

// EnderDest contém o endereço do destinatário
type EnderDest struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl,omitempty"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
}

// Det contém os detalhes dos produtos/serviços
type Det struct {
	NItem   string  `xml:"nItem,attr"`
	Prod    Prod    `xml:"prod"`
	Imposto Imposto `xml:"imposto"`
}

// Prod contém as informações do produto
type Prod struct {
	CProd    string  `xml:"cProd"`
	CEAN     string  `xml:"cEAN,omitempty"`
	XProd    string  `xml:"xProd"`
	NCM      string  `xml:"NCM"`
	CFOP     string  `xml:"CFOP"`
	UCom     string  `xml:"uCom"`
	QCom     float64 `xml:"qCom"`
	VUnCom   float64 `xml:"vUnCom"`
	VProd    float64 `xml:"vProd"`
	CEANTrib string  `xml:"cEANTrib,omitempty"`
	UTrib    string  `xml:"uTrib"`
	QTrib    float64 `xml:"qTrib"`
	VUnTrib  float64 `xml:"vUnTrib"`
	IndTot   string  `xml:"indTot"`
}

// Imposto contém as informações de impostos
type Imposto struct {
	ICMS   *ICMS   `xml:"ICMS,omitempty"`
	IPI    *IPI    `xml:"IPI,omitempty"`
	PIS    *PIS    `xml:"PIS,omitempty"`
	COFINS *COFINS `xml:"COFINS,omitempty"`
}

// ICMS representa as informações do ICMS
type ICMS struct {
	ICMS00    *ICMS00    `xml:"ICMS00,omitempty"`
	ICMS10    *ICMS10    `xml:"ICMS10,omitempty"`
	ICMS20    *ICMS20    `xml:"ICMS20,omitempty"`
	ICMS30    *ICMS30    `xml:"ICMS30,omitempty"`
	ICMS40    *ICMS40    `xml:"ICMS40,omitempty"`
	ICMS51    *ICMS51    `xml:"ICMS51,omitempty"`
	ICMS60    *ICMS60    `xml:"ICMS60,omitempty"`
	ICMS70    *ICMS70    `xml:"ICMS70,omitempty"`
	ICMS90    *ICMS90    `xml:"ICMS90,omitempty"`
	ICMSPart  *ICMSPart  `xml:"ICMSPart,omitempty"`
	ICMSST    *ICMSST    `xml:"ICMSST,omitempty"`
	ICMSSN101 *ICMSSN101 `xml:"ICMSSN101,omitempty"`
	ICMSSN102 *ICMSSN102 `xml:"ICMSSN102,omitempty"`
	ICMSSN201 *ICMSSN201 `xml:"ICMSSN201,omitempty"`
	ICMSSN202 *ICMSSN202 `xml:"ICMSSN202,omitempty"`
	ICMSSN500 *ICMSSN500 `xml:"ICMSSN500,omitempty"`
	ICMSSN900 *ICMSSN900 `xml:"ICMSSN900,omitempty"`
}

// ICMS00 representa ICMS tributado integralmente
type ICMS00 struct {
	Orig  string  `xml:"orig"`
	CST   string  `xml:"CST"`
	ModBC string  `xml:"modBC"`
	VBC   float64 `xml:"vBC"`
	PICMS float64 `xml:"pICMS"`
	VICMS float64 `xml:"vICMS"`
}

// ICMS10 representa ICMS tributado e com cobrança do ICMS por substituição tributária
type ICMS10 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	ModBC   string  `xml:"modBC"`
	VBC     float64 `xml:"vBC"`
	PICMS   float64 `xml:"pICMS"`
	VICMS   float64 `xml:"vICMS"`
	ModBCST string  `xml:"modBCST"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBC  float64 `xml:"pRedBC,omitempty"`
	VBCST   float64 `xml:"vBCST"`
	PICMSST float64 `xml:"pICMSST"`
	VICMSST float64 `xml:"vICMSST"`
}

// ICMS20 representa ICMS com redução de base de cálculo
type ICMS20 struct {
	Orig   string  `xml:"orig"`
	CST    string  `xml:"CST"`
	ModBC  string  `xml:"modBC"`
	PRedBC float64 `xml:"pRedBC"`
	VBC    float64 `xml:"vBC"`
	PICMS  float64 `xml:"pICMS"`
	VICMS  float64 `xml:"vICMS"`
}

// ICMS30 representa ICMS isento ou não tributado e com cobrança do ICMS por substituição tributária
type ICMS30 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	ModBCST string  `xml:"modBCST"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBC  float64 `xml:"pRedBC,omitempty"`
	VBCST   float64 `xml:"vBCST"`
	PICMSST float64 `xml:"pICMSST"`
	VICMSST float64 `xml:"vICMSST"`
}

// ICMS40 representa ICMS isento
type ICMS40 struct {
	Orig   string  `xml:"orig"`
	CST    string  `xml:"CST"`
	VICMSDeson float64 `xml:"vICMSDeson,omitempty"`
	MotDesICMS string  `xml:"motDesICMS,omitempty"`
}

// ICMS51 representa ICMS diferido
type ICMS51 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	ModBC   string  `xml:"modBC,omitempty"`
	PRedBC  float64 `xml:"pRedBC,omitempty"`
	VBC     float64 `xml:"vBC,omitempty"`
	PICMS   float64 `xml:"pICMS,omitempty"`
	VICMSOp float64 `xml:"vICMSOp,omitempty"`
	PDif    float64 `xml:"pDif,omitempty"`
	VICMSDif float64 `xml:"vICMSDif,omitempty"`
	VICMS   float64 `xml:"vICMS,omitempty"`
}

// ICMS60 representa ICMS cobrado anteriormente por substituição tributária
type ICMS60 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	VBCSTRet float64 `xml:"vBCSTRet,omitempty"`
	VICMSSTRet float64 `xml:"vICMSSTRet,omitempty"`
}

// ICMS70 representa ICMS com redução de base de cálculo e cobrança do ICMS por substituição tributária
type ICMS70 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	ModBC   string  `xml:"modBC"`
	PRedBC  float64 `xml:"pRedBC"`
	VBC     float64 `xml:"vBC"`
	PICMS   float64 `xml:"pICMS"`
	VICMS   float64 `xml:"vICMS"`
	ModBCST string  `xml:"modBCST"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBCST float64 `xml:"pRedBCST,omitempty"`
	VBCST   float64 `xml:"vBCST"`
	PICMSST float64 `xml:"pICMSST"`
	VICMSST float64 `xml:"vICMSST"`
}

// ICMS90 representa ICMS outros
type ICMS90 struct {
	Orig    string  `xml:"orig"`
	CST     string  `xml:"CST"`
	ModBC   string  `xml:"modBC,omitempty"`
	VBC     float64 `xml:"vBC,omitempty"`
	PRedBC  float64 `xml:"pRedBC,omitempty"`
	PICMS   float64 `xml:"pICMS,omitempty"`
	VICMS   float64 `xml:"vICMS,omitempty"`
	ModBCST string  `xml:"modBCST,omitempty"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBCST float64 `xml:"pRedBCST,omitempty"`
	VBCST   float64 `xml:"vBCST,omitempty"`
	PICMSST float64 `xml:"pICMSST,omitempty"`
	VICMSST float64 `xml:"vICMSST,omitempty"`
}

// ICMSPart representa ICMS partilha
type ICMSPart struct {
	Orig      string  `xml:"orig"`
	CST       string  `xml:"CST"`
	ModBC     string  `xml:"modBC"`
	VBC       float64 `xml:"vBC"`
	PRedBC    float64 `xml:"pRedBC,omitempty"`
	PICMS     float64 `xml:"pICMS"`
	VICMS     float64 `xml:"vICMS"`
	ModBCST   string  `xml:"modBCST"`
	PMVAST    float64 `xml:"pMVAST,omitempty"`
	PREDBCST  float64 `xml:"pRedBCST,omitempty"`
	VBCST     float64 `xml:"vBCST"`
	PICMSST   float64 `xml:"pICMSST"`
	VICMSST   float64 `xml:"vICMSST"`
	PBCOp     float64 `xml:"pBCOp"`
	UFST      string  `xml:"UFST"`
}

// ICMSST representa ICMS substituição tributária
type ICMSST struct {
	Orig      string  `xml:"orig"`
	CST       string  `xml:"CST"`
	VBCSTRet  float64 `xml:"vBCSTRet"`
	VICMSSTRet float64 `xml:"vICMSSTRet"`
	VBCSTDest float64 `xml:"vBCSTDest"`
	VICMSSTDest float64 `xml:"vICMSSTDest"`
}

// ICMSSN101 representa ICMS Simples Nacional tributado pelo Simples Nacional com permissão de crédito
type ICMSSN101 struct {
	Orig    string  `xml:"orig"`
	CSOSN   string  `xml:"CSOSN"`
	PCredSN float64 `xml:"pCredSN"`
	VCredICMSSN float64 `xml:"vCredICMSSN"`
}

// ICMSSN102 representa ICMS Simples Nacional tributado pelo Simples Nacional sem permissão de crédito
type ICMSSN102 struct {
	Orig  string `xml:"orig"`
	CSOSN string `xml:"CSOSN"`
}

// ICMSSN201 representa ICMS Simples Nacional tributado pelo Simples Nacional com permissão de crédito e com cobrança do ICMS por substituição tributária
type ICMSSN201 struct {
	Orig    string  `xml:"orig"`
	CSOSN   string  `xml:"CSOSN"`
	ModBCST string  `xml:"modBCST"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBCST float64 `xml:"pRedBCST,omitempty"`
	VBCST   float64 `xml:"vBCST"`
	PICMSST float64 `xml:"pICMSST"`
	VICMSST float64 `xml:"vICMSST"`
	PCredSN float64 `xml:"pCredSN"`
	VCredICMSSN float64 `xml:"vCredICMSSN"`
}

// ICMSSN202 representa ICMS Simples Nacional tributado pelo Simples Nacional sem permissão de crédito e com cobrança do ICMS por substituição tributária
type ICMSSN202 struct {
	Orig    string  `xml:"orig"`
	CSOSN   string  `xml:"CSOSN"`
	ModBCST string  `xml:"modBCST"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBCST float64 `xml:"pRedBCST,omitempty"`
	VBCST   float64 `xml:"vBCST"`
	PICMSST float64 `xml:"pICMSST"`
	VICMSST float64 `xml:"vICMSST"`
}

// ICMSSN500 representa ICMS Simples Nacional ICMS cobrado anteriormente por substituição tributária (substituído) ou por antecipação
type ICMSSN500 struct {
	Orig       string  `xml:"orig"`
	CSOSN      string  `xml:"CSOSN"`
	VBCSTRet   float64 `xml:"vBCSTRet,omitempty"`
	VICMSSTRet float64 `xml:"vICMSSTRet,omitempty"`
}

// ICMSSN900 representa ICMS Simples Nacional outros
type ICMSSN900 struct {
	Orig    string  `xml:"orig"`
	CSOSN   string  `xml:"CSOSN"`
	ModBC   string  `xml:"modBC,omitempty"`
	VBC     float64 `xml:"vBC,omitempty"`
	PRedBC  float64 `xml:"pRedBC,omitempty"`
	PICMS   float64 `xml:"pICMS,omitempty"`
	VICMS   float64 `xml:"vICMS,omitempty"`
	ModBCST string  `xml:"modBCST,omitempty"`
	PMVAST  float64 `xml:"pMVAST,omitempty"`
	PREDBCST float64 `xml:"pRedBCST,omitempty"`
	VBCST   float64 `xml:"vBCST,omitempty"`
	PICMSST float64 `xml:"pICMSST,omitempty"`
	VICMSST float64 `xml:"vICMSST,omitempty"`
	PCredSN float64 `xml:"pCredSN,omitempty"`
	VCredICMSSN float64 `xml:"vCredICMSSN,omitempty"`
}

// IPI representa as informações do IPI
type IPI struct {
	CNPJProd string   `xml:"CNPJProd,omitempty"`
	CEnq     string   `xml:"cEnq"`
	IPITrib  *IPITrib `xml:"IPITrib,omitempty"`
	IPINT    *IPINT   `xml:"IPINT,omitempty"`
}

// IPITrib representa IPI tributado
type IPITrib struct {
	CST   string  `xml:"CST"`
	VBC   float64 `xml:"vBC,omitempty"`
	PIPI  float64 `xml:"pIPI,omitempty"`
	QUnid float64 `xml:"qUnid,omitempty"`
	VUnid float64 `xml:"vUnid,omitempty"`
	VIPI  float64 `xml:"vIPI"`
}

// IPINT representa IPI não tributado
type IPINT struct {
	CST string `xml:"CST"`
}

// PIS representa as informações do PIS
type PIS struct {
	PISAliq *PISAliq `xml:"PISAliq,omitempty"`
	PISQtde *PISQtde `xml:"PISQtde,omitempty"`
	PISNT   *PISNT   `xml:"PISNT,omitempty"`
	PISOutr *PISOutr `xml:"PISOutr,omitempty"`
}

// PISAliq representa PIS tributado pela alíquota
type PISAliq struct {
	CST  string  `xml:"CST"`
	VBC  float64 `xml:"vBC"`
	PPIS float64 `xml:"pPIS"`
	VPIS float64 `xml:"vPIS"`
}

// PISQtde representa PIS tributado por quantidade
type PISQtde struct {
	CST     string  `xml:"CST"`
	QBCProd float64 `xml:"qBCProd"`
	VAliqProd float64 `xml:"vAliqProd"`
	VPIS    float64 `xml:"vPIS"`
}

// PISNT representa PIS não tributado
type PISNT struct {
	CST string `xml:"CST"`
}

// PISOutr representa PIS outras operações
type PISOutr struct {
	CST       string  `xml:"CST"`
	VBC       float64 `xml:"vBC,omitempty"`
	PPIS      float64 `xml:"pPIS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VPIS      float64 `xml:"vPIS"`
}

// COFINS representa as informações do COFINS
type COFINS struct {
	COFINSAliq *COFINSAliq `xml:"COFINSAliq,omitempty"`
	COFINSQtde *COFINSQtde `xml:"COFINSQtde,omitempty"`
	COFINSNT   *COFINSNT   `xml:"COFINSNT,omitempty"`
	COFINSOutr *COFINSOutr `xml:"COFINSOutr,omitempty"`
}

// COFINSAliq representa COFINS tributado pela alíquota
type COFINSAliq struct {
	CST     string  `xml:"CST"`
	VBC     float64 `xml:"vBC"`
	PCOFINS float64 `xml:"pCOFINS"`
	VCOFINS float64 `xml:"vCOFINS"`
}

// COFINSQtde representa COFINS tributado por quantidade
type COFINSQtde struct {
	CST       string  `xml:"CST"`
	QBCProd   float64 `xml:"qBCProd"`
	VAliqProd float64 `xml:"vAliqProd"`
	VCOFINS   float64 `xml:"vCOFINS"`
}

// COFINSNT representa COFINS não tributado
type COFINSNT struct {
	CST string `xml:"CST"`
}

// COFINSOutr representa COFINS outras operações
type COFINSOutr struct {
	CST       string  `xml:"CST"`
	VBC       float64 `xml:"vBC,omitempty"`
	PCOFINS   float64 `xml:"pCOFINS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VCOFINS   float64 `xml:"vCOFINS"`
}

// Total contém os valores totais da NF-e
type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

// ICMSTot contém os totais relativos ao ICMS
type ICMSTot struct {
	VBC       float64 `xml:"vBC"`
	VICMS     float64 `xml:"vICMS"`
	VICMSDeson float64 `xml:"vICMSDeson"`
	VFCPUFDest float64 `xml:"vFCPUFDest,omitempty"`
	VICMSUFDest float64 `xml:"vICMSUFDest,omitempty"`
	VICMSUFRemet float64 `xml:"vICMSUFRemet,omitempty"`
	VFCP      float64 `xml:"vFCP,omitempty"`
	VBCST     float64 `xml:"vBCST"`
	VST       float64 `xml:"vST"`
	VFCPST    float64 `xml:"vFCPST,omitempty"`
	VFCPSTRet float64 `xml:"vFCPSTRet,omitempty"`
	VProd     float64 `xml:"vProd"`
	VFrete    float64 `xml:"vFrete"`
	VSeg      float64 `xml:"vSeg"`
	VDesc     float64 `xml:"vDesc"`
	VII       float64 `xml:"vII"`
	VIPI      float64 `xml:"vIPI"`
	VIPIDevol float64 `xml:"vIPIDevol,omitempty"`
	VPIS      float64 `xml:"vPIS"`
	VCOFINS   float64 `xml:"vCOFINS"`
	VOutro    float64 `xml:"vOutro"`
	VNF       float64 `xml:"vNF"`
	VTotTrib  float64 `xml:"vTotTrib,omitempty"`
}

// Transp contém as informações de transporte
type Transp struct {
	ModFrete string       `xml:"modFrete"`
	Transportadora *Transportadora `xml:"transporta,omitempty"`
	VeicTransp *VeicTransp `xml:"veicTransp,omitempty"`
	Vol        []Vol       `xml:"vol,omitempty"`
}

// Transportadora contém as informações da transportadora
type Transportadora struct {
	CNPJ      string `xml:"CNPJ,omitempty"`
	CPF       string `xml:"CPF,omitempty"`
	XNome     string `xml:"xNome,omitempty"`
	IE        string `xml:"IE,omitempty"`
	XEnder    string `xml:"xEnder,omitempty"`
	XMun      string `xml:"xMun,omitempty"`
	UF        string `xml:"UF,omitempty"`
}

// VeicTransp contém as informações do veículo de transporte
type VeicTransp struct {
	Placa string `xml:"placa"`
	UF    string `xml:"UF"`
	RNTC  string `xml:"RNTC,omitempty"`
}

// Vol contém as informações de volumes
type Vol struct {
	QVol  int     `xml:"qVol,omitempty"`
	Esp   string  `xml:"esp,omitempty"`
	Marca string  `xml:"marca,omitempty"`
	NVol  string  `xml:"nVol,omitempty"`
	PesoL float64 `xml:"pesoL,omitempty"`
	PesoB float64 `xml:"pesoB,omitempty"`
}

// Cobr contém as informações de cobrança
type Cobr struct {
	Fat  *Fat  `xml:"fat,omitempty"`
	Dup  []Dup `xml:"dup,omitempty"`
}

// Fat contém as informações da fatura
type Fat struct {
	NFat  string  `xml:"nFat,omitempty"`
	VOrig float64 `xml:"vOrig,omitempty"`
	VDesc float64 `xml:"vDesc,omitempty"`
	VLiq  float64 `xml:"vLiq,omitempty"`
}

// Dup contém as informações das duplicatas
type Dup struct {
	NDup  string    `xml:"nDup,omitempty"`
	DVenc time.Time `xml:"dVenc,omitempty"`
	VDup  float64   `xml:"vDup"`
}

// Pag contém as informações de pagamento
type Pag struct {
	DetPag []DetPag `xml:"detPag"`
	VTroco float64  `xml:"vTroco,omitempty"`
}

// DetPag contém os detalhes do pagamento
type DetPag struct {
	IndPag string  `xml:"indPag,omitempty"`
	TPag   string  `xml:"tPag"`
	XPag   string  `xml:"xPag,omitempty"`
	VPag   float64 `xml:"vPag"`
	Card   *Card   `xml:"card,omitempty"`
}

// Card contém as informações do cartão
type Card struct {
	TpIntegra string `xml:"tpIntegra"`
	CNPJ      string `xml:"CNPJ,omitempty"`
	TBand     string `xml:"tBand,omitempty"`
	CAut      string `xml:"cAut,omitempty"`
}

// InfAdic contém as informações adicionais
type InfAdic struct {
	InfAdFisco string `xml:"infAdFisco,omitempty"`
	InfCpl     string `xml:"infCpl,omitempty"`
}

// InfNFeSupl contém as informações suplementares da NF-e
type InfNFeSupl struct {
	QrCode   string `xml:"qrCode"`
	UrlChave string `xml:"urlChave"`
}

// ProtNFe contém as informações do protocolo de autorização
type ProtNFe struct {
	Versao  string  `xml:"versao,attr"`
	InfProt InfProt `xml:"infProt"`
}

// InfProt contém as informações do protocolo
type InfProt struct {
	TpAmb    string    `xml:"tpAmb"`
	VerAplic string    `xml:"verAplic"`
	ChNFe    string    `xml:"chNFe"`
	DhRecbto time.Time `xml:"dhRecbto"`
	NProt    string    `xml:"nProt"`
	DigVal   string    `xml:"digVal"`
	CStat    string    `xml:"cStat"`
	XMotivo  string    `xml:"xMotivo"`
}

// IsNFCe verifica se a NF-e é uma NFC-e (modelo 65)
func (nfe *NFeProc) IsNFCe() bool {
	return nfe.NFe.InfNFe.Ide.Mod == "65"
}

// GetQRCode retorna o QR Code da NFC-e
func (nfe *NFeProc) GetQRCode() string {
	if nfe.NFe.InfNFeSupl != nil {
		return nfe.NFe.InfNFeSupl.QrCode
	}
	return ""
}

// GetChaveAcesso retorna a chave de acesso da NF-e
func (nfe *NFeProc) GetChaveAcesso() string {
	return nfe.ProtNFe.InfProt.ChNFe
}

// GetNumeroNF retorna o número da NF-e
func (nfe *NFeProc) GetNumeroNF() string {
	return nfe.NFe.InfNFe.Ide.NNF
}

// GetSerieNF retorna a série da NF-e
func (nfe *NFeProc) GetSerieNF() string {
	return nfe.NFe.InfNFe.Ide.Serie
}

// GetDataEmissao retorna a data de emissão da NF-e
func (nfe *NFeProc) GetDataEmissao() time.Time {
	return nfe.NFe.InfNFe.Ide.DHEmi
}

// GetValorTotal retorna o valor total da NF-e
func (nfe *NFeProc) GetValorTotal() float64 {
	return nfe.NFe.InfNFe.Total.ICMSTot.VNF
}

// FormatCNPJ formata um CNPJ
func FormatCNPJ(cnpj string) string {
	if len(cnpj) != 14 {
		return cnpj
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s",
		cnpj[0:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:14])
}

// FormatCPF formata um CPF
func FormatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return fmt.Sprintf("%s.%s.%s-%s",
		cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

// FormatCEP formata um CEP
func FormatCEP(cep string) string {
	if len(cep) != 8 {
		return cep
	}
	return fmt.Sprintf("%s-%s", cep[0:5], cep[5:8])
}

// FormatCurrency formata um valor monetário
func FormatCurrency(value float64) string {
	return fmt.Sprintf("R$ %.2f", value)
}

// FormatQuantity formata uma quantidade
func FormatQuantity(value float64) string {
	// Remove zeros desnecessários
	formatted := strconv.FormatFloat(value, 'f', -1, 64)
	return formatted
}

// ParseXML faz o parse do XML da NF-e
func ParseXML(xmlContent []byte) (*NFeProc, error) {
	var nfe NFeProc
	if err := xml.Unmarshal(xmlContent, &nfe); err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do XML: %w", err)
	}
	return &nfe, nil
}

// GetPaymentMethodDescription retorna a descrição do método de pagamento
func GetPaymentMethodDescription(tPag string) string {
	switch tPag {
	case "01":
		return "Dinheiro"
	case "02":
		return "Cheque"
	case "03":
		return "Cartão de Crédito"
	case "04":
		return "Cartão de Débito"
	case "05":
		return "Crédito Loja"
	case "10":
		return "Vale Alimentação"
	case "11":
		return "Vale Refeição"
	case "12":
		return "Vale Presente"
	case "13":
		return "Vale Combustível"
	case "14":
		return "Duplicata Mercantil"
	case "15":
		return "Boleto Bancário"
	case "16":
		return "Depósito Bancário"
	case "17":
		return "Pagamento Instantâneo (PIX)"
	case "18":
		return "Transferência bancária, Carteira Digital"
	case "19":
		return "Programa de fidelidade, Cashback, Crédito Virtual"
	case "90":
		return "Sem pagamento"
	case "99":
		return "Outros"
	default:
		return "Não informado"
	}
}