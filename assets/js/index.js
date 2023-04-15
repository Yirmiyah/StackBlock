export function loadHomePage() {
    // Afficher la page d'accueil

    let mainContainer = document.querySelector('.main_container');
    mainContainer.innerHTML = `
    <section class="block_navbar">
        <div class="navbar">
            <div class="navbar_logo">
                <img src="assets/images/logo.png" alt="logo">
            </div>
            <div class="navbar_menu">
                <ul id="menu">
                    <li><a href="#">Home</a></li>
                    <li><a href="#">Profil</a></li>
                    <li><a href="#">Wanted</a></li>
                    <li><a href="#">Sold</a></li>
                    <li><a href="#">Logout</a></li>
                    <li><a href="#">Login</a></li>
                </ul>
            </div>
            <div class="navbar_user">
                <img src="assets/images/user.png" alt="user image">
            </div>
        </div>
    </section>

    <section class="blockCentral" id="userSectionCode">

        <img src="assets/images/user.png" alt="user image">
        <h3 class="userName">User Name</h3>

        <textarea name="userDescription" id="userDescription" cols="30" rows="5"
            placeholder="Describe your code..."></textarea>

        <textarea name="userCodeText" id="userCodeText" cols="30" rows="30"
            placeholder="Write your code here..."></textarea>
            <textarea name="userDescription" id="userDescription" cols="30" rows="5"
            placeholder="Describe your code..."></textarea>

        <textarea name="userCodeText" id="userCodeText" cols="30" rows="30"
            placeholder="Write your code here..."></textarea>

        <button id="btn_sell" type="submit">Sell</button>

    </section>

    <section class="blockCentral" id="userCodeToSold">
        <div class="codeDescription">The Code does this...</div>
        <div class="userWantedCode">User code's goes here...</div>
        <br>

        <div class="likes-comments">
            <div class="Likes">
                <button>Like</button>
                <span>10</span>
            </div>
            <div class="Comments">
                <button>Write a comment </button>
                <span>5</span>
            </div>
        </div>
    </section>
    `;

    let btnSell = document.querySelector('#btn_sell');
    btnSell.addEventListener('click', function () {
        loadSoldPage();
    });

}