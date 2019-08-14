(async () => {
  let res = await fetch('/api/board');
  res = await res.json();

  // Set Sidebar
  const doubleThreeInput = document.querySelector('#double_three');
  doubleThreeInput.checked = res.double_three;

  const captureRule = document.querySelector('#capture_rule');
  captureRule.checked = res.capture_rule;

  document.querySelector('#blue_capture').innerHTML = res.captures[1];
  document.querySelector('#red_capture').innerHTML = res.captures[2];

  const depthSelector = document.querySelector('#depth');
  for (let i = 1; i <= maxDepth; i++) {
    const option = newElem('option', null, i.toString());
    depthSelector.appendChild(option);
  }
  depthSelector.value = res.depth;

  const movesSelector = document.querySelector('#moves');
  for (let i = 1; i <= maxMoves; i++) {
    const option = newElem('option', null, i.toString());
    movesSelector.appendChild(option);
  }
  movesSelector.value = res.moves;

  const debugMode = document.querySelector('#debug_mode');
  debugMode.checked = res.debug_mode;
  if (!res.debug_mode) {
    document.querySelector('.debug-wrapper').style.display = 'none';
  }

  // Set Board
  const boardDiv = document.querySelector('#board');

  for (let y = 0; y < boardHeight; y++) {
    const row = newElem('div', 'row');

    for (let x = 0; x < boardWidth; x++) {
      const cell = newElem('div', 'cell');
      if (res.board[y][x] === 1) {
        cell.classList.toggle('blue');
      } else if (res.board[y][x] === 2) {
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
      cell.appendChild(newElem('span'));
      row.appendChild(cell);
    }

    boardDiv.appendChild(row);
  }

  if (res.win || res.win_by_capture !== 0) {
    cellClick = null;
    showWinner(res);
  }
})();

// document.querySelectorAll('.board').forEach(board => {
//   const turnScore = document.createElement('div');
//   turnScore.innerHTML = 'turnScore: 0.12345678';
//   turnScore.className = 'score-num';
//   board.appendChild(turnScore);
//
//   const bestScore = document.createElement('div');
//   bestScore.innerHTML = 'bestScore: 0.12345678';
//   bestScore.className = 'score-num';
//   board.appendChild(bestScore);
//
//   for (let y = 0; y < boardHeight; y++) {
//     const row = document.createElement('div');
//     row.className = 'row';
//
//     for (let x = 0; x < boardWidth; x++) {
//       const cell = document.createElement('div');
//       const circle = document.createElement('span');
//       cell.className = 'cell small';
//       // cell.id = `y${y}x${x}`;
//       // cell.onclick = function () {
//       //   if (cellClick &&
//       //     this.classList.value.search('blue') === -1 &&
//       //     this.classList.value.search('red') === -1) {
//       //     this.classList.toggle('blue');
//       //     cellClick(y, x)
//       //   }
//       // };
//       cell.appendChild(circle);
//       row.appendChild(cell);
//     }
//
//     board.appendChild(row);
//   }
// });
