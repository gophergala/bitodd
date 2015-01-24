$(document).ready(function() {
	console.log("Document loaded");

	if (window.WebSocket) {
		log("WS supported");
		initConnection();
	} else {
		log("WS not supported");
	}
});

var debug = true;

function log(message) {
	if (debug) {
		console.log(message);
	}
}

function initConnection() {

	var loc = window.location;
	var wsUrl = "ws://" + loc.hostname;
	if (loc.port != 80 && loc.port != 443) {
		wsUrl += ":" + loc.port;
	}
	wsUrl += loc.pathname + "ws";

	var ws = new WebSocket(wsUrl);
	ws.onopen = function() {
		log("WS connection opened");
	};
	ws.onmessage = function(evt) {
		var msg = JSON.parse(evt.data);
		log("WS message received: " + evt.data);
	};
	ws.onclose = function() {
		log("WS connection close");
	};
}