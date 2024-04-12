<script setup>
import { setCookie, getCookie, deleteCookie } from '../utils'; // Import cookie functions

// Retrieve and parse user data from session storage
const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
console.log(userDataParsed.profilePicture);

// Logout function
function logout() {
  // Remove user data from session storage
  sessionStorage.removeItem('userData');
  
  // Remove token cookie
  setCookie('token', '', -1);

  // Redirect to login page
  window.open('login.html', '_self');
}
</script>

<template>
  <div class='logo'><a href='homepage.html'><img src='../assets/background/chatters-logo.png'></a></div>
  <ul class="nav-links">
    <div class="menu">
      <li><a href="topic.html">Catalog</a></li>
      <li><a href="about.php">About</a></li>
      <li><a @click="logout">Log Out</a></li> <!-- Add event handler for logout -->
      <div class="profile-picture">
        <a href="editProfile.html"><img class="profile-picture" :src="userDataParsed.profilePicture"></a>
      </div>
    </div> 
  </ul>
</template>


<style scoped>
body {
    margin: 0;
    padding: 0;
}
* {
    font-family: Roboto;
}

.menu li a {
    color: #fff; /* Menetapkan warna teks menjadi putih */
}

a {
    text-decoration: none;
}
li {
    list-style: none;
}
nav {
    top: 0;
}
.navbar {
    height: 80px;
    margin: 0;
    padding: 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 15px;
    background-color: #1c1c1c;
    color: #fff;
}
.nav-links a {
    color: #fff;
}
.logo img{
    height: 60px;
    width: 240px;
}
.logo a {
    text-decoration: none;
}
.menu {
    display: flex;
    gap: 1em;
    font-size: 18px;
}
.menu li:hover:not(:last-child) {
    background-color: #666;
    border-radius: 10px;
    transition: 0.3s ease;
}
.menu li {
    padding: 10px 10px;
}
li a {
    display: inline-block;
}
.profile-picture {
    height: 45px;
    width: 45px;
    border-radius: 50%;
}

</style>

