import {CreateNavbar} from "../Components/Navbar.js";

export const AboutPage = () => {
    const container = document.querySelector(".root")

    container.innerHTML = ''
    console.log('about Page');
    CreateNavbar()
}