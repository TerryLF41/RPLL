<!DOCTYPE php>
<?php
	include "connectDatabase.php";
	$query = "SELECT * FROM user_list";
	if(isset($_GET['name']) && $_GET['name'] != ""){
		$query .= " WHERE username='".$_GET['name']."'";
		if(isset($_GET['ban'])){
			$query .= " AND status_ban = 1";
		}
		if(isset($_GET['admin'])){
			$query .= " AND status_admin = 1";
		}
	}
	else {
		if(isset($_GET['ban'])){
			$query .= " WHERE status_ban = 1";
			if(isset($_GET['admin'])){
				$query .= " AND status_admin = 1";
			}
		}
		else if(isset($_GET['admin'])){
			$query .= " WHERE status_admin = 1";
			if(isset($_GET['banned'])){
				$query .= " AND status_ban = 1";
				
			}
		}
	}
	$result = mysqli_query($con, $query);
	if(mysqli_num_rows($result) == 0){
		echo "<h2>User not found</h2>";
	} 
	else {
		echo 
		'<tr style="text-align:left;">
			<th>Id user</th>
			<th>Username</th>
			<th>Profile Picture</td>
			<th>Email</th>
			<th>Biodata</th>
			<th>Status</th>
			<th>Tanggal Join</th>
			<th>Action</th>
		</tr>';
		while($data = mysqli_fetch_array($result)) {
			$id = $data['id_user'];
			$nama = $data['username'];
			$img = $data['profile_picture'];
			$email = $data["email"];
			$bio = $data["biodata"];
			$admin =  $data["status_admin"];
			$ban = $data["status_ban"];
			$date = $data['tanggal_pembuatan'];
			
			echo "<tr>";
			echo "<td>$id</td>";
			echo "<td>$nama</td>";
			if($img == null){
				echo "<td>No Profile Picture</td>";
			}
			else {
				echo "<td><a href='$img' target='_blank'><img height='150px' width='150ox' src='$img'></a></td>";
			}
			echo "<td>$email</td>";
			echo "<td>$bio</td>";
			$status = "";
			if($ban){
				$status .="Banned ";
			}
			if($admin){
				$status .="Admin";
			}
			else {
				$status .="Member";
			}
			echo "<td>$status</td>";
			echo "<td>$date</td>";

			$action = "";
			if($ban){
				$action .="<a href='function/unbanUser.php?id=$id'>Unban User</a>";
			}
			else {
				$action .="<a href='function/banUser.php?id=$id'>Ban User</a>";
			}
			if($admin){
				$action .="<br><br><a href='function/removeAdmin.php?id=$id'>Revoke Admin</a>";
			}
			else {
				$action .="<br><br><a href='function/setAdmin.php?id=$id'>Grant Admin</a>";
			}
			echo "<td>$action</td></tr>";
		}
	}
	mysqli_close($con);
?>