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
        alert(city)

        $.get("/client/" + city)
            .then(function(data) {
                console.log(data)
            $('#firstQuery').append("<thead><tr><th>Shelter Name</th><th>City</th></tr></thead>");
       // <th>Address</th>" +
       //          "<th>City</th><th>Resources</th><th>Amenities</th>
            $('#firstQuery').append('<tr><td>' + data.City + '</td></tr></tbody>')
            $('#firstQuery').append("<tbody>");
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