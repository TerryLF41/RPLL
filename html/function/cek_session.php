<?php 
    if(isset($_SESSION['username']) == false && isset($_SESSION['id_user']) == false){
        header ("Location: loginPage.php");
    }
?>