$(document).ready(function() {
  $.fn.serializeObject = function()
  {
      var o = {};
      var a = this.serializeArray();
      $.each(a, function() {
          if (o[this.name] !== undefined) {
              if (!o[this.name].push) {
                  o[this.name] = [o[this.name]];
              }
              o[this.name].push(this.value || '');
          } else {
              o[this.name] = this.value || '';
          }
      });
      return o;
  };

  function getAllShelters(successCallback, errorCallback) {
    $.get({
      url: "/shelter/1",
      success: successCallback,
      error: errorCallback
    })
  }

  getAllShelters(function success(data) {
    console.log(data);
    $("#name").text(data.name);
    $("#desc").text(data.desc);
    $("#phoneEmail").text(data.phone + " | " + data.email);
    $("#url").text(data.url);
    shelter = data;
  }, function error(data) {
    console.log(error);
  });
  var editModal = $("#editModal");
  $("#editButton").click(function() {
    console.log(shelter);
    if(shelter != null) {
        editModal.find("[name=Id]").val(shelter.id);
        editModal.find("[name=Name]").val(shelter.name);
        editModal.find("[name=Desc]").text(shelter.desc)
        editModal.find("[name=Phone]").val(shelter.phone)
        editModal.find("[name=Email]").val(shelter.email)
        editModal.find("[name=Url]").val(shelter.url)
        editModal.modal('show');
    }
  })

  editModal.find("#submit").click(function() {
    console.log($("#shelterSubmit").serializeObject());
    $.post({
      url: "/shelter/1",
      data: $("#shelterSubmit").serializeObject(),
      success: function(data) {
        console.log("Success!");
        console.log(data)
        editModal.modal('hide');
        location.reload();
      },
      error: function(data) {
        console.log(data)
      }
    });


  });


});
