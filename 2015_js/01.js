const aoc = require('./lib/aocHelper');

const solveChallenge1 = () => {
  const input = aoc.readInput('./input_01.txt');

  aoc.validate(
    input.length !== 1,
    'File contains an invalid amount of lines'
  );

  const steps = [...input[0]];
  let currPos = 0;

  steps.forEach(s => {
    if(s === '(') {
      currPos += 1;
    } else if(s === ')') {
      currPos -= 1;
    } else {
      aoc.handleError('Invalid character found')
    }
  });

  return currPos;
}

const solveChallenge2 = () => {
  const input = aoc.readInput('./input_01.txt');

  aoc.validate(
    input.length !== 1,
    'File contains an invalid amount of lines'
  );

  const steps = [...input[0]];
  let currPos = 0;

  for(let i = 0; i < steps.length; i++) {
    if(steps[i] === '(') {
      currPos += 1;
    } else if(steps[i] === ')') {
      currPos -= 1;
    } else {
      aoc.handleError('Invalid character found')
    }

    if(currPos < 0) {
      return i + 1;
    }
  }

  return -1;
}

console.log('Result-1: ' + solveChallenge1());
console.log('Result-2: ' + solveChallenge2());
