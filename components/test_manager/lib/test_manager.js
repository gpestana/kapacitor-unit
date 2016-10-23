'use strict';

const request 	= require('request');
const Task 		= require('data.task')
const _			= require('ramda');
const fs 		= require('./fs.js');
const report	= require('./report.js');
const kapacitor = require('./kapacitor.js');

const load_test_data	= fs.load_test_data;
const KAPACITOR_HOST 	= 'http://kapacitor:9092'


const start = cb => { // cb expects err or nothing
	const test_data = load_test_data();
	console.log(test_data.merge());
};


module.exports.start = start;
