const std = (function () {
  function println(s) {
    const encoder = new TextEncoder();
    const data = encoder.encode(s + "\n");
    Deno.stdout.write(data).then(() => {});
  }
  function print(s) {
    const encoder = new TextEncoder();
    const data = encoder.encode(s);
    Deno.stdout.write(data).then(() => {});
  }
  function lengthStr(s) {
    return s.length;
  }
  function trim(s) {
    return s.trim();
  }
  function split(s, delim) {
    return s.split(delim);
  }
  function join(arr, delim) {
    return arr.join(delim);
  }
  function substring(s, start, end) {
    return s.substring(start, end);
  }
  function repeat(s, times) {
    return s.repeat(times);
  }
  function abs(x) {
    return Math.abs(x);
  }
  function sqrt(x) {
    return Math.sqrt(x);
  }
  function pow(x, y) {
    return Math.pow(x, y);
  }
  function sin(x) {
    return Math.sin(x);
  }
  function cos(x) {
    return Math.cos(x);
  }
  function tan(x) {
    return Math.tan(x);
  }
  function floor(x) {
    return Math.floor(x);
  }
  function pi() {
    return Math.PI;
  }
  function random() {
    return Math.random();
  }
  function now() {
    return Date.now();
  }
  function readFile(path) {
    return require("fs").readFileSync(path, "utf8");
  }
  function exit(code) {
    process.exit(code);
  }
  function setIntervalCallback() {
    setInterval(() => donut.frame(), 50);
  }
  function bitwiseOrZero(x) {
    return x | 0;
  }
  function lower(x) {
    return x.toLowerCase();
  }
  function upper(x) {
    return x.toUpperCase();
  }
  function len(x) {
    return x.length;
  }
  return {
    random,
    bitwiseOrZero,
    upper,
    trim,
    split,
    sin,
    tan,
    pi,
    exit,
    lower,
    println,
    floor,
    now,
    setIntervalCallback,
    len,
    print,
    substring,
    sqrt,
    pow,
    readFile,
    lengthStr,
    join,
    repeat,
    abs,
    cos,
  };
})();
const DemoCLI = (function () {
  var greeting = `Hello, R2D2 CLI!`;
  var x = 10;
  var y = 5;
  function stringOp() {
    console.log(greeting);
    console.log("String length:" + std.len(greeting));
    console.log("Uppercase:" + std.upper(greeting));
    console.log("Lowercase:" + std.lower(greeting));
  }
  function mathOp() {
    console.log("Sum:" + x + y);
    console.log("Difference:" + (x - y));
    console.log("Product:" + x * y);
    console.log("Division:" + x / y);
  }
  function conditionOp() {
    if (x > y) {
      console.log("x is greater than y");
    } else {
      console.log("y is greater than or equal to x");
    }
    switch (x) {
      case 10:
        console.log("x is 10");
        break;
      case 20:
        console.log("x is 20");
        break;
      default:
        console.log("x is not 10 or 20");
    }
  }
  function loopOp() {
    console.log("Counting from 1 to 5:");
    let i = 1;
    while (i <= 5) {
      console.log(i);
      i = i + 1;
    }
  }
  function main() {
    stringOp();
    mathOp();
    conditionOp();
    loopOp();
  }
  return { main, x };
})();
DemoCLI.main();
console.log("DemoCLI.x:", DemoCLI.x);
