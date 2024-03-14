<!DOCTYPE php>
<html>
	<head>
		<title>Register</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="css/registerpage.css">
    </head>
    <body>
        <?php 
            session_start();
            if(isset($_SESSION['statusRegister'])){
                if($_SESSION['statusRegister'] == 0){
                    echo "<script>alert('Password dengan konfirmasi password tidak sama!')</script>";
                }
                unset($_SESSION['statusRegister']);
            }
        ?>
    <main>
        <form name="register" action="function/register.php" method="post">
            <div class="register">
                <h2>Register</h2>
                <label for="username">Username</label><br>
                <input type="text" name="username" id="username" required placeholder="Input Username"><br>
                <label for="username">Password</label><br>
                <input type="password" name="password" id="password" required placeholder="Input Password"><br>
                <label for="username">Confirm Password</label><br>
                <input type="password" name="confirm" id="confirm" required placeholder="Confirm Password" onkeyup="check()"><br>
                <label for="email">Email</label><br>
                <input type="email" name="email" id="email" required placeholder="example123@gmail.com"><br>
                <button id="submit" type="submit">Register</button><br>
            </div>
        </form>
        </form>
    </main>
    </body>
</html>