<!DOCTYPE html>
<html>
	<head>
		<title>Moderate Topik</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/topikadmin.css">
		<script src="javascript/suggestadmin.js"></script>
	</head>
	<body>
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "function/cekAdmin.php";
			include "header.php";
			if(isset($_SESSION['approve'])){
				echo "<script>alert('Topik telah diapprove')</script>";
				unset($_SESSION['approve']);
			}
			if(isset($_SESSION['delete'])){
				echo "<script>alert('Topik telah didelete')</script>";
				unset($_SESSION['delete']);
			}
			if(isset($_SESSION['createXML'])){
				echo "<script>alert('File XML telah berhasil dibuat dan disimpan di folder XML')</script>";
				unset($_SESSION['createXML']);
			}
		?>
		<main>
			<form method="POST" action="topikAdmin.php">
				<input class='search' list='txtHint' name='name' onkeyup='showHint(this.value)'  placeholder='Search topik' autocomplete='on'/>
				<datalist id='txtHint'>
					
				</datalist>
				<button type="submit">Search</button>
				<button type="submit">Reset</button>
			</form>
			<br>
			<form method='POST' action='approval.php'>
				<table style="border-collapse:collapse;width:100%">
					<tr style="text-align:left;">
						<th>Id Topik</th>
						<th>Nama</th>
						<th>Deskripsi</td>
						<th>Tanggal Dibuat</th>
						<th>Aksi</th>
					</tr>
					<?php
						include "connectDatabase.php";
						if(isset($_POST['name'])){
							if($_POST['name'] != ""){
								$name = $_POST['name'];
								$query = "SELECT * FROM topik WHERE nama LIKE '%$name%'";
							}
							else {
								$query = "SELECT * FROM topik";
							}
						}
						else {
							$query = "SELECT * FROM topik";
						}
						$result = mysqli_query($con,$query);
						while($data = mysqli_fetch_array($result)) {
							$id = $data['id_topik'];
							$nama = $data['nama'];
							$desc = $data["deskripsi"];
							$date = $data['tanggal_dibuat'];
							echo "<tr>";
							echo "<td>$id</td>";
							echo "<td>$nama</td>";
							echo "<td>$desc</td>";
							echo "<td>$date</td>";
							echo "<td> <a href='threadAdmin.php?id=$id&amp;topik=$nama'>Open</a> / <a href='function/deleteTopik.php?id=$id'>Delete</a>";
							if ($data['approve'] == '0') {
								echo " / <a href='function/approval.php?id=$id'>Approve</a></td>";
							}
							echo "</tr>";
						}
						mysqli_close($con);
					?>
				</table>
			</form>
			<br>
			<a href="function/convertxml.php?tipe=topik">Convert to XML</a>
		</main>
		<?php include "footer.php"; ?>
	</body>
</html>