function changeDifficulty() {
  const depth = document.querySelector('#depth');
  const moves = document.querySelector('#moves');
  const depthValue = depth.options[depth.selectedIndex].value;
  const movesValue = moves.options[moves.selectedIndex].value;

  fetch(`/api/difficulty?depth=${depthValue}&moves=${movesValue}`)
    .catch(console.error)
}
