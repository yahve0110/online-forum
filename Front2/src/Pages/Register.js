import { CreateNavbar } from "../Components/Navbar.js";
import {
  CreateElAddClassAddTextAddhref,
  CreateInputTypePlaceholderName,
  CreateButtonTypeInnerTextClass,
} from "../Helpers/CreateElements.js";
import { Router } from "../index.js";

export function RenderRegisterPage() {
  // Get the main container
  const container = document.querySelector(".root");

  // Clear the page
  container.innerHTML = "";

  // Add the navbar
  CreateNavbar();

  // Create the registration page container
  const registerPageContainer = CreateElAddClassAddTextAddhref(
    "div",
    "registerPageContainer",
    ""
  );

  // Create input fields for registration form
  const usernameInput = CreateInputTypePlaceholderName(
    "input",
    "text",
    "Username",
    "nickname"
  );

  const ageInput = CreateInputTypePlaceholderName(
    "input",
    "text",
    "Age",
    "age"
  );

  const genderSelect = document.createElement("select");
  genderSelect.name = "gender";

  // Add options for the gender select dropdown
  const maleOption = document.createElement("option");
  maleOption.value = "male";
  maleOption.text = "Male";

  const femaleOption = document.createElement("option");
  femaleOption.value = "female";
  femaleOption.text = "Female";

  const otherOption = document.createElement("option");
  otherOption.value = "other";
  otherOption.text = "Other";

  // Add options to the select dropdown
  genderSelect.add(maleOption);
  genderSelect.add(femaleOption);
  genderSelect.add(otherOption);

  const firstNameInput = CreateInputTypePlaceholderName(
    "input",
    "text",
    "First Name",
    "first_name"
  );

  const lastNameInput = CreateInputTypePlaceholderName(
    "input",
    "text",
    "Last Name",
    "last_name"
  );

  const emailInput = CreateInputTypePlaceholderName(
    "input",
    "email",
    "Email",
    "email"
  );

  const passwordInput = CreateInputTypePlaceholderName(
    "input",
    "password",
    "Password",
    "password"
  );

  // Create a submit button for the registration form
  const submitFormBtn = CreateButtonTypeInnerTextClass(
    "button",
    "submit",
    "Register"
  );

    // Add an event listener to the form for handling registration submission
    const registrationForm = document.createElement("form");
    registrationForm.addEventListener("submit", submitRegistrationForm);

  // Add form elements to the registration page container
  registrationForm.append(
    usernameInput,
    ageInput,
    genderSelect,
    firstNameInput,
    lastNameInput,
    emailInput,
    passwordInput,
    submitFormBtn
  );

  // Add the registration page container to the main container
  container.append(registerPageContainer);


  registerPageContainer.append(registrationForm);
}

// Function to handle registration form submission
async function submitRegistrationForm(event) {
  event.preventDefault();

  // Retrieve form data
  const username = document.querySelector('input[name="nickname"]').value;
  const age = document.querySelector('input[name="age"]').value;
  const gender = document.querySelector('select[name="gender"]').value;
  const firstName = document.querySelector('input[name="first_name"]').value;
  const lastName = document.querySelector('input[name="last_name"]').value;
  const email = document.querySelector('input[name="email"]').value;
  const password = document.querySelector('input[name="password"]').value;

  // Create a user object with the form data
  const newUser = {
    nickname: username,
    age: age,
    gender: gender,
    first_name: firstName,
    last_name: lastName,
    email: email,
    password: password,
  };

  try {
    // Make a request to the server's registration endpoint
    const response = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newUser),
    });

    if (response.ok) {
        console.log("Registration successful");
      
        // Log before and after Router() call
        console.log("Before Router() call");
        Router();
        console.log("After Router() call");
      
        window.location.hash = '#home';
      } else {
      // Registration failed
      console.error("Registration failed:", response.statusText);
      // Optionally, display an error message to the user
    }
  } catch (error) {
    console.error("Error during registration:", error);
    // Optionally, display an error message to the user
  }
}
