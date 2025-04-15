# goexpert-stress-test
Pós Go Expert - Desafio Técnico de Conclusão - Stress Test
## Testando a aplicação
1. Clonar o repositório
```
git clone https://github.com/flaviojohansson/goexpert-stress-test
cd goexpert-stress-test
```
2. Criar a imagem do container
```
docker build -t goexpert-stress-test .
```
3. Executar a aplicação com uma chamada via docker
```
docker run goexpert-stress-test stress --url=https://tms.ssw.inf.br --requests=100 --concurrency=10

# https://tms.ssw.inf.br é um bom exemplo pois tem controle de requests por segundo,
# e retornará além de status 200, 429 (Too Many Requests).
```
