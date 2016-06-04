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
      url: "/shelter",
      success: successCallback,
      error: errorCallback
    })
  }

  function getShelter(id, successCallback, errorCallback) {
    $.get({
      url: "/shelter/" + id,
      success: successCallback,
      error: errorCallback
    })
  }

  getAllShelters(function success(data) {
    for(s in data) {
      var shelter = data[s];
      var rowHtml = $("#shelter-row-tpl").clone();
      rowHtml.find("#name").text(shelter.Name);
      rowHtml.find("#desc").text(shelter.Desc);
      rowHtml.find("#phone").text(shelter.Phone);
      rowHtml.find("#email").text(shelter.Email);
      rowHtml.find("#url").text(shelter.Url);
      var btn = rowHtml.find(".editButton");
      btn.data("shelterId", shelter.Id);
      btn.click(function() {
        btnAction($(this));
      });
      rowHtml.removeAttr("id");
      rowHtml.removeAttr("id");
      rowHtml.removeClass("hidden");
      $("#shelter-table").append(rowHtml);
    }


    $("#name").text(data.name);
    $("#desc").text(data.desc);
    $("#phoneEmail").text(data.phone + " | " + data.email);
    $("#url").text(data.url);
    shelter = data;
  }, function error(data) {
    console.log(error);
  });

  var editModal = $("#editModal");
  editModal.find("#submit").click(function() {
    var shelterId = $("#shelterSubmit").find("[name=Id]").val();
    console.log("Updated shelter with Id:" + shelterId);
    $.post({
      url: "/shelter/"+shelterId,
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

  function btnAction(el) {
    var id = el.data("shelterId");
    if(id != undefined) {
      getShelter(id, function success(shelter) {
        console.log("Success getting shelter with Id: " + id);
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
      }, function error(data) {
        console.log("Error getting shelter with Id: " + id);
        console.log(data);

      });
    }
  }

});
