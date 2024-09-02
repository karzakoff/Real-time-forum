import { isNotConnected } from "./noConnected.js";
import { resetBody } from "./header.js";

export function logout(){
    let cookieValue = document.cookie
        .split('; ')
        .find(row => row.startsWith('sessionToken='))
        ?.split('=')[1];
   
    fetch('/logout', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify({ cookie: cookieValue })
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message)
        return
    })
    .catch(error =>{
        resetBody();
        isNotConnected();
        location.reload()
    })
}   
