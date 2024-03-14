<?php
	session_start();
	include "connectDatabase.php";
	$id = $_POST['id_topik'];
	$nama = $_POST['nama'];
	$deskripsi = $_POST['deskripsi'];
	$topik = $_POST['topik'];
	$sql="INSERT INTO thread (id_topik,nama,deskripsi)
	VALUES ($id,'$nama','$deskripsi')";
	if(mysqli_query($con,$sql)){
		$_SESSION['addthread'] = true;
		header("Location: ../thread.php?id=$id&topik=$topik");
	}
	mysqli_close($con);
?>