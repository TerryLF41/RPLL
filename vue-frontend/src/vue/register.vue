<template>
  <main>
    <form @submit.prevent="register" name="register" method="post">
      <div class="register-container">
        <h2>Register</h2>
        <div class="input-group">
          <label for="email">Email</label>
          <input type="email" name="email" id="email" required placeholder="example123@gmail.com">
        </div>
        <div class="input-group">
          <label for="username">Username</label>
          <input type="text" name="username" id="username" required placeholder="Input Username">
        </div>
        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" name="password" id="password" required placeholder="Input Password">
        </div>
        <div class="input-group">
          <label for="confirmPassword">Confirm Password</label>
          <input type="password" name="confirmPassword" id="confirmPassword" required placeholder="Confirm Password">
        </div>
        <button id="submit">Register</button>
      </div>
    </form>
  </main>
</template>

<script setup>
  // Check password dengan confirm password
  function checkPassword(password, confirmPassword) {
    return Boolean(password === confirmPassword);
  }

  // Post data dari form ke API untuk register user
  async function register() {
    // Ambil data dari form
    var email = document.getElementById("email").value;
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var confirmPassword = document.getElementById("confirmPassword").value;
    

    // Check kesamaan password
    var passwordMatches = checkPassword(password, confirmPassword);

    // Kalau sesuai, kirim request
    if (passwordMatches) {
      fetch("http://localhost:8181/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: new URLSearchParams({
          username: username,
          password: password,
          email: email,
        }),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((response) => {
          // Print response di console
          console.log("Response:", Response);
        })
        .catch((error) => {
          // Print error di console
          console.error("Error:", error);
        });
      alert("Register berhasil!");
      window.open("/login.html", "_self");
    } else {
      // Kalau tidak, tampilkan alert box
      alert("Password dengan konfirmasi password tidak sama!");
    }
  }
</script>

<style scoped>
  body {
    background: url("../src/assets/background/bg-register.jpg") no-repeat center center fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
    margin: 0;
    padding: 0;
    font-family: "Roboto", sans-serif;
  }

  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    font-family: "Roboto", sans-serif;
    font-weight: bolder;
  }

  .register-container {
    background-color: rgba(0, 0, 0, 0.8);
    border-radius: 8px;
    box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.1);
    padding: 40px;
    width: 350px;
    max-width: 90%;
    color: #fff;
    
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
    color: #fff;
    margin-bottom: 5px;
  }

  .input-group input {
    width: 94%;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }

  button {
    width: 100%;
    padding: 12px 20px;
    font-size: 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
    background-color: #2196f3;
    color: #fff;
  }

  button:hover {
    background-color: #0b7dda;
  }
</style>

