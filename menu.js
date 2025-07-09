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
  }}return {lengthStr, trim, split, substring, tan, floor, now, println, repeat, sin, pi, random, readFile, setIntervalCallback, bitwiseOrZero, readline, print, abs, sqrt, exit, sleep, join, pow, lower, upper, len, cos}; })();const coffeeMenu = (function () {var options = ["Espresso", "Cappuccino", "Latte"];function menu(options) {console.log("What do you want to drink?");var i = 0;for (; ; ) {console.log(options[i]);i++;if (i==std.len(options)) {break;}}return std.readline();}function main() {var opt = menu(options);console.clear();console.log("You chose "+options[opt-1]);}return {main}; })();coffeeMenu.main();