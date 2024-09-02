import { isNotConnected } from "./noConnected.js";
import { createHeader } from "./header.js";
import { homePage } from "./home.js";
import { logout } from "./logout.js";


checkifCookieIsgood()

createHeader();

export function cookieRedirection(){
    if (checkCommentPermission()){
        homePage();
    } else {
        isNotConnected();
    }
}



function checkifCookieIsgood(){
    fetch('/checkCookie', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

cookieRedirection()

changeHeader(checkCommentPermission())


export function checkCommentPermission() {
    let cookieValue = document.cookie
        .split('; ')
        .find(row => row.startsWith('sessionToken='))
        ?.split('=')[1];
   
    if (cookieValue && cookieValue !== '') {
        return true
    } else {
        return false
    }
}

function changeHeader(bool){
    const loginButton = document.getElementById("loginButtonHeader")
    const chat = document.getElementById("chatPage")
    if (bool==false){
        loginButton.innerHTML = "Login"
        chat.innerHTML=""
        chat.addEventListener("click", ()=>{
            alert("You need to be connected to access the chat")
        })
    } else {
        loginButton.innerHTML = "Logout"
        chat.innerHTML="Chat"
        loginButton.addEventListener("click", ()=>{
            logout()
            isNotConnected()
            changeHeader(false)
        })
    }
}