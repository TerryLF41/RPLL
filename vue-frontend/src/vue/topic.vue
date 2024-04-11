<script>
import Header from '../components/header.vue'
</script>

<template>
        <main>
            <nav class="navbar">
                <Header />
            </nav>
            <h1>Daftar Topik</h1>
            <button type="button" @click="dis">Get Topic</button>     
            <div class="container d-flex justify-content-center">
                <ul class="list-group mt-5 text-white">
                    <li class="list-group-item d-flex justify-content-between align-content-center" @click="goToThread" v-for="item in temp">
                        <div class="d-flex flex-row">
                            <img src="../assets/userUploadedFiles/userProfile/default.png" width="100" />
                            <div class="ml-2 topicDesc">
                                <h6 class="mb-0">{{item.topicTitle}}</h6>
                                <div class="about">
                                    <span>{{item.topicDesc}}</span>
                                </div>
                            </div>
                        </div>
                    </li>
                </ul>
            </div>
        </main>
</template>

<script setup>
let temp =[];
  async function getTopic() {
    const response = await fetch('http://localhost:8181/topic', {
        method: "GET",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
    });
    if (response.ok) {
      const data = await response.json()
      if (data.status == '200'){
            temp = Object.assign([], data.data);
      } else {
            console.error("Failed!", data.message);
      }
    }
    console.log(temp[0])
  }

  function displayTopic(){
    const data = getTopic();
    console.log(temp[0])
    //var textedTopic = JSON.stringify(topicJSON.data);
  }

  function goToThread() {
    window.open('homepage.html?threadNo=1','_self');
  }

  window.onload = displayTopic();
</script>

<style scoped>
.list-group{
	width: 100% !important;
    border-radius: 0;
}
.list-group-item{
	margin-top:15px;
	cursor: pointer;
    border-style: none;
	transition: all 0.3s ease-in-out;
    background-color: #ADA7A7;
}
.list-group-item:hover{
	opacity: 0.8;
}
.about span{
	font-size: 12px;
	margin-right: 10px;
}
.topicDesc{
    margin-left: 10px;
}

body {
    color: white;
    background-image:url("../upload/media/bg-topic.jpg");
    background-repeat: no-repeat;
    background-size: cover;
}
main {
    min-height: 80vh;
    width: 66%;
    margin: auto;
    padding: 0;
}
h1,h2,#desc {
    text-align: center;
}
.list{
    padding: 5px;
    width: 90%;
    margin:auto;
    border: 2px solid grey;
    border-radius: 10px;
}
table tr td a {
    margin-top: 10px;
    display: block;
    width: 100%;
    height: 100%;
    text-decoration: none;
}
a{
    color: white;
    text-decoration: none;
}
td,th{
    border-bottom:1px inset grey;
    height:30px;
}
tr:hover{
    background-color:#5c757e;
    color: Black;
}
h1{
    color: white;
    padding-left: 5%;
}
.topik{
    float:right;
    padding-right: 5%;
}
.add {
    text-align: center;
    cursor: pointer;
}
::placeholder {
    color: #cccccc;
}
.modal {
    position: fixed;
    z-index: 999;
    border: 2px solid black;
    padding: 2.5%;
    width: 35%;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    display: none; 
    background-color: #454545;
    font-family: Roboto;
    transform: translate(-50%, -50%);
    overflow: auto;
    top: 50%;
    left: 50%; 
}
.modal input, textarea {
    background-color: #696969;
    color: white;
    width: 100%;
    padding: 10px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #3f3f3f;
    box-sizing: border-box;
    font-family: Roboto;
}
.modal button {
    border: 2px solid black;
    background-color: #303030;
    color: white;
    padding: 12px 20px;
    border: none;
    cursor: pointer;
    float: left;
    width: 50%;
}
button:hover {
    opacity: 75%;
}
#cancel {
    background-color: #7a0800;
}
h2.title {
    text-align: center;
    margin-top: 0;
}
</style>