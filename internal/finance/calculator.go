package finance

import (
	"github.com/diillson/calculador-de-plr/internal/domain"
	"github.com/sirupsen/logrus"
)

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

//func (c *Calculator) CalcularPLR(dados domain.PLRDados) (float64, error) {
//	return dados.Calcular()
//}

func (c *Calculator) CalcularPLR(dados domain.PLRDados) (float64, error) {
	plr, err := dados.Calcular()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"salario":                 dados.Salario,
			"multiplicador":           dados.Multiplicador,
			"porcentagemParticipacao": dados.PorcentagemParticipacao,
			"mesesTrabalhados":        dados.MesesTrabalhados,
		}).Errorf("Erro ao calcular PLR: %v", err)
		return 0, err
	}
	return plr, nil
}
