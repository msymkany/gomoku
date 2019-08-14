const newElem = (elemName, className = null, innerHTML = null) => {
  const el = document.createElement(elemName);
  el.className = className;
  el.innerHTML = innerHTML;
  return el
};

let cellClick = selectCell;
const maxMoves = 150;
const maxDepth = 4;

const boardWidth = 19;
const boardHeight = 19;
