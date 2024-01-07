import { CreateNavbar } from "../Components/Navbar.js";
import {
  CreateElAddClassAddTextAddhref,
  CreateInputTypePlaceholderName,
  CreateButtonTypeInnerTextClass,
  CreateTextareaPlaceholderName,
} from "../Helpers/CreateElements.js";
import { Router } from "../index.js";

export function CreatePostPage() {
  // Get the main container
  const container = document.querySelector(".root");

  // Clear the page
  container.innerHTML = "";

  
  // Get data if the user is logged in
  const isLogged = localStorage.getItem("isLogged");

  // Add the navbar
  CreateNavbar(isLogged);

  // Create the post creation page container
  const createPostPageContainer = CreateElAddClassAddTextAddhref(
    "div",
    "createPostPageContainer",
    ""
  );

  // Create input fields for the post creation form
  const titleInput = CreateInputTypePlaceholderName(
    "input",
    "text",
    "Title",
    "title"
  );

  const contentTextarea = CreateTextareaPlaceholderName("Content", "content");

  // Create a submit button for the post creation form
  const submitPostBtn = CreateButtonTypeInnerTextClass(
    "button",
    "submit",
    "Create Post"
  );

  // Create the post creation form
  const postCreationForm = document.createElement("form");

  // Add an event listener to the form for handling post creation submission
  postCreationForm.addEventListener("submit", async function (event) {
    event.preventDefault();
    await submitPostForm(postCreationForm);
  });

  // Add form elements to the post creation page container
  postCreationForm.append(titleInput, contentTextarea, submitPostBtn);

  // Add the post creation page container to the main container
  container.append(createPostPageContainer);

  createPostPageContainer.append(postCreationForm);
}

// Function to handle post creation form submission
async function submitPostForm(form) {
  // Retrieve form data
  const title = form.querySelector('input[name="title"]').value;
  const content = form.querySelector('textarea[name="content"]').value;

  // Create a post object with the form data
  const postData = {
    title: title,
    content: content,
  };

  try {
    // Make a request to the server's create post endpoint
    const response = await fetch("http://localhost:8080/create-post", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(postData),
    });

    if (response.ok) {
      // Post creation successful
      console.log("Post created successfully");
      window.location.hash = '#home';
      Router();
    } else {
      // Post creation failed
      console.error("Post creation failed:", response.statusText);
      // Optionally, display an error message to the user
    }
  } catch (error) {
    console.error("Error during post creation:", error);
    // Optionally, display an error message to the user
  }
}
