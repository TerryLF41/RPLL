<script setup>
import { setCookie, getCookie, deleteCookie } from '../utils'; // Import cookie functions
import { logUserActivity } from '../activityLogger'; // Import user activity logger

// Retrieve and parse user data from session storage
const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
// If user data is not present, redirect to the login page
if (!userDataParsed) {
  alert('You haven\'t logged in yet.');
  window.location.href = 'login.html'; // Redirect to the login page
}
// console.log(userDataParsed.profilePicture);
const profilePicture = userDataParsed.profilePicture;
// console.log(profilePicture);
// Logout function
function logout() {
  // Remove user data from session storage
  sessionStorage.removeItem('userData');
  
  // Remove token cookie
  setCookie('token', '', -1);

  // Log logout activity as "Ban thread"
  logUserActivity("Logout",userDataParsed.userId);

  //Alert user
  alert("Goodbye")

  // Redirect to login page
  window.open('login.html', '_self');
}
</script>

<template>
  <div class='logo'><a href='homepage.html'><img src='../assets/background/chatters-logo.png'></a></div>
  <ul class="nav-links">
    <div class="menu">
      <li><a href="homepage.html">Home</a></li>
      <li><a href="topic.html">Catalog</a></li>
      <li><a href="about.php">About</a></li>
      <li id="logoutbutton"><a @click="logout">Log Out</a></li> <!-- Add event handler for logout -->
      <div class="profile-picture">
        <a href="profile.html"><img class="profile-picture" :src="profilePicture"></a>
      </div>
    </div> 
  </ul>
</template>


<style scoped>
body {
    margin: 0;
    padding: 0;
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
    font-family: Roboto;
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
    margin-top: auto;
    height: 40px;
    width: 40px;
    border-radius: 50%;
}
.menu #logoutbutton :hover{
    cursor: pointer;
}

.menu #logoutbutton a{
    color: #fff;
}

</style>

