
import { registerPage } from "./register.js";
import { loginPage } from "./login.js"
import { resetBody } from "./header.js";




export function isNotConnected() {
    resetBody()
    document.body.style.overflow = 'hidden'
    const createTitleElement = (classList, innerHTML, h) => {
        const element = document.createElement(h);
        element.classList = classList;
        element.innerHTML = innerHTML;
        return element; 
    };

    const UpTitle = document.createElement('div');
    UpTitle.classList = "titleFirstPage";


    const BigTitle = document.createElement('div');
    BigTitle.classList = "BigTitle"


    const Title = createTitleElement("title", "welcome to", "h1");
    const Culture = createTitleElement("titleCulture", "Culture.", "h1");
    const UnderTitle = createTitleElement("underTitle", "place where you can learn about everything", "h2")

    const square = document.createElement('div');
    square.classList.add('button');

    const createButtonElement = (text) =>{
        const element = document.createElement('div')
        element.classList = "button"
        const squareText = document.createElement('span');
        squareText.classList = "text-button";
        squareText.textContent = text
        element.appendChild(squareText)
        
        element.addEventListener('click', function() {
        resetBody()
        if (text == "login") {
            loginPage()
        } else {
            registerPage()
            
        }
    
        });
        return element

    }

    UpTitle.append(Title, Culture);
    BigTitle.append(UpTitle, UnderTitle, createButtonElement("login"), createButtonElement("register"))
    document.body.appendChild(BigTitle);
}





