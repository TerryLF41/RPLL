<script setup>
import Header from '../components/header.vue'

// Retrieve user data from session storage
const userData = JSON.parse(sessionStorage.getItem('userData'));

// If user data is not present, redirect to the login page
if (!userData) {
  alert('You haven\'t logged in yet.');
  window.location.href = 'login.html'; // Redirect to the login page
}

// Ambil usertype dari session
const userType = userData.userType;
</script>

<template>
  <main>
    <nav class="navbar">
      <Header />
    </nav>
    <h1>Welcome to</h1>
    <img src='../assets/background/chatters-logo.png'>
    <h2>Diskusi Tanpa Batas</h2>
    <p id='desc'>Chatters Forum adalah sebuah tempat dimana Anda dapat mendiskusikan apa saja yang Anda inginkan tanpa batasan topik </p>

    <!-- Tombol-tombol untuk admin -->
    <div v-if="userType === 1" class="option">
      <a href="userlist.html" class="show">User List</a>
      <a href="reportlist.html" class="show">Report List</a>
    </div>
  </main>
</template>

<style scoped>
main {
  min-height: 80vh;
  width: 66%;
  margin: auto;
  padding: 0;
}

.catalog {
  border-radius: 25px;
  border: 2px solid;
  height: 20px;
  padding-top: 5px;
  width: 70px;
  text-align: center;
  background-color: #A9A9A9;
}

a, h2, p {
  color: white;
  text-decoration: none;
}

input[type=submit] {
  background-color: grey;
  border: 1px solid;
  cursor: pointer;
  text-decoration: none;
  color: white;
}

h1, h2, #desc {
  text-align: center;
}

h2 {
  margin-top: 0.5em;
  margin-bottom: 1.5em;
}

img {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}

form {
  float: left;
}

.option {
  display: flex;
  /* justify-content: space-around; Tombol-tombol akan berderet dari kanan ke kiri */
  justify-content: center;
}

.show {
  background-color: grey;
  border: none;
  color: white;
  padding: 15px 15px;
  text-decoration: none;
  margin: 4px 2px;
  cursor: pointer;
  font-size: 16px;
  width: 30%; /* Lebar tombol */
  text-align: center; /* Teks tengah */
}

.show:hover {
  opacity: 0.75;
}
</style>
