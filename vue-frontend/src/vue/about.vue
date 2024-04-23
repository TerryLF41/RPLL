<script setup>
import Header from '../components/header.vue'

// Retrieve user data from session storage
const userData = JSON.parse(sessionStorage.getItem('userData'));

// If user data is not present, redirect to the login page
if (!userData) {
  alert('You haven\'t logged in yet.');
  window.location.href = 'login.html'; // Redirect to the login page
}

</script>

<template>
    <main>
        <nav class="navbar">
            <Header />
        </nav>
        <div class="content-top">
            <img src='../assets/background/chatters-logo.png'>
            <h2>Diskusi Tanpa Batas</h2>
        </div>
        <div class="content-about">
            <h3>Selamat datang di Chatters Forum!</h3>
            <p>Chatters Forum adalah sebuah tempat dimana Anda dapat mendiskusikan apa saja yang Anda inginkan tanpa batasan topik </p>
            <p>Ingin membahas mengenai suatu topik baru? Tinggal buat saja!</p>
            <p>Kirimkan request topik beserta deskripsinya dan tunggu approval Admin</p>
            <p>Setelah disetujui, maka Anda bisa mengepost apapun yang berkaitan dengan topik tersebut</p>
            <h3><b>Visi-Misi</b></h3>
            <p>Visi</p>
            <p>Forum kami mempunyai visi untuk menjadi wadah diskusi dimana pengguna dapat mengekspresikan pendapat mereka mengenai suatu topik dengan bebas</p>
            <p>Misi</p>
            <p>1. Menyediakan wadah diskusi bagi para pengguna forum kami</p>
            <p>2. Menjadi sarana penyebaran informasi bagi para pengguna forum kami</p>
            <p>3. Mempromosikan bentuk media sosial forum yang sudah mulai berkurang eksistensinya</p>
            <h3>Tim Penyusun</h3>
            <p>1121004 - Jericho Kuskanto</p>
            <p>1121015 - Terry Louis Fujinami</p>
            <p>1121018 - Friendly Sejati Bunardi</p>
            <p>1121024 - Jonathan Hadiwijaya Ang</p>
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

h3, h2, p {
  color: white;
  text-decoration: none;
}

h2 {
  margin-top: 0.5em;
  margin-bottom: 1.5em;
  text-align: center;
}


img {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}

</style>