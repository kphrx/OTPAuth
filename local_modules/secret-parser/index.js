require('dotenv').config();

module.exports = function(secretsString) {
    let secrets = secretsString.split(' ')
      , secretsDict = secrets.reduce((previous, current) => {
          let secret = current.split(':');
          previous[secret[0]] = secret[1];
          return previous;
        }, {});

    return secretsDict;
}

