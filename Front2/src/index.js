import { RenderHomePage } from "./Pages/HomePage.js";
import { RenderLoginPage } from "./Pages/LoginPage.js";
import { CreatePostPage } from "./Pages/CreatePostPage.js";
import { RenderRegisterPage } from "./Pages/Register.js";
import { RenderPostPage } from "./Pages/PostPage.js";


export const Router = () => {
  const path = window.location.hash.slice(1);

  if (path.startsWith("/post/")) {
    const postId = path.split("/")[2];
    RenderPostPage(postId);
  } else {
    switch (path) {
      case "/home":
        RenderHomePage();
        break;
      case "/login":
        RenderLoginPage();
        break;
      case "register":
        RenderRegisterPage();
        break;
      case "create-post":
        CreatePostPage();
        break;
      default:
        RenderHomePage();
    }
  }
};


// Call the router initially and on every navigation
window.addEventListener("DOMContentLoaded", Router);

window.addEventListener("hashchange", Router);

// Handle link clicks to prevent default navigation
document.addEventListener("click", (e) => {
  const target = e.target;

  if (target.tagName === "A" && target.origin === window.location.origin) {
    e.preventDefault();
    const href = target.getAttribute("href");
    window.location.hash = href; // Update the hash instead of pushState
    Router();
  }
});
