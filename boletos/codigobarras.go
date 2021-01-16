package boletos

// CodigoBarras - representa a estrutura para o código de barras
type CodigoBarras struct {
	CodigoBanco          string
	CodigoMoeda          string
	DV                   string
	FatorVencimento      string
	Valor                string
	CampoLivre           string
	CodigoBarrasCompleto string
}
