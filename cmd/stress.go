package cmd

import (
	"log"
	"time"

	"github.com/flaviojohansson/goexpert-stress-test/internal/pkg/report"
	"github.com/flaviojohansson/goexpert-stress-test/internal/pkg/requester"

	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var runCmd = &cobra.Command{
	Use:   "stress",
	Short: "Executa o teste de carga",
	Long:  `Executa um teste de carga contra a URL especificada com os parâmetros fornecidos.`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()

		// Validação dos parâmetros
		if url == "" {
			log.Fatal("URL não pode ser vazia")
		}
		if requests <= 0 {
			log.Fatal("Número de requests deve ser maior que zero")
		}
		if concurrency <= 0 {
			log.Fatal("Concorrência deve ser maior que zero")
		}

		// Executa os requests
		results := requester.RunLoadTest(url, requests, concurrency)

		// Gera o relatório
		report.GenerateReport(results, time.Since(startTime))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado (obrigatório)")
	runCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Número total de requests")
	runCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Número de chamadas simultâneas")

	runCmd.MarkFlagRequired("url")
}
