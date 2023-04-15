
socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function () {
    socket.send("Hello World");
}

socket.onmessage = function (event) {
    console.log(event.data);
}

socket.onclose = function (event) {
    if (event.wasClean) {
        alert('Connection closed cleanly');
    } else {
        alert('Connection died');
    }
    alert('Code: ' + event.code + ' reason: ' + event.reason);
}

socket.onerror = function (error) {
    alert("Error " + error.message);
}

