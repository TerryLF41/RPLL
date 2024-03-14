<!DOCTYPE php>
<html>
	<head>
		<title>About Us</title>
		<meta charset= "UTF-8">
		<meta name="keywords" content= "HTML, CSS, PHP, XML, JS">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="css/header.css">		
		<link rel="stylesheet" href="css/footer.css">
        <link rel="stylesheet" href="css/about.css">
        <script type="text/javascript" src="javascript/register.js"></script>
    </head>
    <body>
        <?php 
            session_start();
            include "function/cek_session.php";
			include "header.php";
        ?>
        <main>
            <div class="content top">
                <img src="upload/others/chattergrey.png">
                <h2>Diskusi Tanpa Batas</h2>
            </div>
            <br>
            <div class="content">
                <h3>Selamat datang di Chatters Forum!</h3>
                <p>Chatters Forum adalah sebuah tempat dimana Anda dapat mendiskusikan apa saja yang Anda inginkan tanpa batasan topik </p>
                <p>Ingin membahas mengenai suatu topik baru? Tinggal buat saja!</p>
                <p>Kirimkan request topik beserta deskripsinya dan tunggu approval Admin</p>
                <p>Setelah disetujui, maka Anda bisa mengepost apapun yang berkaitan dengan topik tersebut</p>
                <h3><b>Visi-Misi</b></h3>
                <p>Visi</p>
                <p>Forum kami mempunyai visi untuk menjadi wadah diskusi dimana pengguna dapat mengekspresikan pendapat mereka mengenai suatu topik dengan bebas</p>
                <p>Misi</p>
                <p>1. Menyediakan wadah diskusi bagi para pengguna forum kami</p>
                <p>2. Menjadi sarana penyebaran informasi bagi para pengguna forum kami</p>
                <p>3. Mempromosikan bentuk media sosial forum yang sudah mulai berkurang eksistensinya</p>
                </ul>
            </div>
            <div class="content">
                <h3>Tim Penyusun</h3>
                <p>1121004 - Jericho Kuskanto</p>
                <p>1121008 - Enriek Naldo Patera</p>
                <p>1121018 - Friendly Sejati Bunardi</p>
            </div>
        </main><br>
        <?php include "footer.php"?>
    </body>
</html>