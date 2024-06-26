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
        <div class="container my-5">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <h1 class="display-4">Topic List</h1>
                <button type="button" @click="showModal" class="btn btn-primary">Add New Topic</button>
            </div>
            <div class="d-flex align-items-left mb-4 ml-10"> 
                <button type="button" @click="viewPopularTopic" class="btn btn-primary" id="popularButton">Sort by Popular Topic</button>&nbsp;
                <button type="button" @click="viewLatestTopic" class="btn btn-primary" id="latestButton">Sort by Latest Topic</button><br>
            </div>
            <div class="row">
                <div class="col-md-6 col-lg-4" v-for="item in temp">
                    <div class="card mb-4 text-light bg-dark">
                        <img :src="item.topicPicture" class="card-img-top img-thumbnail img-fluid" alt="Topic Image">
                        <div class="card-body">
                            <h5 class="card-title">{{ item.topicTitle }}</h5>
                            <p class="card-text">{{ item.topicDesc }}</p>
                            <p class="card-text">{{ item.threadCount + ' thread(s)'}} </p>
                            <p class="card-text text-light"><small class="">{{ item.createDate }}</small></p>
                            <div class="d-flex justify-content-between align-items-center">
                                <button @click="goToThread(item.topicNo)" class="btn btn-primary">View Thread</button>
                                <div v-if="userType == 1">
                                    <button v-if="!item.banstatus" @click="banTopic(item.topicNo)"
                                        class="btn btn-danger ml-2">Ban Topic</button>
                                    <button v-else @click="unbanTopic(item.topicNo)" class="btn btn-info ml-2">Unban
                                        Topic</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
    <div class="modal-topic" id="modal">
        <form class="form" method="POST" id="postTopic" onsubmit="event.preventDefault();">
            <h2 class="title">Add New Topic</h2>
            <label><b>Nama Topik</b></label><br>
            <input type="text" name="topicName" id="topicName" placeholder="Input nama topik" required><br>
            <label><b>Deskripsi</b></label><br>
            <textarea name="topicDesc" id="topicDesc" placeholder="Input deskripsi topik" required></textarea><br>
            <label><b>Foto Topik</b></label><br>
            <input type="file" id="topicPicture" name="topicPicture" accept=".jpg, .jpeg, .png">
            <button type="submit" id="submit" @click="postTopic" class="btn btn-primary">Submit</button>
            <button type="reset" id="cancel" @click="closeModal" class="btn btn-secondary">Cancel</button>
        </form>
    </div>
</template>
<script setup>
    import { logUserActivity } from '../activityLogger'; // Import user activity logger
    const temp = ref([]);
    // Retrieve and parse user data from session storage
    const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
    const profilePicture = userDataParsed.profilePicture;
    // Ambil usertype dari session
    var userType = userDataParsed.userType;

    // Catat tipe order topic yang digunakan
    const orderByPopularity = true

    function viewLatestTopic(){
        window.open('topic.html?orderBy=latest', '_self')
    }

    function viewPopularTopic(){
        window.open('topic.html?orderBy=popularity', '_self')
    }

    async function getTopic() {
        // Ambil nomor topic sekarang dari URL param
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        var orderBy = urlParams.get('orderBy')
        const response = await fetch('http://localhost:8181/topic/' + orderBy, {
            method: "GET",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            }
        });
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200'){
                for (const key in data.data) {
                    // Format datetime agar lebih mudah dibaca
                    data.data[key].createDate = dateTimeFormatter(data.data[key].createDate)
                    temp.value.push(data.data[key]);
                }
                console.log(data.data)
            } else {
                console.error("Failed!", data.message);
            }
        }
    }

    // Format tanggal agar lebih mudah dibaca
    function dateTimeFormatter(timestamp){
        var date = timestamp.substring(0, 10)
        var hour = timestamp.substring(11, 19)
        return date + " " + hour
    }

    onMounted(getTopic());

    function goToThread(threadNo) {
        var url = 'thread.html?topicNo='+threadNo;
        window.open(url,'_self');
    }

    // Post topic
    async function postTopic() {
        // Ambil data dari form
        var topicName = document.getElementById("topicName").value;
        var topicDesc = document.getElementById("topicDesc").value;

        // Handle gambar yang diupload
        var fileInput = document.getElementById('topicPicture');

        // Buat path url untuk image yang akan diupload ke database
        var urlTopicPicture = "../src/assets/userUploadedFiles/topicPicture/"

        // Nama file gambar. Default adalah default.png
        var filename = "default.png"

        // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
        if(fileInput.files.length == 0){
            urlTopicPicture += filename
        } else {
            // Ambil ekstensi file dari nama file
            var selectedFile = fileInput.files[0].name;
            var selectedFileExtension = selectedFile.split('.').pop();

            // Buat nama file berdasarkan unix time
            filename = Date.now() + "." + selectedFileExtension 
            urlTopicPicture += filename
        }
        const response = await fetch('http://localhost:8181/topic', {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },    
            body: new URLSearchParams({
                'topicTitle': topicName,
                'topicDesc': topicDesc,
                'topicPicture': urlTopicPicture
            })
        })
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200'){
                // Log create topic activity as "Create topic"
                logUserActivity("Create topic",userDataParsed.userId);
                // Jika terdapat file yang diupload, handle gambar dan simpan gambar secara lokal
                if(fileInput.files.length > 0){
                    // Panggil fungsi untuk save gambar di Go
                    saveTopicImage(filename)
                }
                else {
                    alert("Topic berhasil ditambahkan!")
                    // Reload page
                    location.reload();
                }
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengajukan request topic!")
            }
        }
    }

    async function saveTopicImage(filename){
        // Siapkan form data yang akan menampung data gambar dan masukkan data form
        var formData = document.querySelector("#postTopic")
        var pictureFormData  = new FormData(formData);

        // Input data gambar kedalam formdata
        pictureFormData.append("filename", filename);

        const responsePicture = await fetch('http://localhost:8181/topic/picture', {
            method: 'POST',
            // Header akan diset otomatis oleh browser sebagai multipart/form-data
            // Body yang memuat data gambar
            body: pictureFormData
        })
        if (responsePicture.ok) {
            const data = await responsePicture.json()
            if (data.status == '200'){
                // Log create thread activity as "Uploaded topic picture"
                logUserActivity("Upload topic picture",userDataParsed.userId);
                alert("Topic berhasil ditambahkan!")
                location.reload();
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengajukan request topic!")
                location.reload();
            }
        }
    }

    // Kode ban topic
    // Admin only
    async function banTopic(topicNo){
        // Tampilkan prompt konfirmasi ban topic atau tidak
        if (confirm('Apakah Anda yakin ingin memban topic?')) {
            // Jika ya, ban topic
            var url = "http://localhost:8181/topic/ban/" + topicNo;
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
                if (data.status == '200'){
                    // Log ban topic activity as "Ban topic"
                    logUserActivity("Ban topic",userDataParsed.userId);
                    alert("Topic berhasil diban!")
                    location.reload();
                } else {
                    console.error("Failed!", data.message);
                    alert("Gagal memban topic!")
                }
            }
        }
    }

    async function unbanTopic(topicNo){
        // Tampilkan prompt konfirmasi ban topic atau tidak
        if (confirm('Apakah Anda yakin ingin memunban topic?')) {
            // Jika ya, unban topic
            var url = "http://localhost:8181/topic/ban/" + topicNo;
            const response = await fetch(url, {
                method: 'PUT',
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },    
                body: new URLSearchParams({
                    'banStatus': 0,
                })    
            })
            if (response.ok) {
                const data = await response.json()
                if (data.status == '200'){
                    // Log ban topic activity as "Unban topic"
                    logUserActivity("Unban topic",userDataParsed.userId);
                    alert("Topic berhasil diunban!")
                    location.reload();
                } else {
                    console.error("Failed!", data.message);
                    alert("Gagal memunban topic!")
                }
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
.card-img-top {
    height: 200px;
    /* Ubah nilai sesuai kebutuhan */
    object-fit: cover;
}

.card-img-top {
    position: relative;
    overflow: hidden;
}

.card-img-top::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    opacity: 0;
    transition: opacity 0.3s ease;
}

.card:hover .card-img-top::before {
    opacity: 1;
}

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

.list-group-item:hover {
    opacity: 0.8;
}

.about span {
    font-size: 12px;
    margin-right: 10px;
}

.topicDesc {
    margin-left: 10px;
}

img {
    transition: transform .2s;
}

.card-img-top:hover {
    transform: scale(1.25);
    z-index: 999;
    height: 20%;
    object-fit:cover;
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

a {
    color: white;
    text-decoration: none;
}

h1 {
    color: white;
    padding-left: 5%;
}

.add {
    text-align: center;
    cursor: pointer;
}

::placeholder {
    color: #cccccc;
}

.modal-topic {
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

.modal-topic input,
textarea {
    background-color: #696969;
    color: white;
    width: 100%;
    padding: 10px 10px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #3f3f3f;
    box-sizing: border-box;
}

.modal-topic button {
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

#banButton {
    background-color: #ff0000;
    font-size: small;
}

#unbanButton {
    background-color: #0e9fde;
    font-size: small;
}

#popularButton {
    background-color: #0f9020;
    border: none;
}
#latestButton {
    background-color: #0e31de;
    border: none;
}
</style>