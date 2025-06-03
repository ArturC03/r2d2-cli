const std = (function () {function println(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s+"\n");
      Deno.stdout.write(data).then(() => {});
    }function print(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s);
      Deno.stdout.write(data).then(() => {});
    }function lengthStr(s) { return s.length; }function trim(s) { return s.trim(); }function split(s, delim) { return s.split(delim); }function join(arr, delim) { return arr.join(delim); }function substring(s, start, end) { return s.substring(start, end); }function repeat(s, times) { return s.repeat(times); }function abs(x) { return Math.abs(x); }function sqrt(x) { return Math.sqrt(x); }function pow(x, y) { return Math.pow(x, y); }function sin(x) { return Math.sin(x); }function cos(x) { return Math.cos(x); }function tan(x) { return Math.tan(x); }function floor(x) { return Math.floor(x); }function pi() { return Math.PI; }function random() { return Math.random(); }function now() { return Date.now(); }function readFile(path) { return require('fs').readFileSync(path, 'utf8'); }function exit(code) { process.exit(code); }function setIntervalCallback() { setInterval(() => donut.frame(), 50); }function bitwiseOrZero(x) { return x | 0; }function lower(x) { return x.toLowerCase(); }function upper(x) { return x.toUpperCase(); }function len(x) { return x.length; }return {substring, pi, readFile, lower, len, println, lengthStr, split, repeat, pow, sin, cos, floor, sqrt, random, now, setIntervalCallback, upper, print, trim, abs, tan, exit, bitwiseOrZero, join}; })();const WebDemo = (function () {function main() {let greeting = "Hello from R2D2 Web!";
            document.getElementById('demoResult').innerHTML += 
                '<p>Original: ' + greeting + '</p>' +
                '<p>Length: ' + greeting.length + '</p>' +
                '<p>Uppercase: ' + greeting.toUpperCase() + '</p>';
        let x = 10;let y = 5;
            document.getElementById('demoResult').innerHTML += 
                '<p>Sum: ' + (x + y) + '</p>' +
                '<p>Difference: ' + (x - y) + '</p>';
        if (x>y) {
                document.getElementById('demoResult').innerHTML += 
                    '<p>x is greater than y</p>';
            } else {
                document.getElementById('demoResult').innerHTML += 
                    '<p>y is greater than or equal to x</p>';
            }
            document.getElementById('demoResult').innerHTML += '<p>Counting from 1 to 5:</p>';
            let i = 1;
            while (i <= 5) {
                document.getElementById('demoResult').innerHTML += 
                    '<p>' + i + '</p>';
                i = i + 1;
            }
        }return {main}; })();WebDemo.main();
