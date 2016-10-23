'use strict';

const report = [];


const prints_report = (log) => {
	log.info(report);
}

const add_no_test = (task_id) => {
	report.push({
		type: 'warning',
		task_id: task_id,
		reason: 'No test defined.'
	});
}

const add_err = (task_id, err) => {
	report.push({
		type: 'err',
		task_id: task_id,
		reason: err
	});
}

module.exports = {
	prints_report: prints_report,
	add_no_test: add_no_test,
	add_err: add_err,
}