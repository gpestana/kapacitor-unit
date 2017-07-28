start-kapacitor:
	docker-compose -f infra/docker-compose.yml up -d

run-sample:
	go install .
	kapacitor-unit -dir ./sample/ -tests ./sample/test_case.yaml

tests:
	 go test -cover ./...
	
