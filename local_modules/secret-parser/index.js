require('dotenv').config();

module.exports = function(issuer) {
    let secrets = process.env.TOTP_SECRET.split(' ')
      , secretsDict = secrets.reduce((previous, current) => {
          let secret = current.split(':');
          previous[secret[0]] = secret[1];
          return previous;
        }, {});

    return secretsDict;
}

