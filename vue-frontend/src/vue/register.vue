<template>
    <main>
    <form name="register" method="post">
        <div class="register">
            <h2>Register</h2>
            <label for="username">Username</label><br>
            <input type="text" name="username" id="username" required placeholder="Input Username" ><br>
            <label for="username">Password</label><br>
            <input type="password" name="password" id="password" required placeholder="Input Password"><br>
            <label for="username">Confirm Password</label><br>
            <input type="password" name="confirmPassword" id="confirmPassword" required placeholder="Confirm Password"><br>
            <label for="email">Email</label><br>
            <input type="email" name="email" id="email" required placeholder="example123@gmail.com"><br>
            <button id="submit" @click="register">Register</button><br>
        </div>
    </form>
    <button @click="registerTest">RegisterTest</button><br>
    </main>
</template>

<script setup>
    // Check password dengan confirm password
    function checkPassword(password,confirmPassword) {
        return Boolean(password === confirmPassword)
    }

    async function registerTest(){
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;
        var confirmPassword = document.getElementById("confirmPassword").value;
        var email = document.getElementById("email").value;

        var passwordMatches = checkPassword(password,confirmPassword)

        if(passwordMatches){
            console.log("Matches");
        } else {
            console.log("Doesn't match");
        }
    }
    // Post data dari form ke API untuk register user
    async function register() {
        // Ambil data dari form
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;
        var confirmPassword = document.getElementById("confirmPassword").value;
        var email = document.getElementById("email").value;

        // Check kesamaan password
        var passwordMatches = checkPassword(password,confirmPassword)

        // Kalau sesuai, kirim request
        if(passwordMatches){
            fetch('http://localhost:8181/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },    
                body: new URLSearchParams({
                    'username': username,
                    'password': password,
                    'email': email
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
            alert('Register berhasil!')
        } else {
            // Kalau tidak, tampilkan alert box
            alert('Password dengan konfirmasi password tidak sama!')
        }
    }
</script>

<style scoped>
body{
    background-image:url("../upload/media/bg-register.jpg");
    background-repeat: no-repeat;
    background-size: cover;
}
.register {
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
.register input {
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
.register button {
    border: 2px solid #303030;
    background-color: rgb(22, 192, 79);
    color: white;
    padding: 12px 20px;
    border: none;
    cursor: pointer;
    float: left;
    width: 100%;
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

