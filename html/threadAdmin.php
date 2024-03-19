<!DOCTYPE html>
<html>
	<head>
		<title>Moderate Thread</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/threadadmin.css">
		<script src="javascript/suggestadmin.js"></script>
	</head>
	<body>
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "function/cekAdmin.php";
			include "header.php";
			if(isset($_SESSION['delete'])){
				echo "<script>alert('Thread telah didelete')</script>";
				unset($_SESSION['delete']);
			}
			if(isset($_SESSION['createXML'])){
				echo "<script>alert('File XML telah berhasil dibuat dan disimpan di folder XML')</script>";
				unset($_SESSION['createXML']);
			}
			
		?>
		<main>
			<form method="POST" action="threadAdmin.php?id=<?php echo $_GET['id'];?>&topik=<?php echo $_GET['topik']; ?>" >
				<input class='search' list='txtHint' name='name' onkeyup='showHintThread(this.value)'  placeholder='Search thread' autocomplete='on'/>
				<datalist id='txtHint'>
					
				</datalist>
				<button type="submit">Search</button>
				<button type="submit">Reset</button>
			</form>
			<br>
			<form method='POST' action='approval.php'>
				<table style="border-collapse:collapse;width:100%">
					<tr style="text-align:left;">
						<th>Id Thread</th>
						<th>Topik</th>
						<th>Nama Thread</th>
						<th>Deskripsi</td>
						<th>Tanggal dibuat</th>
						<th>Aksi</th>
					</tr>
					<?php
						include "connectDatabase.php";
						$id_topik = $_GET['id'];
						$nama_topik = $_GET['topik'];
						if(isset($_POST['name'])){
							if($_POST['name'] != ""){
								$name = $_POST['name'];
								$query = "SELECT * FROM thread WHERE nama LIKE '%$name%' AND id_topik = $id_topik";
							}
							else {
								$query = "SELECT * FROM thread WHERE id_topik = $id_topik";
							}
						}
						else {
							$query = "SELECT * FROM thread WHERE id_topik = $id_topik";
						}
						$result = mysqli_query($con, $query);
						while($data = mysqli_fetch_array($result)) {
							$id = $data['id_thread'];
							$nama = $data['nama'];
							$desc = $data["deskripsi"];
							$date = $data['tanggal_dibuat'];
							echo "<tr>";
							echo "<td> $id </td>";
							echo "<td> $nama_topik </td>";
							echo "<td> $nama </td>";
							echo "<td> $desc </td>";
							echo "<td> $date </td>";
							echo "<td><a href='postAdmin.php?id=$id&amp;thread=$nama'>Open</a> / <a href='function/deleteThread.php?id=$id&amp;id-topik=$id_topik&amp;topik=$nama_topik'>Delete</a>";
							echo "</tr>";
						}
						mysqli_close($con);
					?>
				</table>
			</form>
			<br>
			<a href ="topikAdmin.php">Back to Topik</a><br><br>
			<a href="function/convertxml.php?tipe=thread&id-topik=<?php echo $id_topik; ?>&topik=<?php echo $nama_topik; ?>">Convert to XML</a>
			<div id="hidden"><?php echo $id_topik; ?></div>
		</main>
		<?php include "footer.php"; ?>
	</body>
</html>