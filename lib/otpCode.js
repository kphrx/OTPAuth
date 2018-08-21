const OTPAuth = require('otpauth')
    , secrets = require("secret-parser")

module.exports = function(issuer) {
    let secret = secrets()[issuer];

    if (!secret) return;

    let totp = new OTPAuth.TOTP({
        secret: OTPAuth.Secret.fromB32(secret),
        hmacAlgorithm: 'sha1'
    });

    return totp.generate();
}

