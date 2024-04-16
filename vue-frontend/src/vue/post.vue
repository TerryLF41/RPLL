<script>
import Header from '../components/header.vue'
import { ref } from 'vue';
import { onMounted } from 'vue';
</script>

<template>
    <main>
        <nav class="navbar">
            <Header />
        </nav>
        <h1>Daftar Post</h1>
        <button type="button" @click="showModal">Add New Thread</button> 
        <div class="container d-flex justify-content-center">
            <ul class="list-group mt-5 text-white" >
                <li class="list-group-item d-flex justify-content-between align-content-center" v-for="post in postList">
                    <div class="post d-flex flex-row">
                        <div class="profileUser">
                            <span>nama user</span><br>
                            <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                        </div>
                        <div class="ml-2 PostDesc">
                            <h6 class="mb-0"> {{ post.postDate }}</h6>
                            <div class="about">
                                <span>{{ post.postText }}</span>
                            </div>
                        </div>
                        <div class="ml-2 PostImage">
                            <h6 class="mb-0">No. {{ post.postNo }}</h6>
                            <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                        </div>
                    </div>
                    <ul class="list-group mt-5 text-white" >
                        <li class="list-group-item d-flex justify-content-between align-content-center" v-for="post in postList">
                            <div class="d-flex flex-row">
                                <div class="profileUser">
                                    <span>nama user</span><br>
                                    <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                                </div>
                                <div class="ml-2 PostDesc">
                                    <h6 class="mb-0"> {{ post.postDate }}</h6>
                                    <div class="about">
                                        <span>{{ post.postText }}</span>
                                    </div>
                                </div>
                                <div class="ml-2 PostImage">
                                    <h6 class="mb-0">No. {{ post.postNo }}</h6>
                                    <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                                </div>
                            </div>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </main>
    <!-- <div class="modal-thread" id="modal">
        <form class="form" method="POST">
            <h2 class="title">Add New Thread</h2>
            <label><b>Nama Topik</b></label><br>
            <input type="text" name="threadName" id="threadName" placeholder="Input nama thread" required><br>
            <label><b>Deskripsi</b></label><br>
            <textarea name="threadDesc" id="threadDesc" placeholder="Input deskripsi thread" required></textarea><br>
            <button type="submit" id="submit" @click="postThread">Submit</button>
            <button type="reset" id="cancel" @click="closeModal">Cancel</button>
        </form>
    </div> -->
</template>

<script setup>
const postList = ref([]);

  async function getPost() {
    // Ambil nomor threadNo sekarang dari URL param
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var threadNo = urlParams.get('threadNo')
    var query = 'http://localhost:8181/post/' + threadNo;
    const response = await fetch(query, {
        method: "GET",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        }
    });
    if (response.ok) {
      const data = await response.json()
      if (data.status == '200'){
        for (const key in data.data) {
            postList.value.push(data.data[key]);
        }
        console.log(postList);
      } else {
        console.error("Failed!", data.message);
      }
    }
  }
  
  onMounted(getPost);

  // Post thread
  async function postThread() {
    // Ambil nomor topic sekarang dari URL param
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var topicNo = urlParams.get('topicNo')
    // Ambil data dari form
    var threadName = document.getElementById("threadName").value;
    var threadDesc = document.getElementById("threadDesc").value;

    const response = fetch('http://localhost:8181/thread', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },    
        body: new URLSearchParams({
            'topicNo': topicNo,
            'threadTitle': threadName,
            'threadDesc': threadDesc
        })
    })
    if (response.ok) {
        const data = response.json()
      if (data.status == '200'){
        alert("Thread berhasil ditambahkan!")
      } else {
        console.error("Failed!", data.message);
        alert("Thread mengajukan request topic!")
      }
    }
  }

  // Kode untuk modal
  function showModal(){
        var modal = document.getElementById("modal");
        modal.style.display = "block";
    }
    
    // Close modal
    function closeModal() {
        var modal = document.getElementById("modal");
        modal.style.display = "none"
    }

    // Close modal kalau klik diluar modal
    window.onclick = function(event) {
        var modal = document.getElementById("modal");
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
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
.PostDesc{
    margin-left: 10px;
}
.PostImage{
    margin-left: 10px;
}
.list-group .list-group-item .list-group-item {
    background-color: #959090; /* Adjust color if needed */
    margin-top: 5px; /* Add margin to create indentation */
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
.modal-thread {
    position: fixed;
    z-index: 999;
    border: 2px solid black;
    padding: 2.5%;
    width: 35%;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    background-color: #454545;
    color: white;
    display: none;
    font-family: Arial, Helvetica, sans-serif;
    transform: translate(-50%, -50%);
    overflow: auto;
    top: 50%;
    left: 50%; 
}
.modal-thread input, textarea {
    background-color: #696969;
    color: white;
    width: 100%;
    padding: 10px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #3f3f3f;
    box-sizing: border-box;
}
.modal-thread button {
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