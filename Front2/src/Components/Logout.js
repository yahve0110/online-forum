export async function LogoutFunction() {
  try {
    // Make a request to the server's logout endpoint
    const response = await fetch("http://localhost:8080/logout", {
      method: "POST", // You might want to use 'GET' or 'DELETE' depending on your server implementation
      credentials: "include", // Include credentials (cookies) in the request
    });

    if (response.ok) {
      // Logout successful on the server side
      // Perform any client-side cleanup, if necessary
      console.log("Logout successful");
      localStorage.removeItem("isLogged");
      localStorage.removeItem("UserNickname");
      return true; // Indicate success
    } else {
      // Handle logout failure
      console.error("Logout failed:", response.statusText);
      return false; // Indicate failure
    }
  } catch (error) {
    console.error("Error during logout:", error);
    return false; // Indicate failure
  }
}
