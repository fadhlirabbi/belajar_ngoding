// panggil fungsi
function greet(name) {
    console.log("Hello, " + name);
}

greet("April");
greet("Leslie");

// We can store the return value in a variable too
// kalo ga ada return age; maka akan (undefined)
function userAge(number) {
    const age = "User age: " + number;
    return age; // return bisa any value, single value  
    }

const result = userAge(29);
console.log(result);

// manggil array pakai [0]
function getListOfNames(names) {
    return names[0] + ", " + names[1];
    
    }
    const students = ["Vera", "Rubin"];
    const list = getListOfNames(students);
    console.log(list);

// contoh memanggil fungsi tanpa variable
function addTen(number) {
    const addedTen = 10 + number;
    return addedTen;
    }
    console.log(addTen(20));

function getAverage(grades) {
    return (grades[0] + grades[1] + grades[2]) / 3;
    }
    const grades = [10, 5, 8];
    const average = getAverage(grades);
    console.log(average);
    // bisa panggil array dengan arr saja 

    function showWinners(first, second, third) {
    console.log("1st: " + first);
    console.log("2nd: " + second);
    console.log("3rd: " + third);   
    }
    showWinners("Elaine", "Helene", "Joe");

// // blackbox bcs we dont know that is inside the function
// const author = "Jautonh Cena";
// // const title = "supervisor";
// // const date = 1980;
// displayAuthor(author);
// // displayTitle(title);
// // displayYear(date);

