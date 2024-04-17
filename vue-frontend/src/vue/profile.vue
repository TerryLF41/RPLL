<script>
import Header from '../components/header.vue'
import { ref } from 'vue';
import { onMounted } from 'vue';
const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));
</script>

<template>
    <main>
        <nav class="navbar">
            <Header />
        </nav>
        <h1>User Profile</h1>  
        <div class="container rounded bg-black mt-5 mb-5" >
    <div class="row">
        <div class="col-md-3 border-right">
            <div class="d-flex flex-column align-items-center text-center p-3 py-5"><img class="rounded-circle mt-5" width="150px" v-bind:src=userDataParsed.profilePicture><span class="font-weight-bold">{{ userDataParsed.username }}
            </span><span class="text-white-50">{{userDataParsed.email}}
            <input class="masukan" type="file" id="avatar" name="avatar" accept="image/png, image/jpeg">
            </span><span> </span></div>
        </div>
        <div class="col-md-5 border-right">
            <div class="p-3 py-5">
                <div class="d-flex justify-content-between align-items-center mb-3">
                    <h4 class="text-right">Profile Settings</h4>
                </div>
                <div class="row mt-2">
                    <div class="col-md-12"><label class="labels">Name</label><input type="text" id="userName" class="form-control" v-bind:value=userDataParsed.username></div>
                </div>
                <div class="row mt-3">
                    <div class="col-md-12"><label class="labels">Email</label><input type="text" id="email" class="form-control" v-bind:value=userDataParsed.email></div>
                    <div class="col-md-12"><label class="labels">Profile Desc</label><textarea id="description" class="form-control height:fit-content; form-control-sm" v-bind:value=userDataParsed.profileDesc></textarea></div>
                </div>
                <div class="mt-5 text-center"><button class="btn btn-primary profile-button" type="button" @click="changeProfile">Save Profile</button></div>
            </div>
        </div>
        <div class="col-md-4">
            <div class="p-3 py-5">
                <div class="d-flex justify-content-between align-items-center experience"><span>Change Password</span><span class="border px-3 p-1 add-experience"><i class="fa fa-plus"></i>&nbsp;change!</span></div><br>
                <div class="col-md-12"><label class="labels">old password</label><input type="password" id="oldPass" class="form-control" placeholder="old pass" value=""></div> <br>
                <div class="col-md-12"><label class="labels">new password</label><input type="password" id="newPass" class="form-control" placeholder="new pass" value=""></div>
            </div>
        </div>
    </div>
</div>


    </main>
    
</template>

<script setup>
    const temp = ref([]);
    // Retrieve and parse user data from session storage
    
    // Ambil usertype dari session

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

    async function changeProfile() {
        // Ambil data dari form
        var userName = document.getElementById("userName").value;
        var email = document.getElementById("email").value;
        var description = document.getElementById("description").value;

        // Handle gambar yang diupload
        var propifePicture = document.getElementById('avatar');

        // Buat path url untuk image yang akan diupload ke database
        var urlProfilePicture = "../src/assets/userUploadedFiles/topicPicture/"

        // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
        if(fileInput.files.length == 0){
            urlProfilePicture += userDataParsed.profileProfile
        } else {
            // Ambil nama file yang diupload
            var selectedFile = fileInput.files[0].name;
            urlProfilePicture += selectedFile
        }
        const response = await fetch('http://localhost:8181/user/', {
            method: 'SET',
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

</script>

<style scoped>
h1{
    color: white;
    padding-left: 5%;
    text-align: center;
}
body {
    color: white;
    background: rgb(99, 39, 120)
}
.col-md-12{
}
.form-control:focus {
    box-shadow: none;
    border-color: #BA68C8
}
.masukan{
    font-size: .8rem;
    padding-left: 4rem;
}
.form-control-sm {
    height: calc(3.5em + 2.5rem + 2px);
    
    padding: .25rem .5rem;
    font-size: .875rem;
    line-height: 1.5;
    border-radius: .2rem;
}

.profile-button {
    background: rgb(99, 39, 120);
    box-shadow: none;
    border: none
}

.profile-button:hover {
    background: #682773
}

.profile-button:focus {
    background: #682773;
    box-shadow: none
}

.profile-button:active {
    background: #682773;
    box-shadow: none
}

.back:hover {
    color: #682773;
    cursor: pointer
}

.labels {
    font-size: 11px
}
.container {
    color: white;
}
.add-experience:hover {
    background: #BA68C8;
    color: #fff;
    cursor: pointer;
    border: solid 1px #BA68C8
}

</style>