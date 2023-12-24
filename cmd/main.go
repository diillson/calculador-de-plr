package main

import (
	"github.com/diillson/calculador-de-plr/internal/domain"
	"github.com/diillson/calculador-de-plr/internal/finance"
	"github.com/diillson/calculador-de-plr/pkg/input"
	"github.com/diillson/calculador-de-plr/pkg/validation"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "plr_calculator",
		Short: "Calcula a PLR",
		Run:   runPLRCalculator,
	}

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func runPLRCalculator(cmd *cobra.Command, args []string) {
	plrData := domain.PLRDados{
		Salario:                 input.LeFloat64("Digite o salário: "),
		Multiplicador:           input.LeFloat64("Digite o multiplicador da PLR: "),
		PorcentagemParticipacao: input.LeFloat64("Digite a porcentagem de participação nos lucros: "),
		MesesTrabalhados:        input.LeInt("Digite o número de meses trabalhados: "),
	}

	if err := validation.ValidarPLRInput(plrData); err != nil {
		logrus.Errorf("Erro de validação: %v", err)
		return
	}

	calculadora := finance.NewCalculator()
	plr, err := calculadora.CalcularPLR(plrData)
	if err != nil {
		logrus.Errorf("Erro no cálculo da PLR: %v", err)
		return
	}

	logrus.Infof("A PLR bruta é: R$ %.2f", plr)
}
