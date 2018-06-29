#! /usr/bin/node env

const fs = require('fs');
const path = require('path');
const execSync = require('child_process').execSync;

const cwd = process.cwd();
const work = __dirname;

process.chdir(work);

const file = fs.readFileSync(path.resolve('..', 'package.json'));
const packageJSON = JSON.parse(file);
delete packageJSON.devDependencies;
fs.writeFileSync('../app/package.json', JSON.stringify(packageJSON));

process.chdir(path.resolve('..', 'app'));

const result = execSync('yarn install');
console.log('result of yarn install: ', result.toString());

process.chdir(cwd);
