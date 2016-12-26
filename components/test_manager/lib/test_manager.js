'use strict';

const _ 		    = require('ramda');
const fs 		    = require('./fs.js');
const Task    	= require('data.task');
const utils			= require('./utils')
const report	  = require('./report');
const kapacitor = require('./kapacitor');

const trace = utils.trace;
const convert = utils.convert
const traverse = utils.traverse

const load_test_data	= fs.load_test_data;
const load_task = kapacitor.load_task;
const delete_task = kapacitor.delete_task;
const run_test = kapacitor.run_test;

const start = () => {
	const jobs = load_test_data()
		.map(_.toPairs)
		.chain(_.map(load_task))
		.map(_.chain(run_test)) 
		.map(_.chain(delete_task))

		jobs.forEach(t => {
			t.fork(
				err => console.log(err),
				ok => { console.log(ok); console.log('--')}
			)
		});
}


module.exports.start = start;
