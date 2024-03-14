<?php
	session_start();
	include "connectDatabase.php";
	$id = $_GET['id'];
	$sql = "UPDATE topik SET approve = 1 WHERE id_topik = '$id'";
	if(mysqli_query($con,$sql)){
		$_SESSION['approve'] = true;
		header("Location: ../topikAdmin.php");
	} else{
		echo "Error Approval Topik : ".mysqli_error($con);
	}
	mysqli_close($con);
	
?>