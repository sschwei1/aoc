const fs = require('fs');
const readInput = (path = './input_01.txt') => {
  const content = fs.readFileSync(path, 'utf-8');
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

module.exports = {
  readInput,
  handleError,
  validate
}