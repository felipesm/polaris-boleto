# Polaris Boleto

Biblioteca desenvolvida em Golang para geração de informações como código de barras e linha digitável.

A aplicação permite a geração de dados como código barras e linha digitável para os bancos **Bradesco** e **Santander**. Ela foi desenvolvida para ser utilizada como dependência em diferentes aplicações. Atualmente é usada na [`polaris-api`](https://github.com/felipesm/polaris-api), API em Golang que atende requisições para geração do código de barras e linha digitável.

### Download e Iniciando

As instruções abaixo possibilitará você fazer o download e configuração do projeto para sua máquina para fins de estudo, desenvolvimento, testes.

Fazer o clone do projeto do github:
> git clone https://github.com/felipesm/polaris-boleto.git

No VSCode, acessar o menu `Terminal`, `New Terminal` e na aba terminal digitar:
> go run boleto-run-t.go
 
**Obs**: A classe acima `boleto-run-t.go` trata-se apenas de uma classe criada para testar a geração do código de barras e linha digitável, validando que as informações retornadas estão de acordo.

### Utilizando a Biblioteca

Para utilização da biblioteca, seguir as etapas descritas abaixo:

- Instanciar o banco através do método `InstanciarBoleto(string)`
- Chamar o método `SetCodigo()`
- Chamar o método `SetAgencia(integer)`
  - Obs: para o Santander, não é necessário chamar esse método
- Chamar o método `SetCarteira(integer)`
- Chamar o método `SetFatorVencimento(string, bool)`
  - Obs: se deseja que o fator vencimento seja zerado, informar true no segundo argumento do método 
- Chamar o método `SetValorBoleto(float, bool)`
  - Obs: se deseja que o valor do boleto seja zerado, informar true no segundo argumento do método 
- Chamar o método `SetNossoNumero(string)`
- Chamar o método `SetCodigoBeneficiario(integer)`

> ***Observar o arquivo `boleto-run-t.go`, nele existem exemplos de como devem ser feitas essas chamadas dos métodos***.

### Geração do Código de Barras

Chamar o método `GetCodigoBarras()`. Esse método retorna um struct do tipo `CodigoBarras` com os seguintes campos:

  | Nome  |  Tipo de dado  |
  | :---: | :---: |
  |  CodigoBanco |  string |
  |  CodigoMoeda |  string |
  |  DV |  string |
  |  FatorVencimento |  string |
  |  Valor |  string |
  |  CampoLivre |  string |
  |  CodigoBarrasCompleto |  string |
    
### Geração da Linha Digitável

Chamar o método `GetLinhaDigitavel(string)`. O argumento desse método deve ser uma string com o código de barras completo. Essa informação do código de barras completo pode ser obtido através do campo `CodigoBarrasCompleto` citado acima. O método de geração da linha digitável também retorna um struct do tipo `LinhaDigitavel` com os seguintes campos:

  | Nome  |  Tipo de dado  |
  | :---: | :---: |
  |  LinhaDigitavel |  string |
  |  LinhaDigitavelFormatada |  string |
