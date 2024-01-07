import { CreateNavbar } from "../Components/Navbar.js";
import { Post } from "../Components/Post.js";
import { CreateElAddClassAddTextAddhref } from "../Helpers/CreateElements.js";
import { Router } from "../index.js";

export const RenderHomePage = () => {
  const container = document.querySelector(".root");
  container.innerHTML = "";
  console.log("Home Page");
  CreateNavbar();

  getMainPageData();
};

function getMainPageData() {
  // Make a request to the backend to check if the user is logged in
  fetch("http://localhost:8080/", {
    method: "GET",
    credentials: "include",
  })
    .then((response) => response.json())
    .then((data) => {
      // Check the response from the server
      console.log(data);
      if (data.isLogged) {
        showLoggedData(data);
      } else {
        showNotLoggedData();
        console.log("User is not logged in");
      }
    })
    .catch((error) => {
      console.error("Error checking login status:", error);
    });
}

function showLoggedData(data) {
  const container = document.querySelector(".root");
  container.innerHTML = "";
  if (data.isLogged) {
    // User is logged in, set a flag in localStorage
    localStorage.setItem("isLogged", "true");
    localStorage.setItem("UserNickname",data.user.nickname)
  }
  CreateNavbar(data.isLogged);
  const homeContainer = CreateElAddClassAddTextAddhref(
    "div",
    "homepage-container"
  );
  container.append(homeContainer);

if(data.posts){
  data.posts.reverse().forEach((post) => {
    const postElement = Post(post.id, post.author_id, post.title, post.content, post.created_at);
    homeContainer.appendChild(postElement);
  });

  // Attach a single click event listener to the homeContainer
  homeContainer.addEventListener("click", (e) => {
    const clickedPost = e.target.closest('.post-container');
    if (clickedPost) {
      // Extract post ID from the clicked post and navigate to the post page
      const postId = clickedPost.dataset.postId;
      window.location.hash = `#/post/${postId}`;

      Router()
      console.log(postId);
    }
  });
}

}

function showNotLoggedData() {
  const container = document.querySelector(".root");
  container.innerHTML = "";
  console.log("Home Page");
  CreateNavbar();

  const homeContainer = CreateElAddClassAddTextAddhref(
    "div",
    "homepage-container-notLogged"
  );
  const notLoggedContainer = CreateElAddClassAddTextAddhref(
    "div",
    "notLogged-container"
  );
  const notLoggedInfo = CreateElAddClassAddTextAddhref(
    "p",
    "notLoggedText",
    "Plese login to see the content"
  );
  const loginLink = CreateElAddClassAddTextAddhref(
    "a",
    "Homepage-login-link",
    "Login",
    "/login"
  );

  notLoggedContainer.append(notLoggedInfo,loginLink);
  homeContainer.append(notLoggedContainer);
  container.append(homeContainer);
}
