<!DOCTYPE html>
<html>
	<head>
		<title>Add New Thread</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
	</head>
	<body>
		<?php
			session_start();
			include "function/cek_session.php";
			if(isset($_POST['insert'])){
				include ("connectDatabase.php");
				$id = $_POST['id_topik'];
				$nama = $_POST['nama'];
				$deskripsi = $_POST['deskripsi'];
				$sql="INSERT INTO thread (id_topik,nama,deskripsi)
				VALUES ($id,'$nama','$deskripsi')";
				if(mysqli_query($con,$sql)){
					$_SESSION['addthread'] == true;
					header("Location: topik.php");
				}
				mysqli_close($con);
			} else{
				$id = $_GET['id'];
				echo'<form method="POST" action="addThread.php">';
				echo 'Nama Thread : <input type="text" name="nama" value=""/><br>';
				echo 'Deskripsi : <textarea name="deskripsi"></textarea><br>';
				echo "<input type='hidden' name='id_topik' value='$id'></input>";
				echo '<input type="submit" name="insert" value="Insert"/>';
				echo '</form>';
			}
		?>
	</body>
</html>