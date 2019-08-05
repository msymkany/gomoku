function loadScenario1() {
  fetch(`/api/scenario1`)
    .then(response => response.json())
    .then(data => resetBoard(data))
    .catch(console.error)
}

function loadScenario2() {
  fetch(`/api/scenario2`)
    .then(response => response.json())
    .then(data => resetBoard(data))
    .catch(console.error)
}
