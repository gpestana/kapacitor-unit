'use strict';

const request 			= require('request');
const Task				= require('data.task');
const bunyan			= require('bunyan');
const test_manager 		= require('./lib/test_manager.js');
const report			= require('./lib/report.js');

const DEFAULT_KAPACITOR = 'http://kapacitor:9092/kapacitor/v1/ping';
const log = bunyan.createLogger({name: 'kapacitor-unit'});

const waitForKapacitor = () => {
	request(DEFAULT_KAPACITOR, (err, res, body) => {
		if(!err && res.statusCode == 204) {
			clearInterval(intervalId);
			log.info('=> Kapacitor API ready. Starting test manager.');

			test_manager.start(err => {
				if(!err) {
					report.prints_report(log);
				} else log.info({err: err}, 'waitForKapacitor');
			});

		} else { 
			log.info('=> Kapacitor API not up yet, trying again in 3s.');
		}
	});
}

const intervalId = setInterval(waitForKapacitor, 3000);
