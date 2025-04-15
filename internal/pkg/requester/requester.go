package requester

import (
	"io"
	"net/http"
	"sync"

	"github.com/flaviojohansson/goexpert-stress-test/internal/pkg/result"
)

func RunLoadTest(url string, totalRequests, concurrency int) []result.Result {
	results := make([]result.Result, totalRequests)
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
