<template>
  <main>
    <form @submit.prevent="login" name="login" method="post">
      <div class="login-container">
        <h2>Login</h2>
        <div class="input-group">
          <label for="email">Email</label>
          <input type="text" name="email" id="email" required placeholder="Input Email" :value="rememberedEmail">
        </div>
        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" name="password" id="password" required placeholder="Input Password" :value="rememberedPassword">
        </div>
        <div class="remember-me">
          <input type="checkbox" id="remember" name="remember" value="true">
          <label for="remember">Remember Me</label>
        </div>
        <button type="submit" id="login">Login</button>
        <button type="button" id="register" @click="goToRegister">Register</button>
      </div>
    </form>
  </main>
</template>

<script setup>
import { setCookie, getCookie, deleteCookie } from '../utils'; // Import cookie functions
import { logUserActivity } from '../activityLogger'; // Import user activity logger

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
    if (data.status == '200') {
      // Set persistent cookies if remember me is checked
      setCookie('token', data.token, 7);
      if (rememberMe) {
        setCookie('email', email, 7); // Set email cookie for 7 days
        setCookie('password', password, 7); // Set password cookie for 7 days
      }

      // Access user data from the response and do something with it
      const userData = data.data; // Assuming the user data is returned under 'data' key
      console.log(userData); // Log the user data

      // Store user data in session storage
      sessionStorage.setItem('userData', JSON.stringify(userData));
      // Replace the current history state with the homepage URL
      history.replaceState(null, '', '/homepage.html');
      // Add a new history entry pointing to the homepage URL
      history.pushState(null, '', '/homepage.html');

      // Log user login activity as "Login"
      logUserActivity("Login", userData.userId);

      console.log("Login successful!");
      // alert('Login successful!');
      
      window.open('/homepage.html', '_self');
    } else {
      console.error("Login failed:", data.message);
      alert('Login gagal, password/username salah'); // Handle error message
    }
  }
}

// Listen for the popstate event to handle navigation
window.addEventListener('popstate', function (event) {
  // Get the URL of the current page
  const currentPage = window.location.pathname;
  // Check if the current page is the login page
  if (currentPage === '/login.html') {
    // Redirect the user to the homepage
    window.location.replace('/homepage.html');
  }
});

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
  margin: 0;
  padding: 0;
  /* font-family: 'Roboto', sans-serif; */
}

main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  font-family: 'Roboto', sans-serif;
  font-weight: bolder;
}

.login-container {
  background-color: rgba(0, 0, 0, 0.8);
  border-radius: 8px;
  box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.1);
  padding: 40px;
  width: 350px;
  max-width: 90%;
  color:#fff;
}

h2 {
  text-align: center;
  margin-top: 0;
  margin-bottom: 20px;
}

.input-group {
  margin-bottom: 20px;
}

.input-group label {
  display: block;
  font-size: 14px;
  color:#fff;
  margin-bottom: 5px;
}

.input-group input {
  width: 94%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.remember-me {
  margin-bottom: 20px;
}

.remember-me label {
  font-size: 14px;
  color:#fff;
  margin-left: 5px;
}

button {
  width: 100%;
  padding: 12px 20px;
  font-size: 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

#login {
  background-color: #3f51b5;
  color: #fff;
}

#login:hover {
  background-color: #2a3f9d;
}

#register {
  background-color: #8bc34a;
  color: #fff;
}

#register:hover {
  background-color: #72a436;
}

</style>
