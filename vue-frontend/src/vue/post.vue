<script>
import Header from '../components/header.vue'
import { ref } from 'vue';
import { onMounted } from 'vue';
import { computed } from 'vue';
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
                <li class="list-group-item d-flex justify-content-between align-content-center" v-for="(post, index) in userAndPost" :key="index" style="height: calc(30.5em + 2.5rem + 2px);">
                    <div class="post d-flex flex-row" >
                        <div class="profileUser">
                            <span>{{ post.user.username }}</span><br>
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
                    <ul class="list-group mt-5 text-white" v-for="(reply, index) in userAndReply" :key="index">
                        <li class="list-group-item d-flex justify-content-between align-content-center" v-if="reply.replyTo === post.postNo">
                            <div class="d-flex flex-row">
                                <div class="profileUser">
                                    <span>{{ reply.user.username }}</span><br>
                                    <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                                </div>
                                <div class="ml-2 PostDesc">
                                    <h6 class="mb-0"> {{ reply.postDate }}</h6>
                                    <div class="about">
                                        <span>{{ reply.postText }}</span>
                                    </div>
                                </div>
                                <div class="ml-2 PostImage">
                                    <h6 class="mb-0">No. {{ reply.postNo }}</h6>
                                    <span><img src="../assets/userUploadedFiles/userProfile/default.png" width="100" /></span>
                                </div>
                            </div>
                        </li>
                    </ul>
                    <div class="reply" style="position: absolute; bottom: 0; margin-bottom: 15px">
                        <form class="formReply" method="POST">
                            <input :name="'textReply' + post.postNo" :id="'textReply' + post.postNo" type="textbox" placeholder="your comment here">
                            <input :name="'textReply' + post.postNo" :id="'replyImage' + post.postNo" type="file" accept="image/png, image/gif, image/jpeg"><br><br>
                            <button type="submit" id="reply" @click="replyPost(post.postNo)">Comment</button>
                        </form>
                    </div>
                </li>
                <li class="list-group-item d-flex justify-content-between align-content-center">
                    <div class="d-flex flex-row">
                        <form class="formPost" method="POST">
                            <input name="textComment" id="textComment" type="textbox" placeholder="your comment here">
                            <input name="textImage" id="textImage" type="file" accept="image/png, image/gif, image/jpeg"><br><br>
                            <button type="submit" id="post" @click="newPost">Comment</button>
                        </form>
                    </div>
                </li>
            </ul>
        </div>
    </main>
    
</template>

<script setup>
  const postList = ref([]);
  const replyList = ref([]);
  const userList = ref([]);

  const userAndPost = computed(() => {
    return postList.value.map(post => {
        // Find the corresponding user for this post
        const user = userList.value.find(user => user.userId === post.userId);
        return {
            ...post,
            user
        };
    });
  });

  const userAndReply = computed(() => {
    return replyList.value.map(reply => {
        // Find the corresponding user for this post
        const user = userList.value.find(user => user.userId === reply.userId);
        return {
            ...reply,
            user
        };
    });
  });

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
            if (data.data[key].replyTo == null) {
                postList.value.push(data.data[key]);
            }
        }
      } else {
        console.error("Failed!", data.message);
      }
    }
  }

  async function getReply() {
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
            if (data.data[key].replyTo != null) {
                replyList.value.push(data.data[key]);
            }
        }
      } else {
        console.error("Failed!", data.message);
      }
    }
  }
  async function getUsers() {
        const response = await fetch('http://localhost:8181/user', {
            method: "GET",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            }
        });
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200'){
                for (const key in data.data) {
                    userList.value.push(data.data[key]);
                }
            } else {
                console.error("Failed!", data.message);
            }
        }
    }
  onMounted(getPost);
  onMounted(getReply);
  onMounted(getUsers);

  // Post
  async function newPost() {
    // Ambil threadNo dari URL param
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var threadNo = urlParams.get('threadNo')

    // Ambil data dari form
    var postText = document.getElementById("textComment").value;
    // var textImage = document.getElementById("textImage").value;
    var postImage = "../assets/userUploadedFiles/userProfile/default.png"

    // Ambil user Id dari sesion
    const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
    const userId = userDataParsed.userId;

    const response = fetch('http://localhost:8181/post', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },    
        body: new URLSearchParams({
            'threadNo': threadNo,
            'postText': postText,
            'postImage': postImage,
            'userId': userId,
            'replyTo': ""
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
  async function replyPost(postNo) {
    // Ambil threadNo dari URL param
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var threadNo = urlParams.get('threadNo')

    // Ambil data dari form
    var textReply = document.getElementById("textReply"+postNo).value;
    // var textImage = document.getElementById("textImage"+postNo).value;
    var replyImage = "../assets/userUploadedFiles/userProfile/default.png"
    alert(textReply)
    // Ambil user Id dari sesion
    const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
    const userId = userDataParsed.userId;

    const response = fetch('http://localhost:8181/post', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },    
        body: new URLSearchParams({
            'threadNo': threadNo,
            'postText': textReply,
            'postImage': replyImage,
            'userId': userId,
            'replyTo': postNo
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