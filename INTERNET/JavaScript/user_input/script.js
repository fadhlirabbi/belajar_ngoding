const username = document.getElementById("usernameInput").value;

function register() {
    const paragraph = document.querySelector("p");
    const username = document.getElementById("usernameInput").value;
    paragraph.innerHTML = username;
}
// make a text with have a return u arwe signup 
function register() {
    const paragraph = document.querySelector("p");
    const username = document.getElementById("usernameInput").value;
    paragraph.innerHTML = username + ", you're signed up!";
}

// const review = document.getElementById("reviewText").value;
// that code make tha rawtext, like teks yang bisa di atur kolonm  panjang dan lebarnya 


// membuat text username sebelum diisi menjadi hilang dan mengganti kepada input yang dimasukkan 
function register() {
    document.getElementById("usernameInput").value = "";
}

// menyimpan innerHTML kedalam variable post 
function share() {
    const text = document.querySelector("p");
    const post = document.getElementById("postText").value;
    text.innerHTML = post;
}

function search() {
    const searchWord = document.getElementById("searchInput").value;
    document.getElementById("message").innerHTML = searchWord;
}

function showAttribute() {
    const el = document.querySelector("img");
    const paragraph = document.querySelector("p");
    paragraph.innerText = el.src;
}

function showAttribute() {
    const el = document.querySelector("a");
    const paragraph = document.querySelector("p");
    paragraph.innerText = el.href;
}

// merubah tampilan website kedalam style.css
function addStyleSheet() {
    const el = document.querySelector("link");
    el.href = "style.css";
}

function changeQuestionType() {
    const el = document.querySelector("input");
    el.type = "checkbox";
    el.type = "email";
}

const el = document.querySelector("input");
    el.type = "password";
    el.placeholder = "1234";

    function changeColor() {
    const el = document.querySelector("p");
    el.style.color = "#d7465f";
}

function changeColor() {
    const el = document.querySelector("p");
    el.style.backgroundColor = "#d7465f";
}

function changeStyle() {
    const el = document.querySelector("img");
    el.style.borderRadius
    = "45px";
}

// kita bisa juga merubah style dari sebuah file css dari file yang telah dihubungkan kedalamnya 

function changeStyles() {
    const el = document.querySelector("p");
    el.style.backgroundColor = "aliceBlue";
    el.style.borderRadius = "45px";
}

// el.style.border = "5px solid lightGray";

function displayAttribute() {
    const img = document.querySelector("img");
    const imageLink = img.getAttribute("src");
    const p = document.querySelector('p');
    p.innerText = imageLink;
}

function displayAttribute() {
    const img = document.querySelector("img");
    const imageLink = img.getAttribute("src");
    const p = document.querySelector('p');
    p.innerText = imageLink;
    img.setAttribute();
}

const el = document.querySelector("a");
const imageStyle = el.getAttribute("style");
document.querySelector("p").innerText = imageStyle;

function changeImage() {
    const el = document.querySelector("img");
    const imageLink = el.getAttribute("src");
    const p = document.querySelector('p');  
    p.innerText = imageLink;
}

function showAttribute() {
    const el = document.querySelector("a");
    const elStyle = el.getAttribute("style");
    const elLink = el.getAttribute("href");
    const p = document.querySelector('p');
    p.innerText = elStyle + " " + elLink;
}

// set a date
function changeType() { 
    const el = document.querySelector("input"); 
    el.setAttribute("type","range"); 
}

function updateElement() { 
    const el = document.querySelector("input"); 
    el.setAttribute("placeholder", "1234"); 
}

// .italic {
//     font-style: italic;
// }
// .bold {
//     font-weight: bold;
// }

    function turnBold() {
        const el = document.querySelector("p");
        el.setAttribute("class", "bold");
    }
    
    function turnItalic() {
        const el = document.querySelector("p");
        el.setAttribute("class", "italic");
    }

function addBold() {
    const el = document.querySelector("p");
    el.classList.add("bold");
    // el.classList.remove("bold");
    // el.classList.toggle("bold");
    // el.classList.add("highlight", "underline");
    // penambahan class dari add atau apapun itu harus dibatasi dengan "x" dan tanda koma 
}

// What is the difference between using classList and setAttribute() to modify an element's classes?
// setAttribute set overwrites elements classes while classList does not 

// set password id with js

function showAttribute() {
    const el = document.querySelector("input");
    el.id = "password";
    console.log(el.id);
}

// set source setAttribute
// el.setAttribute("src", "https://mimo.app/r/kittles.png");