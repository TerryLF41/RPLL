<?php
	session_start();
	include "connectDatabase.php";
	$nama = $_POST['topik'];
	$deskripsi = $_POST['deskripsi'];
	$query ="SELECT * FROM topik";
	$result = mysqli_query($con,$query);
	while($row = mysqli_fetch_array($result)){
		if($row['nama'] == $nama){
			$_SESSION['status_topic'] = false;
			header("Location: ../daftarTopik.php");
			exit();
		}
	}
	if(isset($_SESSION['admin'])){
		if($_SESSION['admin'] == true){
			$query = "INSERT INTO topik (nama,deskripsi,approve) 
			VALUES ('$nama','$deskripsi',1)";
		}
	}
	else {
		$query = "INSERT INTO topik (nama,deskripsi) 
		VALUES ('$nama','$deskripsi')";
	}
	if(mysqli_query($con,$query)){
		$_SESSION['status_topic'] = true;
	}
	header("Location: ../daftarTopik.php");
	mysqli_close($con);
?>