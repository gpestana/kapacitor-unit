'use strict';

const load_task = (task_id, task, cb) => {
	let err;
	cb(err);
}

const load_data = (task_id, test_def, cb) => {
	let err;
	cb(err);
}

const get_stats = (task_id, cb) => {
	let err;
	cb(err, {});
}

const purge = () => {}

module.exports = {
	load_task: load_task,
	load_data: load_data,
	get_stats: get_stats,
	purge: purge,
}