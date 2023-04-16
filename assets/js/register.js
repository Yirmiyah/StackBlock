function loadRegister() {
    // Load the register form
    let mainContainer = document.getElementById('main_container');
    mainContainer.innerHTML = `

    <section class="blockCentral" id="userSectionCode">

    <input type="email" name="email" id="email">
    <input type="text" name="username" id="username">
    <input type="password" name="password" id="password">
    <input type="password" name="passwordVerification" id="passwordVerification">

    <button type="submit">Sign Up</button>

</section>
<script type="text/javascript">

    let password = document.getElementById('password');
    let passwordVerification = document.getElementById('passwordVerification');

    password.addEventListener('keyup', function() {
        if (password.value !== passwordVerification.value) {
            passwordVerification.setCustomValidity("Passwords Don't Match");
        } else {
            passwordVerification.setCustomValidity('');
        }
    });
    
</script>
`;



}