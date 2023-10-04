const fahrenheit = [72, 68, 70, 74, 77, 75, 79];
const celsius = [];

for (const element of fahrenheit) {
    let c = (element - 32) * (5/9);
    celsius.push(c);
}
console.log(celsius);

console.log("==============================================")

// const fahrenheit = [72, 68, 70, 74, 77, 75, 79];
// const celsius = fahrenheit.map();

const states = ['ak', 'md', 'va', 'ca', 'or'];
const uppercaseStates = states.map((element) => element.toUpperCase());
console.log(uppercaseStates);

console.log("==============================================")

const values = [1, 3, 4, 7];
const squares = values.map(element => element*element);
console.log(squares);

console.log("==============================================")

const queue = ["Sarah", "Hank", "Anna", "Beatrice"];
const displayQueue = queue.map(function(element, index) {
    return index + ": " + element;
});
console.log(displayQueue);

console.log("==============================================")

// .map() doesn't only work on integer arrays, but on other value types, like strings, as well.

const drivers = ["LeClerc", "Sainz", "Hamilton"];
const places = drivers.map(function(element, index) { 
    return index+1 + ". " + element;
});
// index berfungsi untuk menunjukkan nomor dari array tersebut
console.log(places);

console.log("==============================================")

const participants = ["Ben", "Sarah", "Eli", "Henry", "Sean", "Annabel"];
const sNames = participants.filter(function(element) {
    return element[0] === "S";
});
console.log(sNames);

console.log("==============================================")

const week = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
const workWeek = week.filter(function(element, index) {
    return (index > 0 && index < 6);
});
console.log(workWeek);

console.log("==============================================")

// What type of value does the inline function of .filter() return ? = a boleaan

const prices = [19.99, 20.99, 18.99, 17.99, 21.99];
const inBudget = prices.filter(function(price) {
    return price < 20;
});
console.log(inBudget);

console.log("==============================================")

const finishers = ["Sarah", "Sally", "Jake", "Alex", "Tori"];
const top3 = finishers.filter(function(finisher, place) {
    return place < 3;
});
console.log(top3);

console.log("==============================================")

// const finishers = ["Sarah", "Sally", "Jake", "Alex", "Tori"];
// const top3 = finishers.filter((finisher, place) => place < 3);
// console.log(top3);

// Which JavaScript function can we use to convert a series of array values to a single value = .reduce()
// What does the inline function's first parameter give reduce() access to? = the result of the previous reduction
const miles = [33, 95, 79];
const total = miles.reduce(function(prev, curr) {
    return prev + curr;
});
console.log(total);

console.log("==============================================")

const x = [17.50, 20.25, 24.75, 13.50, 8.50];
const y = [0.9, 0.9, 1.0, 1.0, 0.9];

const totalx = x.reduce(function(prevVal, currVal, index) {
    return prevVal + (currVal * y [index]);
});
console.log(total);

console.log("==============================================")

const words = ["The", "house", "is", "red"];
const sentence = words.reduce(function(prev, curr) {
    return prev + " " + curr;
});
console.log(sentence);

console.log("==============================================")

const xd = [17.50, 20.25, 24.75, 13.50, 8.50];
const totalxy = xd.reduce((prevVal, currVal) => prevVal + currVal);
console.log(totalxy);


console.log("======================");

const pricescd = [1.99, 2.49, 0.99, 4.99];
const taxed = pricescd.map(price => price * 1.1);
console.log(taxed);

console.log("======================");


const names = ["Tony", "Tania", "Vince", "Terry"];
const T = names.filter(element => (element[0] === "T"));
console.log(T);

