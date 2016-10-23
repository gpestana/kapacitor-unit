## Kapacitor-unit

Kapacitor-unit is a framework perform sane unit tests for TICK scripts. Kapacitor-unit spins up docker containers as test environment. The test environment consists of an influxdb and kapacitor services and the test manager. The test manager is resposible for loading the tasks in kapacitor, loading the data in influxdb and compare the triggered alerts with the defined test cases.

### Running kapacitor-unit:

1) `git clone github.com/gpestana/kapacitor-unit`
2) `./kapacitor-unit/bin/kapacitor-unit --tick-dir <*.tick directory> --tests <test cases definition> --output <file> [optional]

### Test case definition:

#### Recording based test cases

The recording based test cases consist on loading a set of previously recorded data by kapacitor and replay it against the tick scripts. Kapacitor-unit loads the recordings, loads the tasks and inspects the results. At the end of the test cases, Kapacitor-unit provides a report with the test results.

Example:

```yaml
# TODO

```

#### File based test cases

The file based test cases are simple test cases in which the user defines which tick script to be tested and data that should eventually trigger kapacitor alerts.

The test cases are defined in YAML format. Each test case is defined as YAML mapping and it must define the data set and the expected results after kapacitor had ran the tick script against the data set. The data set for each testi case is defined in the Influx Line Protocol syntax.
The test case name must match the tick script that it is suppose to test.

Example:

```yaml

# Test case for alert_weather.tick
alert_weather:

  warn_trigger_test: ## Name of the test case, for report purposes
    data_set:
      - weather,location=us-midwest temperature=80
      - weather,location=us-midwest temperature=82
    expects:
      - warn: temperature>80 

  crit_trigger_test:
    data_set:
      - weather,location=us-midwest temperature=80
      - weather,location=us-midwest temperature=86
    expects:
      - crit: temperature>80 

  crit_trigger_test:
    data_set:
      - weather,location=us-midwest temperature=88
      - weather,location=us-midwest temperature=80
    expects:
      - ok: temperature>80 
```  

### Requirements:

- Docker version 1.12.1 or above
- docker-compose version 1.8.0 or above


### Contributions:

Contributions are welcome. Fork and PR and use issues for bug reports, feature requests and general comments.

### Author:

[gpestana](http://gpestana.github.com)
