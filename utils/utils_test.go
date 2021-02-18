package utils

import (
	"testing"
)

const erroPadrao = "Valor esperado -> %v - Valor recebido -> %v."

/*
	Teste para uma data de vencimento válida.
	Para uma data válida, precisa ser maior que 06/10/1997 no formato yyyy-MM-dd.
*/
func TestCalcularFatorVencimento(t *testing.T) {
	t.Parallel()

	data := "2021-01-23"
	zerarVencimento := false

	valorSucesso := int16(8509)

	valorRetorno, _ := CalcularFatorVencimento(data, zerarVencimento)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste para uma data de vencimento anterior a 07/10/1997 no formato yyyy-MM-dd.
*/
func TestCalcularFatorVencimentoAnterior(t *testing.T) {
	t.Parallel()

	data := "1997-01-01"
	zerarVencimento := false

	valorSucesso := int16(0)

	valorRetorno, _ := CalcularFatorVencimento(data, zerarVencimento)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste para uma data de vencimento com argumento zerarVencimento 'true'.
	O retorno deverá ser '0'.
*/
func TestCalcularFatorVencimentoZerado(t *testing.T) {
	t.Parallel()

	data := "1997-01-01"
	zerarVencimento := true

	valorSucesso := int16(0)

	valorRetorno, _ := CalcularFatorVencimento(data, zerarVencimento)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste para uma data de vencimento inválida.
*/
func TestCalcularFatorVencimentoInvalido(t *testing.T) {
	t.Parallel()

	data := "1997-30-50"
	zerarVencimento := false

	valorSucesso := int16(0)

	valorRetorno, _ := CalcularFatorVencimento(data, zerarVencimento)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste quando o argumento 'zerarValor' for false.
*/
func TestFormatarValorBoleto(t *testing.T) {
	t.Parallel()

	valor := 123.45
	tamanho := 10
	zerarValor := false

	valorSucesso := "0000012345"

	valorRetorno := FormatarValorBoleto(valor, tamanho, zerarValor)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste quando o argumento 'zerarValor' for true -> o retorno sempre deve ser uma string preenchida com zeros.
*/
func TestFormatarValorBoletoZerado(t *testing.T) {
	t.Parallel()

	valor := 80.35
	tamanho := 6
	zerarValor := true

	valorSucesso := "000000"

	valorRetorno := FormatarValorBoleto(valor, tamanho, zerarValor)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

func TestFormatarNossoNumero(t *testing.T) {
	t.Parallel()

	valor := "4860"
	tamanho := 7

	valorSucesso := "0004860"

	valorRetorno := FormatarNossoNumero(valor, tamanho)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste quando nosso número possui algum caractere não numérico.
*/
func TestFormatarNossoNumeroInvalido(t *testing.T) {
	t.Parallel()

	valor := "f4860"
	tamanho := 7

	valorSucesso := ""

	valorRetorno := FormatarNossoNumero(valor, tamanho)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	Teste quando nosso número é maior que o tamanho máximo permitido.
*/
func TestFormatarNossoNumeroMaior(t *testing.T) {
	t.Parallel()

	valor := "12345678"
	tamanho := 7

	valorSucesso := "12345678"

	valorRetorno := FormatarNossoNumero(valor, tamanho)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

func TestCalcularBaseCodigoBarras(t *testing.T) {
	t.Parallel()

	valorSucesso := 72
	codigo := "701483"
	min := 2
	max := 5

	valorRetorno := CalcularBaseCodigoBarras(codigo, min, max)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

func TestCalcularBaseLinhaDigitavel(t *testing.T) {
	t.Parallel()

	valorSucesso := 76
	codigo := "287394198269"
	min := 1
	max := 2
	somarAcima := true

	valorRetorno := CalcularBaseLinhaDigitavel(codigo, min, max, somarAcima)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}

/*
	funcao ->		(valorSubtrair - (valor % mod))

	valor = 8
	valorSubtrair = 11
	mod = 3

	retorno sucesso = 9
*/
func TestCalcularMod(t *testing.T) {
	t.Parallel()
	valorSucesso := 9

	valor := 8
	valorSubtrair := 11
	mod := 3
	valorRetorno := CalcularMod(valor, valorSubtrair, mod)

	if valorSucesso != valorRetorno {
		t.Errorf(erroPadrao, valorSucesso, valorRetorno)
	}
}
