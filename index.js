#!/usr/bin/env node
'use strict';

process.chdir(__dirname)

const program = require("commander")
    , qrcode = require('qrcode-terminal')
    , secretParser = require("secret-parser")
    , otpCode = require("./lib/otpCode.js")
    , { version } = require('./package.json');

program
    .version(version, '-v, --version')

program
    .command('list')
    .description('List secrets name')
    .action((cmd) => {
        let secrets = secretParser(process.env.TOTP_SECRET);

        Object.keys(secrets).forEach(function (name) {
            console.log(name);
        });
    });

program
    .command('secret-key [name]')
    .description('Get secret key')
    .action((name) => {
        let secrets = secretParser(process.env.TOTP_SECRET)
          , secret = secrets[name];

        console.log(secret);
    });

program
    .command('qr [name]')
    .description('Get QRCode')
    .action((name) => {
        let secrets = secretParser(process.env.TOTP_SECRET)
          , secret = secrets[name];

        qrcode.generate(otpCode(name).toString(), {small: true}, function (qrcode) {
            console.log(qrcode);
        });
    });

program
    .command('token [name]')
    .description('Get token')
    .action((name) => {
        console.log(otpCode(name).generate());
    });

program.parse(process.argv);
