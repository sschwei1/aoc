const aoc = require('../lib/aocHelper');

const solveChallenge1 = () => {
  const input = aoc.readInput('./input_01.txt');

  let totalPaper = 0;

  input.forEach((line, i) => {
    const lineSplit = line.split('x');
    aoc.validate(
      lineSplit.length !== 3,
      `Line ${i+1} contains an invalid input`
    );

    totalPaper += getPackageArea(lineSplit);
  });

  return totalPaper;
}

const getPackageArea = ([a, b, c]) => {
  const s1 = a*b;
  const s2 = b*c;
  const s3 = a*c;

  return 2*s1 + 2*s2 + 2*s3 + getSmallestValue(s1,s2,s3);
}

const solveChallenge2 = () => {
  const input = aoc.readInput('./input_01.txt');

  let totalRibbonLength = 0;

  input.forEach((line, i) => {
    const lineSplit = line.split('x');
    aoc.validate(
      lineSplit.length !== 3,
      `Line ${i+1} contains an invalid input`
    );

    totalRibbonLength += getRibbonLength(lineSplit);
  });

  return totalRibbonLength;
}

const getRibbonLength = ([a, b, c]) => {
  const l1 = 2*a;
  const l2 = 2*b;
  const l3 = 2*c;

  const bow = a*b*c;

  return l1 + l2 + l3 + bow - getLargestValue(l1,l2,l3);
}

const getSmallestValue = (...nums) => {
  return Math.min(...nums);
}

const getLargestValue = (...nums) => {
  return Math.max(...nums);
}

console.log('Result-1: ' + solveChallenge1());
console.log('Result-2: ' + solveChallenge2());