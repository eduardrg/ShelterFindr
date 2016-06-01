$(function(){
    function getAllShelters(successCallback, errorCallback) {
      $.get({
        url: "/shelters",
        success: successCallback,
        error: errorCallback
      })
    }

    getAllShelters(function success(data) {
      console.log(data);
    }, function error(data) {
      console.log(error);
    })
})
