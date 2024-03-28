<template>
    <main>
        <form name="login" method="post">
            <div class="login">
            <h2>Login</h2>
                <label for="username">Username</label><br>
                <input type="text" name="username" id="username" required placeholder="Input Username"><br>
                <label for="password">Password</label><br>
                <input type="password" name="password" id="password" required placeholder="Input Password"><br>
                <input id="remember" type="checkbox" name="remember" value="true">Remember Me<br>
                <button @click="login">Login</button>
                <button id='register'type="reset" onclick="window.open('register.html','_self');">Register</button>
            </div>
        </form>
    </main>
</template>

<script>
async function login() {
    // Ambil data dari form
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    // Panggil API login
    fetch('http://localhost:8181/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },    
        body: new URLSearchParams({
            'username': username,
            'password': password,
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(response => {
        // Print response di console
        console.log('Response:', Response);
    })
    .catch(error => {
        // Print error di console
        console.error('Error:', error);
    });
    alert('Login berhasil!')
}
</script>

<style scoped>
body {
    background: url("../src/assets/background/bg-login.jpg") no-repeat center center
    fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
}
.login {
    position: fixed;
    z-index: 999;
    border: 2px solid black;
    padding: 2.5%;
    width: 35%;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    background-color: #1c1c1c;
    color: white;
    font-family: Roboto;
    transform: translate(-50%, -50%);
    overflow: auto;
    top: 50%;
    left: 50%; 
}
.login input {
    background-color: #1c1c1c;
    width: 100%;
    padding: 10px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #454545;
    box-sizing: border-box;
    font-family: Roboto;
    color: white;
}
.login button {
    border: 2px solid #303030;
    background-color: rgb(22, 192, 79);
    color: white;
    padding: 12px 20px;
    border: none;
    cursor: pointer;
    float: left;
    width: 50%;
}
#register {
    background-color: #1938be;
}
button:hover {
    opacity: 75%;
}
h2 {
    text-align: center;
    margin-top: 0;
}
#no-acc {
    display: block;
}
#remember {
    width: 20px;
}
</style>

