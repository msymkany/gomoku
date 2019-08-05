function resetGame() {
  fetch(`/api/reset`)
    .then(response => response.json())
    .then(data => resetBoard(data))
    .catch(console.error)
}

function resetBoard(data) {
  document.querySelector('.notification').innerHTML = '';
  for (let y = 0; y < 19; y++) {
    for (let x = 0; x < 19; x++) {
      const cell = document.querySelector(`#y${y}x${x}`);
      cell.className = 'cell';
      if (data[y][x] === 1) {
        cell.classList.toggle('blue');
      } else if (data[y][x] === 2) {
        cell.classList.toggle('red');
      }
    }
  }
  cellClick = selectCell;

}
