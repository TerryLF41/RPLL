<!DOCTYPE php>
<?php
    session_start();
    include "connectDatabase.php";  
    if(isset($_POST['username']) && isset($_POST['password']) && isset($_POST['email']) && isset($_POST['confirm'])){
        $uname = $_POST['username'];
        $password = $_POST['password'];
        if($password != $_POST['confirm']){
            $_SESSION['statusRegister'] = 0;
            header("Location: ../registerpage.php");
            exit();
        }
        $password = password_hash($password, PASSWORD_DEFAULT);
        $email = $_POST['email'];
        $query = "INSERT INTO user_list(username, password, email) 
        VALUES ('$uname','$password','$email')";
        if(mysqli_query($con,$query)){
            $_SESSION['statusRegister'] = 1;
            header("Location: ../loginpage.php");
            exit();
        }
        else {
            $_SESSION['statusRegister'] = 0;
            header("Location: ../loginpage.php"); 
        }
    }
?>