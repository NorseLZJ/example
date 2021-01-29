// nodejs exec shell script

var callfile = require('child_process');

callfile.execFile('./version.sh', null, function (err, stdout, stderr) {
    if (err || stderr) {
        console.log(err);
        console.log(stderr);
        return;
    }
    console.log(stdout);
});
