function getNewID(e) {
    if (e.keyCode === 13) {
        var link_input = document.getElementById("link");
        var ajax = new XMLHttpRequest();
        ajax.open("POST", "/create", true);
        ajax.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        ajax.send("url=" + link_input.value);
        ajax.onreadystatechange = function () {
            if (ajax.readyState === 4 && ajax.status === 200) {
                link_input.value = document.URL + "link/" + ajax.responseText
                link_input.disabled = true
            }
            return false;
        }
    }
}