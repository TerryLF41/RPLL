<?php
    include "connectDatabase.php";
    $query = "SELECT * FROM topik WHERE approve = 1 ORDER BY id_topik DESC LIMIT 5";
    $result = mysqli_query($con,$query);
    while($row = mysqli_fetch_assoc($result)) {
        if($row['approve'] == true){
            $id = $row['id_topik'];
            $nama = $row['nama'];
            echo '<form name="topik" action="thread.php" method="get">';
            echo "<button class='show' type='submit' name ='id' value='$id'>$nama</button>";
            echo "<input type='hidden' name='topik'value = $nama></input>";
            echo "</form>";
        }
    }
    mysqli_close($con);
?>
