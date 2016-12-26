'use strict';

const _					= require('ramda');
const fs 				= require('fs');
const Task      = require('data.task');
const Either 		= require('data.either');
const yaml 			= require('js-yaml');
const path_lib	= require('path');

const Left = Either.Left;
const Right = Either.Right;

// Defined in docker-compose.yml
const SCRIPTS_DIRECTORY = '/scripts_dir';
const TEST_DEF_DIRECTORY = '/tests_def_dir';
const TESTS_DEF_FILE = process.env.TESTS_DEF_FILE;

const SCRIPTS_DIRECTORY = '/Users/home/dev/kapacitor-unit/sample';
const TEST_DEF_DIRECTORY = '/Users/home/dev/kapacitor-unit/sample';
const TESTS_DEF_FILE = 'test_case.yaml';


const load_test_data = () => {
	return load_tests_def()
		.map(tests_def => {
      _.keys(tests_def).map(key => {
        const script = load_script(key).getOrElse({});
        tests_def[key].script = script;
      });
      return Right(tests_def);
    });
}

// load_tests_def: -> Either(Error, Object)
const load_tests_def = () => {
  try {
    return Right(yaml.safeLoad(
    	fs.readFileSync(`${TEST_DEF_DIRECTORY}/${TESTS_DEF_FILE}`, 'utf8')));
  } catch (e) {
    return Left(e);
  }
};

// load_script: String -> Either(Error, Object);
const load_script = (script_name) => {
  try {
    return Right(fs.readFileSync(`${SCRIPTS_DIRECTORY}/${script_name}.tick`, 'utf8'));
  } catch (e) {
    return Left(e);
  }
};


module.exports  = {
	load_test_data: load_test_data
};
