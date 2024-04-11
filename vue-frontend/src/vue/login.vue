<template>
  <main>
    <form @submit.prevent="login" name="login" method="post">
      <div class="login">
        <h2>Login</h2>
        <label for="email">Email</label><br>
        <input type="text" name="email" id="email" required placeholder="Input Email" :value="rememberedEmail"><br>
        <label for="password">Password</label><br>
        <input type="password" name="password" id="password" required placeholder="Input Password" :value="rememberedPassword"><br>
        <input type="checkbox" id="remember" name="remember" value="true">Remember Me<br>
        <button type="submit" id="login">Login</button>
        <button type="button" @click="goToRegister">Register</button>
      </div>
    </form>
  </main>
</template>

<script setup>
import { setCookie, getCookie, deleteCookie } from '../utils'; // Import cookie functions

const rememberedEmail = getCookie('email') || ''; // Retrieve remembered email from cookie
const rememberedPassword = getCookie('password') || ''; // Retrieve remembered password from cookie

async function login() {
  var email = document.getElementById("email").value;
  var password = document.getElementById("password").value;
  const rememberMe = document.getElementById("remember").checked;

  const response = await fetch('http://localhost:8181/login', {
    method: "POST",
    body: new URLSearchParams({
      "email": email,
      "password": password
    }),
    headers: {
      "Content-Type": "application/x-www-form-urlencoded"
    }
  });

  if (response.ok) {
    const data = await response.json();
    console.log(data);
    if (data.status == '200') {
      // Set persistent cookies if remember me is checked
      if (rememberMe) {
        setCookie('email', email, 7); // Set email cookie for 7 days
        setCookie('password', password, 7); // Set password cookie for 7 days
      }

      console.log("Login successful!");
      alert('Login successful!');
      
      window.open('/homepage.html', '_self');
    } else {
      console.error("Login failed:", data.message);
      alert('Login gagal, password/username salah'); // Handle error message
    }
  }
}

function goToRegister() {
  window.open('register.html', '_self');
}
</script>


<style scoped>
body {
  background: url("../src/assets/background/bg-login.jpg") no-repeat center center fixed;
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
