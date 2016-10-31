'use strict';

const _ 		    = require('ramda');
const fs 		    = require('./fs.js');
const report	  = require('./report.js');
const kapacitor = require('./kapacitor.js');

const load_test_data	= fs.load_test_data;
const load_task = kapacitor.load_task;
const run_test = kapacitor.run_test;


const start = cb => {
	const t = load_test_data()
		.chain(t => {return _.map(load_task, t.getOrElse())})

		for(let task in t) {
			t[task].fork(
				err => {console.log(err)},
				ok => {console.log(`ok: ${ok}`);}
			)
		}
};


module.exports.start = start;
