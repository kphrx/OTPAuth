const OTPAuth = require('otpauth')
    , secretParser = require("secret-parser")

module.exports = function(issuer) {
    let secrets = secretParser(process.env.TOTP_SECRET)
      , secret = secrets[issuer];

    if (!secret) return;

    let totp = new OTPAuth.TOTP({
        issuer: issuer,
        algorithm: 'SHA1',
        digits: 6,
        period: 30,
        secret: OTPAuth.Secret.fromB32(secret),
    });

    return totp;
}

