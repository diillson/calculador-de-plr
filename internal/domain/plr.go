package domain

import "errors"

type PLRDados struct {
	Salario                 float64
	Multiplicador           float64
	PorcentagemParticipacao float64
	MesesTrabalhados        int
}

type FaixaIRPF struct {
	LimiteInferior float64
	LimiteSuperior float64
	Aliquota       float64
	ParcelaDeduzir float64
}

type ResultadoIRPF struct {
	Aliquota       float64
	ParcelaDeduzir float64
	ImpostoApurado float64
	ValorLiquido   float64
}

// TabelaIRPF retorna a tabela de IRPF atualizada para 2024/2025 (valores anuais)
func TabelaIRPF() []FaixaIRPF {
	return []FaixaIRPF{
		{0.0, 28467.20, 0.0, 0.0},                // Isento
		{28467.21, 33919.80, 7.5, 2135.04},       // Primeira faixa
		{33919.81, 45012.60, 15.0, 4679.03},      // Segunda faixa
		{45012.61, 55976.16, 22.5, 8054.97},      // Terceira faixa
		{55976.17, 999999999.99, 27.5, 10853.78}, // Quarta faixa
	}
}

func (p PLRDados) Calcular() (float64, error) {
	if p.Salario < 0 || p.Multiplicador < 0 || p.PorcentagemParticipacao < 0 || p.MesesTrabalhados < 0 || p.MesesTrabalhados > 12 {
		return 0, errors.New("valores inválidos para o cálculo da PLR")
	}

	plrBruta := p.Salario * p.Multiplicador * p.PorcentagemParticipacao
	return plrBruta * (float64(p.MesesTrabalhados) / 12.0), nil
}
