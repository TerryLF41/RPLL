<!DOCTYPE html>
<html>
	<head>
		<title>Thread</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
		<link rel="stylesheet" href="css/thread.css">
		<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
		<script type="text/javascript" src="javascript/thread.js"></script>
		<script type="text/javascript" src="javascript/modal.js"></script>
	</head>
	<body>
		<?php 
			session_start();
			include "function/cekcookies.php";
			include "function/cek_session.php";
			include "header.php";
			if(isset($_SESSION['addthread'])){
				echo "<script>alert('Thread telah berhasil ditambahkan')</script>";
				unset($_SESSION['addthread']);
			}
			$_SESSION['id_topik'] = $_GET['id'];
			$_SESSION['nama_topik'] = $_GET['topik'];
		?>
		<main>
			<form id='form-search' method='POST' action='thread.php?id=<?php echo $_GET['id'] ?>&topik=<?php echo $_GET['topik'] ?>'>
				<input class='search' list='txtHint' name='name' onkeyup='showHint(this.value)'  placeholder='Search thread' autocomplete='on'/>
				<datalist id='txtHint'>
					
				</datalist>
				<input type='hidden' value=".$_GET['id']." name='id_topik'/>
				<input type='hidden' value= ".$_GET['topik']." name='nama_topik'/>
				<input type='submit' value='Search' name='search'/>
				<input type='submit' value='Reset' name='reset' />
			</form>
		<br>
		<div class ="list">
		<table style="border-collapse:collapse;width:100%">
			<?php
			include "connectDatabase.php";
				if(isset($_POST['search'])){
					if(isset($_POST['name'])){
						$name = $_POST['name'];
						$query =mysqli_query( $con,"SELECT * FROM thread WHERE nama LIKE '%$name%' AND id_topik = '".$_GET['id']."'");
						if (mysqli_num_rows($query) == 0){
							echo "<script>alert('Thread dengan nama tersebut tidak ditemukan')</script>";
							$query = mysqli_query($con, "SELECT * FROM thread WHERE id_topik = '".$_GET['id']."'");
						}
					}
				} else{
					$query = mysqli_query($con, "SELECT * FROM thread WHERE id_topik = '".$_GET['id']."'");
				}
				if(mysqli_num_rows($query) == 0){
					echo "<tr><th colspan='4'><h2>Thread kosong</h2></th></tr> ";				
				}
				else {
					echo'<tr style="text-align:left;">
					<th>Topik</th>
					<th>Nama Thread</th>
					<th>Deskripsi</td>
					<th>Tanggal dibuat</th>
					<th>&nbsp</th></tr>';
				}
				while($data = mysqli_fetch_array($query)) {
					$id = $data['id_thread'];
					$nama = $data['nama'];
					$desc = $data["deskripsi"];
					$date = $data['tanggal_dibuat'];
					$countsql = "SELECT * FROM post WHERE id_thread = $id";
					$count = mysqli_num_rows(mysqli_query($con, $countsql)); 
					echo "<tr>";
					echo "<td> <a href='post.php?id=$id&amp;topik=$nama'>". $_GET['topik']."</a></td>";
					echo "<td><a href='post.php?id=$id&amp;topik=$nama'>  $nama </a></td>";
					echo "<td><a href='post.php?id=$id&amp;topik=$nama'> $desc </a></td>";
					echo "<td><a href='post.php?id=$id&amp;topik=$nama'>  $date </a></td>";
					echo "<td>$count posts</td>";
					echo "</tr>";
				}
				mysqli_close($con);
			?>
			<tr class="add"><td colspan="5">Add New Thread</td></tr>
		</table>
		</div>
		</main>
		<div class="modal">
			<form class="form" method="POST" action="function/addThread.php" >
				<h2 class="title">Add New Thread</h2>
				<label><b>Nama Thread</b></label><br>
				<input type="text" name="nama" placeholder="Input nama thread"><br>
				<label><b>Deskripsi</b></label><br>
				<textarea name="deskripsi" placeholder="Input deskripsi thread"></textarea><br>
				<input type='hidden' name='id_topik' value='<?php echo $_GET['id'] ?>'></input>
				<input type='hidden' name='topik' value='<?php echo $_GET['topik'] ?>'></input>
				<button type="submit" id="submit">Submit</button>
				<button type="reset" id="cancel">Cancel</button>
			</form>
		</div>
		<div id="hidden"><?php echo $_GET['id']; ?></div>
		<?php include "footer.php"; ?>
	</body>
</html>