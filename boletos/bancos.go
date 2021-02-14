package boletos

// Banco - interface que representa diferentes bancos
type Banco interface {
	SetCodigo()
	SetAgencia(agencia int32)
	SetCarteira(carteira int16)
	SetCodigoBeneficiario(contaBeneficiario int32)
	SetFatorVencimento(dataVencimento string, zerarVencimento bool)
	SetValorBoleto(valor float64, zerarValor bool)
	SetNossoNumero(nossoNumero string)
	getDVCodigoBarras(codigoBarras CodigoBarras) string
	getCampoLivre() string
	validarDados() Erro
	GetCodigoBarras() (CodigoBarras, Erro)
	GetLinhaDigitavel(codigoBarras string) (LinhaDigitavel, Erro)
}
