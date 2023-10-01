// innerHTML adalah properti yang digunakan untuk memperoleh atau merubah konten dari elemen itu sendiri.
// getElementById berfungsi untuk menyeleksi id ddari sebuah id html
const element = document.getElementById("messages");
element.innerHTML = "No Messages"; 

// const title = document.getElementById("reviews"); 

const lyrics = document.getElementById("lyrics");
lyrics.innerHTML = "When the sun is gone...";

// querySelector() works with all selectors, like tags, classes, and ID's. 
// querySelector() work on the first element that matches it's search

    // ("p"); for paragraf tag
    // (".prompt"); for a class 
    // ("#promp"); for a id

// If we're looking for an ID, we add the # at the beginning.
function publish() {
    const el = document.querySelector("#prompt"); 
    el.innerHTML = "Website published!";
}
// p.tech berfunmgsi untuk membuat tag p yang ada tech nya 

function user() {
    const el = document.querySelector("p.prompt");
    el.innerHTML = "Website published!";
}

// we can also chane the second class into a second paragraph like this 
// To access an element with two classes, we include them together. Here, it's .prompt.update or .update.prompt .
// Now we don't need to specify the HTML tag because there's only one element with those two classes.

function publish() {
    const el = document.querySelector(".prompt");
    el.innerHTML = "Website published!"; 
}

function refresh() {
    const element = document.querySelector(".order.info");
    element.innerHTML = "Your order has been shipped. Happy self-assembly!";
}

// a p element of class prompt?
// publish.prompt

/* <p class="prompt"></p>
<script src="script.js"></script> */

function displayItem() {
    const el = document.getElementsByTagName("li");
    const p = document.querySelector("p");
    p.innerText = el[0].innerHTML;
}
// bisa juga dibuat untuk mengetahui jumlah dari sebuah length 
// p.innerText = el.length;
// el[0] berfungsi untuk menyekleksi id atau class yang pertama, sama seperti index array

function displayItem() {
    const el = document.getElementsByClassName("urgent");
    // getelementbyclass name berfungsi untuk menyeleksi nama berdasarkan cllass name
    const p = document.querySelector("p");
    p.innerText = "You have " + el.length + " urgent tasks";
}

// document.getElementsByClassName berfungsi untuk menyeleksi class yang sama 
// document.getElementsByTagName befungsi untuk menyeleksi tag yang sama 

// const p = document.querySelector("p"); ngebuat innerHTML (merubah teks) kedalam tag p
// document.getElementsByTagName("h3"); into a h3 heading 
// p.innerText = el[3].innerHTML;

// There's also another powerful way of getting array-like lists of elements, and that's with the querySelectorAll() method.

function displayItem() {
    const el = document.querySelectorAll("h3"); // (.urgent) untuk menyeleksi class name 
    const p = document.querySelector("p");
    p.innerText = el[2].innerHTML;
}

// ("h3.news"); untuk menyeleksi tag h3 dan class news 
// document.querySelectorAll(".movie, .tech"); untuk menyeleksi class name of movie and tech

function displayItems() {
    const el = document.querySelectorAll(".movie, .sports, img"); // bisa menampilkan image juga didalam 
    const p = document.querySelector("p");
    p.innerText = el[2].src;
}

function connect() {
    document.getElementById("listen")
    .innerHTML="Listen to music and podcasts.";
}

// paragraph[0].innerHTML = el[0].innerHTML;
// menjelaskan mengambil tag paragraph pertama dan menulis .innerHTML pada el pertama 

function displayItems() {
    const el = document.querySelectorAll(".sports.news");
    const paragraph = document.querySelectorAll("p");
    paragraph[0].innerHTML = el[0].innerHTML;
}

// <h3 class="movie gossip">Oscar Nominations Drama</h3>
// <h3 class="sports news">The Next Olympic Sport?</h3>
// <h3 class="tech news">The New Y-Phone</h3> 

// TUGASS.. KENAPA DITULIS .SPORT.NEWS padahal dia satu tanda kutip
// dan ketika dia beda class, dia b=dibatasi dengan tanda (,)