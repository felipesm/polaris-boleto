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

func codigoBarrasInvalido(codigoBarras string) Erro {
	msg := fmt.Sprintf("Ocorreu um erro ao tentar gerar a linha digitável, pois o código de barras é inválido! O código %s não tem o tamanho correto.", codigoBarras)
	return Erro{Titulo: "Código de Barras Inválido", Mensagem: msg, Status: 400}
}