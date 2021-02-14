package boletos

import "fmt"

// Erro - representa estrutura para mensagens de erro
type Erro struct {
	Titulo   string `json:"titulo"`
	Mensagem string `json:"mensagem"`
	Status   int    `json:"status"`
}

func codigoBancoInvalido(codigoBanco string) Erro {
	msg := fmt.Sprintf("Ocorreu um erro ao tentar identificar o banco, pois não existe serviço disponível para o código %s! Informe um código de banco válido.", codigoBanco)
	return Erro{Titulo: "Código do Banco Inválido", Mensagem: msg, Status: 400}
}

func agenciaInvalida(agencia string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para a agência é inválido. O código deve possuir no máximo %d dígitos.", agencia, tamanho)
	return Erro{Titulo: "Agência Inválida", Mensagem: msg, Status: 400}
}

func carteiraInvalida(carteira string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para a carteira é inválido. O código deve possuir no máximo %d dígitos.", carteira, tamanho)
	return Erro{Titulo: "Carteira Inválida", Mensagem: msg, Status: 400}
}

func codigoBeneficiarioInvalido(codigoBeneficiario string, tamanho int8) Erro {
	msg := fmt.Sprintf("O código %s informado para o código beneficiário é inválido. O código deve possuir no máximo %d dígitos.", codigoBeneficiario, tamanho)
	return Erro{Titulo: "Código Beneficiário Inválido", Mensagem: msg, Status: 400}
}

func nossoNumeroInvalido(nossoNumero string, tamanho int8) Erro {
	msg := fmt.Sprintf("O nosso número %s é inválido, pois ele deve possuir no máximo %d dígitos.", nossoNumero, tamanho)
	return Erro{Titulo: "Código Beneficiário Inválido", Mensagem: msg, Status: 400}
}

func codigoBarrasInvalido(codigoBarras string) Erro {
	msg := fmt.Sprintf("Ocorreu um erro ao tentar gerar a linha digitável, pois o código de barras é inválido! O código %s não tem o tamanho correto.", codigoBarras)
	return Erro{Titulo: "Código de Barras Inválido", Mensagem: msg, Status: 400}
}
