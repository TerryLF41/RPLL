<!DOCTYPE php>
<?php 
    session_start();
    include "connectDatabase.php";
    $filename = time();
    if($_GET['tipe'] == 'topik'){
        $query = "SELECT * FROM topik";
        $result = mysqli_query($con,$query);
        $xml = "<tabel_topik>";
        while($row = mysqli_fetch_assoc($result)) {
            $xml .= "<topik>";
            $xml .= "<id_topik>".$row['id_topik']."</id_topik>";
            $xml .= "<nama>".$row['nama']."</nama>";
            $xml .= "<deskripsi>".$row['deskripsi']."</deskripsi>";
            $xml .= "<tanggal_dibuat>".$row['tanggal_dibuat']."</tanggal_dibuat>";
            $xml .= "<approval>".$row['approve']."</approval>";
            $xml .= "</topik>";
        }
        $xml .= "</tabel_topik>";
        $xmlFile = new SimpleXMLElement($xml);
        $xmlFile->asXML("../xml/topik/$filename-topik.xml");
        $_SESSION['createXML'] = true;
        header("location: ../topikAdmin.php");
    }
    else if($_GET['tipe'] == 'thread'){
        $idtopik = $_GET['id-topik'];
	    $topik = $_GET['topik'];
        $query = "SELECT * FROM thread WHERE id_topik = $idtopik";
        $result = mysqli_query($con,$query);
        $xml = "<tabel_thread>";
        while($row = mysqli_fetch_assoc($result)) {
            $xml .= "<thread>";
            $xml .= "<id_thread>".$row['id_thread']."</id_thread>";
            $xml .= "<id_topik>".$row['id_topik']."</id_topik>";
            $xml .= "<nama>".$row['nama']."</nama>";
            $xml .= "<deskripsi>".$row['deskripsi']."</deskripsi>";
            $xml .= "<tanggal_dibuat>".$row['tanggal_dibuat']."</tanggal_dibuat>";
            $xml .= "</thread>";
        }
        $xml .= "</tabel_thread>";
        $xmlFile = new SimpleXMLElement($xml);
        $xmlFile->asXML("../xml/thread/$filename-$topik.xml");
        $_SESSION['createXML'] = true;
        header("Location: ../threadAdmin.php?id=$idtopik&topik=$topik");
    }
    else if($_GET['tipe'] == 'post'){
        $id_thread = $_GET['id-thread'];
	    $nama_thread = $_GET['nama-thread'];
        $query = "SELECT * FROM post WHERE id_thread = $id_thread";
        $result = mysqli_query($con,$query);
        $xml = "<tabel_post>";
        while($row = mysqli_fetch_assoc($result)) {
            $xml .= "<post>";
            $xml .= "<id_post>".$row['id_post']."</id_post>";
            $xml .= "<id_thread>".$row['id_thread']."</id_thread>";
            $xml .= "<teks>".$row['teks']."</teks>";
            $xml .= "<image>".$row['image']."</image>";
            $xml .= "<tanggal>".$row['tanggal']."</tanggal>";
            $xml .= "<nama_poster>".$row['nama_poster']."</nama_poster>";
            $xml .= "</post>";
        }
        $xml .= "</tabel_post>";
        $xmlFile = new SimpleXMLElement($xml);
        $xmlFile->asXML("../xml/post/$filename-$nama_thread.xml");
        $_SESSION['createXML'] = true;
        header("location: ../postAdmin.php?id=$id_thread&thread=$nama_thread");
    }
    else {
        $query = "SELECT * FROM user_list"; 
        $result = mysqli_query($con,$query);
        $xml = "<user_list>";
        while($row = mysqli_fetch_assoc($result)) {
            $xml .= "<user>";
            $xml .= "<id>".$row['id_user']."</id>";
            $xml .= "<username>".$row['username']."</username>";
            $xml .= "<password>".$row['password']."</password>";
            $xml .= "<email>".$row['category']."</email>";
            $xml .= "<biodata>".$row['price']."</biodata>";
            $xml .= "<profile_picture>".$row['profile_picture']."</profile_picture>";
            $xml .= "<ban>".$row['status_ban']."</ban>";
            $xml .= "<admin>".$row['status_admin']."</admin>";
            $xml .= "<tanggal_pembuatan>".$row['tanggal_pembuatan']."</tanggal_pembuatan>";
            $xml .= "</user>";
        }
        $xml .= "</user_list>";
        $xmlFile = new SimpleXMLElement($xml);
        $xmlFile->asXML("../xml/user/$filename-user.xml");
        $_SESSION['createXML'] = true;
        header("Location: ../userAdmin.php");
    }
?>
</html>