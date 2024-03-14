<html>
	<head>
		<title>Profile</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<script src="javascript/user.js" type="text/javascript"></script>
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/editUser.css">
    </head>
	<body>
		<?php
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "header.php";
			if(isset($_SESSION['update_profile'])){
				//Set Alert Box
				if($_SESSION['update_profile'] == true){
					echo "<script>alert('Profile berhasil diupdate')</script>";
					
				}
				else {
					echo "<script>alert('Anda hanya bisa menggunakan file .png, .jpg/jpeg, dan .gif untuk foto profile')</script>";
				}
				unset($_SESSION['update_profile']);
			}
			include "connectDatabase.php";
			$poster = $_SESSION['id_user'];
			$query = mysqli_query($con, "SELECT * FROM user_list WHERE id_user = '$poster'");
			$data = mysqli_fetch_assoc($query);
			$nama = $data['username'];
			$email = $data['email'];
			$biodata = $data['biodata'];
			$profilePict = $data['profile_picture'];
			$pass = $data['password'];	
		?>
		<div class="konten">
			<h2>Edit Profile</h2><br>
			<hr>
			<form method="post" action="function/profileEditFungsi.php" enctype="multipart/form-data" id="form1">
			Username<br>
			<input type='text' name='nama' value='<?php echo $nama ?>'><br>
			Email<br>
			<input type='email' name='Email' value='<?php echo $email ?>'><br>
			Biodata<br>
			<textarea rows='4' cols='50' name='biodata' value ='' form='form1'><?php echo $biodata ?></textarea><br>
			Profile Picture:<br>
			<?php
				if($profilePict == null){
					echo "No Profile Picture yet<br>";
				}
				else {
					echo "<img height='200px' width='200px' src='$profilePict'><br><br>";
				}
			?>
			<input type='file' name='profile' value=''><br><br>
			<input type='submit' value='Update'>
			</form>
			<?php
				mysqli_close($con);	
			?>
		</div>
		<?php include "footer.php"; ?>
	</body>
</html>
