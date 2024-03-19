<?php
	session_start();
	include "connectDatabase.php";
	$id = $_GET['id'];
	$sql = "DELETE FROM topik WHERE id_topik = '$id'";
	if(mysqli_query($con,$sql)){
		$_SESSION['delete'] = true;
		header("Location: ../topikAdmin.php?");
	} else{
		echo "Error Approval Topik : ".mysqli_error($con);
	}
	mysqli_close($con);
?>