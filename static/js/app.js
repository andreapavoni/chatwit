$(function() {
  var conn;
  var msg = $("#msg");
  var log = $("#log");

  function appendLog(msg) {
    var d = log[0]
    var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
    msg.appendTo(log)

    if (doScroll) {
      d.scrollTop = d.scrollHeight - d.clientHeight;
    }
  }

  function timestamp() {
    return (new Date().toTimeString().replace(/.*(\d{2}:\d{2}:\d{2}).*/, "$1"));
  }

  function formatMsg(data) {
    switch(data.Event) {
      case 0:
        appendLog($("<div><b>[" + timestamp() + "] " + data.Arguments.Nickname + ":</b> " + data.Arguments.Body + "</div>"));
        break;
      case 1:
        appendLog($("<div><b>[" + timestamp() + "] " + "*** " + data.Arguments.Nickname + " has joined ***</b></div>"));
        break;
      case 2:
        appendLog($("<div><b>[" + timestamp() + "] " + "*** " + data.Arguments.Nickname + " has left ***</b></div>"));
        break;
    }
  }

  $("#form").submit(function() {
    if (!conn) {
      return false;
    }
    if (!msg.val()) {
      return false;
    }
    conn.send(msg.val());
    msg.val("");
    return false
  });

  if (window["WebSocket"]) {
    conn = new WebSocket($("body").data("socket"));

    conn.onclose = function(evt) {
      appendLog($("<div><b>*** Connection closed. ***</b></div>"))
    }

    conn.onmessage = function(evt) {
      console.log(evt.data);

      o = jQuery.parseJSON(evt.data)
      console.log(o["Event"]);
      console.log(o.Event);

      formatMsg(o);
    }
  } else {
    appendLog($("<div><b>*** Your browser does not support WebSockets. ***</b></div>"))
  }
});
