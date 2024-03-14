<html>
<body>
<?php
	 $id_thread = $_GET['id'];
	 include ("connectDatabase.php");
	 $query = mysqli_query($link, 'SELECT * FROM post');
	 echo "<table border='1'>";
	 echo "<tr><td>id</td><td>teks</td><td>id_thread</td><td>tanggal</td><td>image</td><td>nama poster</td></tr>";
	 while($data = mysqli_fetch_array($query)) {
		 echo "<tr>";
		 $id =$data['id_post'] ;
		 echo "<td> $id </td>";
		 echo "<td> ". $data['teks'] ."</td>";
		 echo "<td> ". $data['id_thread'] ."</td>";
		 echo "<td> ". $data['tanggal'] ."</td>";
		 $img = $data['image'];
		 echo "<td> ". "<img height='200' width='150' src='$img'>" ."</td>";
		 echo "</tr> ";
	 }
	 echo "</table>";
	 mysqli_close($link);
	 echo '<form method="post" action="postFungsi.php" enctype="multipart/form-data">
	Text:
	<input type="text" name="text" value=""/><br>	
	IMG/File : 
	<input type="file" name="file" value=""/><br>
	<input type="hidden" name="id_thread" value="$id_thread">
	<input type="submit" value="Insert"/>
	</form>';
?>

</body>
</html>
