<!DOCTYPE php>
<html>
	<head>
		<title>Home</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">		
		<link rel="stylesheet" href="css/homepage.css">		
    </head>
    <body>
        <?php 
			session_start(); 
			include "function/cekcookies.php";
            include "function/cek_session.php";
			include "header.php";
            if(isset($_SESSION['logged_in'])){
            //Display alert box di pada saat pertama kali log in dan masuk ke home
                if($_SESSION['logged_in'] == true){
                    $uname = $_SESSION['username'];
                    echo '<script>alert("Welcome '.$uname.'")</script>';
                    $_SESSION['logged_in'] = false;
                }
            }
        ?>
		<main>
			<h1>Welcome to</h1>
			<img src='upload/others/chattergrey.png'>
			<h2>Diskusi Tanpa Batas</h2><br><br>
			<p id='desc'>Chatters Forum adalah sebuah tempat dimana Anda dapat mendiskusikan apa saja yang Anda inginkan tanpa batasan topik </p><br>
			<h2>Topik Terbaru</h2>
			<?php 
				include "function/getTopik.php";
			?>
			<?php 
				if(isset($_SESSION['admin'])){
					if($_SESSION['admin'] == true){
						include "menuadmin.php";
					}
				}
			?>      
		</main>
	<?php include "footer.php"; ?>
    </body>
</html>