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

// Tabela de IRPF especifica de PLR (valores anuais).

// De janeiro a abril de 2025
func TabelaIRPFPLRJanAbr2025() []FaixaIRPF {
	return []FaixaIRPF{
		{0.0, 7640.80, 0.0, 0.0},
		{7640.81, 9922.28, 7.5, 573.06},
		{9922.29, 13167.00, 15.0, 1317.23},
		{13167.01, 16380.38, 22.5, 2304.76},
		{16380.39, 999999999.99, 27.5, 3123.78},
	}
}

// A partir de maio de 2025
func TabelaIRPFPLRAPartirDeMaio2025() []FaixaIRPF {
	return []FaixaIRPF{
		{0.0, 8214.40, 0.0, 0.0},
		{8214.41, 9922.28, 7.5, 616.08},
		{9922.29, 13167.00, 15.0, 1360.25},
		{13167.01, 16380.38, 22.5, 2347.78},
		{16380.39, 999999999.99, 27.5, 3166.80},
	}
}

// TabelaIRPFPLR retorna por padrao a tabela mais recente (a partir de maio/2025).
func TabelaIRPFPLR() []FaixaIRPF {
	return TabelaIRPFPLRAPartirDeMaio2025()
}

func (p PLRDados) Calcular() (float64, error) {
	if p.Salario < 0 || p.Multiplicador < 0 || p.PorcentagemParticipacao < 0 || p.MesesTrabalhados < 0 || p.MesesTrabalhados > 12 {
		return 0, errors.New("valores invalidos para o calculo da PLR")
	}

	plrBruta := p.Salario * p.Multiplicador * p.PorcentagemParticipacao
	return plrBruta * (float64(p.MesesTrabalhados) / 12.0), nil
}
