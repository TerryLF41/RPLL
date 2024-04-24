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
                <h1 class="display-4">Thread List</h1>
                <button type="button" @click="showModal" class="btn btn-primary">Add New Thread</button>
            </div>
            <div class="row" onload="getThread()">
                <div class="col-md-6" v-for="item in temp">
                    <div class="card mb-4">
                        <div class="card-body text-light bg-dark">
                            <h5 class="card-title">{{ item.threadTitle }}</h5>
                            <p class="card-text">{{ item.threadDesc }}</p>
                            <p class="card-text">{{ item.postCount + ' post(s)'}} </p>
                            <p class="card-text"><small>{{ item.createDate }}</small></p>
                            <div class="d-flex justify-content-between align-items-center">
                                <button @click="goToPost(item.threadNo)" class="btn btn-primary">View Posts</button>
                                <div v-if="userType == 1">
                                    <button v-if="!item.banStatus" @click="banThread(item.threadNo)"
                                        class="btn btn-danger ml-2">Ban Thread</button>
                                    <button v-else @click="unbanThread(item.threadNo)" class="btn btn-info ml-2">Unban
                                        Thread</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
 
    <div class="modal-thread" id="modal">
    <form class="form" method="POST" id="postTopic" onsubmit="event.preventDefault();">
        <h2 class="title">Add New Thread</h2>
        <label><b>Nama Thread</b></label><br>
        <input type="text" name="threadName" id="threadName" placeholder="Input nama thread" required><br>
        <label><b>Deskripsi Thread</b></label><br>
        <textarea name="threadDesc" id="threadDesc" placeholder="Input deskripsi thread" required></textarea><br>
        <button type="submit" id="submit" @click="postThread" class="btn btn-primary">Submit</button>
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

    async function getThread() {
        // Ambil nomor topic sekarang dari URL param
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        var topicNo = urlParams.get('topicNo')
        var query = 'http://localhost:8181/thread/' + topicNo;
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
                    // Format datetime agar lebih mudah dibaca
                    data.data[key].createDate = dateTimeFormatter(data.data[key].createDate)
                    temp.value.push(data.data[key]);
                }
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

    onMounted(getThread);

    function goToPost(threadNo) {
        var url = 'post.html?threadNo=' + threadNo;
        window.open(url, '_self');
    }

    // Post thread
    async function postThread() {
        // Ambil nomor topic sekarang dari URL param
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        var topicNo = urlParams.get('topicNo')
        // Ambil data dari form
        var threadName = document.getElementById("threadName").value;
        var threadDesc = document.getElementById("threadDesc").value;

        const response = await fetch('http://localhost:8181/thread', {
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
            const data = await response.json()
            if (data.status == '200') {
                // Log create thread activity as "Create thread"
                logUserActivity("Create thread", userDataParsed.userId);
                alert("Thread berhasil ditambahkan!")
                location.reload();
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengajukan request topic!")
            }
        }
    }

    // Kode ban thread
    // Admin only
    async function banThread(threadNo) {
        // Tampilkan prompt konfirmasi ban topic atau tidak
        if (confirm('Apakah Anda yakin ingin memban thread?')) {
            // Jika ya, ban topic
            var url = "http://localhost:8181/thread/ban/" + threadNo;
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
                    // Log ban thread activity as "Ban thread"
                    logUserActivity("Ban thread", userDataParsed.userId);

                    alert("Thread berhasil diban!")
                    // Ambil nomor topic sekarang dari URL param
                    const queryString = window.location.search;
                    const urlParams = new URLSearchParams(queryString);
                    var topicNo = urlParams.get('topicNo')
                    var url = '/thread.html?topicNo=' + topicNo;
                    window.open(url, '_self');
                } else {
                    console.error("Failed!", data.message);
                    alert("Gagal memban thread!")
                }
            }
        }
    }

    async function unbanThread(threadNo) {
        // Tampilkan prompt konfirmasi ban topic atau tidak
        if (confirm('Apakah Anda yakin ingin memunban thread?')) {
            // Jika ya, ban topic
            var url = "http://localhost:8181/thread/ban/" + threadNo;
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
                if (data.status == '200') {
                    // Log ban thread activity as "Ban thread"
                    logUserActivity("Unban thread", userDataParsed.userId);

                    alert("Thread berhasil diunban!")
                    // Ambil nomor topic sekarang dari URL param
                    const queryString = window.location.search;
                    const urlParams = new URLSearchParams(queryString);
                    var topicNo = urlParams.get('topicNo')
                    var url = '/thread.html?topicNo=' + topicNo;
                    window.open(url, '_self');
                } else {
                    console.error("Failed!", data.message);
                    alert("Gagal memban thread!")
                }
            }
        }
    }

    // Kode untuk modal
    function showModal() {
        var modal = document.getElementById("modal");
        modal.style.display = "block";
    }

    // Close modal
    function closeModal() {
        var modal = document.getElementById("modal");
        modal.style.display = "none"
    }

    // Close modal kalau klik diluar modal
    window.onclick = function (event) {
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