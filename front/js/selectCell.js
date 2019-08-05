function selectCell(selectedY, selectedX) {
  fetch(`/api?y=${selectedY}&x=${selectedX}`)
    .then(response => response.json())
    .then(data => {
      const {y, x, win, notification} = data;

      addNotification(notification);

      const placeRed = () => {
        const cell = document.querySelector(`#y${y}x${x}`);
        cell.classList.toggle('red');
      };

      if (win !== undefined && win !== null) {
        cellClick = null;
        showWinner(win);
        if (y !== -1 && x !== -1) {
          placeRed();
        }
      } else if (y === -1 && x === -1) {
        const cell = document.querySelector(`#y${selectedY}x${selectedX}`);
        cell.classList.toggle('blue');
      } else if (y !== undefined && x !== undefined) {
        placeRed();
      }
    })
    .catch(console.error)
}
