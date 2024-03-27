<script setup>
import Header from '../components/header.vue'
</script>

<template>
    <main>
        <nav class="navbar">
            <Header />
        </nav>
        <h1>Welcome to</h1>
        <img src='../assets/background/chatters-logo.png'>
        <h2>Diskusi Tanpa Batas</h2><br><br>
        <p id='desc'>Chatters Forum adalah sebuah tempat dimana Anda dapat mendiskusikan apa saja yang Anda inginkan tanpa batasan topik </p><br>
        <h2>Topik Terbaru</h2>
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
a, h2{
    color: white;
    text-decoration: none;
}

input[type=submit]{
    background-color: grey;
    border: 1px solid;
    cursor: pointer;
    text-decoration: none;
    color: white;
}
h1,h2,#desc {
    text-align: center;
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
.show{
    background-color: grey;
    border: none;
    color: white;
    padding: 15px 15px;
    text-decoration: none;
    margin: 4px 2px;
    cursor: pointer;
    font-size: 16px;
}
.show:hover {
    opacity: 75%;
}
button.show {
    width: 165px;
}
.option{
    margin-left: 33%;
    padding-top: 30px;
    text-align: left;
    clear: both;
    display: block;
    padding-bottom: 30px;
}
h2 {
    clear: both;
}


</style>

