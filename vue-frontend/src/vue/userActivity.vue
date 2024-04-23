<script>
import { ref } from 'vue';
import { onMounted } from 'vue';
</script>

<template>
    <main>
    <div class="container">
        <div class="row">
            <div class="col-12 mb-3 mb-lg-5">
                <div class="overflow-hidden card table-nowrap table-card">
                    <div class="card-header d-flex justify-content-between align-items-center">
                        <h5 class="mb-0">User Activity</h5>
                    </div>
                    <div class="table-responsive">
                        <table class="table mb-0">
                            <thead class="small text-uppercase bg-body text-muted">
                                <tr id="tableContent">
                                    <th>No</th>
                                    <th>Activity</th>
                                    <th>Activity Time</th>
                                    <th>IP Address</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr class="align-middle" id="tableContent" v-for="item in temp">
                                    <td>{{ item.logNo }}</td>
                                    <td>{{ item.logType }}</td>
                                    <td>{{ item.logTime }}</td>
                                    <td>{{ item.ipAddress }}</td>
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
    const temp = ref([]);

    async function getUserActivity() {
        // Ambil nomor user sekarang dari URL param
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        var userId = urlParams.get('userId')

        var url = "http://localhost:8181/userLog/" + userId

        const response = await fetch(url, {
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
                    data.data[key].logTime = dateTimeFormatter(data.data[key].logTime)
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

    // Load function saat page dimuat
    onMounted(getUserActivity);

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
.list-group{
	width: 100%;
    border-radius: 0;
}
.card-header{
    background-color: #ADA7A7;
}
body {
    color: rgb(31, 25, 25);
    background-image:url("../upload/media/bg-topic.jpg");
    background-repeat: no-repeat;
    background-size: cover;
}
main {
    min-height: 80vh;
    width: 90%;
    margin: auto;
    padding: 0;
}
#tableContent {
    background-color: #ADA7A7;
}
h1 {
    text-align: center;
}
.list{
    padding: 5px;
    width: 90%;
    margin:auto;
    border: 2px solid grey;
    border-radius: 10px;
}
a{
    color: white;
    text-decoration: none;
}
td,th{
    border-bottom:1px inset grey;
    height:30px;
}
h1{
    color: white;
    padding-left: 5%;
}
::placeholder {
    color: #cccccc;
}
</style>