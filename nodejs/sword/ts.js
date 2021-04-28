//原始字符串
var str = "欢迎访问!\r\nhangge.com    做最好的开发者知识平台";


console.log(str);

//去掉所有的换行符
str = str.replace(/\r\n/g, "")
str = str.replace(/\n/g, "");
str = str.replace(/\s/g, "");

//输出转换后的字符串
console.log(str);