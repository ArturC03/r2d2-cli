const sad = (function () {
  function main() {
    var x = 0;
    if (x > 0) {
      console.log("positive");
    }
  }
  return { main };
})();
sad.main();

