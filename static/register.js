
import { resetBody } from "./header.js";
import { homePage } from "./home.js";
import { loginPage } from "./login.js";

export function registerPage() {
    const container = document.createElement('div')
    container.classList.add('register-background')
   

    const TextRegister = document.createElement('span')
    TextRegister.innerHTML = "REGISTER"
    TextRegister.classList = "register-text"
    
    const createInput = (text, id, type) =>{
        const element = document.createElement('input')
        element.classList = "register-input"
        element.placeholder = text
        element.id = id
        element.name = id
        element.autocomplete = true
        element.required = true
        element.type = type
        return element;
    };

    const registerButton = document.createElement('button')
    registerButton.classList = "buttonRegister"
    registerButton.textContent = "Register now"
  
    const form = document.createElement('form');
    form.id = "registerForm"

    const passwordVerif = document.createElement('div')
    passwordVerif.className = "verifpassword"
    passwordVerif.style.display = "none"

    const createDenyPassword = (sentence, keyword) => {
        const element = document.createElement('div')
        element.id = "test" + keyword
        element.innerHTML = "❌" + sentence
        return element
    }

    passwordVerif.append(
        createDenyPassword("majuscule", "maj"),
        createDenyPassword("minuscules", "min"),
        createDenyPassword("chiffres", "num"),
        createDenyPassword("8 characteres", "len")
    )

    const registerNames = document.createElement('div')
    registerNames.classList.add('registerNames')
    registerNames.append(
        createInput("First Name", "FirstName", "name"),
        createInput("Last Name", "LastName", "name"),
    )

    const createGender = (text, id, type) => {
        const genders = document.createElement('input')
        genders.type = type
        genders.id = id
        genders.text = text
        genders.name = "lucasmoncrush"
        if (id === "male") {
            genders.checked = true
        }
        return genders
    }

    const createLabel = (textContent, id) => {
        const label = document.createElement('label')
        label.textContent = textContent
        label.id = id
        return label
    }

    const registerLabelRadio = document.createElement('fieldset')
    registerLabelRadio.classList.add('registerLabelRadio')
    registerLabelRadio.append(
        createGender("Male", "male", "radio" ),
        createLabel("Male", "maleLabel"),
        createGender("Female", "female", "radio"),
        createLabel("Female", "femaleLabel"),
    )

    const registerInfo = document.createElement('div')
    registerInfo.classList.add('registerInfo')
    registerInfo.append(
        createInput("Age", "age", "age"),
        registerLabelRadio
    )

    const registerTopDiv = document.createElement('div')
    registerTopDiv.classList.add('registerTopDiv')
    registerTopDiv.append(
        registerNames,
        registerInfo
    )

    const registerDiv = document.createElement('div')
    registerDiv.classList ="register-buttons-div"
    registerDiv.append(
        registerTopDiv,
        createInput("Pseudo", "name", "name"), 
        createInput("Email", "mail", "mail"),
        createInput("Password", "password", "password"),
        passwordVerif,
        createInput("Password (again)", "password2", "password")
    )

    form.append(registerDiv, registerButton)

    form.addEventListener('submit', function(event) {
        event.preventDefault();
    
        const username = document.getElementById('name').value;
        const password = document.getElementById('password').value;
        const password2 = document.getElementById('password2').value;
        const mail = document.getElementById('mail').value;

        const passwordregex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.{8,})/;

        const mailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

        const usernameRegex = /^[0-9a-zA-Z]{4,16}$/

        if (!usernameRegex.test(username)) {
            alert("username isn't valid")
            return
        }
        
        if (password != password2) {
            alert('Passwords do not match!');
            return
        }

        if (!mailRegex.test(mail)) {
            alert("mail isn't valid")
            return
        }

        if (!passwordregex.test(password)){
            alert("password isn't valid")
            return
        }

    
        fetch('/register', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username: username, password: password, mail: mail })
        })
        .then(response => response.json())
        .then(data => {
            alert(data.message)
            return
        })
        .catch(error =>{
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
            });
        });
    });
    


    



    container.append(TextRegister,form);
    
    document.body.appendChild(container);

    document.getElementById('password').addEventListener('input', function() {
        passwordVerif.style.display = "flex"
        if (this.value == ""){
            passwordVerif.style.display = "none"
        }
        const password = this.value;
        if (/[A-Z]/.test(password)) {
            document.getElementById("testmaj").innerHTML = "✅ majuscule"
        } else {
            document.getElementById("testmaj").innerHTML = "❌ majuscule"
        } 
        if (/[a-z]/.test(password)) {
            document.getElementById("testmin").innerHTML = "✅ minuscule"
        } else {
            document.getElementById("testmin").innerHTML = "❌ minuscule"
        }
        if (/[0-9]/.test(password)) {
            document.getElementById("testnum").innerHTML = "✅ number"
        } else {
            document.getElementById("testnum").innerHTML = "❌ number"
        }
        if (/^.{8,}$/.test(password)) {
            document.getElementById("testlen").innerHTML = "✅ 8 character"
        } else {
            document.getElementById("testlen").innerHTML = "❌ 8 character"
        }
    });
}


