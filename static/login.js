
import { resetBody } from "./header.js";
import { homePage } from "./home.js";


export function loginPage() {
    const container = document.createElement('div')
    container.classList.add('login-background')

    const TextLogin = document.createElement('span')
    TextLogin.innerHTML = "LOGIN"
    TextLogin.classList = "login-text"

    const createInput = (text, id, type) => {
        const element = document.createElement('input')
        element.classList = "login-input"
        element.placeholder = text
        element.id = id
        element.type = type
        element.autocomplete = true
        element.name = id
        element.required = true

        if (id=="email") {
            element.name = ""
        } else {
            element.name = id
        }
        element.required = true
        return element;
    }

    const loginButton = document.createElement('button')
    loginButton.classList = "buttonLogin"
    loginButton.textContent = "Login"

    const form = document.createElement('form')
    form.id = "loginForm"

    
    const loginDiv = document.createElement('div')
    loginDiv.classList = "login-button-div"
    loginDiv.append(
        createInput("Pseudo or Email", "name", "name"), 
        createInput("Password", "password", "password"),
    )

    form.append(loginDiv, loginButton)



    form.addEventListener('submit', function(event) {
        event.preventDefault();
    
        const username = document.getElementById('name').value;
        const password = document.getElementById('password').value;

    
            fetch('/login', {
                method: 'POST',
                headers: {
                'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username: username, password: password })
            })
            .then(response => response.json())
            .then(data => {
                alert(data.message)
                return
            })
            .catch(error =>{
                resetBody();
                homePage();
                location.reload()

        })
    });


    container.append(TextLogin, form)

    document.body.appendChild(container)
}