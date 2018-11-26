const OTPAuth = require('otpauth')
    , config = require('config');

module.exports = function(issuer) {
    let issuerName = "TOTP-Secrets." + issuer;

    if (!config.has(issuerName)) return;

    let secret = config.get(issuerName + ".secret")
      , defaultParams = Object.assign({issuer: issuer}, config.get("TOTP-Params"))
      , params = Object.assign(defaultParams, config.get(issuerName))
      , secretParams = Object.assign(defaultParams, {secret: OTPAuth.Secret.fromB32(secret)});

    let totp = new OTPAuth.TOTP(secretParams);

    return totp;
}

