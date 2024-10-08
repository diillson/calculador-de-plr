package domain

import "errors"

type PLRDados struct {
	Salario                 float64
	Multiplicador           float64
	PorcentagemParticipacao float64
	MesesTrabalhados        int
}

type FaixaIRPF struct {
	Aliquota       float64
	ParcelaDeduzir float64
}

type ResultadoIRPF struct {
	Aliquota       float64
	ParcelaDeduzir float64
	ImpostoApurado float64
	ValorLiquido   float64
}

func TabelaIRPF() []FaixaIRPF {
	return []FaixaIRPF{
		{0.0, 0.0},     // Isento
		{7.5, 169.44},  // Primeira faixa
		{15.0, 381.44}, // Segunda faixa
		{22.5, 662.77}, // Terceira faixa
		{27.5, 896.00}, // Quarta faixa
	}
}

func (p PLRDados) Calcular() (float64, error) {
	if p.Salario < 0 || p.Multiplicador < 0 || p.PorcentagemParticipacao < 0 || p.MesesTrabalhados < 0 || p.MesesTrabalhados > 12 {
		return 0, errors.New("valores inválidos para o cálculo da PLR")
	}

	plrBruta := p.Salario * p.Multiplicador * p.PorcentagemParticipacao
	return plrBruta * (float64(p.MesesTrabalhados) / 12.0), nil
}
