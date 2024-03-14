<!DOCTYPE php>
    <?php
        session_start();
        $_SESSION = array();
        session_destroy();
        if(isset($_COOKIE['id_user']) && isset($_COOKIE['username'])){
            setcookie("id_user","", time()-7200*12,"/");
            setcookie("username", "", time()-7200*12,"/");
            if(isset($_COOKIE['admin'])){
                setcookie("admin", "", time()-7200*12,"/");
            }
        }
        header("Location: ../loginPage.php");
    ?>
</html>