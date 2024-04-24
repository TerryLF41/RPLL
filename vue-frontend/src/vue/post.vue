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
      <div class="container mt-5">
        <h1 class="text-center mb-5">Posts</h1>
        <div v-for="(post, index) in userAndPost" :key="index">
          <div class="card mb-5 text-light bg-dark" v-if="post.banStatus == 0">
            <div class="row g-0">
              <div class="col-md-2 d-flex align-items-center" style="padding-left: 1em;">
                <img v-bind:src="'..' + post.user.profilePicture" class="img-fluid rounded-circle profile-picture">
              </div>
              <div class="col-md-10">
                <div class="card-body">
                  <div class="d-flex justify-content-between align-items-center mb-2">
                    <h5 class="card-title m-0">{{ post.user.username }}</h5>
                    <div v-if="userType == 1">
                      <button @click="banPost(post.postNo)" class="btn btn-outline-danger btn-sm">Ban Post</button>
                    </div>
                  </div>
                  <h6 class="card-subtitle mb-2">{{ post.postDate }}</h6>
                  <p class="card-text" v-html="post.postText"></p>
                  <img v-bind:src="post.postImage" class="img-fluid mt-3" style="max-width:75%;">
                  <div class="d-flex justify-content-end align-items-center mt-3">
                    <button @click="reportPost(post.postNo)" class="btn btn-outline-warning btn-sm">Report Post</button>
                  </div>
                </div>
              </div>
            </div>
            <div class="card-footer">
              <form class="formReply mb-1" method="POST" :id="'idFormReply' + post.postNo" @submit.prevent="replyPost(post.postNo)">
                  <textarea :name="'textReply' + post.postNo" :id="'textReply' + post.postNo" class="form-contro text-light bg-dark" placeholder="Your comment here" rows="2" style="background-color: #ffffff; color: #000000;"></textarea>
                  <input :name="'idFormReply' + post.postNo" :id="'idFormReply' + post.postNo" type="hidden" :value="'idFormReply' + post.postNo">
                  <input :name="'postImage' + post.postNo" :id="'replyImage' + post.postNo" type="file" class="form-control text-light bg-dark" accept=".jpg, .jpeg, .png"><br>
                  <button class="btn btn-primary" type="submit" :id="'reply' + post.postNo">Comment</button>
              </form>
              <div class="reply-container">
                <div class="wrapper-li" v-for="(reply, index) in userAndReply" :key="index">
                  <div class="d-flex flex-row" v-if="reply.replyTo === post.postNo" style="border-bottom: 2px solid black;">
                    <div class="col-md-2 d-flex align-items-center">
                      <img v-bind:src="'..' + reply.user.profilePicture" class="img-fluid rounded-circle profile-picture">
                    </div>
                    <div class="card-body">
                      <div class="d-flex justify-content-between align-items-center mb-2">
                        <h5 class="card-title m-0">{{ reply.user.username }}</h5>
                      </div>
                      <h6 class="card-subtitle mb-2 text-muted">{{ reply.postDate }}</h6>
                      <p class="card-text" v-html="reply.postText"></p>
                      <img v-bind:src="reply.postImage" class="img-fluid mt-3" style="max-width: 50%;">
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="card mt-5">
          <div class="card-body text-light bg-dark">
            <form class="formPost" id="idFormPost" method="POST" onsubmit="event.preventDefault();">
                <input name="idFormPost" id="idFormPost'" type="hidden" value="idFormPost">
                <div class="form-group">
                <label for="textComment">Comment:</label>
                <textarea name="textComment" id="textComment" class="form-control text-light bg-dark" rows="3" placeholder="Your comment here"></textarea>
              </div>
              <div class="form-group">
                <input name="postImage" id="textImage" type="file" class="form-control text-light bg-dark" accept=".jpg, .jpeg, .png"><br>
              </div>
              <button type="submit" id="post" class="btn btn-outline-primary" @click="newPost">Submit Post</button>
            </form>
          </div>
        </div>
      </div>
    </main>
  </template>

<script setup>
import { logUserActivity } from '../activityLogger'; // Import user activity logger
const postList = ref([]);
const replyList = ref([]);
const userList = ref([]);

// Retrieve and parse user data from session storage
const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
// Ambil usertype dari session
var userType = userDataParsed.userType;

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
        if (data.status == '200') {
            for (const key in data.data) {
                if (data.data[key].replyTo == null) {
                    // Format datetime agar lebih mudah dibaca
                    data.data[key].postDate = dateTimeFormatter(data.data[key].postDate)
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
        if (data.status == '200') {
            for (const key in data.data) {
                if (data.data[key].replyTo != null) {
                    // Format datetime agar lebih mudah dibaca
                    data.data[key].postDate = dateTimeFormatter(data.data[key].postDate)
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
        if (data.status == '200') {
            for (const key in data.data) {
                userList.value.push(data.data[key]);
            }
        } else {
            console.error("Failed!", data.message);
        }
    }
}

// Format tanggal agar lebih mudah dibaca
function dateTimeFormatter(timestamp) {
    var date = timestamp.substring(0, 10)
    var hour = timestamp.substring(11, 19)
    return date + " " + hour
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
    var button = document.getElementById("post");
    var idForm = button.form.id;

    // Ambil user Id dari session
    const userId = userDataParsed.userId;

    // Handle gambar yang diupload
    var fileInput = document.getElementById('textImage');

    // Buat path url untuk image yang akan diupload ke database
    var urlTopicPicture = "../src/assets/userUploadedFiles/postPicture/"

    // Nama file gambar. Default adalah default.png
    var filename = "default.png"

    // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
    if (fileInput.files.length == 0) {
        urlTopicPicture = ""
    } else {
        // Ambil ekstensi file dari nama file
        var selectedFile = fileInput.files[0].name;
        var selectedFileExtension = selectedFile.split('.').pop();

        // Buat nama file berdasarkan unix time
        filename = Date.now() + "." + selectedFileExtension
        urlTopicPicture += filename
    }
    const response = await fetch('http://localhost:8181/post', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: new URLSearchParams({
            'threadNo': threadNo,
            'postText': postText,
            'postImage': urlTopicPicture,
            'userId': userId,
            'replyTo': ""
        })
    })
    if (response.ok) {
        const data = await response.json()
        if (data.status == '200') {
            // Log create topic activity as "Create topic"
            logUserActivity("Create post", userDataParsed.userId);
            // Jika terdapat file yang diupload, handle gambar dan simpan gambar secara lokal
            if (fileInput.files.length > 0) {
                // Panggil fungsi untuk save gambar di Go
                savePostImage(filename, idForm, 0)
            }
            else {
                alert("Post berhasil ditambahkan!")
                // Reload page
                location.reload();
            }
        } else {
            console.error("Failed!", data.message);
            alert("Gagal mengajukan request post!")
        }
    }
}
async function replyPost(postNo) {
    // Ambil threadNo dari URL param
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var threadNo = urlParams.get('threadNo')

    // Ambil data dari form
    var textReply = document.getElementById("textReply" + postNo).value;
    var button = document.getElementById("reply" + postNo);
    var idForm = button.form.id;

    const userId = userDataParsed.userId;

    // Handle gambar yang diupload
    var fileInput = document.getElementById("replyImage" + postNo);
    // Buat path url untuk image yang akan diupload ke database
    var urlTopicPicture = "../src/assets/userUploadedFiles/postPicture/"

    // Nama file gambar. Default adalah default.png
    var filename = "default.png"

    // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
    if (fileInput.files.length == 0) {
        urlTopicPicture = ""
    } else {
        // Ambil ekstensi file dari nama file
        var selectedFile = fileInput.files[0].name;
        var selectedFileExtension = selectedFile.split('.').pop();

        // Buat nama file berdasarkan unix time
        filename = Date.now() + "." + selectedFileExtension
        urlTopicPicture += filename
    }
    const response = await fetch('http://localhost:8181/post', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: new URLSearchParams({
            'threadNo': threadNo,
            'postText': textReply,
            'postImage': urlTopicPicture,
            'userId': userId,
            'replyTo': postNo
        })
    })
    if (response.ok) {
        const data = await response.json()
        if (data.status == '200') {
            // Log create topic activity as "Create topic"
            logUserActivity("Create post", userDataParsed.userId);
            // Jika terdapat file yang diupload, handle gambar dan simpan gambar secara lokal
            if (fileInput.files.length > 0) {
                // Panggil fungsi untuk save gambar di Go
                savePostImage(filename, idForm, postNo)
            }
            else {
                alert("Post berhasil ditambahkan!")
                // Reload page
                location.reload();
            }
        } else {
            console.error("Failed!", data.message);
            alert("Gagal mengajukan request post!")
        }
    }
}

async function savePostImage(filename, idForm, postNo) {
    // Siapkan form data yang akan menampung data gambar dan masukkan data form
    var formData = document.querySelector("#" + idForm);
    var pictureFormData = new FormData(formData);

    // Input data gambar kedalam formdata
    pictureFormData.append("filename", filename);
    pictureFormData.append("postNo", postNo);

    const responsePicture = await fetch('http://localhost:8181/post/picture', {
        method: 'POST',
        // Header akan diset otomatis oleh browser sebagai multipart/form-data
        // Body yang memuat data gambar
        body: pictureFormData
    })
    if (responsePicture.ok) {
        const data = await responsePicture.json()
        if (data.status == '200') {
            // Log create thread activity as "Uploaded topic picture"
            logUserActivity("Upload topic picture", userDataParsed.userId);
            alert("Topic berhasil ditambahkan!")
            location.reload();
        } else {
            console.error("Failed!", data.message);
            alert("Gagal mengajukan request topic!")
            location.reload();
        }
    }
}

// Ban sebuah post
async function banPost(postNo) {
    // Tampilkan prompt konfirmasi ban topic atau tidak
    if (confirm('Apakah Anda yakin ingin memban post?')) {
        // Jika ya, ban topic
        var url = "http://localhost:8181/post/ban/" + postNo;
        console.log(url)
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: new URLSearchParams({
                'banStatus': 1,
            })
        })
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200') {
                // Log ban post activity as "Ban post"
                logUserActivity("Ban post", userDataParsed.userId);

                alert("Post berhasil diban!")
                // Ambil nomor topic sekarang dari URL param
                const queryString = window.location.search;
                const urlParams = new URLSearchParams(queryString);
                var topicNo = urlParams.get('threadNo')
                var url = '/post.html?threadNo=' + topicNo;
                window.open(url, '_self');
            } else {
                console.error("Failed!", data.message);
                alert("Gagal memban post!")
            }
        }
    }
}

// Report sebuah post
async function reportPost(postNo) {
    // Tampilkan prompt konfirmasi ban topic atau tidak
    if (confirm('Apakah Anda ingin melaporkan post ini?')) {
        // Jika ya, minta alasan report via prompt
        var reportText = prompt("Reason for reporting:");
        var url = "http://localhost:8181/reportpost";
        console.log(reportText)
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: new URLSearchParams({
                'userId': userDataParsed.userId,
                'postNo': postNo,
                'reportText': reportText
            })
        })
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200') {
                // Log report post activity as "Report post"
                logUserActivity("Report post", userDataParsed.userId);

                alert("Post berhasil dilaporkan!")
                // Ambil nomor topic sekarang dari URL param
                const queryString = window.location.search;
                const urlParams = new URLSearchParams(queryString);
                var topicNo = urlParams.get('threadNo')
                var url = '/post.html?threadNo=' + topicNo;
                window.open(url, '_self');
            } else {
                console.error("Failed!", data.message);
                alert("Anda sudah pernah melaporkan post ini!")
            }
        }
    }
}

</script>

<style scoped>
.list-group {
    width: 100% !important;
    border-radius: 0;
}

.list-group-item {
    margin-top: 15px;
    cursor: pointer;
    border-style: none;
    transition: all 0.3s ease-in-out;
    background-color: #ADA7A7;
}

.about span {
    font-size: 12px;
    margin-right: 10px;
}

.PostDesc {
    margin-left: 10px;
}

.PostImage {
    margin-left: 10px;
}

.PostImage img {
    margin-left: 10px;
}

.profile-picture {
    width: 150px;
    height: 150px;
    object-fit: cover;
}

img[src=""] {
    display: none;
}

.list-group .list-group-item .list-group-item {
    background-color: #959090;
    /* Adjust color if needed */
    margin-top: 5px;
    /* Add margin to create indentation */
}

body {
    color: white;
    background-image: url("../upload/media/bg-topic.jpg");
    background-repeat: no-repeat;
    background-size: cover;
}

main {
    min-height: 80vh;
    width: 66%;
    margin: auto;
    padding: 0;
}

h1,
h2,
#desc {
    text-align: center;
}

.list {
    padding: 5px;
    width: 90%;
    margin: auto;
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

a {
    color: white;
    text-decoration: none;
}

td,
th {
    border-bottom: 1px inset grey;
    height: 30px;
}



h1 {
    color: white;
    padding-left: 5%;
}

.topik {
    float: right;
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

.modal-thread input,
textarea {
    background-color: #696969;
    color: white;
    width: 100%;
    padding: 10px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid lightgray;
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

#banButton {
    background-color: #ff0000;
}

#reportButton {
    background-color: #eda02d;
}

h2.title {
    text-align: center;
    margin-top: 0;
}

#textComment {
    background-color: #ffffff; /* Change to a lighter color */
    color: #000000; /* Change to a darker text color if necessary */
}
</style>