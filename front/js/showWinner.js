function showWinner(win) {
  if (win.win) {
    win.win.forEach(({y, x}) => {
      const cell = document.querySelector(`#y${y}x${x}`);
      cell.classList.toggle('win');
    })
  } else if (win.win_by_capture !== 0) {
    document.querySelector('.capture-rule')
      .querySelector({1: '.blue', 2: '.red'}[win.win_by_capture])
      .classList.toggle('win')
  }
}
