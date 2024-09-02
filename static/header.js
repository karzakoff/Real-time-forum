import { homePage} from "./home.js";
import {chatPage } from "./chat.js";
import { cookieRedirection } from "./main.js";
import { changeChatRoom } from "./chat.js";


export function resetBody(){
    const elementsToRemove = [...document.body.children].filter(el => el.tagName !== 'HEADER');
    elementsToRemove.forEach(el => el.remove());
}
function launchChatPage(){
    chatPage()
}

var ModalValue = 0

export function createHeader(){
    const header = document.createElement('header');

    const topDiv = document.createElement('div');
    topDiv.classList.add('top');

    const h3 = document.createElement('h3');
    h3.textContent = 'Real Time Forum';
    topDiv.appendChild(h3);

    const topRightDiv = document.createElement('div');
    topRightDiv.classList.add('topRight');

    const homePage = document.createElement('h4');
    homePage.textContent = "Home"
    homePage.addEventListener("click", ()=>{
        ModalValue = 0
        resetBody()
        cookieRedirection()
        changeChatRoom("general")
    })

    const chatPage = document.createElement('h4');
    chatPage.textContent = "Chat"
    chatPage.id = 'chatPage'
    chatPage.addEventListener("click",()=>{
        ModalValue = 0
        resetBody()
        launchChatPage()
    })

    const notif = document.createElement('img')
    notif.classList.add('notif')
    notif.src = "assets/bell.png"
    notif.id = 'notif'

    const notifSignal = document.createElement('div')
    notifSignal.classList.add('notifSignal')
    notifSignal.id = "notifSignal"


    topRightDiv.appendChild(homePage)
    topRightDiv.appendChild(chatPage)
    
    const loginButton = document.createElement('h4')
    loginButton.id = "loginButtonHeader"
    loginButton.textContent = "test"
    
    topRightDiv.appendChild(loginButton)
    notifSignal.style.display = 'none'
    topDiv.append(notif, notifSignal)
    topDiv.appendChild(topRightDiv);
    
    header.appendChild(topDiv);
    document.body.appendChild(header);

    notif.addEventListener('click', () => {
        if (ModalValue == 0) {
            createModal()
            loadMessage()
        }
        notifSignal.style.display = 'none'
    })
}





export function createModal() {
    ModalValue = 1

    const myModal = document.createElement('div')
    myModal.classList.add('myModal')
    myModal.id = 'myModal'

    const buttonClose = document.createElement('img')
    buttonClose.classList.add('buttonClose')
    buttonClose.id = 'close'
    buttonClose.src = 'assets/close.png'
    buttonClose.textContent = "Close"
    buttonClose.addEventListener('click', () => {
        myModal.remove()
        ModalValue = 0
    })

   myModal.append(buttonClose)
   loadMessage()
   document.body.appendChild(myModal)
}

const notifications = (name, content, id) => {
    const newNotif = document.createElement('div')
    newNotif.classList.add('newNotif')
 
    const nameNotif = document.createElement('span')
    nameNotif.classList.add('nameNotif')
    nameNotif.textContent = name
    nameNotif.id = id

    const contentNotif = document.createElement('p')
    contentNotif.classList.add('contentNotif')
    contentNotif.textContent = content

    newNotif.append(nameNotif, contentNotif)
    return newNotif
}


const loadMessage = async () => {
    const messages = await fetchMessage();
    if (messages){
        const myModal = document.getElementById('myModal')
        messages.forEach(element => {

            myModal.append(notifications(element.username,element.content, element.id))
    
        });
    }
};







async function fetchMessage() { 
    try {
        const response = await fetch('/recupNotif') 
        if (!response.ok) {
            throw new Error(`Network response was not ok: ${response.statusText}`);
        }
        const message = await response.json();
        return message;
    } catch (error) {
        console.error('Fetch operation failed:', error);
        return [];
    }
}