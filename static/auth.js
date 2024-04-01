document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from submitting normally

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    if (username === 'admin' && password === 'root') {
        console.log('successful');
        // You can redirect the user or perform other actions here
    } else {
        console.log('Login failed');
        // Handle failed login attempt
    }
});
// rockers