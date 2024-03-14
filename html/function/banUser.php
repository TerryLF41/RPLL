<?php
	session_start();
	include "connectDatabase.php";
	$id = $_GET['id'];
	$sql = "UPDATE user_list SET status_ban= 1 WHERE id_user = '$id'";
	if(mysqli_query($con,$sql)){
		$_SESSION['ban'] = true;
		header("Location: ../userAdmin.php");
	} else{
		echo "Error Approval Topik : ".mysqli_error($con);
	}
	mysqli_close($con);
?>