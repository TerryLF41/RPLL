<!DOCTYPE PHP>
<?php
    session_start();
    include "connectDatabase.php";
    $id_thread = $_GET['id'];
    $idUser = $_SESSION['id_user'];
    $query = "SELECT * FROM post WHERE id_thread = '$id_thread'";
    $result = mysqli_query($con,$query);
    if(mysqli_num_rows($result) == 0){
        echo "<h1>Thread ini masih kosong</h1>";
    }
    echo "<table>";
    while($row = mysqli_fetch_assoc($result)){
        $poster = $row['nama_poster'];
        $post = nl2br($row['teks']);
        $img = $row['image'];
        $date = $row['tanggal'];
        $no = $row['id_post'];

        $queryimg = "SELECT * FROM user_list WHERE username = '$poster' ";
        $resultimg = mysqli_query($con,$queryimg);
        $rowimg = mysqli_fetch_assoc($resultimg);
        $img_profile = $rowimg['profile_picture'];
        if($img_profile == ""){
            $img_profile = "upload/profile/default.png";
        }

        echo "<tr><th class='user'>$poster</th><th class='date'>$date   No.$no</th></tr>";
        echo "<tr><td class='pict'><a href='$img_profile' target='_blank'><img height='100' width='100' src='$img_profile'></a></td><td class='post'>$post";
        if($img == null){
            echo "&nbsp</td>";
        }
        else {
            echo"<br><br><a href='$img' target='_blank'><img height='200' width='200' src='$img'></a></td>";
        }
    }
    echo "</table>";
    mysqli_close($con);
?>