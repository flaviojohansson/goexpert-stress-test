package report

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flaviojohansson/goexpert-stress-test/internal/pkg/result"
)

func GenerateReport(results []result.Result, duration time.Duration) {
	statusCount := make(map[int]int)
	totalRequests := len(results)
	successCount := 0

	for _, result := range results {
		if result.Error == nil {
			statusCount[result.StatusCode]++
			if result.StatusCode == http.StatusOK {
				successCount++
			}
		} else {
			statusCount[0]++ // 0 representa erros de requisição
		}
	}

	fmt.Println("\n=== Relatório do Teste de Carga ===")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Requests bem-sucedidos (Status 200): %d (%.2f%%)\n", successCount, float64(successCount)/float64(totalRequests)*100)

	fmt.Println("\nDistribuição de Status HTTP:")
	for status, count := range statusCount {
		if status == 0 {
			fmt.Printf("Erros de requisição: %d (%.2f%%)\n", count, float64(count)/float64(totalRequests)*100)
		} else {
			fmt.Printf("Status %d: %d (%.2f%%)\n", status, count, float64(count)/float64(totalRequests)*100)
		}
	}

	fmt.Println("\nRequisições por segundo:", float64(totalRequests)/duration.Seconds())
}
