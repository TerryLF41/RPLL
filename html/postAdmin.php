<!DOCTYPE html>
<html>
	<head>
		<title>Moderate Post</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/postadmin.css">
	</head>
	<body>
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "function/cekAdmin.php";
			include "header.php";
			if(isset($_SESSION['delete'])){
				echo "<script>alert('Post telah didelete')</script>";
				unset($_SESSION['delete']);
			}
			if(isset($_SESSION['createXML'])){
				echo "<script>alert('File XML telah berhasil dibuat dan disimpan di folder XML')</script>";
				unset($_SESSION['createXML']);
			}
		?>
		<main>
			<form method="POST" action="postAdmin.php?id=<?php echo $_GET['id'];?>&thread=<?php echo $_GET['thread'];?>">
				<input class="search" type ="text" name="post" placeholder="Search post"/>
				<button type="submit">Search</button>
				<button type="submit">Reset</button>
			</form>
			<br>
			<form method='POST' action='approval.php'>
				<table style="border-collapse:collapse;width:100%">
					<tr style="text-align:left;">
						<th>No. Post</th>
						<th>Thread</th>
						<th>Poster</th>
						<th>Post</th>
						<th>Image</td>
						<th>Tanggal Post</th>
						<th>Aksi</th>
					</tr>
					<?php
						include "connectDatabase.php";
						$id_thread = $_GET['id'];
						$nama_thread = $_GET['thread'];
						if(isset($_POST['post'])){
							if($_POST['post'] != ""){
								$post = $_POST['post'];
								$query = "SELECT * FROM post WHERE teks LIKE '%$post%' AND id_thread = $id_thread";
							}
							else {
								$query = "SELECT * FROM post WHERE id_thread = $id_thread";
							}
						}
						else {
							$query = "SELECT * FROM post WHERE id_thread = $id_thread";
						}
						$result = mysqli_query($con,$query);
						while($data = mysqli_fetch_array($result)) {
							$id = $data['id_post'];
							$post = nl2br($data['teks']);
							$date = $data["tanggal"];
							$poster = $data['nama_poster'];
							$img = $data['image'];
							echo "<tr>";
							echo "<td> $id </td>";
							echo "<td id='thread'>$nama_thread</td>";
							echo "<td> $poster </td>";
							echo "<td id='post'> $post </td>";
							if($img == null){
								echo "<td>No Image</td>";
							}
							else {
								echo"<td id='img'><a href='$img' target='_blank'><img height='200' width='200' src='$img'></a></td>";
							}
							echo "<td id='date'> $date </td>";
							echo "<td id='aksi'> <a href='function/deletePost.php?id=$id&amp;id-thread=$id_thread&amp;nama-thread=$nama_thread'>Delete</a>";
							echo "</tr>";
						}
						mysqli_close($con);
					?>
				</table>
			</form>
			<br>
			<a href ="topikAdmin.php">Back to Topik</a><br><br>
			<a href="function/convertxml.php?tipe=post&id-thread=<?php echo$id_thread; ?>&nama-thread=<?php echo $nama_thread; ?>">Convert to XML</a>
		</main>
		<?php include "footer.php"; ?>
	</body>
</html>