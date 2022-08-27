


benchmark:
	go test -benchmem -run=^# -benchtime=1s -bench=. ./benchmarks > ./reports/latest.txt

build:
	go build ./...

test:
	go test ./...

generate_reports:
	go run ./cmd/reports/main.go