function editLocationForm(app) {

    var form = document.forms['locationEditForm'];
    form.elements["LocationId"].value = app.locationId;
    form.elements["Name"].value = app.name;
    form.elements["Longitude"].value = app.longitude;
    form.elements["Latitude"].value = app.latitude;
}