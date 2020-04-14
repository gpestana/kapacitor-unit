# Kapacitor-unit

**A test framework for TICKscripts**

[![Build Status](https://travis-ci.org/gpestana/kapacitor-unit.svg?branch=master)](https://travis-ci.org/gpestana/kapacitor-unit) ![Release Version](https://img.shields.io/badge/release-0.8-blue.svg)


Kapacitor-unit is a testing framework to make TICK scripts testing easy and
automated. Testing with Kapacitor-unit is as easy as defining the test configuration saying which alerts are expected to trigger when the TICK script processes specific data. 


Read more about the idea and motivation behind kapacitor-unit in 
[this blog post](http://www.gpestana.com/blog/kapacitor-unit/)


## Show me Kapacitor-unit in action!
![usage-example](https://media.giphy.com/media/xT0xetJEkloDtbVHSU/giphy.gif)


## Features

:heavy_check_mark: Run tests for **stream** TICK scripts using protocol line data input 

:heavy_check_mark: Run tests for **batch** TICK scripts using protocol line data input 

:soon: Run tests for **stream** and **batch** TICK scripts using recordings 


## Requirements:

In order for all features to be supported, the Kapacitor version running the tests must be v1.3.4 or higher.

## Running kapacitor-unit:


1) Install kapacitor-unit and run

```
 $ make install
 $ make build-cmd

 $ make run  	# same as ./cmd/kapacitor-unit/kapacitor-unit
```

You can add `./cmd/kapacitor-unit/kapacitor-unit` to your `$PATH` so you can
easily call the `kapacitor-unit` executable  anywhere.

2) Define the test configuration file (see below) 

3) Run the tests

```
kapacitor-unit --dir <*.tick directory> --kapacitor <kapacitor host> --influxdb <influxdb host> --tests <test configuration path>
```

### Test case definition:

```yaml

# Test case for alert_weather.tick
tests:
  
   # This is the configuration for a test case. The 'name' must be unique in the
   # same test configuration. 'description' is optional

  - name: Alert weather:: warning
    description: Task should trigger Warning when temperature raises about 80 

    # 'task_name' defines the name of the file of the tick script to be loaded
    # when running the test
    task_name: alert_weather.tick

    db: weather
    rp: default 
    type: stream

     # 'data' is an array of data in the line protocol
    data:
      - weather,location=us-midwest temperature=75
      - weather,location=us-midwest temperature=82

    # Alert that should be triggered by Kapacitor when test data is running 
    # against the task
    expects:
      ok: 0
      warn: 1
      crit: 0


  - name: Alert no. 2 using recording
    task_id: alert_weather.tick
    db: weather
    rp: default 
    type: stream
    recordind_id: 7c581a06-769d-45cb-97fe-a3c4d7ba061a
    expects:
      ok: 0
      warn: 1
      crit: 0


  - name: Alert no. 3 - Batch
    task_id: alert_weather.tick
    db: weather
    rp: default 
    type: batch
    data:
      - weather,location=us-midwest temperature=80
      - weather,location=us-midwest temperature=82
    expects:
      ok: 0
      warn: 1
      crit: 0

```  

## Contributions:

Fork and PR and use issues for bug reports, feature requests and general comments.

:copyright: MIT
