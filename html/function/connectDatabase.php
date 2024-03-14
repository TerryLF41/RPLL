<?php
	date_default_timezone_set("Asia/Jakarta");
	$con = mysqli_connect("localhost","root","","my_forum");
	if($con === FALSE){
		die ("Koneksi Eror : ".mysqli_connect_error());
	}
?>