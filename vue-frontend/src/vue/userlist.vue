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
        <h1>Daftar User</h1>

        <div class="container">
            <div class="row">
                <div class="col-12 mb-3 mb-lg-5">
                    <div class="overflow-hidden card table-nowrap table-card">
                        <div class="card-header d-flex justify-content-between align-items-center">
                            <h5 class="mb-0">User List</h5>
                        </div>
                        <div class="table-responsive">
                            <table class="table mb-0">
                                <thead class="small text-uppercase bg-body text-muted">
                                    <tr>
                                        <th>Name</th>
                                        <th>Email</th>
                                        <th>Created Date</th>
                                        <th class>View Activity</th>
                                        <th class="text-end">Ban Status</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr class="align-middle" v-for="item in temp">
                                        <td>
                                            <div class="d-flex align-items-center">
                                                <img v-bind:src="item.profilePicture"
                                                    class="avatar sm rounded-pill me-3 flex-shrink-0" alt="Customer">
                                                <div>
                                                    <div class="h6 mb-0 lh-1">{{ item.username }}</div>
                                                </div>
                                            </div>
                                        </td>
                                        <td>{{ item.email }}</td>
                                        <td>{{ item.joinDate }}</td>
                                        <td @click="viewUserActivity(item.userId)">View Activity</td>
                                        <td @click="banUser(item.banstatus, item.userId)" v-bind:class="userStatusClass"
                                            style="background-color: grey !important; text-align: center; color: red">
                                            {{ item.banstatus ? 'Unban User' : 'Ban user' }}
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>

<script setup>
    import { logUserActivity } from '../activityLogger'; // Import user activity logger
    const temp = ref([]);
    // Retrieve and parse user data from session storage
    const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));

    async function getUsers() {
        // Ambil nomor topic sekarang dari URL param
        var query = 'http://localhost:8181/user';
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
                    data.data[key].joinDate = dateTimeFormatter(data.data[key].joinDate)
                    temp.value.push(data.data[key]);
                    console.log(data.data[key])
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

    async function viewUserActivity(userId, username) {
        // Buat link ke halaman useractivity
        var url = '/useractivity.html?userId=' + userId

        // Buka window baru yang menampilkan log activity user
        newwindow = window.open(url, 'User Activity', 'height=450,width=800');
        if (window.focus) {
            newwindow.focus()
        }
        return false;

    }

    async function banUser(status, id) {
        if (status == 1) {
            status = 0
            console.log(status)
        } else {
            status = 1
            console.log(status)
        }
        const unbanresponse = await fetch('http://localhost:8181/user/ban/' + id, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: new URLSearchParams({
                'banStatus': status,
            })
        })
            .then(unbanresponse => {
                if (unbanresponse.ok) {
                    return unbanresponse.json(); // Only proceed if the fetch is successful
                } else {
                    console.error("Failed to fetch unban response:", unbanresponse.status);
                    throw new Error("Failed to unban user (fetch issue)"); // Or handle the error differently
                }
            })
            .then(data => {
                console.log(data.status)
                if (data.status == '200') {
                    // Log change user ban status activity as "Change user ban status"
                    logUserActivity("Change user ban status", userDataParsed.userId);
                    alert("User is Banned/unbanned!");
                    window.location.reload();
                } else {
                    console.error("Failed to unban user:", data.message);
                    alert("Failed to Ban/unban the user!");
                }
            })
            .catch(error => {
                console.error("Error banning/unbanning user:", error);
                alert("An error occurred while unbanning the user."); // Handle errors more gracefully
            });
    }

    onMounted(getUsers);

</script>

<style scoped>
.card {
    box-shadow: 0 20px 27px 0 rgb(0 0 0 / 5%);
}

.avatar.sm {
    width: 2.25rem;
    height: 2.25rem;
    font-size: .818125rem;
}

.list-group {
    width: 100%;
    border-radius: 0;
}

.card-header {
    background-color: #ADA7A7;
}

body {
    color: rgb(31, 25, 25);
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

h1 {
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

td,
th {
    border-bottom: 1px inset grey;
    height: 30px;
    cursor: pointer;
}

h1 {
    color: white;
    padding-left: 5%;
}

::placeholder {
    color: #cccccc;
}
</style>