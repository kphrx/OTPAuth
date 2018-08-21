#!/usr/bin/env node
'use strict';

const program = require("commander")
    , secrets = require("secret-parser")
    , otpCode = require("./lib/otpCode.js")
    , { version } = require('./package.json');

program
    .version(version, '-v, --version')

program
    .command('list')
    .description('List secrets name')
    .action((cmd) => {
        Object.keys(secrets()).forEach(function (name) {
            console.log(name);
        });
    });

program
    .command('secret-key [name]')
    .description('Get secret key')
    .action((name) => {
        console.log(secrets()[name]);
    });

program
    .command('code [name]')
    .description('Get code')
    .action((name) => {
        console.log(otpCode(name));
    });

program.parse(process.argv);
