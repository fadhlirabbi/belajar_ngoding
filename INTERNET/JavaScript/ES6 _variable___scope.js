// ES6 is stand for EMAScript 2015
// to fix issues and improve js syntax

var number = 20;
const player1 = 'Sonic';
let badWeather = "rainy";
console.log(number);
console.log(player1);
console.log(badWeather);

console.log("================================");
// var can be declare multiple times, while let and const can't

var surname = "Smit";
var surname = "Smith";
console.log(surname);

console.log("================================");
// how to declare a let 

let year = 2021;
year = 2022;
console.log(year);

console.log("================================");
// we can push string into a const

const newPresent = "Trousers";
const wardrobe = ["Shirt"];
wardrobe.push(newPresent);

console.log(wardrobe);

console.log("================================");
// we can't update the const, but we can update it's properties 

const myClothes = {
    shirts: 1,
};

myClothes.shirts++;
console.log(myClothes.shirts);

console.log("================================");

// we can store data before declare it
zodiacSign = "Aries";
var zodiacSign;
console.log(zodiacSign);

console.log("================================");
// we can't acces myDiary because it's private 

// function readDiary() {
//     const myDiary = "Dear diary, today...";
// }
// console.log(myDiary);

console.log("it will display error");

console.log("================================");
// The last scope type is called block scope. 
//It's the part of our code in between the braces ( {} ) that belong to loops or conditionals.

if (10 === 10) {
    console.log("This space here is block scope");
}

//if we declare let and const in a block scope, we only can acces that only in that block scope

// for (let index = 1; index <= 2; index++) {
    //     console.log("Let's count to three");
    // }
    // // console.log(index); // this will display error

console.log("================================");
// we can't acces bath bcs its not in a block space
console.log("acces const ouside the code block")
// let neighborhood = {
    //     apartments: 13,
    //     garden: 1
    // };
    
// function fixMyHouse() { 
    //     const BATH = "Fix Mirror";
    // }
    // console.log(BATH);
console.log("================================"); 
// We use this (=>) operator right before the code block or expression we want the function to execute.

// shorthand for a function
function getGreeting (){
    return 'Greetings!';
}
const greeting = () => {
    return 'Hi';
}; // jangan lupa ; diakhir scope

console.log(getGreeting());
console.log(greeting());

console.log("================================"); 

const getVehicleCount = () => {
    const count = 4003;
    return count;
};
console.log(getVehicleCount());

console.log("================================"); 

const getPrice = (discount) => 500 - (discount * 500);
console.log(getPrice(0.25));

console.log("================================"); 
// In arrow functions, we don't have to add the parentheses if we have (ONLY ONE) parameter.
// example =

const get = total => {
    const discount = 0.10;
    return total - (discount * total);
};
console.log(get(500));

console.log("================================"); 
// Parameter values are set to undefined by default. If a function call doesn't pass a value, the function will use undefined as the value. NaN mean undefine value.

const price = (total) => {
    const tax = 0.16;
    return total + (tax * total);
};
console.log(price());
// will show error, because that total doesn't have a value

// to avoid that error use (total = 200) to make it work
// we can use that like (total = 200, discount = 10)

// Arrow functions let us have more compact and concise code. How we write them depends on the parameters we use and the code block size.

console.log("================================"); 
// we can code like this too because value fill the nom, and denom have own declare value   

const getRatio = (nom, denom = 2) => {
    return nom / denom;
};
console.log(getRatio(40));

// we can alse change the value of denom 

// const getRatio = (nom, denom = 2) => {
//     return nom / denom;
//     };
// console.log(getRatio(40));


