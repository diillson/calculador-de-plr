package finance

import (
	"fmt"
	"github.com/diillson/calculador-de-plr/internal/domain"
	"github.com/sirupsen/logrus"
	"strings"
)

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func formatarNumero(num float64) string {
	s := fmt.Sprintf("%.2f", num)
	partes := strings.Split(s, ".")
	inteiro := partes[0]
	decimal := partes[1]
	inteiroInvertido := reverseString(inteiro)
	var comVirgulas strings.Builder
	for i, char := range inteiroInvertido {
		if i > 0 && i%3 == 0 {
			comVirgulas.WriteString(",")
		}
		comVirgulas.WriteRune(char)
	}
	inteiroFormatado := reverseString(comVirgulas.String())
	return inteiroFormatado + "." + decimal
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (c *Calculator) CalcularPLR(dados domain.PLRDados) (string, error) {
	if dados.PorcentagemParticipacao > 1 {
		dados.PorcentagemParticipacao /= 100
	}
	plr, err := dados.Calcular()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"salario":                 dados.Salario,
			"multiplicador":           dados.Multiplicador,
			"porcentagemParticipacao": dados.PorcentagemParticipacao,
			"mesesTrabalhados":        dados.MesesTrabalhados,
		}).Errorf("Erro ao calcular PLR: %v", err)
		return "", err
	}
	plrFormatado := formatarNumero(plr)
	return plrFormatado, nil
}

func determinarFaixaIRPF(valorAnual float64, tabela []domain.FaixaIRPF) int {
	for i, faixa := range tabela {
		if valorAnual >= faixa.LimiteInferior && valorAnual <= faixa.LimiteSuperior {
			return i
		}
	}
	return len(tabela) - 1
}

func calcularImposto(baseCalculo float64, faixa domain.FaixaIRPF) float64 {
	return (baseCalculo * faixa.Aliquota / 100) - faixa.ParcelaDeduzir
}

func (c *Calculator) CalcularIRPF(plr float64, tabela []domain.FaixaIRPF) (*domain.ResultadoIRPF, error) {
	faixaIndex := determinarFaixaIRPF(plr, tabela)
	if faixaIndex < 0 || faixaIndex >= len(tabela) {
		return nil, fmt.Errorf("erro ao determinar a faixa de IRPF para o valor: R$ %.2f", plr)
	}
	faixa := tabela[faixaIndex]
	impostoApurado := calcularImposto(plr, faixa)
	if impostoApurado < 0 {
		impostoApurado = 0
	}
	valorLiquido := plr - impostoApurado
	logrus.WithFields(logrus.Fields{
		"plrBruta":       plr,
		"faixaIndex":     faixaIndex,
		"aliquota":       faixa.Aliquota,
		"parcelaDeduzir": faixa.ParcelaDeduzir,
		"impostoApurado": impostoApurado,
		"valorLiquido":   valorLiquido,
	}).Debug("Calculo de IRPF realizado")
	return &domain.ResultadoIRPF{
		Aliquota:       faixa.Aliquota,
		ParcelaDeduzir: faixa.ParcelaDeduzir,
		ImpostoApurado: impostoApurado,
		ValorLiquido:   valorLiquido,
	}, nil
}
