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

var salario float64
var multiplicador float64
var porcentagemParticipacao float64
var mesesTrabalhados int

func main() {
	var rootCmd = &cobra.Command{
		Use:   "plr_calculator",
		Short: "Calcula a PLR",
		Run:   runPLRCalculator,
	}

	// Definindo as flags
	rootCmd.Flags().Float64VarP(&salario, "salario", "s", 0.0, "Salário do funcionário")
	rootCmd.Flags().Float64VarP(&multiplicador, "multiplicador", "m", 0.0, "Multiplicador da PLR")
	rootCmd.Flags().Float64VarP(&porcentagemParticipacao, "participacao", "p", 0.0, "Porcentagem de participação nos lucros")
	rootCmd.Flags().IntVarP(&mesesTrabalhados, "meses", "t", 0, "Número de meses trabalhados")

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func runPLRCalculator(cmd *cobra.Command, args []string) {
	// Verificar se as flags foram fornecidas
	cmd.Flags().Parse(args)
	salarioFlag, _ := cmd.Flags().GetFloat64("salario")
	multiplicadorFlag, _ := cmd.Flags().GetFloat64("multiplicador")
	porcentagemParticipacaoFlag, _ := cmd.Flags().GetFloat64("participacao")
	mesesTrabalhadosFlag, _ := cmd.Flags().GetInt("meses")

	// Se nenhuma flag foi fornecida, usar o prompt interativo
	if salarioFlag == 0 && multiplicadorFlag == 0 && porcentagemParticipacaoFlag == 0 && mesesTrabalhadosFlag == 0 {
		salarioFlag = input.LeFloat64("Digite o salário: ")
		multiplicadorFlag = input.LeFloat64("Digite o multiplicador da PLR: ")
		porcentagemParticipacaoFlag = input.LeFloat64("Digite a porcentagem de participação nos lucros: ")
		mesesTrabalhadosFlag = input.LeInt("Digite o número de meses trabalhados: ")
	}

	plrData := domain.PLRDados{
		Salario:                 salarioFlag,
		Multiplicador:           multiplicadorFlag,
		PorcentagemParticipacao: porcentagemParticipacaoFlag,
		MesesTrabalhados:        mesesTrabalhadosFlag,
	}

	if err := validation.ValidarPLRInput(plrData); err != nil {
		logrus.Errorf("Erro de validação: %v", err)
		return
	}

	calculadora := finance.NewCalculator()
	plrFormatada, err := calculadora.CalcularPLR(plrData)
	if err != nil {
		logrus.Errorf("Erro no cálculo da PLR: %v", err)
		return
	}

	logrus.Infof("A PLR bruta é: R$ %s", plrFormatada)
}
