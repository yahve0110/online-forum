import {CreateNavbar} from "../Components/Navbar.js";
import {
    CreateButtonTypeInnerTextClass,
    CreateElAddClassAddTextAddhref,
    CreateInputTypePlaceholderName
} from "../Helpers/CreateElements.js";
import { Router } from "../index.js";


export function RenderLoginPage() {
    //get main container
    const container = document.querySelector(".root")

    //clear page
    container.innerHTML = ""
    //add navbar
    CreateNavbar()
    //create container for login page
    const loginPageContainer = CreateElAddClassAddTextAddhref('div', "loginPageContainer", "")

    const loginForm = CreateElAddClassAddTextAddhref("form","loginForm")
     const userNameInput =  CreateInputTypePlaceholderName("input","text","Username","identifier")
    const passwordInput =  CreateInputTypePlaceholderName("input","password","Password","password")

    //add to login form
    loginForm.append(userNameInput)
    loginForm.append(passwordInput)


//create submit form btn
  const submitFormBtn =  CreateButtonTypeInnerTextClass("button","submit","Log in")

    loginForm.append(submitFormBtn)

    //add form to login container
    loginPageContainer.append(loginForm)
    //add login container to main container
    container.append(loginPageContainer)

    //add listener to form
    loginForm.addEventListener('submit', (e) => {
        e.preventDefault();
        submitLoginFormFn(userNameInput.value, passwordInput.value);
    });
}

async function submitLoginFormFn(identifier, password) {
    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: "include",
            body: JSON.stringify({
                identifier: identifier,
                password: password,
            }),
            
        });

        window.location.hash = '#home'
        Router()

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
       
    

        // For example, redirect to another page on successful login
        if (data.isLogged) {
        console.log("logged");
        } else {
            console.log("ne zaregan")
        }
    } catch (error) {
        console.error('Fetch Error:', error);
    }
}

