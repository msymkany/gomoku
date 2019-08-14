function selectCell(selectedY, selectedX) {

  document.querySelector('#capture_rule').disabled = true;

  fetch(`/api?y=${selectedY}&x=${selectedX}`)
    .then(response => response.json())
    .then(data => {
      const {blue_pos, red_pos, win, notification} = data;

      document.querySelector('#blue_capture').innerHTML = data.captures[1];
      document.querySelector('#red_capture').innerHTML = data.captures[2];

      addNotification(notification);

      const capture = (pos, className) => {
        const cell1 = document.querySelector(`#y${pos[0].y}x${pos[0].x}`);
        const cell2 = document.querySelector(`#y${pos[1].y}x${pos[1].x}`);
        cell1.classList.toggle(className);
        cell2.classList.toggle(className);
      };

      const placePos = () => {
        const cell = document.querySelector(`#y${red_pos.y}x${red_pos.x}`);
        cell.classList.toggle('red');
        if (red_pos.capture) {
          capture(red_pos.capture.pos, 'blue');
        }
        if (blue_pos.capture) {
          capture(blue_pos.capture.pos, 'red');
        }
      };

      if (data.debug) {
        updateDebug(data.debug);
      }

      if (win !== undefined && win !== null) {
        cellClick = null;
        showWinner(data);
        if (red_pos) {
          placePos();
        }
      } else if (data.win_by_capture) {
        cellClick = null;
        showWinner(data);
        if (red_pos) {
          placePos()
        } else {
          capture(blue_pos.capture.pos, 'red');
        }
      } else if (!red_pos) {
        const cell = document.querySelector(`#y${selectedY}x${selectedX}`);
        cell.classList.toggle('blue');
      } else {
        placePos();
      }

    })
    .catch(console.error)
}
