package result

// Result representa o resultado de uma requisição HTTP
type Result struct {
	StatusCode int
	Error      error
}
