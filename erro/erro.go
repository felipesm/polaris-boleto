package errorutil

import "fmt"

// Erro - representa estrutura para mensagens de erro
type Erro struct {
	Titulo   string `json:"titulo"`
	Mensagem string `json:"mensagem"`
	Status   int    `json:"status"`
}

// CodigoBancoInvalido - erro para código de banco inválido
func CodigoBancoInvalido(codigoBanco string) Erro {
	msg := fmt.Sprintf("Ocorreu um erro ao tentar identificar o banco, pois não existe serviço disponível para o código %s! Informe um código de banco válido.", codigoBanco)
	return Erro{Titulo: "Código do Banco Inválido", Mensagem: msg, Status: 400}
}

// AgenciaInvalida - erro para agência inválida
func AgenciaInvalida(agencia string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para a agência é inválido. O código deve possuir no máximo %d dígitos.", agencia, tamanho)
	return Erro{Titulo: "Agência Inválida", Mensagem: msg, Status: 400}
}

// CarteiraInvalida - erro para carteira inválida
func CarteiraInvalida(carteira string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para a carteira é inválido. O código deve possuir no máximo %d dígitos.", carteira, tamanho)
	return Erro{Titulo: "Carteira Inválida", Mensagem: msg, Status: 400}
}

// VencimentoInvalido - erro  para vencimento inválido
func VencimentoInvalido() Erro {
	msg := "O vencimento informado é inválido. A data de vencimento deve ser no formato yyyy-MM-dd e precisa ser superior a 06/10/1997."
	return Erro{Titulo: "Vencimento Inválido", Mensagem: msg, Status: 400}
}

// CodigoBeneficiarioInvalido - erro para código beneficiário inválido
func CodigoBeneficiarioInvalido(codigoBeneficiario string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para o código beneficiário é inválido. O código deve possuir no máximo %d dígitos.", codigoBeneficiario, tamanho)
	return Erro{Titulo: "Código Beneficiário Inválido", Mensagem: msg, Status: 400}
}

// NossoNumeroInvalido - erro para nosso número inválido
func NossoNumeroInvalido(nossoNumero string, tamanho int8) Erro {
	msg := fmt.Sprintf("O nosso número %s é inválido, pois ele deve possuir no máximo %d dígitos.", nossoNumero, tamanho)
	return Erro{Titulo: "Código Beneficiário Inválido", Mensagem: msg, Status: 400}
}

// CodigoBarrasInvalido - erro para código barras inválido
func CodigoBarrasInvalido(codigoBarras string) Erro {
	msg := fmt.Sprintf("Ocorreu um erro ao tentar gerar a linha digitável, pois o código de barras é inválido! O código %s não tem o tamanho correto.", codigoBarras)
	return Erro{Titulo: "Código de Barras Inválido", Mensagem: msg, Status: 400}
}
