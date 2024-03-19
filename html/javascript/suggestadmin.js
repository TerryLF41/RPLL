function showHint(str){
    if(str.length == 0){
        document.getElementById("txtHint").innerHTML="";
    }
    else {
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("txtHint").innerHTML = xmlhttp.responseText;
        }
        };
        xmlhttp.open("GET","function/suggestTopik.php?q="+str,true);
        xmlhttp.send();
    }
}
function showHintThread(str){
    var id = document.getElementById("hidden").innerHTML;
    if(str.length == 0){
        document.getElementById("txtHint").innerHTML="";
    }
    else {
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("txtHint").innerHTML = xmlhttp.responseText;
        }
        };
        xmlhttp.open("GET","function/suggestThread.php?q="+str+"&id="+id,true);
        xmlhttp.send();
    }
}