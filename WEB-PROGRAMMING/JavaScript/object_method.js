const myCereal = {
    name: "A",

    changeAndDisplay: function() {
        this.name = "B";
        console.log(this.name);
    }
};

myCereal.changeAndDisplay();
console.log(myCereal.name);
