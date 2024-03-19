<?php
	$con = mysqli_connect("localhost","root","","my_forum");
	if($con === FALSE){
		die ("Koneksi Eror : ".mysqli_connect_error());
	}
?>