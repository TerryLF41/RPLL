<?php
	session_start();
	include "connectDatabase.php";
	$idThread = $_POST['id_thread'];
	$isi = $_POST["text"];
	$isi = htmlspecialchars($isi);
	$poster = $_SESSION['username'];
	if(strlen($isi) > 1000){
		$_SESSION['posted'] = 'panjang';
		header("Location: ../post.php?id=$idThread");
		exit();
	}
	if(file_exists($_FILES["image"]["tmp_name"])){
		$allowed = ['png','jpg','jpeg','webm','gif'];
		$filetype = explode('.',$_FILES["image"]["name"]);
		$filetype = end($filetype);
		$filetype = strtolower($filetype);
		if(!in_array($filetype,$allowed)){
			$_SESSION['posted'] = 'extension';
			header("Location: ../post.php?id=$idThread");
			exit();
		}
		$path = "upload/post/";
		$name = $_FILES["image"]["name"];
		$name = htmlspecialchars($name);
		$name = str_replace(" ","",$name);
		$path = $path.basename($name);
		$query = "INSERT INTO post(id_thread, teks, image, nama_poster)
		VALUES ('$idThread', '$isi', '$path', '$poster')";
		if(mysqli_query($con,$query)){
			$_SESSION['posted'] = true;
			move_uploaded_file($_FILES["image"]["tmp_name"], "../".$path);
			header("Location: ../post.php?id=$idThread");
		}	
	}
	else {
		$query = "INSERT INTO post(id_thread, teks, nama_poster)
		VALUES ('$idThread', '$isi', '$poster')";
		if(mysqli_query($con,$query)){
			$_SESSION['posted'] = 'sukses';
			header("Location: ../post.php?id=$idThread");
		}	
	}
	
	mysqli_close($link);
?>