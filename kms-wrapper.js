var execSync = require('child_process').execSync;

try {
  const runCommand = execSync(`cli --app web --env ${process.env.NODE_ENV} list`);

  module.exports = JSON.parse(runCommand);

  console.log(blah);
} catch (e) {
  console.log(e);
}
