
const category = `pie`;
const cake = `Apple ${category}`;
console.log(cake);

console.log("================================");

const quantity = 20;
const price = 10;
const cost = `Total cost : ${quantity * price}`;
console.log(cost);

console.log("================================");

const greetings = (name) => {
    return `Bonjour ${name}!`;
};

console.log(greetings`Jane`);

console.log("================================");

let old = `Superman`;
let latest = `Startrek`;
let movies = `${old}, ${latest}`;
console.log(movies);

console.log("================================");

const likedBy = person => {
    return `Movie liked by: ${person}`;
};

console.log(likedBy`Hellen`);