let A = 0, B = 0;

function renderFrame() {
  const z = new Array(1760).fill(0);
  const b = new Array(1760).fill(' ');

  for (let j = 0; j < 6.28; j += 0.07) {
    for (let i = 0; i < 6.28; i += 0.02) {
      const c = Math.sin(i),
            d = Math.cos(j),
            e = Math.sin(A),
            f = Math.sin(j),
            g = Math.cos(A),
            h = d + 2,
            D = 1 / (c * h * e + f * g + 5),
            l = Math.cos(i),
            m = Math.cos(B),
            n = Math.sin(B),
            t = c * h * g - f * e,
            x = Math.floor(40 + 30 * D * (l * h * m - t * n)),
            y = Math.floor(12 + 15 * D * (l * h * n + t * m)),
            o = x + 80 * y,
            N = Math.floor(8 * ((f * e - c * d * g) * m - c * d * e - f * g - l * d * n));

      if (y > 0 && y < 22 && x > 0 && x < 80 && D > z[o]) {
        z[o] = D;
        b[o] = ".,-~:;=!*#$@"[N > 0 ? N : 0];
      }
    }
  }

  console.clear();
  let output = "";
  for (let i = 0; i < 1760; i++) {
    output += i % 80 ? b[i] : "\n";
  }
  console.log(output);
  A += 0.04;
  B += 0.02;
}

// Render loop
setInterval(renderFrame, 50);
