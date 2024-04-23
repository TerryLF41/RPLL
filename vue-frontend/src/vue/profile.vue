<script>
import Header from '../components/header.vue'
import { ref } from 'vue';
import { onMounted } from 'vue';
import { setCookie, getCookie, deleteCookie } from '../utils'; // Import cookie functions
import { logUserActivity } from '../activityLogger'; // Import user activity logger
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
                    <form class="form" method="POST" id="postProfilePicture" onsubmit="event.preventDefault();">
                        <input type="file" class="masukan" id="avatar" name="avatar" accept=".jpg, .jpeg, .png">
                    </form>
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
                            <div class="col-md-12"><label class="labels">Profile Description</label><textarea id="description" class="form-control height:fit-content; form-control-sm" v-bind:value=userDataParsed.profileDesc></textarea></div>
                        </div>
                        <div class="mt-5 text-center"><button class="btn btn-primary profile-button" type="button" @click="changeProfile">Save Profile</button></div>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="p-3 py-5">
                        <div class="d-flex justify-content-between align-items-center experience" @click="changePassword"><span>Change Password</span><span class="border px-3 p-1 add-experience"><i class="fa fa-plus"></i>&nbsp;Change!</span></div><br>
                        <div class="col-md-12"><label class="labels">Old password</label><input type="password" id="oldPass" class="form-control" placeholder="Input old password" value=""></div> <br>
                        <div class="col-md-12"><label class="labels">New password</label><input type="password" id="newPass" class="form-control" placeholder="Input new password" value=""></div>
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
        console.log(userDataParsed)
    }

    onMounted(getProfile);

    async function changeProfile() {
        // Ambil data dari form
        var userName = document.getElementById("userName").value;
        var email = document.getElementById("email").value;
        var description = document.getElementById("description").value;

        // Handle gambar yang diupload
        var profilePicture = document.getElementById('avatar');
        var urlProfilePicture = "/src/assets/userUploadedFiles/userProfile/"

        // Nama file gambar. Default adalah default.png
        var filename = "default.png"

        // Cek apakah input file kosong, kalau kosong, kasih path ke foto default
        if(profilePicture.files.length == 0){
           urlProfilePicture = userDataParsed.profilePicture
        } else {
            // Ambil ekstensi file dari nama file
            var selectedFile = profilePicture.files[0].name;
            var selectedFileExtension = selectedFile.split('.').pop();

            // Buat nama file berdasarkan unix time
            filename = Date.now() + "." + selectedFileExtension 
            urlProfilePicture += filename
        }
        const response = await fetch('http://localhost:8181/profile/'+ userDataParsed.userId, {
            method: 'PUT',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },    
            body: new URLSearchParams({
                'username': userName,
                'email': email,
                'description': description,
                'profilePicture': urlProfilePicture
            })
        })
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200'){
                // Log change profile activity as "Change profile"
                logUserActivity("Change profile",userDataParsed.userId);

                // Jika terdapat file yang diupload, handle gambar dan simpan gambar secara lokal
                if(profilePicture.files.length > 0){
                    // Panggil fungsi untuk save gambar di Go
                    saveProfileImage(filename)
                }
                userDataParsed.username = userName
                userDataParsed.email = email
                userDataParsed.profileDesc = description
                userDataParsed.profilePicture = urlProfilePicture
                sessionStorage.setItem('userData', JSON.stringify(userDataParsed));
                alert("Profile berhasil diganti!")
                window.location.reload();
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengganti profil!")
            }
        }
    }

    async function saveProfileImage(filename){
        // Siapkan form data yang akan menampung data gambar dan masukkan data form
        var formData = document.querySelector("#postProfilePicture")
        var pictureFormData  = new FormData(formData);

        // Input data gambar kedalam formdata
        pictureFormData.append("filename", filename);

        const responsePicture = await fetch('http://localhost:8181/user/picture', {
            method: 'POST',
            // Header akan diset otomatis oleh browser sebagai multipart/form-data
            // Body yang memuat data gambar
            body: pictureFormData
        })
        if (responsePicture.ok) {
            const data = await responsePicture.json()
            if (data.status == '200'){
                // Log create thread activity as "Upload profile picture"
                logUserActivity("Upload profile picture",userDataParsed.userId);
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengajukan request topic!")
                location.reload();
            }
        }
    }

    async function changePassword(){
        var oldPass = document.getElementById("oldPass").value;
        var newPass = document.getElementById("newPass").value;
        const response = await fetch('http://localhost:8181/password/'+ userDataParsed.userId, {
            method: 'PUT',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },    
            body: new URLSearchParams({
                'old_password': oldPass,
                'new_password': newPass,
            })
        })
        if (response.ok) {
            const data = await response.json()
            if (data.status == '200'){
                 // Log change profile activity as "Change password"
                 logUserActivity("Change password",userDataParsed.userId);
                alert("Password berhasil diganti!")
            } else {
                console.error("Failed!", data.message);
                alert("Gagal mengganti password!")
            }
        }
    }
</script>

<style scoped>
main {
    min-height: 80vh;
    width: 66%;
    margin: auto;
    padding: 0;
}
h1{
    color: white;
    padding-left: 5%;
    text-align: center;
}
body {
    color: white;
    background: rgb(99, 39, 120)
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