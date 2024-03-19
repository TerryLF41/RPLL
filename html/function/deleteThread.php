<?php
	session_start();
	include "connectDatabase.php";
	$id = $_GET['id'];
	$idtopik = $_GET['id-topik'];
	$topik = $_GET['topik'];
	$sql = "DELETE FROM thread WHERE id_thread = '$id'";
	if(mysqli_query($con,$sql)){
		$_SESSION['delete'] = true;
		header("Location: ../threadAdmin.php?id=$idtopik&topik=$topik");
	} else{
		echo "Error Approval Topik : ".mysqli_error($con);
	}
	mysqli_close($con);
?>