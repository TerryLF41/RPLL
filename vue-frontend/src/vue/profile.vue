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

    async function getProfile() {
        // const response = await fetch('http://localhost:8181/topic', {
        //     method: "GET",
        //     headers: {
        //         "Content-Type": "application/x-www-form-urlencoded"
        //     }
        // });
        // if (response.ok) {
        //     const data = await response.json()
        //     if (data.status == '200'){
        //         for (const key in data.data) {
        //             temp.value.push(data.data[key]);
        //         }
        //         console.log(data.data)
        //     } else {
        //         console.error("Failed!", data.message);
        //     }
        // }
        console.log(userDataParsed)
    }

    onMounted(getProfile);

    // async function postTopic() {
    //     // Ambil data dari form
    //     var topicName = document.getElementById("topicName").value;
    //     var topicDesc = document.getElementById("topicDesc").value;

    //     // Handle gambar yang diupload
    //     var fileInput = document.getElementById('topicPicture');

    //     // Buat path url untuk image yang akan diupload ke database
    //     var urlTopicPicture = "../src/assets/userUploadedFiles/topicPicture/"

    //     // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
    //     if(fileInput.files.length == 0){
    //         urlTopicPicture += "default.png"
    //     } else {
    //         // Ambil nama file yang diupload
    //         var selectedFile = fileInput.files[0].name;
    //         urlTopicPicture += selectedFile
    //     }
    //     const response = await fetch('http://localhost:8181/topic', {
    //         method: 'POST',
    //         headers: {
    //             "Content-Type": "application/x-www-form-urlencoded"
    //         },    
    //         body: new URLSearchParams({
    //             'topicTitle': topicName,
    //             'topicDesc': topicDesc,
    //             'topicPicture': urlTopicPicture
    //         })
    //     })
    //     if (response.ok) {
    //         const data = await response.json()
    //         if (data.status == '200'){
    //             alert("Topic berhasil ditambahkan!")
    //         } else {
    //             console.error("Failed!", data.message);
    //             alert("Gagal mengajukan request topic!")
    //         }
    //     }
    // }

</script>

<style scoped>
h1{
    color: white;
    padding-left: 5%;
    text-align: center;
}

</style>