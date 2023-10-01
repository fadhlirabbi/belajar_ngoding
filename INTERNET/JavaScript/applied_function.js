// satu codeblck is running

function isRentingAge(age) {
    if (age >= 21) {
        console.log("Approved");
    }
    console.log("Check done");
}

isRentingAge(15);

console.log("==============================================");

// running 2 codeblock
function isLying(boolean) {
    if (boolean === true) {
        console.log("He's not lying");
    }
    console.log("Done");
}

isLying(true);

console.log("==============================================");

function isGreater(a, b) {
    if (a > b) {
        console.log(a + " is greater than " + b);
    }
}
isGreater(13, 10);

console.log("==============================================");

function isGreater(a, b) {
    if (a > b) {
        console.log(a + " is greater than " + b);
    }
}
isGreater(13, 10);

console.log("==============================================");

function checkPassword(word) {
    if (word === "12345") {
        console.log("Logging in");
    } else {
        console.log("Wrong password");
    }
}
checkPassword("swordfish");

console.log("==============================================");

// Let's learn how to use return statements to exit functions at any time we want, even in the middle of the code.

function multiplyByTen(number) {
    return number * 10;
    return "This line doesn't execute";
}

const result = multiplyByTen(29);
console.log(result);

console.log("==============================================");

function multiplyByTen(number) {
    console.log("Math started");
    return number * 10;
    console.log("Math is done!"); // this line wont execute
}

const result1 = multiplyByTen(30);
console.log(result1);

console.log("==============================================");

function displayNumbers(message) {

    for (let i = 0; i < 3; i++) {
        console.log(i + message);
    }
}

displayNumbers(" ginger beers on the wall");

console.log("==============================================");

function displayAltitude(start, end) {
    for (let i = start; i < end; i++) {
        console.log(i + " ft");
    }
}
displayAltitude(1300, 1305);

console.log("==============================================");
console.log("looping the array");

const grades = [92, 66, 77, 84];
const grades2 = [50, 60, 70, 80];

function searchGrade(grades, grades2) {
    console.log(grades);
    console.log(grades2);
}
searchGrade(grades, grades2);

console.log("==============================================");

const group = 4;
const total = 500;
const minimum = 150;

function stockCheck(groupSize, totalAvailable, minPerUser) {
    if (minPerUser > totalAvailable) {
        return "Insufficient for one user";
    } 
    else if (groupSize * minPerUser > totalAvailable) {
        return "Insufficient for all users";
    }
    else {
        return "Sufficient for all users";
    }
}

console.log(stockCheck(group, total, minimum));
