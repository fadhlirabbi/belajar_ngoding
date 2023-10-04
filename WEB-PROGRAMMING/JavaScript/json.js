// JSON stands for JavaScript Object Notation. 
// We can use this data format to exchange data between web applications and servers.

const story = '{"story": "Life","times": 2,"fiction": true}';

// console.log("==============================================");
// We can convert a JS object to a JSON with the JSON.stringify() method.
// Here we can convert the concert JS object into a JSON object if we use JSON.stringify(concert).

console.log("merubah string object ke JSON");

const concert = {
    band: "Super Carrots",
    music: "Indie"
    };
    
console.log(concert);
console.log(JSON.stringify(concert));

console.log("==============================================");

console.log("merubah JSON ke object");

const dog = '{"name":"Rocko","age":3}';
console.log(JSON.parse(dog));

console.log("==============================================");

const flower = '{"name": "tulip", "petals": 30, "city": "Amsterdam"}';
const flowerConversion = JSON.parse(flower);
console.log(flowerConversion.name);

console.log("==============================================");

const tajMahalJSON = '{ "cuisine": "Indian", "takeout": false, "starRating": 4.5}';

const tajMahal = JSON.parse(tajMahalJSON);
const tenRating = tajMahal.starRating * 2;
console.log(tenRating);
console.log(100)