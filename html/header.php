<?php 
    include "connectDatabase.php";
    $query = mysqli_query($con,"SELECT * FROM user_list WHERE id_user = '".$_SESSION['id_user']."'");
    while($row = mysqli_fetch_assoc($query)){
        $profileimg = $row['profile_picture'];
    }
    if($profileimg == ""){
        $profileimg = "upload/profile/default.png";
    }
?>
<nav class="navbar">
    <div class='logo'><a href='homepage.php'><img src='upload/others/chattergrey.png'></a></div>
    <ul class="nav-links">
        <div class="menu">
            <li><a href="daftarTopik.php">Catalog</a></li>
            <li><a href="about.php">About</a></li>
            <li><a href="function/logout.php">Log Out</a></li>
            <div class="profile-picture">
                <a href="editUser.php"><img class="profile-picture" src="<?php echo $profileimg ?>"></a>
            </div>
        </div> 
    </ul>
</nav>