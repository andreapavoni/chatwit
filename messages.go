type Message struct {
  // The text of the message
  Text string

  // Connection relative to the current Room
  connection *Connection
}
