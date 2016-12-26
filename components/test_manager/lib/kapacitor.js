'use strict';

const _       = require('ramda');
const Task    = require('data.task');
const request = require('request');

const KAPACITOR_HOST = 'http://kapacitor:9092'

const load_task = test_def => {
    return new Task((reject, resolve) => {
        const req_url = `${KAPACITOR_HOST}/kapacitor/v1/tasks`;
        const success_http_codes = [200, 204];

        if(!test_def) return reject('Test definition not provided');
        let task = _.pick(['id', 'type', 'dbrps', 'script'], test_def[1]);
        request.post({
            url: req_url,
            body: task,
            json: true
        }, (err, res, body) => {
            if(err || !_.contains(res.statusCode, success_http_codes)) {
                return reject(err || res.body);
            } else {
                return resolve(task);
            }
        });
    });
}

const delete_task = task => {
    return new Task((reject, resolve) => {
        const req_url = `${KAPACITOR_HOST}/kapacitor/v1/tasks/${task.id}`;
        const success_http_codes = [200, 204];
        request.delete(req_url, (err, res, body) => {
            if(err || !_.contains(res.statusCode, success_http_codes)) {
                return reject(err || body);
            } else return resolve(task);
        });        
    })
};

const run_test = task => {
    console.log('load_task')
    return new Task((reject, resolve) => {
        let err;
        if(err) return reject(err)
        else return resolve(task);
    });
}

module.exports = {
	load_task: load_task,
    delete_task: delete_task,
    run_test: run_test,
}
