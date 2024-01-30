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
	// Formatar o número com duas casas decimais
	s := fmt.Sprintf("%.2f", num)

	// Dividir em partes inteira e decimal
	partes := strings.Split(s, ".")
	inteiro := partes[0]
	decimal := partes[1]

	// Inverter a string do inteiro para facilitar a inserção das vírgulas
	inteiroInvertido := reverseString(inteiro)

	// Adicionar as vírgulas
	var comVirgulas strings.Builder
	for i, char := range inteiroInvertido {
		if i > 0 && i%3 == 0 {
			comVirgulas.WriteString(",")
		}
		comVirgulas.WriteRune(char)
	}

	// Inverter novamente para a ordem correta
	inteiroFormatado := reverseString(comVirgulas.String())

	// Juntar a parte inteira e decimal
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
	// Ajustar a porcentagem de participação se for maior que 1 (como 83 em vez de 0.83)
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
	// Formatação do valor da PLR para melhor legibilidade
	plrFormatado := formatarNumero(plr)
	return plrFormatado, nil
}

func determinarFaixaIRPF(valor float64, tabela []domain.FaixaIRPF) int {
	// Esta é uma função simplificada.
	// A lógica exata depende das faixas de valores da tabela do IRPF.
	// Você deve implementar a lógica de acordo com as faixas atuais.
	if valor <= 7407.11 {
		return 0
	} else if valor <= 9922.28 {
		return 1
	} else if valor <= 13167.00 {
		return 2
	} else if valor <= 16380.38 {
		return 3
	} else {
		return 4
	}
}

func calcularImposto(baseCalculo float64, faixa domain.FaixaIRPF) float64 {
	return (baseCalculo * faixa.Aliquota / 100) - faixa.ParcelaDeduzir
}

func (c *Calculator) CalcularIRPF(plr float64, tabela []domain.FaixaIRPF) (*domain.ResultadoIRPF, error) {
	// Defina a tabela IRPF aqui
	tabelaIRPF := domain.TabelaIRPF()

	// Determinar a faixa de IRPF
	faixa := determinarFaixaIRPF(plr, tabelaIRPF)
	if faixa < 0 || faixa >= len(tabelaIRPF) {
		return nil, fmt.Errorf("Erro ao determinar a faixa de IRPF")
	}

	// Calcular o imposto
	impostoApurado := calcularImposto(plr, tabelaIRPF[faixa])
	valorLiquido := plr - impostoApurado

	return &domain.ResultadoIRPF{
		Aliquota:       tabelaIRPF[faixa].Aliquota,
		ParcelaDeduzir: tabelaIRPF[faixa].ParcelaDeduzir,
		ImpostoApurado: impostoApurado,
		ValorLiquido:   valorLiquido,
	}, nil
}
