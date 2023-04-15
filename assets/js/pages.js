

document.querySelectorAll('a').forEach(link => {
    link.addEventListener('click', event => {
        event.preventDefault();
        const href = link.getAttribute('href');
        history.pushState(null, null, href);
        updateContent(); // Fonction pour mettre à jour le contenu de la page
    });
});


window.addEventListener('popstate', event => {
    history.replaceState(null, null, document.location.pathname);
    updateContent(); // Fonction pour mettre à jour le contenu de la page
});

function updateContent() {
    const path = window.location.pathname;
    if (path === '/') {
        // Afficher la page d'accueil
        loadHomePage();
    } else if (path === '/login') {
        // Afficher la page Login
        loadLoginPage();
    } else {
        // Afficher une page d'erreur 404
    }
}











