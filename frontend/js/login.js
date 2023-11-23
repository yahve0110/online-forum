const loginForm = document.querySelector(".loginform")


function login() {
    const loginData = {
        identifier: document.getElementById('nickname').value,
        password: document.getElementById('password').value,
    };

    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(loginData),
        credentials: 'include',
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Login failed');
        }
        console.log('Login successful');
        // Handle successful login, e.g., redirect to another page
    })
    .catch(error => {
        console.error('Login error:', error);
    });
}

loginForm.addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission
    login();
    window.location.href = "../index.html";
});


