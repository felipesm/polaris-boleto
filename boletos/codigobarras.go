package boletos

// CodigoBarras - representa a estrutura para o código de barras
type CodigoBarras struct {
	CodigoBanco          string `json:"codigobanco"`
	CodigoMoeda          string `json:"codigomoeda"`
	DV                   string `json:"digitoverificador"`
	FatorVencimento      string `json:"fatorvencimento"`
	Valor                string `json:"valor"`
	CampoLivre           string `json:"campolivre"`
	CodigoBarrasCompleto string `json:"codigobarras"`
}
