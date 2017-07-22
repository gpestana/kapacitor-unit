# Kapacitor-unit

**A test framework for TICKscripts**


**Note**: kapacitor-unit is still work in progress. Contributions and ideas
are more then welcome!

Read more about the idea and motivation behind kapacitor-unit in 
[this blog post](http://www.gpestana.com/blog/post/kapacitor-unit/)


*DRAFT*

### Running kapacitor-unit:


1) Install kapacitor-unit

```
go install github.com/gpestana/kapacitor-unit/kapacitor-unit 
````

2) Define the test configuration file (see below) 

3) Run the tests

```
kapacitor-unit --dir <*.tick directory> --kapacitor <kapacitor host> --tests <test configuration path>
```

### Test case definition:

The test cases are defined in YAML format. Each test case is defined as YAML 
mapping and it must define the data set and the expected results after kapacitor
had ran the tick script against the data set. The data set for each test case 
is defined in the Influx Line Protocol syntax.
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



### Contributions:

Fork and PR and use issues for bug reports, feature requests and general comments.

### Author:

[gpestana](http://gpestana.github.com)
