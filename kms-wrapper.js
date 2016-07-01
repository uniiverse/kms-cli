var execSync = require('child_process').execSync;

const runCommand = execSync(`cli --app web --env ${process.env.NODE_ENV} list`);

var blah = JSON.parse(runCommand);

console.log(blah);
