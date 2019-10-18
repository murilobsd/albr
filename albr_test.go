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
package albr

import "testing"

func TestAlertaHoje(t *testing.T) {
	uf := "sp"
	cidade := "ribeirão preto"

	err, _ := AlertaHoje(uf, cidade)

	if err != nil {
		t.Errorf("TestAlertaHoje(%s, %s), err != nil, queremos nil", uf, cidade)
	}
}

func TestValidarArgumentos(t *testing.T) {
	uf := " sp "
	cidade := "ribeirão preto"
	ufVazio := ""
	ufPequeno := "s"
	ufGrande := "sps"
	ufInvalida := "aa"
	cidadePequena := "ri"

	err, obtiveUF := validarArgumentos(uf, true)

	if err != nil {
		t.Errorf("TestvalidarArgumentos(%s, true), err != nil, queremos nil", uf)
	}

	if obtiveUF != "SP" {
		t.Errorf("TestvalidarArgumentos(%s, true), %q, queremos 'sp'", uf, obtiveUF)
	}

	err, obtiveCidade := validarArgumentos(cidade, false)

	if err != nil {
		t.Errorf("TestvalidarArgumentos(%s, false), err != nil, queremos nil", cidade)
	}

	if obtiveCidade != "Ribeirão Preto" {
		t.Errorf("TestvalidarArgumentos(%s, false), %q, queremos 'sp'", cidade, obtiveCidade)
	}

	err, _ = validarArgumentos(ufVazio, true)

	if err == nil {
		t.Errorf("TestvalidarArgumentos(%s, true), err == nil, queremos erro uf vazio", ufVazio)
	}

	err, _ = validarArgumentos(ufPequeno, true)

	if err == nil {
		t.Errorf("TestvalidarArgumentos(%s, true), err == nil, queremos erro uf pequeno", ufPequeno)
	}

	err, _ = validarArgumentos(ufGrande, true)

	if err == nil {
		t.Errorf("TestvalidarArgumentos(%s, true), err == nil, queremos erro uf grande", ufGrande)
	}

	err, _ = validarArgumentos(ufInvalida, true)

	if err == nil {
		t.Errorf("TestvalidarArgumentos(%s, true), err == nil, queremos erro uf invalida", ufInvalida)
	}

	err, _ = validarArgumentos(cidadePequena, false)

	if err == nil {
		t.Errorf("TestvalidarArgumentos(%s, false), err == nil, queremos erro cidade pequena", cidadePequena)
	}

}

func TestRequisitar(t *testing.T) {
	body, err := requisitar(URL_AL_ESTADOS)
	if err != nil {
		t.Errorf("TestRequisitar(%s), err != nil, queremos erro igual nulo", URL_AL_ESTADOS)
	}

	if len(body) == 0 {
		t.Errorf("TestRequisitar(%s), len body, queremos tamanho do body maior que 0", URL_AL_ESTADOS)
	}
}

func TestExtrairURLRelatorio(t *testing.T) {
	urls, err := extrairURLRelatorio(URL_AL_ESTADOS)
	if err != nil {
		t.Errorf("TestExtrairURLRelatorio(%s), err != nil, queremos erro igual nulo", URL_AL_ESTADOS)
	}

	if len(urls) == 0 {
		t.Errorf("TestRequisitar(%s), len urls, queremos tamanho urls maior que 0", URL_AL_ESTADOS)
	}
}

func TestExtrairRelatorio(t *testing.T) {
	extrairRelatorio("")
	/*
		urls, err := extrairURLRelatorio(URL_AL_ESTADOS)
		if err != nil {
			t.Errorf("TestExtrairURLRelatorio(%s), err != nil, queremos erro igual nulo", URL_AL_ESTADOS)
		}

		if len(urls) == 0 {
			t.Errorf("TestRequisitar(%s), len urls, queremos tamanho urls maior que 0", URL_AL_ESTADOS)
		}
	*/
}
