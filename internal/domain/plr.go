package domain

import "errors"

type PLRDados struct {
	Salario                 float64
	Multiplicador           float64
	PorcentagemParticipacao float64
	MesesTrabalhados        int
}

func (p PLRDados) Calcular() (float64, error) {
	if p.Salario < 0 || p.Multiplicador < 0 || p.PorcentagemParticipacao < 0 || p.MesesTrabalhados < 0 || p.MesesTrabalhados > 12 {
		return 0, errors.New("valores inválidos para o cálculo da PLR")
	}

	plrBruta := p.Salario * p.Multiplicador * p.PorcentagemParticipacao
	return plrBruta * (float64(p.MesesTrabalhados) / 12.0), nil
}
