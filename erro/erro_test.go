package errorutil

import (
	"fmt"
	"testing"
)

const erroPadrao = "Valor Esperado %v - Valor recebido foi %v."

func TestCodigoBancoInvalido(t *testing.T) {
	t.Parallel()
	codigoBanco := "123"
	msg := fmt.Sprintf("Ocorreu um erro ao tentar identificar o banco, pois não existe serviço disponível para o código %s! Informe um código de banco válido.", codigoBanco)
	erroTest := Erro{Titulo: "Código do Banco Inválido", Mensagem: msg, Status: 400}

	erro := CodigoBancoInvalido(codigoBanco)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}

func TestAgenciaInvalida(t *testing.T) {
	t.Parallel()
	agencia := "12345"
	tamanho := int8(4)
	msg := fmt.Sprintf("O código %s informado para a agência é inválido. O código deve possuir no máximo %d dígitos.", agencia, tamanho)
	erroTest := Erro{Titulo: "Agência Inválida", Mensagem: msg, Status: 400}

	erro := AgenciaInvalida(agencia, tamanho)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}

func TestCarteiraInvalida(t *testing.T) {
	t.Parallel()
	carteira := "123"
	tamanho := int8(3)
	msg := fmt.Sprintf("O código %s informado para a carteira é inválido. O código deve possuir no máximo %d dígitos.", carteira, tamanho)
	erroTest := Erro{Titulo: "Carteira Inválida", Mensagem: msg, Status: 400}

	erro := CarteiraInvalida(carteira, tamanho)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}

func TestVencimentoInvalido(t *testing.T) {
	t.Parallel()
	msg := "O vencimento informado é inválido. A data de vencimento deve ser no formato yyyy-MM-dd e precisa ser superior a 06/10/1997."
	erroTest := Erro{Titulo: "Vencimento Inválido", Mensagem: msg, Status: 400}

	erro := VencimentoInvalido()

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}

func TestCodigoBeneficiarioInvalido(t *testing.T) {
	t.Parallel()
	codigoBeneficiario := "1234567"
	tamanho := int8(7)
	msg := fmt.Sprintf("O código %s informado para o código beneficiário é inválido. O código deve possuir no máximo %d dígitos.", codigoBeneficiario, tamanho)
	erroTest := Erro{Titulo: "Código Beneficiário Inválido", Mensagem: msg, Status: 400}

	erro := CodigoBeneficiarioInvalido(codigoBeneficiario, tamanho)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}

}

func TestNossoNumeroInvalido(t *testing.T) {
	t.Parallel()
	nossoNumero := "12345678901"
	tamanho := int8(11)
	msg := fmt.Sprintf("O nosso número %s é inválido, pois ele deve possuir no máximo %d dígitos.", nossoNumero, tamanho)
	erroTest := Erro{Titulo: "Nosso Número Inválido", Mensagem: msg, Status: 400}

	erro := NossoNumeroInvalido(nossoNumero, tamanho)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}

func TestCodigoBarrasInvalido(t *testing.T) {
	t.Parallel()
	codigoBarras := "01234567890"
	msg := fmt.Sprintf("Ocorreu um erro ao tentar gerar a linha digitável, pois o código de barras é inválido! O código %s não tem o tamanho correto.", codigoBarras)
	erroTest := Erro{Titulo: "Código de Barras Inválido", Mensagem: msg, Status: 400}

	erro := CodigoBarrasInvalido(codigoBarras)

	if erroTest != erro {
		t.Errorf(erroPadrao, erroTest, erro)
	}
}
