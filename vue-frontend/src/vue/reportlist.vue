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
        <h1>Daftar Report Post</h1>

        <div class="container">
            <div class="row">
                <div class="col-12 mb-3 mb-lg-5">
                    <div class="overflow-hidden card table-nowrap table-card">
                        <div class="card-header d-flex justify-content-between align-items-center">
                            <h5 class="mb-0">Report Post List</h5>
                        </div>
                        <div class="table-responsive">
                            <table class="table mb-0">
                                <thead class="small text-uppercase bg-body text-muted">
                                    <tr>
                                        <th>User ID</th>
                                        <th>Username</th>
                                        <th>Post No</th>
                                        <th>Report Text</th>
                                        <th>Report Date</th>
                                        <th>Report Status</th>
                                        <th class="text-end">Action</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr class="align-middle" v-for="item in reportPostList">
                                        <td>{{ item.userId }}</td>
                                        <td>{{ item.username }}</td>
                                        <td>{{ item.postNo }}</td>
                                        <td>{{ item.reportText }}</td>
                                        <td>{{ item.reportDate }}</td>
                                        <td>{{ item.reportStatus ? 'Solved' : 'Open' }}</td>
                                        <td class="text-end">
                                            <div v-if="item.reportStatus == 0">
                                                <button @click="openPost(item.threadNo)"
                                                    class="btn btn-sm btn-primary">Open Post</button>
                                                <button @click="resolveReport(item.postNo)"
                                                    class="btn btn-sm btn-success">Resolve</button>
                                            </div>
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
const reportPostList = ref([]);

// Retrieve and parse user data from session storage
const userDataParsed = JSON.parse(sessionStorage.getItem('userData'));

// If user data is not present, redirect to the login page
if (!userDataParsed) {
    alert('You haven\'t logged in yet.');
    window.location.href = 'login.html'; // Redirect to the login page
}

async function getReportPosts() {
    const response = await fetch('http://localhost:8181/reportpostu', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    if (response.ok) {
        const data = await response.json();
        if (data.status == '200') {
            for (const key in data.data) {
                // Format datetime agar lebih mudah dibaca
                data.data[key].reportDate = dateTimeFormatter(data.data[key].reportDate)
                reportPostList.value.push(data.data[key]);
                console.log(data.data[key])
            }
        } else {
            console.error('Failed!', data.message);
        }
    } else {
        console.error('Failed to fetch report posts');
    }
}

// Format tanggal agar lebih mudah dibaca
function dateTimeFormatter(timestamp) {
    var date = timestamp.substring(0, 10)
    var hour = timestamp.substring(11, 19)
    return date + " " + hour
}

async function resolveReport(postNo) {
    const response = await fetch(`http://localhost:8181/reportpost/resolve/${postNo}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    if (response.ok) {
        const data = await response.json();
        if (data.status == '200') {
            // Log resolve report activity as "Resolve report"
            logUserActivity("Resolve report", userDataParsed.userId);
            alert('Report resolved successfully!');
            window.location.reload();
        } else {
            console.error('Failed to resolve report:', data.message);
            alert('Failed to resolve report');
        }
    } else {
        console.error('Failed to resolve report');
        alert('Failed to resolve report');
    }
}


// async function getUsername(userId) {
//     const response = await fetch(`http://localhost:8181/user/username/${userId}`, {
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json'
//         }
//     });
//     if (response.ok) {
//         const data = await response.json();
//         if (data.status == '200') {
//             console.log(data.data);
//             return data.data;
            
//         } else {
//             console.error('Failed to fetch username:', data.message);
//             return null;
//         }
//     } else {
//         console.error('Failed to fetch username');
//         return null;
//     }
// }

function openPost(threadNo) {
    // Redirect to the post page with postNo and userId
    window.location.href = `post.html?threadNo=${threadNo}`;
}

function authorization() {
    // Ambil usertype dari session
    const userType = userDataParsed.userType;
    console.log(userType)
    
    if (userType != 1) {
        window.location.href = 'homepage.html';
    }
    if (userType == 1) {
        getReportPosts();
    }
}
onMounted(authorization);

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
}

h1 {
    color: white;
    padding-left: 5%;
}

::placeholder {
    color: #cccccc;
}
</style>
