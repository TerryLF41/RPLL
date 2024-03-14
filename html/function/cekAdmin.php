<?php 
    if(isset($_SESSION['admin']) == false){
        header ("Location: loginPage.php");
    }
?>