export function loadLoginPage(){
    //afficher la page login
    let mainContainer = document.getElementById('main_container');
    mainContainer.innerHTML = `
    <section class="blockCentral" id="userSectionCode">
    
            <img src="assets/images/user.png" alt="user image">
            <h3 class="userName">User Name</h3>
    
            <input type="email" name="login" id="loginEmail">
            <input type="password" name="password" id="loginPassword">
            <button id="btn_sell" type="submit">LogIn</button>
    
        </section>
        `;
        
        
}

