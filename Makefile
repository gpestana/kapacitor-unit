tests:
	go get .
	go test -cover ./...
	
start-kapacitor:
	docker-compose -f infra/docker-compose.yml up -d

sample1:
	go install .
	kapacitor-unit -dir ./sample/ -tests ./sample/test_case.yaml -stderrthreshold=INFO

