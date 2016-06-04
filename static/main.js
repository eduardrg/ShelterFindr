$(function(){
    var city;

    var search = function(data) {
        console.log(data);
        alert(data);
        
    }

    $('#searching').submit(function(e) {
        e.preventDefault();
    });

    $.get("/ping", function(data){
        if(data.error == "true"){
            $("#results").prepend("<div class='alert alert-danger'><strong>Error!</strong> "+ data.message +"</div>");
        }
    }, "json")


    $('#submit').click(function(e) {
        city = $("#city").val();

        $.get("/client/" + city)
            .then(function(data) {
                console.log(data)
            $('#firstQuery').append("<thead><tr><th>Shelter Name</th><th>City</th><<th>Description</th><th>Phone</th><th>StreetAddress</th><th>StateAbbrev</th><th>Zip</th>/tr></thead><tbody>");
       // <th>Address</th>" +
       //          "<th>City</th><th>Resources</th><th>Amenities</th>
            for (var i = 0; i < data.length; i++) {
                $('#firstQuery').append('<tr><td>' + data[i].ShelterName + '</td><td>' + data[i].City + '</td><td>' + data[i].Description + '</td><td>' + data[i].Phone + '</td><td>' + data[i].StreetAddress + '</td><td>' + data[i].StateAbbrev + '</td><td>' + data[i].Zip + '</td></tr>')
            }
            $('#firstQuery').append("</tbody>");
        })
        // .then((function(data) { 
        // $('#submit').click(function(e) {
        //     city = $("#city").val();
        //     alert(city)

        //     $.get("/query1", function(data) {
        //         console.log(data);
        //         $('#firstQuery').append(data);
        //     });
        // });
    })
});