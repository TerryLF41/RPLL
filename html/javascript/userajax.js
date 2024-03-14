function search(){
    var xmlhttp = new XMLHttpRequest();
    var php = "function/getUser.php";
    if(document.getElementById("username").value == ""){
        if(document.getElementById("ban").checked){
            php += "?ban=1";
            if(document.getElementById("admin").checked){
                php += "&admin=1";
            }
        }
        if(document.getElementById("admin").checked){
            php += "?admin=1";
            if(document.getElementById("ban").checked){
                php += "&ban=1"
            }
        }
        
    }
    else {
        php += "?name=" + document.getElementById("username").value;
        if(document.getElementById("ban").checked){
            php += "&ban=1";
        }
        if(document.getElementById("admin").checked){
            php += "&admin=1";
        }
    }
    xmlhttp.onload = function() {
    if (this.readyState == 4 && this.status == 200) {
        document.getElementById("tabel-user").innerHTML=xmlhttp.responseText;
    }
    };
    xmlhttp.open("GET", php, true);
    xmlhttp.send();
}

function reset(){
    document.getElementById("username").value = "";
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onload = function() {
    if (this.readyState == 4 && this.status == 200) {
        document.getElementById("tabel-user").innerHTML=xmlhttp.responseText;
    }
    };
    xmlhttp.open("GET", "function/getUser.php", true);
    xmlhttp.send();
}

function suggestion(str){
    if(str.length == 0){
        document.getElementById("names").innerHTML = "";
    }
    else {
        var xmlhttp = new XMLHttpRequest();
        xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("names").innerHTML=xmlhttp.responseText;
        }
        };
        xmlhttp.open("GET", "function/suggestUser.php?q="+str, true);
        xmlhttp.send();
    }
}