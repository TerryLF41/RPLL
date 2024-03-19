<!DOCTYPE php>
<?php
    session_start(); 
    if(isset($_POST['username']) && isset($_POST['password'])){
        include "connectDatabase.php";
        $uname = $_POST['username'];
        $password = $_POST['password'];
        $success = false;
        $query = "SELECT * FROM user_list";
        $result = mysqli_query($con,$query);

        while($row = mysqli_fetch_assoc($result)) {
            if($uname == $row['username'] && password_verify($password, $row['password'])){
                if($row['status_ban'] == true){
                    $_SESSION['statusLogin'] = "banned";
                    header("Location: ../loginpage.php");
                    exit();
                }
                $success = true;
                $_SESSION['logged_in'] = true;
                $_SESSION['id_user'] = $row['id_user'];
                $_SESSION['username'] = $row['username'];
                if(isset($_POST['remember'])){
                    if($_POST['remember'] == "true"){
                        setcookie("id_user", $row['id_user'], time() + 86400, "/");
                        setcookie("username", $uname, time() + 86400, "/");
                        if($row['status_admin'] == true){
                            setcookie("admin", true, time() + 86400, "/");
                        }
                    }
                }
                if($row['status_admin'] == true){
                    $_SESSION['admin'] = true;
                }
                header("Location: ../homepage.php");
            }
        }
        if($success == false){
            $_SESSION['statusLogin'] = "gagal";
            header("Location: ../loginpage.php");
        }
        mysqli_close($con);
    }
?>