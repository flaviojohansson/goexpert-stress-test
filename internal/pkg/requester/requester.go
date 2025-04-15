package requester

import (
	"io"
	"net/http"
	"sync"

	"github.com/flaviojohansson/goexpert-stress-test/internal/pkg/result"
)

func RunLoadTest(url string, totalRequests, concurrency int) []result.Result {
	results := make([]result.Result, totalRequests)

	// TIL: channel de struct{} é chamado de Semaphore pattern, e serve para limitar a concorrência.
	//
	// struct{} é o menor tipo possível em Go (ocupa 0 bytes)
	// Usamos o channel apenas como um semáforo, não para transportar dados
	// A única informação importante é a presença/ausência de um valor no channel
	ch := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i := range totalRequests {
		wg.Add(1)
		ch <- struct{}{}

		go func(index int) {
			defer wg.Done()
			defer func() { <-ch }()

			resp, err := http.Get(url)
			if err != nil {
				results[index] = result.Result{Error: err}
				return
			}
			defer resp.Body.Close()

			// A necessidade disso só aprendi pesquisando na internet.
			// Basicamente evita vazamento de memória e simula uma requisição real.
			_, _ = io.Copy(io.Discard, resp.Body)

			results[index] = result.Result{StatusCode: resp.StatusCode}
		}(i)
	}

	wg.Wait()
	return results
}
