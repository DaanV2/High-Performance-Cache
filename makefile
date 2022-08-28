


benchmark-latest:
	go test -benchmem -run=^# -benchtime=1s -bench=. ./benchmarks > ./reports/latest.txt

benchmark:
	go test -benchmem -run=^# -benchtime=1s -bench=. ./benchmarks

build:
	go build ./...

test:
	go test ./...

generate-reports:
	go run ./cmd/reports/main.go