var onSignIn = function (googleUser) {
    user = {};
    var profile = googleUser.getBasicProfile();
    user.Id = profile.getId();
    user.name = profile.getName();
    user.email = profile.getEmail();

     $.post("/login",user,function(data,status){
            document.location.href = data;
    })

};

var onError = function(error){
    alert(JSON.stringify(error, undefined, 2));
}

var googleUser = {};
  var startApp = function() {
    gapi.load('auth2', function(){
        $.get("/clientId","",function(clientId,status){
              auth2 = gapi.auth2.init({
                    client_id: clientId,
                    cookiepolicy: 'single_host_origin',
              });
              attachSignin(document.getElementById('customBtn'));
        })
    });
  };

  function attachSignin(element) {
    auth2.attachClickHandler(element, {},onSignIn,onError);
  }


  $(document).ready(function(){
    startApp();
  })