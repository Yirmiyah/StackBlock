function loadRegister(){
    // Load the register form
    $.get('register.html', function(data){
        $('#content').html(data);
    });
}