setInterval(post, 10000);
function post(){
    var id = document.getElementById("id").innerHTML;
    var xmlhttp = new XMLHttpRequest();
    var php = "function/includepost.php";
    xmlhttp.onload = function() {
    if (this.readyState == 4 && this.status == 200) {
        document.getElementById("post").innerHTML=xmlhttp.responseText;
        
    }
    };
    xmlhttp.open("GET", "function/includepost.php?id="+id, true);
    xmlhttp.send();
}