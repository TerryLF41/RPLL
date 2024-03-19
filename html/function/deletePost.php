<?php
	session_start();
	include "connectDatabase.php";
	$id = $_GET['id'];
	$id_thread = $_GET['id-thread'];
	$nama_thread = $_GET['nama-thread'];
	$sql = "DELETE FROM post WHERE id_post = '$id'";
	if(mysqli_query($con,$sql)){
		$_SESSION['delete'] = true;
		header("Location: ../postAdmin.php?id=$id_thread&thread=$nama_thread");
	} else{
		echo "Error Approval Topik : ".mysqli_error($con);
	}
	mysqli_close($con);
?>