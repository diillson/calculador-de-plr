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
	if num < 0 {
		return "-" + formatarNumero(-num)
	}

	s := fmt.Sprintf("%.2f", num)
	partes := strings.Split(s, ".")
	inteiro := partes[0]
	decimal := partes[1]

	var resultado strings.Builder
	cont := 0

	for i := len(inteiro) - 1; i >= 0; i-- {
		if cont > 0 && cont%3 == 0 {
			resultado.WriteString(",")
		}
		resultado.WriteByte(inteiro[i])
		cont++
	}

	var resultadoFinal strings.Builder
	for i := len(resultado.String()) - 1; i >= 0; i-- {
		resultadoFinal.WriteByte(resultado.String()[i])
	}

	if len(decimal) > 0 {
		resultadoFinal.WriteString(".")
		resultadoFinal.WriteString(decimal)
	}

	return resultadoFinal.String()
}

func (c *Calculator) CalcularPLR(dados domain.PLRDados) (string, error) {
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
