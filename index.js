#!/usr/bin/env node
'use strict';

process.chdir(__dirname)
require('dotenv').config();

const program = require("commander")
    , config = require('config')
    , qrcode = require('qrcode-terminal')
    , otpCode = require("./lib/otpCode.js")
    , { version } = require('./package.json');

program
    .version(version, '-v, --version')

program
    .command('list')
    .description('List secrets name')
    .action((cmd) => {
        let secrets = config.get("TOTP-Secrets");

        Object.keys(secrets).forEach(function (name) {
            console.log(name);
        });
    });

program
    .command('secret-key [name]')
    .description('Get secret key')
    .action((name) => {
        let secret = config.get("TOTP-Secrets." + issuer);

        console.log(secret);
    });

program
    .command('qr [name]')
    .description('Get QRCode')
    .action((name) => {
        qrcode.generate(otpCode(name).toString(), {small: true}, function (qrcode) {
            console.log(qrcode);
        });
    });

program
    .command('url [name]')
    .description('Get otp url')
    .action((name) => {
        console.log(otpCode(name).toString());
    });

program
    .command('token [name]')
    .description('Get token')
    .action((name) => {
        console.log(otpCode(name).generate());
    });

program.parse(process.argv);
