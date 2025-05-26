const std = (function () {function println(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s+"\n");
      Deno.stdout.write(data).then(() => {});
    }function print(s) {
      const encoder = new TextEncoder();
      const data = encoder.encode(s);
      Deno.stdout.write(data).then(() => {});
    }function lengthStr(s) { return s.length; }function trim(s) { return s.trim(); }function split(s, delim) { return s.split(delim); }function join(arr, delim) { return arr.join(delim); }function substring(s, start, end) { return s.substring(start, end); }function repeat(s, times) { return s.repeat(times); }function abs(x) { return Math.abs(x); }function sqrt(x) { return Math.sqrt(x); }function pow(x, y) { return Math.pow(x, y); }function sin(x) { return Math.sin(x); }function cos(x) { return Math.cos(x); }function tan(x) { return Math.tan(x); }function floor(x) { return Math.floor(x); }function pi() { return Math.PI; }function random() { return Math.random(); }function now() { return Date.now(); }function readFile(path) { return require('fs').readFileSync(path, 'utf8'); }function exit(code) { process.exit(code); }function setIntervalCallback() { setInterval(() => donut.frame(), 50); }function bitwiseOrZero(x) { return x | 0; }function lower(x) { return x.toLowerCase(); }function upper(x) { return x.toUpperCase(); }function len(x) { return x.length; }return {substring, cos, println, pow, sin, tan, random, exit, setIntervalCallback, bitwiseOrZero, split, sqrt, floor, pi, readFile, lower, len, print, trim, repeat, abs, now, upper, lengthStr, join}; })();const WebUtils = (function () {function getElementById(id) {
            return document.getElementById(id);
        }function setInnerHTML(element, html) {
            element.innerHTML = html;
        }function appendHTML(element, html) {
            element.innerHTML += html;
        }function createElement(tag) {
            return document.createElement(tag);
        }function addClass(element, className) {
            element.classList.add(className);
        }function createTextNode(text) {
            return document.createTextNode(text);
        }function appendChild(parent, child) {
            parent.appendChild(child);
        }function createButton(text) {let button = createElement("button");let textNode = createTextNode(text);appendChild(button, textNode);return button;}function createParagraph(text) {let p = createElement("p");let textNode = createTextNode(text);appendChild(p, textNode);return p;}return {}; })();