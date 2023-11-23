const header = document.querySelector('.header')
const loggedinfo =  document.querySelector('.loggedinfo')
const notloggedinfo =  document.querySelector('.notloggedinfo')
const logoutButton = document.querySelector(".logoutButton");


// function getMainPageData(){
//     fetch('http://localhost:8080/')
//     .then((response)=>{
//         if (!response.ok) {
//             throw new Error('Network response was not ok');
//         }
//         const contentType = response.headers.get('Content-Type');
//         if (!contentType || !contentType.includes('application/json')) {
//             throw new Error('Expected JSON response from the server');
//         }
//         return response.json();
//     })
//     .then((data)=>{
//         console.log(data);
//         header.innerHTML = data.message
//     })
//     .catch((error) => {
//         console.error('Error during fetch operation:', error);
//     });


// }

async function getMainPageData(){
    try{
        const response = await fetch('http://localhost:8080/', {
            credentials: 'include',})

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const contentType = response.headers.get('Content-Type');
        if (!contentType || !contentType.includes('application/json')) {
            throw new Error('Expected JSON response from the server');
        }

        const data = await response.json()
        if (data.isLoggedIn) {
            // Пользователь авторизован
           
            notloggedinfo.style.display = 'none';
            loggedinfo.style.display = 'block';
            header.innerHTML = data.userDetails
        } else {
            // Пользователь не авторизован
            notloggedinfo.style.display = 'block';
            loggedinfo.style.display = 'none';
        }
        console.log(data);
}catch (error){
    console.log('Error during fetch operation:', error);
}

}

getMainPageData()

logoutButton.addEventListener("click", function() {
    console.log("click");
    // Выполнение запроса на сервер для logout
    fetch('http://localhost:8080/logout', {
        method: 'POST', // или другой метод, который вы используете
        credentials: 'include' // важно для отправки куки
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Logout request failed');
        }
        
        // Перенаправление на главную страницу после успешного выхода
        location.reload()
    })
    .catch(error => {
        console.error('Logout error:', error);
    });
});
