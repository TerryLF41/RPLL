<?php
	include ("connectDatabase.php");
	$q = $_GET['q'];
	$arrNama = [];

	$sql ="SELECT * FROM topik";
	$result = mysqli_query($con,$sql);
	while($row = mysqli_fetch_assoc($result)) {
		array_push($arrNama,$row['nama']);
	}
	mysqli_close($con);

	if(strlen($q) > 0){
		$result = "";
		for($i = 0; $i < count($arrNama);$i++){
			if(strtolower($q) == strtolower(substr($arrNama[$i],0,strlen($q)))){
				$result = $result."<option value='".$arrNama[$i]."'>";
			}
		}
	}
	echo $result;
?>