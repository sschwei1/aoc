const fs = require('fs');
const readInput = () => {
  const content = fs.readFileSync('./input_01.txt', 'utf-8');
  return content.split(/\r?\n/);
}

const handleError = (msg = '') => {
  console.error(msg)
  process.exit(1);
}
const validate = (validateError = false, msg = '') => {
  if(validateError) {
    handleError(msg)
  }
}

const solveChallenge1 = () => {
  const input = readInput();

  validate(
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
      handleError('Invalid character found')
    }
  });

  return currPos;
}

const solveChallenge2 = () => {
  const input = readInput();

  validate(
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
      handleError('Invalid character found')
    }

    if(currPos < 0) {
      return i + 1;
    }
  }

  return -1;
}

console.log('Result-1: ' + solveChallenge1());
console.log('Result-2: ' + solveChallenge2());