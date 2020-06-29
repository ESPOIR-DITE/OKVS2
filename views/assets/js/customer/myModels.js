function DetailForm(app) {
    var form = document.forms['myTable'];
    var i;
    for (i = 0; i < app.length; i++) {
        form.elements["Name"].value = app[i].Item.name;
        form.elements["Description"].value = app[i].Item.description;
        form.elements["Quantity"].value = app[i].quantity;
        form.elements["Price"].value = app[i].price;
    }


}

function myreview() {
    var form = document.forms['contactForm'];
    alert("we are in")
}