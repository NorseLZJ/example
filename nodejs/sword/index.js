const fs = require('fs')

let curDir = ""
let txtDir = "/Placements"


let arr = fs.readdirSync('./')
let buf;
let strOrg;
let str;
let repStr;
let regExp = /mc":{.*]}}}/g; //使用g选项

for (let idx in arr) {
    let file = arr[idx]
    if (file.indexOf(".json") != -1 && file != 'package.json') {
        buf = fs.readFileSync(file)

        str = buf.toString()
        str = str.replace(/\r\n/g, "")
        str = str.replace(/\n/g, "");
        str = str.replace(/\s/g, "");

        strOrg = str
        //console.log(str)
        let res = regExp.exec(str).toString();
        console.log(res)

        break
    }
}