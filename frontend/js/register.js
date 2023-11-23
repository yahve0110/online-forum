const form = document.querySelector('.registerform')


function register() {
    // Get form values
    var nickname = document.getElementById("nickname").value;
    var age = document.getElementById("age").value;
    var firstName = document.getElementById("firstName").value;
    var lastName = document.getElementById("lastName").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    var gender = document.getElementById("gender").value;

    console.log("Nickname:", nickname);
    console.log("Age:", age);
    console.log("First Name:", firstName);
    console.log("Last Name:", lastName);
    console.log("Email:", email);
    console.log("Password:", password);
    console.log("Gender:", gender);

    // Create a user object with the form values

    var user = {
        "nickname": nickname,
        "age": age,
        "first_name": firstName,
        "last_name": lastName,
        "email": email,
        "password": password,
        "gender": gender,
    };

    console.log("User Object:", user);

    // Send a POST request to your registration endpoint
    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify(user),
    })
    .then(response => {
        if (!response.ok) {
            // Вывести текст ошибки, если есть
            return response.text().then(text => { throw new Error(`HTTP error! Status: ${response.status}, ${text}`); });
        }
        return response.json();
  
    })
    .then(data => {
        // Обрабатываем ответ, например, выводим сообщение об успехе
        console.log(data);
    })
    .catch((error) => {
        // Выведем ошибку в консоль
        console.error('Error during registration:', error);
    });
}

form.addEventListener('submit', (event) => {
    event.preventDefault(); // Prevent the default form submission
    register();
});
