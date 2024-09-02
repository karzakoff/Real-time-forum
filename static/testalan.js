const messageContainer = document.getElementById('message-container');

// Fonction pour simuler le chargement des messages
function loadMessages(startIndex, count) {
    // Simule la récupération des messages depuis un serveur ou une source de données
    const messages = [];
    for (let i = startIndex; i < startIndex + count; i++) {
        messages.push(`Message ${i + 1}`);
    }
    return messages;
}

// Fonction pour afficher les messages dans le conteneur
function displayMessages(messages) {
    messages.forEach(message => {
        const div = document.createElement('div');
        div.textContent = message;
        messageContainer.appendChild(div);
    });
}

// Charger les messages initiaux
let lastIndex = 0;
const messagesPerPage = 10;
displayMessages(loadMessages(lastIndex, messagesPerPage));
lastIndex += messagesPerPage;

// Utilitaires pour le throttling et le debounce
function throttle(func, limit) {
    let lastFunc;
    let lastRan;
    return function() {
        const context = this;
        const args = arguments;
        if (!lastRan) {
            func.apply(context, args);
            lastRan = Date.now();
        } else {
            clearTimeout(lastFunc);
            lastFunc = setTimeout(function() {
                if ((Date.now() - lastRan) >= limit) {
                    func.apply(context, args);
                    lastRan = Date.now();
                }
            }, limit - (Date.now() - lastRan));
        }
    };
}

function debounce(func, delay) {
    let timer;
    return function() {
        const context = this;
        const args = arguments;
        clearTimeout(timer);
        timer = setTimeout(() => func.apply(context, args), delay);
    };
}

// Gérer l'événement de défilement
const handleScroll = throttle(() => {
    if (messageContainer.scrollTop === 0) {
       
        const newMessages = loadMessages(lastIndex - messagesPerPage, messagesPerPage);
        displayMessages(newMessages);
        lastIndex -= messagesPerPage;
        messageContainer.scrollTop = 50; 
    }
}, 200); 

messageContainer.addEventListener('scroll', handleScroll);
