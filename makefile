


benchmark-latest:
	go test -benchmem -benchtime=5s -run=^$ -bench ^Benchmark github.com/DaanV2/High-Performance-Cache/benchmarks > ./reports/latest.txt

benchmark:
	go test -benchmem -benchtime=5s -run=^$ -bench ^Benchmark github.com/DaanV2/High-Performance-Cache/benchmarks

build:
	go build ./...

test:
	go test ./...

generate-reports:
	go run ./cmd/reports/main.go