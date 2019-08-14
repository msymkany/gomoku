function changeDoubleThreeRule() {
  const doubleThree = document.querySelector('#double_three');

  fetch(`/api/settings/double_three?check=${doubleThree.checked}`)
    .catch(console.error)
}

function changeCaptureRule() {
  const capture = document.querySelector('#capture_rule');

  fetch(`/api/settings/capture?check=${capture.checked}`)
    .catch(console.error)
}

function changeDifficulty() {
  const depth = document.querySelector('#depth');
  const moves = document.querySelector('#moves');
  const depthValue = depth.options[depth.selectedIndex].value;
  const movesValue = moves.options[moves.selectedIndex].value;

  fetch(`/api/settings/difficulty?depth=${depthValue}&moves=${movesValue}`)
    .catch(console.error)
}

function changeDebugMode() {
  const debug = document.querySelector('#debug_mode');
  const wrapper = document.querySelector('.debug-wrapper');
  if (debug.checked) {
    wrapper.style.display = 'flex';
  } else {
    wrapper.style.display = 'none';
  }

  fetch(`/api/settings/debug?check=${debug.checked}`)
    .catch(console.error)
}

