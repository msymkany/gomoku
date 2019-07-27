function showWinner(line) {
  if (line) {
    line.forEach(({y, x}) => {
      const cell = document.querySelector(`#y${y}x${x}`);
      cell.classList.toggle('win');
    })
  }
}
