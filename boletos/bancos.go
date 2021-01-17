package boletos

// Banco - interface que representa diferentes bancos
type Banco interface {
	SetCodigo()
	SetAgencia(agencia int32)
	SetCarteira(carteira int8)
	SetCodigoBeneficiario(contaBeneficiario int32)
	SetFatorVencimento(dataVencimento string, zerarVencimento bool)
	SetValorBoleto(valor float64)
	SetNossoNumero(nossoNumero string)
	getDVCodigoBarras(codigoBarras CodigoBarras) string
	getCampoLivre() string
	GetCodigoBarras() CodigoBarras
	GetLinhaDigitavel(codigoBarras string) LinhaDigitavel
}
