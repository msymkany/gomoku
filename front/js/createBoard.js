(async () => {
  let res = await fetch('/api/board');
  res = await res.json();

  const {board, win} = res;

  const boardDiv = document.querySelector('.board');

  for (let y = 0; y < 19; y++) {
    const row = document.createElement('div');
    row.className = 'row';

    for (let x = 0; x < 19; x++) {
      const cell = document.createElement('div');
      const circle = document.createElement('span');
      cell.classList.toggle('cell');
      if (board[y][x] === 1) {
        cell.classList.toggle('blue');
      } else if (board[y][x] === 2) {
        cell.classList.toggle('red');
      }
      cell.id = `y${y}x${x}`;
      cell.onclick = function () {
        if (cellClick &&
          this.classList.value.search('blue') === -1 &&
          this.classList.value.search('red') === -1) {
          this.classList.toggle('blue');
          cellClick(y, x)
        }
      };
      cell.appendChild(circle);
      row.appendChild(cell);
    }

    boardDiv.appendChild(row);
  }

  if (win) {
    cellClick = null;
    showWinner(win);
  }
})();
