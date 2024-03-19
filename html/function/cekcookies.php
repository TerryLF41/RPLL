<?php 
    if(isset($_COOKIE['id_user']) && isset($_COOKIE['username'])){
        $_SESSION['id_user'] = $_COOKIE['id_user'];
        $_SESSION['username'] = $_COOKIE['username'];
        if(isset($_COOKIE['admin'])){
            $_SESSION['admin'] = true;
        }
    }
?>