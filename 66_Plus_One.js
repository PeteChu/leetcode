// https://leetcode.com/problems/plus-one/

var plusOne = function (digits) {
  let count = digits.length - 1;
  digits[count] += 1;
  while (digits[count] / 10 >= 1) {
    digits[count] = 0;
    if (count === 0) {
      digits = [1, ...digits];
      return digits;
    }
    digits[--count] += 1;
  }
  return digits;
};
