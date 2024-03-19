<!DOCTYPE html>
<html>
	<head>
		<title>Moderate User</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">	
		<link rel="stylesheet" href="css/manageUser.css">
		<script type="text/javascript" src="javascript/userajax.js"></script>
		<style>
		</style>
	</head>
	<body onload="search()">
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "function/cekAdmin.php";
			include "header.php";
			if(isset($_SESSION['ban'])){
				if($_SESSION['ban']){
					echo "<script>alert('User telah diban')</script>";
				}
				else {
					echo "<script>alert('User telah diunban')</script>";
				}
				unset($_SESSION['ban']);
			}
			if(isset($_SESSION['set_admin'])){
				if($_SESSION['set_admin']){
					echo "<script>alert('User telah diberikan status admin')</script>";
				}
				else {
					echo "<script>alert('Status admin user telah dicabut')</script>";
				}
				
				unset($_SESSION['set_admin']);
			}
			if(isset($_SESSION['createXML'])){
				echo "<script>alert('File XML telah berhasil dibuat dan disimpan di folder XML')</script>";
				unset($_SESSION['createXML']);
			}
		?>
		<main>
			<h1>Manage User</h1>
			<input id="username" list="names" name="names" onkeyup="suggestion(this.value)" value="" placeholder="Insert Username">
				<datalist id="names">
					
				</datalist>
			<input id="ban" type="checkbox" name="ban">Banned User
			<input id="admin" type="checkbox" name="admin">Admin
			<button name="search" onclick="search()">Search</button>
			<button name="reset" onclick="reset()">Reset</button>
			<hr>
			<table id="tabel-user" style="border-collapse:collapse;width:100%">
			</table>
			<a id='xml' href="function/convertxml.php?tipe=user">Convert to XML</a><br><br>
		</main>
		<?php include "footer.php"; ?>
	</body>
</html>