function short() {
    let url = document.getElementById("long_url").value

    if (url === "") {return}

    $.ajax({
        type: "POST",
        url: "/api/create",
        data: {"url": url},
        success: success,
        dataType: "json"
      });
}

function success(e) {
    document.getElementById("long_url").value = e.url
}