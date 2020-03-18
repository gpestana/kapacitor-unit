tests:
	go test -cover ./cmd/kapacitor-unit ./io ./task ./test

setup:
	go get ./cmd/kapacitor-unit ./io ./task ./test

install:
	go install ./cmd/kapacitor-unit

build-cmd:
	go build cmd/kapacitor-unit/main.go 

run:
	./main

start-kapacitor:
	docker-compose -f infra/docker-compose.yml up -d

sample1:
	go run cmd/kapacitor-unit/main.go -dir ./sample/tick_scripts -tests ./sample/test_cases/test_case.yaml

sample1_debug:
	go run cmd/kapacitor-unit/main.go -dir ./sample/tick_scripts -tests ./sample/test_cases/test_case.yaml -stderrthreshold=INFO

sample1_batch:
	go run cmd/kapacitor-unit/main.go -dir ./sample/tick_scripts -tests ./sample/test_cases/test_case_batch.yaml

sample1_batch_debug:
	go run cmd/kapacitor-unit/main.go -dir ./sample/tick_scripts -tests ./sample/test_cases/test_case_batch.yaml -stderrthreshold=INFO

sample_dir:
	go run cmd/kapacitor-unit/main.go -dir ./sample/tick_scripts -tests ./sample/test_cases
