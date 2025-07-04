const std = (function () {function println(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s+"\n");
     Deno.stdout.write(data).then(() => {});
    }function readline(promptText) {
  const input = promptText ? prompt(promptText) : prompt();
  return input?.trim() ?? "";
  }function print(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s);
      Deno.stdout.write(data).then(() => {});
    }function lengthStr(s) { return s.length; }function trim(s) { return s.trim(); }function split(s, delim) { return s.split(delim); }function join(arr, delim) { return arr.join(delim); }function substring(s, start, end) {
  let result = "";
  for (let i = start; i < end && i < s.length; i++) {
    result += s[i];
  }
  return result;
  }function repeat(s, times) { return s.repeat(times); }function abs(x) { return Math.abs(x); }function sqrt(x) { return Math.sqrt(x); }function pow(x, y) { return Math.pow(x, y); }function sin(x) { return Math.sin(x); }function cos(x) { return Math.cos(x); }function tan(x) { return Math.tan(x); }function floor(x) { return Math.floor(x); }function pi() { return Math.PI; }function random() { return Math.random(); }function now() { return Date.now(); }function readFile(path) { return require('fs').readFileSync(path, 'utf8'); }function exit(code) { process.exit(code); }function setIntervalCallback() { setInterval(() => donut.frame(), 50); }function bitwiseOrZero(x) { return x | 0; }function lower(x) { return x.toLowerCase(); }function upper(x) { return x.toUpperCase(); }function len(x) { return x.length; }function sleep(ms) { const end = Date.now() + ms;
  while (Date.now() < end) {
    // espera ativa
  }}return {sqrt, floor, setIntervalCallback, upper, len, sleep, println, trim, join, tan, pi, random, now, readFile, readline, split, substring, abs, pow, sin, exit, lower, lengthStr, repeat, cos, bitwiseOrZero, print}; })();const DemoCLI = (function () {var greeting = `Hello, R2D2 CLI!`;var x = 10;var y = 5;function stringOp() {console.log(greeting);std.println("String length:"+std.len(greeting));std.println("Uppercase:"+std.upper(greeting));std.println("Lowercase:"+std.lower(greeting));}function mathOp() {std.println("Sum:"+x+y);std.println("Difference:"+(x-y));std.println("Product:"+x*y);std.println("Division:"+x/y);}function conditionOp() {if (x>y) {std.println("x is greater than y");} else {std.println("y is greater than or equal to x");}switch (x) {case 10:std.println("x is 10");break;case 20:std.println("x is 20");break;default:std.println("x is not 10 or 20");}}function loopOp() {std.println("Counting from 1 to 5:");let i = 1;while (i<=5) {std.println(i);i = i+1;}}function main() {var associativeArray = ["key1"=>"value1","key2"=>"value2","key3"=>"value3"];console.log(associativeArray);stringOp();mathOp();conditionOp();loopOp();}return {stringOp, main, x}; })();DemoCLI.main();