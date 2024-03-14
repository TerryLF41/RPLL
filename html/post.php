<!DOCTYPE php>
<html>
	<head>
		<title>Post</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
        <link rel="stylesheet" href="css/post.css">
        <script src="javascript/postajax.js"></script>
        <script>
            
        </script>
    </head>
    <body onload="post()">
        <?php
            session_start();
            include "function/cekcookies.php";
            include "function/cek_session.php";
            include "header.php"; 
        ?>
        <?php
            if(isset($_SESSION['posted'])){
                if($_SESSION['posted'] == 'sukses'){
                    echo "<script>alert('Post Anda telah berhasil dipost')</script>";
                }
                else if($_SESSION['posted'] == 'panjang'){
                    echo "<script>alert('Post Anda melebihi 1000 karakter!')</script>";
                }
                else{
                    echo "<script>alert('Anda hanya bisa mengirimkan file .png, .jpg/jpeg, .webm, dan .gif')</script>";
                }
                unset($_SESSION['posted']);
            }
            include ("connectDatabase.php");
            $id_thread = $_GET['id'];
        ?>
        <div id="id"><?php echo $id_thread;?></div>
        <div class="display" id="post">
        </div>
		<div class ='add'>
            <?php 
                $id_topik = $_SESSION['id_topik'];
                $nama = $_SESSION['nama_topik'];
                echo "<a id='back' href='thread.php?id=$id_topik&amp;topik=$nama'>Back to Thread</a><br>";
            ?>
            <h2>Reply to This Thread</h2>
            <form method="post" action="function/postFungsi.php" enctype="multipart/form-data">
                Text:<br>
                <textarea id = "text" name="text" value="" placeholder="Maksimal 1000 karakter"></textarea><br>	
                Image:<br> 
                <input type="file" name="image"><br>
                <input type="hidden" name="id_thread" value=<?php echo $id_thread ?>>
                <input type="submit" value="Post">
            </form>
		</div>
        <?php include "footer.php"; ?>
    </body>
</html>