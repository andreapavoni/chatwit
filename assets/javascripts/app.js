//= require javascripts/foundation.min

$(function() {
  $(document).foundation();

  var conn;
  var msg = $("#msg");
  var log = $("#log");
  var form = $("#form");

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
    var nickname = data.Arguments.Nickname;
    var text;

    switch(data.Event) {
      case 0:
        text = " <strong>" + nickname + ":</strong> " + data.Arguments.Body;
        break;
      case 1:
        text = " <strong>*** " + nickname + " has joined ***</strong> ";
				updateStats(data.Arguments.Stats);
        break;
      case 2:
        text = " <strong>*** " + nickname + " has left ***</strong> ";
				updateStats(data.Arguments.Stats);
        break;
    }
    appendLog($("<div class='row'><div class='large-12 columns msg'> <p><strong><span class='timestamp'>" + timestamp() + "</span></strong>" + text + "</p> </div></div>"));
  }

	function updateStats(stats) {
		$("#room-stats").text(stats.UsersCount);
	}

  msg.on('keyup', function(e) {
    e = e || event;
    if (e.keyCode === 13 && !e.shiftKey) {
      form.submit();
    }
    return true;
  });

  form.submit(function() {
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

  if (form.lenght > 0) {
    if (window["WebSocket"]) {
      conn = new WebSocket(form.data("socket"));

      conn.onclose = function(evt) {
        appendLog($("<div class='row'><div class='large-12 columns msg'> <p><strong><span class='timestamp'>" + timestamp() + "</span></strong>   <strong>*** CONNECTION CLOSED ***</strong></p> </div></div>"));
      }

      conn.onmessage = function(evt) {
        o = jQuery.parseJSON(evt.data)
        formatMsg(o);
      }
    } else {
      appendLog($("<div class='row'><div class='large-12 columns msg'><p><strong>*** Your browser doesn't support WebSockets. ***</strong></p> </div></div>"));
    }
  }
});
