<?php 
    $con = mysqli_connect("localhost","root","","my_forum");
    $sql = "CREATE TABLE `user_list` (
      `id_user` int(11) NOT NULL AUTO_INCREMENT,
      `username` varchar(200) NOT NULL,
      `password` varchar(200) NOT NULL,
      `email` varchar(50) NOT NULL,
      `biodata` text DEFAULT NULL,
      `profile_picture` varchar(200) DEFAULT 'upload/profile/default.png',
      `status_ban` tinyint(1) NOT NULL DEFAULT 0,
      `status_admin` tinyint(1) NOT NULL DEFAULT 0,
      `tanggal_pembuatan` datetime DEFAULT current_timestamp(),
      PRIMARY KEY (`id_user`)
    )";
    mysqli_query($con,$sql);

    $sql = "CREATE TABLE `topik` (
      `id_topik` int(11) NOT NULL AUTO_INCREMENT,
      `nama` varchar(100) NOT NULL,
      `deskripsi` varchar(255) NOT NULL,
      `tanggal_dibuat` date NOT NULL DEFAULT current_timestamp(),
      `approve` tinyint(1) NOT NULL DEFAULT 0,
      PRIMARY KEY (`id_topik`)
    )";
    mysqli_query($con,$sql);

    $sql = "CREATE TABLE `thread` (
      `id_thread` int(11) NOT NULL AUTO_INCREMENT,
      `id_topik` int(11) NOT NULL,
      `nama` varchar(50) NOT NULL,
      `deskripsi` varchar(255) NOT NULL,
      `tanggal_dibuat` date NOT NULL DEFAULT curdate(),
      PRIMARY KEY (`id_thread`)
    )";
    mysqli_query($con,$sql);

    $sql = "CREATE TABLE `post` (
      `id_post` int(11) NOT NULL AUTO_INCREMENT,
      `id_thread` int(11) NOT NULL,
      `teks` varchar(1000) NOT NULL,
      `image` varchar(255) DEFAULT NULL,
      `tanggal` datetime DEFAULT current_timestamp(),
      `nama_poster` varchar(255) NOT NULL,
      PRIMARY KEY (`id_post`)
    )";
    mysqli_query($con,$sql);

    $sql = "INSERT INTO `post` (`id_post`, `id_thread`, `teks`, `image`, `tanggal`, `nama_poster`) VALUES
    (1, 1, 'Ask me anything ', 'upload/post/1616214856261.png', '2022-04-17 12:37:12', 'Admin'),
    (2, 2, 'Veh Deh Veh!', 'upload/post/1648079129921.jpg', '2022-04-17 12:42:49', 'John'),
    (4, 3, 'Apakah ini bola ini merupakan sebuah bola? Atau prisma?', 'upload/post/bola.png', '2022-04-17 15:25:25', 'Admin'),
    (5, 3, 'Itu kubus min', NULL, '2022-04-17 15:26:02', 'John'),
    (6, 4, 'Cats', 'upload/post/kot.jpg', '2022-04-17 16:31:17', 'John'),
    (10, 1, 'Siapa kamu?', NULL, '2022-04-17 17:50:10', 'User123'),
    (11, 4, '', 'upload/post/1616370319473.jpg', '2022-04-17 17:53:23', 'User123'),
    (12, 4, 'Cat', 'upload/post/1615537924938.jpg', '2022-04-17 17:55:44', 'Admin'),
    (13, 1, 'Admin forum ini', NULL, '2022-04-17 17:59:02', 'Admin'),
    (14, 1, 'Ok', NULL, '2022-04-18 15:10:47', 'John'),
    (15, 1, 'Haha', 'upload/post/1567081184243.gif', '2022-04-18 16:51:50', 'Admin'),
    (16, 1, 'Hahaha', 'upload/post/1600036985362.gif', '2022-04-18 17:01:43', 'Admin'),
    (17, 1, 'HAHAHA', 'upload/post/schizo.jpg', '2022-04-18 17:02:20', 'Admin');";
    echo "Tabel-tabel telah berhasil dibuat";
    mysqli_query($con,$sql);
    
    $sql = "INSERT INTO `thread` (`id_thread`, `id_topik`, `nama`, `deskripsi`, `tanggal_dibuat`) VALUES
    (1, 1, 'Ask Admin', 'Pertanyaan kepada Admin Chatters Forum', '2022-04-17'),
    (2, 2, 'Veh Deh Veh', 'VDV', '2022-04-17'),
    (3, 2, 'Random', 'Random', '2022-04-17'),
    (4, 5, 'Cat', 'Post cat', '2022-04-17'),
    (15, 1, 'News', 'Berita mengenai Chatters Forum', '2022-04-18')";
    mysqli_query($con,$sql);

    $sql = "INSERT INTO `topik` (`id_topik`, `nama`, `deskripsi`, `tanggal_dibuat`, `approve`) VALUES
    (1, 'Chatters Forum', 'Diskusi untuk Chatters Forum', '2022-04-17', 1),
    (2, 'Random', 'Pure Chaos', '2022-04-17', 1),
    (4, 'Gaming', 'Wadah diskusi untuk game', '2022-04-17', 1),
    (5, 'Animals', 'Diskusi hewan', '2022-04-17', 1),
    (6, 'Politics', 'Diskusi politik', '2022-04-17', 1),
    (7, 'Sport', 'Diskusi olahraga', '2022-04-18', 1),
    (9, 'Musik Rock', 'ABC', '2022-04-18', 1)";
    mysqli_query($con,$sql);

    $sql = "INSERT INTO `user_list` (`id_user`, `username`, `password`, `email`, `biodata`, `profile_picture`, `status_ban`, `status_admin`, `tanggal_pembuatan`) VALUES
    (2, 'Ban', '$2y$10\$pHB3sIDfP7cTQYM6Leg/e.nqd5RVH7YaGQWnpGkMYn9dq2zjlgGCa', 'admin@hotmail.com', NULL, 'upload/profile/default.png', 1, 0, '2022-04-14 12:13:32'),
    (3, 'Ahmad', '$2y$10\$hi5or49tR24sSelLMCHNZ.v.WsSRGvelX6IeY.yyi/s/GtFLDSt7m', 'ABC@gmail.com', NULL, 'upload/profile/default.png', 1, 0, '2022-04-14 12:16:25'),
    (4, 'Admin', '$2y$10\$k4oiovBB6hOaMUCPA9eomeW/NAINWlnsqkC4t7KSwCJI7DjGWs0E2', 'admin@hotmail.com', 'Admin', 'upload/profile/1616214856261.png', 0, 1, '2022-04-14 12:27:10'),
    (5, 'John', '$2y$10\$SL60QGpK3eaHvUJjKVgmt..RDiWY1viWVVy3z6h5ePERIWVzt1TAK', 'uncup@protonmail.com', 'Lorem Ipsum dolor sit amet', 'upload/profile/Profile1.jpg', 0, 0, '2022-04-14 16:57:11'),
    (6, 'Guy', '$2y$10$7zaukAQMoDy7fBweufeXZ.snD7gTL.HRP3mNsvJlEMf.2ZHZyPnxe', 'ABC@gmail.com', 'The man', 'upload/profile/unknown-1.png', 0, 0, '2022-04-16 22:04:23'),
    (7, 'Jaka', '$2y$10\$Jpy5552kspj..56MTts3aOYpXM2xzm8gKFvsmDJJI6j9Yc3a6IoNu', 'jaka@hotmail.com', NULL, 'upload/profile/default.png', 0, 0, '2022-04-16 22:07:08'),
    (8, 'User123', '$2y$10\$cKyZaBBkMH5xq4WRPWSN5.c6mbVjCz8EIIHX.3qqJbJS4nUcWWTQa', 'random1245712@gmail.com', '', 'upload/profile/1619096719295.png', 0, 0, '2022-04-17 17:46:55');";
    mysqli_query($con,$sql);
?>








