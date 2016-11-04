numbers = [12, 37, 5, 42, 8, 3]

const even = numbers.filter(n => !(n % 2));
const odd = numbers.filter(n => n % 2);
console.log(even);
console.log(odd);
