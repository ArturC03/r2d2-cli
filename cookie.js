const std = (function () {function print(s) { console.log(s); }function println(s) { console.log(s + "\n"); }function printf(formatStr, a, b, c, d, e) {
      let args = [a, b, c, d, e];
      let result = formatStr;
      for (const arg of args) {
        result = result.replace(/%[sdifo]/, String(arg));
      }
      console.log(result);
    }function input(prompt) {
      const readline = require('readline').createInterface({
        input: process.stdin,
        output: process.stdout
      });

      return new Promise((resolve) => {
        readline.question(prompt, (answer) => {
          readline.close();
          resolve(answer);
        });
      });
    }function lengthStr(s) { return s.length; }function substring(s, start, end) { return s.substring(start, end); }function replaceStr(s, oldStr, newStr) { return s.replace(new RegExp(oldStr, 'g'), newStr); }function split(s, delimiter) { return s.split(delimiter); }function join(arr, delimiter) { return arr.join(delimiter); }function trim(s) { return s.trim(); }function toLower(s) { return s.toLowerCase(); }function toUpper(s) { return s.toUpperCase(); }function abs(x) { return Math.abs(x); }function maxVal(x, y) { return Math.max(x, y); }function minVal(x, y) { return Math.min(x, y); }function pow(base, exponent) { return Math.pow(base, exponent); }function sqrt(x) { return Math.sqrt(x); }function round(x) { return Math.round(x); }function floor(x) { return Math.floor(x); }function ceil(x) { return Math.ceil(x); }function random() { return Math.random(); }function randomInt(min, max) { return Math.floor(Math.random() * (max - min + 1)) + min; }function contains(arr, item) { return arr.includes(item); }function indexOf(arr, item) { return arr.indexOf(item); }function reverse(arr) { return [...arr].reverse(); }function sort(arr) { return [...arr].sort(); }function now() { return Date.now(); }function sleep(ms) { return new Promise(resolve => setTimeout(resolve, ms)); }function formatDate(timestamp, format) {
      const date = new Date(timestamp);
      return format
        .replace('YYYY', date.getFullYear())
        .replace('MM', String(date.getMonth() + 1).padStart(2, '0'))
        .replace('DD', String(date.getDate()).padStart(2, '0'))
        .replace('HH', String(date.getHours()).padStart(2, '0'))
        .replace('mm', String(date.getMinutes()).padStart(2, '0'))
        .replace('ss', String(date.getSeconds()).padStart(2, '0'));
    }function readFile(path) {
      const fs = require('fs');
      return fs.readFileSync(path, 'utf8');
    }function writeFile(path, content) {
      const fs = require('fs');
      try {
        fs.writeFileSync(path, content, 'utf8');
        return true;
      } catch (err) {
        return false;
      }
    }function appendFile(path, content) {
      const fs = require('fs');
      try {
        fs.appendFileSync(path, content, 'utf8');
        return true;
      } catch (err) {
        return false;
      }
    }function fileExists(path) {
      const fs = require('fs');
      return fs.existsSync(path);
    }function exec(command) {
      const { execSync } = require('child_process');
      return execSync(command, { encoding: 'utf8' });
    }function exit(code) { process.exit(code); }function getEnv(name) { return process.env[name] || ""; }function setEnv(name, value) { process.env[name] = value; }return {randomInt, readFile, exec, maxVal, sort, sleep, println, substring, replaceStr, round, indexOf, getEnv, trim, abs, pow, reverse, floor, random, lengthStr, toLower, toUpper, ceil, contains, writeFile, appendFile, setEnv, print, printf, input, sqrt, now, fileExists, exit, split, formatDate, join, minVal}; })();const cookie = (function () {function main() {var cookie = 2;if (cookie=="2") {std.print("Cookie is 2");} else {std.print("Cookie is not 2");}}return {main}; })();cookie.main();