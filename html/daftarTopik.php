<!DOCTYPE html>
<html>
	<head>
		<title>Catalog</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/daftarTopik.css">
		<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
		<script type="text/javascript" src="javascript/modal.js"></script>

	</head>
	<body>
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "header.php";
			if(isset($_SESSION['status_topic'])){
				if($_SESSION['status_topic'] == true){
					if(isset($_SESSION['admin'])){
						echo "<script>alert('Topik berhasil ditambahkan')</script>";
					}	
					else {
						echo "<script>alert('Topik berhasil diajukan')</script>";
					}
				}
				else {
					echo "<script>alert('Topik dengan nama tersebut sudah ada')</script>";
				}
				unset ($_SESSION['status_topic']);
			}
		?>
		<main>
			<h1>Daftar Topik</h1>
			<br>
			<div class="list">
			<table style="border-collapse:collapse;width:100%">
				<?php
					include "connectDatabase.php";
					$sql = "SELECT * FROM topik WHERE approve = 1";
					$result = mysqli_query($con, $sql);
					if(mysqli_num_rows($result) == 0){
						echo "<tr><th colspan='4'><h2>Topik kosong</h2></th></tr> ";		
					}
					else {
						echo '<tr style="text-align:left;">
							<th>Nama</th>
							<th>Deskripsi</td>
							<th>Tanggal Dibuat</th>
							<th>&nbsp</th>
							</tr>';
						while($row = mysqli_fetch_array($result)) {
							$id = $row['id_topik'];
							$nama = $row['nama'];
							$desc = $row["deskripsi"];
							$date = $row['tanggal_dibuat'];
							$countsql = "SELECT * FROM thread WHERE id_topik = $id";
							$count = mysqli_num_rows(mysqli_query($con, $countsql)); 
							echo "<tr>";
							echo "<td> <a href='thread.php?id=$id&amp;topik=$nama'>". $row['nama'] ."</a></td>";
							echo "<td> <a href='thread.php?id=$id&amp;topik=$nama'>".$row["deskripsi"]."</a></td>";
							echo "<td> <a href='thread.php?id=$id&amp;topik=$nama'>". $row['tanggal_dibuat'] ."</a></td>";
							echo "<td> $count threads</td>";
						echo "</tr>";
						}
					}
					mysqli_close($con);
				?>
				<tr class="add"><td colspan="4">Add New Topic</td></tr>
			</table>
			</div>
			<br><br>
		</main>
		<div class="modal">
			<form class="form" method="POST" action="function/pengajuan_topik.php" >
				<h2 class="title">Add New Topic</h2>
				<label><b>Nama Topik</b></label><br>
				<input type="text" name="topik" placeholder="Input nama topik"><br>
				<label><b>Deskripsi</b></label><br>
				<textarea name="deskripsi" placeholder="Input deskripsi topik"></textarea><br>
				<button type="submit" id="submit">Submit</button>
				<button type="reset" id="cancel">Cancel</button>
			</form>
		</div>
		<?php include "footer.php"; ?>
	</body>
</html>