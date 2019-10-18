/*
 * Copyright (c) 2019 Murilo Ijanc' <mbsd@m0x.ru>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
// Package albr possui funções que checam alertas meteorológicos do
// website Alert-AS (http://alert-as.inmet.gov.br), os alertas são do
// dia em que foi realizado a consulta ou alertas futuros.
package albr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
  "C"
)

const (
	// URL da listagem dos alertas + estados
	URL_AL_ESTADOS = "http://alert-as.inmet.gov.br/cv/"
	// Regex
	RE_ESTADOS = "estados_quadro2_barra_cinza2"
)

var (
	// Erro caso os argumentos sejam strings em branco.
	ERR_TAMANHO_ARGUMENTOS = "os argumentos não podem ser embrancos"
	// Erro tamanho da Unidade Federativa.
	ERR_TAMANHO_UF = "unidade federativa deve possuir 2 letras"
	// Erro tamanho da Unidade Federativa.
	ERR_INVALIDA_UF = "unidade federativa inválida"
	// Erro não existe as urls do relatório.
	ERR_NAOEXISTE_URLS = "urls dos relatórios não encontradas"
	// Map das unidades federativas
	unidadesFederativas = map[string]string{
		"AC": "AC",
		"AL": "AL",
		"AM": "AM",
		"AP": "AP",
		"BA": "BA",
		"CE": "CE",
		"DF": "DF",
		"ES": "ES",
		"GO": "GO",
		"MA": "MA",
		"MG": "MG",
		"MS": "MS",
		"MT": "MT",
		"PA": "PA",
		"PB": "PB",
		"PE": "PE",
		"PI": "PI",
		"PR": "PR",
		"RJ": "RJ",
		"RN": "RN",
		"RO": "RO",
		"RR": "RR",
		"RS": "RS",
		"SC": "SC",
		"SE": "SE",
		"SP": "SP",
		"TO": "TO",
	}
)
// parameter
type parametro struct {
	XMLName xml.Name `xml:"parameter"`
  Name string `xml:"valueName"`
  Value string `xml:"value"`
}


// info estrut
type info struct {
	XMLName xml.Name `xml:"info"`
	Event   string   `xml:"event"`
  Parametros []parametro  `xml:"parameter"`
}

// Relatorio XML
type alert struct {
	XMLName xml.Name `xml:"alert"`
	Info    info     `xml:"info"`
}


// AlertaHoje retorna alerta do dia caso ele exista na região que foi
// passada nos argumentos dessa função em caso negativo uma string
// dizendo que não existem alertas.
func AlertaHoje(uf, cidade string) (error, string) {
	err, uf := validarArgumentos(uf, true)
	if err != nil {
		return err, ""
	}

	err, cidade = validarArgumentos(cidade, false)
	if err != nil {
		return err, ""
	}

	return nil, ""
}

// request realiza a requisição.
func requisitar(url string) (string, error) {
	var err error
	var body []byte
	if url != "" {
		body, err = ioutil.ReadFile("data/lista_estados.html")
	} else {
		body, err = ioutil.ReadFile("data/relatorio.xml")
	}
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// extrairRelatorio ...
func extrairRelatorio(url string) {
	body, err := requisitar(url)
	r := alert{}
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", r.XMLName)
	fmt.Printf("Name: %q\n", r.Info.Parametros)
}

// extrairURLRelatorio se existe algum alerta para o estado e em caso
// positivo retorna url
func extrairURLRelatorio(url string) ([]string, error) {
	var urls []string
	body, err := requisitar(url)
	if err != nil {
		return []string{}, nil
	}
	re := regexp.MustCompile(`http:\/\/alert-as\.inmet\.gov\.br\/.*\.xml`)
	urls = re.FindAllString(body, -1)
	if len(urls) == 0 {
		return urls, errors.New(ERR_NAOEXISTE_URLS)
	}
	fmt.Printf("%q\n", re.FindAllString(body, -1))
	return urls, nil
}

// ufValido simplesmente checa na estrutrua map se a sigla é realmente
// de um estado brasileiro.
func ufValido(uf string) bool {
	if _, valido := unidadesFederativas[uf]; valido {
		return true
	}
	return false
}

// validarArgumentos função para retornar o argumento validado para
// passarmos a filtrar as informações da requisição.
func validarArgumentos(arg string, uf bool) (error, string) {
	// pelo menos 2 caracteres.
	if len(arg) < 2 {
		return errors.New(ERR_TAMANHO_ARGUMENTOS), ""
	}

	arg = strings.TrimSpace(arg) // " sp " -> "sp"

	if uf && len(arg) == 2 {
		arg = strings.ToUpper(arg) // uf -> UF
		if !ufValido(arg) {
			return errors.New(ERR_INVALIDA_UF), ""
		}
	} else if !uf && len(arg) > 2 {
		arg = strings.Title(arg) // ribeirão preto -> Ribeirão Preto
	} else {
		return errors.New(ERR_TAMANHO_UF), ""
	}

	return nil, arg
}
