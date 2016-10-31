'use strict';

const request = require('request');
const Task    = require('data.task');
const _				= require('ramda');

const KAPACITOR_HOST = 'http://kapacitor:9092'


const load_task = test_def => {
	return new Task((reject, resolve) => {
  	const req_url = `${KAPACITOR_HOST}/kapacitor/v1/tasks`;
    const success_http_codes = [200, 204];

    if(!test_def) return reject('err');

	let t = _.pick(['id', 'type', 'dbrps', 'script'], test_def);
	const t2 = {"id" : "TASK_ID","type" : "stream","dbrps": [{"db": "DATABASE_NAME", "rp" : "RP_NAME"}],"script": "stream\n|from()\n.measurement('cpu')\n"}

    request.post({
        url: req_url,
        body: t,
        json: true
    }, (err, res, body) => {
    		console.log(t)
        if(err || !_.contains(res.statusCode, success_http_codes)) {
            reject(err || res.body);
        } else resolve(res.statusCode);
    });
	});
}

const run_test = id => {
	return new Task((reject, resolve) => {
		const task_id = id[0];
		console.log(task_id);
	});
}

module.exports = {
	load_task: load_task,
	run_test: {}
}
