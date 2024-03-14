<?php
	session_start();
	include "connectDatabase.php";
	$nama = $_POST["nama"];
	$oldname = $_SESSION['username'];
	$email = $_POST["Email"];
	$bio = $_POST["biodata"];
	$poster = $_SESSION['id_user'];
	if(file_exists($_FILES["profile"]["tmp_name"])){
		$allowed = ['png','jpg','jpeg'];
		$filetype = explode('.',$_FILES["profile"]["name"]);
		$filetype = end($filetype);
		$filetype = strtolower($filetype);
		if(!in_array($filetype,$allowed)){
			$_SESSION['update_profile'] = false;
			header("Location: ../editUser.php");
			exit();
		}
		$path = "upload/profile/";
		$path = $path.basename($_FILES["profile"]["name"]);
		$query = "UPDATE user_list SET username ='$nama', email ='$email', biodata = '$bio',profile_picture = '$path' WHERE id_user = '$poster'";
		if(mysqli_query($con,$query)){
			$_SESSION['update_profile'] = true;
			move_uploaded_file($_FILES["profile"]["tmp_name"], "../".$path);
			header("Location: ../editUser.php");
		}	
	}
	else {
		//Update nama poster di tabel post
		$query = "UPDATE post SET nama_poster = '$nama' WHERE nama_poster ='$oldname' ";
		if(mysqli_query($con,$query)){
			//Update username di session
			$_SESSION['username'] = $nama;
		}
		//Update username di tabel user_list
		$query = "UPDATE user_list SET username ='$nama', email ='$email', biodata='$bio' WHERE id_user = '$poster'";
		if(mysqli_query($con,$query)){
			$_SESSION['username'] = $nama;
			$_SESSION['update_profile'] = true;
			header("Location: ../editUser.php");
		}
		
	}	
?>