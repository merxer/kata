var items = [{price: 10}, {price: 120}, {price: 1000}];

var reducer = function add(sumSoFar, item) { return sumSoFar + item.price; }

var total = items.reduce(reducer, 0);
console.log(total);
