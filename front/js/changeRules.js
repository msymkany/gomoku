function changeDoubleThree() {
  const doubleThree = document.querySelector('#double_three');

  fetch(`/api/rules/double_three?check=${doubleThree.checked}`)
    .catch(console.error)
}
