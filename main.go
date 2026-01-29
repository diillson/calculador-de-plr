package main

import (
	"github.com/diillson/calculador-de-plr/internal/domain"
	"github.com/diillson/calculador-de-plr/internal/finance"
	"github.com/diillson/calculador-de-plr/pkg/input"
	"github.com/diillson/calculador-de-plr/pkg/validation"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var salario float64
var multiplicador float64
var porcentagemParticipacao float64
var mesesTrabalhados int
var periodo string

func main() {
	var rootCmd = &cobra.Command{
		Use:   "plr_calculator",
		Short: "Calcula a PLR",
		Run:   runPLRCalculator,
	}

	rootCmd.Flags().Float64VarP(&salario, "salario", "s", 0.0, "Salário do funcionário")
	rootCmd.Flags().Float64VarP(&multiplicador, "multiplicador", "m", 0.0, "Multiplicador da PLR")
	rootCmd.Flags().Float64VarP(&porcentagemParticipacao, "participacao", "p", 0.0, "Porcentagem de participação nos lucros")
	rootCmd.Flags().IntVarP(&mesesTrabalhados, "meses", "t", 0, "Número de meses trabalhados")
	rootCmd.Flags().StringVar(&periodo, "periodo", "maio-mais", "Período da tabela IRPF: jan-abr ou maio-mais")

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func runPLRCalculator(cmd *cobra.Command, args []string) {
	cmd.Flags().Parse(args)
	salarioFlag, _ := cmd.Flags().GetFloat64("salario")
	multiplicadorFlag, _ := cmd.Flags().GetFloat64("multiplicador")
	porcentagemParticipacaoFlag, _ := cmd.Flags().GetFloat64("participacao")
	mesesTrabalhadosFlag, _ := cmd.Flags().GetInt("meses")
	periodoFlag, _ := cmd.Flags().GetString("periodo")

	if salarioFlag == 0 && multiplicadorFlag == 0 && porcentagemParticipacaoFlag == 0 && mesesTrabalhadosFlag == 0 {
		salarioFlag = input.LeFloat64("Digite o salário: ")
		multiplicadorFlag = input.LeFloat64("Digite o multiplicador da PLR: ")
		porcentagemParticipacaoFlag = input.LeFloat64("Digite a porcentagem de participação nos lucros: ")
		mesesTrabalhadosFlag = input.LeInt("Digite o número de meses trabalhados: ")
		logrus.Info("Escolha o período da tabela IRPF:")
		logrus.Info("1) Janeiro a Abril/2025")
		logrus.Info("2) A partir de Maio/2025 (padrão)")
		escolha := input.LeInt("Digite sua escolha (1 ou 2): ")
		if escolha == 1 {
			periodoFlag = "jan-abr"
		} else {
			periodoFlag = "maio-mais"
		}
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

	plrBruta, err := strconv.ParseFloat(strings.ReplaceAll(plrFormatada, ",", ""), 64)
	if err != nil {
		logrus.Errorf("Erro ao converter PLR formatada: %v", err)
		return
	}

	var tabelaIRPF []domain.FaixaIRPF
	if periodoFlag == "jan-abr" {
		tabelaIRPF = domain.TabelaIRPFPLRJanAbr2025()
		logrus.Info("Usando tabela IRPF: Janeiro a Abril/2025")
	} else {
		tabelaIRPF = domain.TabelaIRPFPLRAPartirDeMaio2025()
		logrus.Info("Usando tabela IRPF: A partir de Maio/2025")
	}

	resultadoIRPF, err := calculadora.CalcularIRPF(plrBruta, tabelaIRPF)
	if err != nil {
		logrus.Errorf("Erro ao calcular o IRPF: %v", err)
		return
	}

	logrus.Infof("A PLR bruta é: R$ %s", plrFormatada)
	logrus.Infof("Alíquota: %.2f%%", resultadoIRPF.Aliquota)
	logrus.Infof("Parcela a Deduzir: R$ %.2f", resultadoIRPF.ParcelaDeduzir)
	logrus.Infof("Imposto de Renda Apurado: R$ %.2f", resultadoIRPF.ImpostoApurado)
	logrus.Infof("Valor Líquido da PLR: R$ %.2f", resultadoIRPF.ValorLiquido)
}
