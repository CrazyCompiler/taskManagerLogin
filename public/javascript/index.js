var onSignIn = function (googleUser) {
        user = {};
        var profile = googleUser.getBasicProfile();
        user.Id = profile.getId();
        user.name = profile.getName();
        user.email = profile.getEmail();

         $.post("/login",user,function(data,status){

        })

      };

$(document).ready(function(){
})