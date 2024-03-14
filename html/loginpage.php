<!DOCTYPE php>
<html>
	<head>
		<title>Index</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="css/login.css">
        <style></style>
    </head>
    <body>
    <main>
        <?php session_start(); ?>
        <?php 
            if(isset($_SESSION['statusLogin'])){
                if($_SESSION['statusLogin'] == 'gagal'){
                    echo "<script>alert('Username/Password Anda salah!')</script>";
                }
                if($_SESSION['statusLogin'] == 'banned') {
                    echo "<script>alert('Anda telah di ban dari forum kami!')</script>";
                }
                unset($_SESSION['statusLogin']);
            }
            if(isset($_SESSION['statusRegister'])){
                if($_SESSION['statusRegister'] == 1){
                    echo "<script>alert('Registrasi berhasil!')</script>";
                }
                else {
                    echo "<script>alert('Registrasi gagal!')</script>";
                }
                unset($_SESSION['statusRegister']);
            }
            if(isset($_COOKIE['id_user']) && isset($_COOKIE['username'])){
                $_SESSION['id_user'] = $_COOKIE['id_user'];
                $_SESSION['username'] = $_COOKIE['username'];
                $_SESSION['logged_in'] = true;
                if(isset($_COOKIE['admin'])){
                    $_SESSION['admin'] = true;
                }
                header("Location: homepage.php");
            }
        ?>
        <form name="login" action="function/login.php" method="post">
            <div class="login">
            <h2>Login</h2>
                <label for="username">Username</label><br>
                <input type="text" name="username" id="username" required placeholder="Input Username"><br>
                <label for="password">Password</label><br>
                <input type="password" name="password" id="password" required placeholder="Input Password"><br>
                <input id="remember" type="checkbox" name="remember" value="true">Remember Me<br>
                <button type="submit">Login</button>
                <button id='register'type="reset" onclick="window.open('registerpage.php','_self');">Register</button>
            </div>
        </form>
    </main>
    </body>
</html>