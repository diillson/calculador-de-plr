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
