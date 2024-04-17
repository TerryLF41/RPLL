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
        <h1>User Profile</h1>  
        
    </main>
    
</template>

<script setup>
    const temp = ref([]);
    // Retrieve and parse user data from session storage
    const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
    // Ambil usertype dari session
    var userType = userDataParsed.userType;

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
                for (const key in data.data) {
                    temp.value.push(data.data[key]);
                }
                console.log(data.data)
            } else {
                console.error("Failed!", data.message);
            }
        }
    }

    onMounted(getTopic);

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

        // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
        if(fileInput.files.length == 0){
            urlTopicPicture += "default.png"
        } else {
            // Ambil nama file yang diupload
            var selectedFile = fileInput.files[0].name;
            urlTopicPicture += selectedFile
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
                alert("Topic berhasil ditambahkan!")
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengajukan request topic!")
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
                    alert("Topic berhasil diban!")
                    window.open('/topic.html','_self');
                } else {
                    console.error("Failed!", data.message);
                    alert("Gagal memban topic!")
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
h1{
    color: white;
    padding-left: 5%;
    text-align: center;
}

</style>