import {
  CreateButtonTypeInnerTextClass,
  CreateElAddClassAddTextAddhref,
} from "../Helpers/CreateElements.js";
import { RenderHomePage } from "../Pages/HomePage.js";
import { LogoutFunction } from "./Logout.js";

const container = document.querySelector(".root");

export function CreateNavbar(isLogged = false) {
  if (isLogged) {
    const navDiv = CreateElAddClassAddTextAddhref("div", "navbar");
    const homeLink = CreateElAddClassAddTextAddhref(
      "a",
      "nav-link",
      "Home",
      "/home"
    );
    const logoutBtn = CreateButtonTypeInnerTextClass(
      "button",
      "",
      "Logout",
      "logout-btn"
    );

    logoutBtn.addEventListener("click", async () => {
      const logoutResult = await LogoutFunction();

      if (logoutResult) {
        // Logout was successful
        // You can redirect the user, update UI, or perform any other action
        RenderHomePage();
      } else {
        // Logout failed
        // You can display an error message or take other actions
        console.error("Logout failed. Please try again.");
      }
    });


    const createPostLink = CreateElAddClassAddTextAddhref("a","create-post-link","Create-post","create-post")

   

    navDiv.append(homeLink);
    navDiv.append(logoutBtn);
    navDiv.append(createPostLink)
    container.append(navDiv);
  } else {
    const navDiv = CreateElAddClassAddTextAddhref("div", "navbar");
    const homeLink = CreateElAddClassAddTextAddhref(
      "a",
      "nav-link",
      "Home",
      "/home"
    );
    const loginLink = CreateElAddClassAddTextAddhref(
      "a",
      "nav-link",
      "Login",
      "/login"
    );

    const registerLink = CreateElAddClassAddTextAddhref(
      "a",
      "nav-link",
      "Register",
      "register"
    );

    navDiv.append(homeLink);
    navDiv.append(loginLink);
 
    navDiv.append(registerLink);
    container.append(navDiv);
  }
}
