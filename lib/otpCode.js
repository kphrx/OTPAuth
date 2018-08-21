const OTPAuth = require('otpauth')
    , secretParser = require("secret-parser")

module.exports = function(issuer) {
    let secrets = secrets(process.env.TOTP_SECRET)
      , secret = secrets[issuer];

    if (!secret) return;

    let totp = new OTPAuth.TOTP({
        secret: OTPAuth.Secret.fromB32(secret),
        hmacAlgorithm: 'sha1'
    });

    return totp.generate();
}

