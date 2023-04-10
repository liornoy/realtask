package scheduler
import (
    "fmt"
    "net"
    "os"
)
type Scheduler struct{}

const (
    CONN_HOST = "localhost"
    CONN_TYPE = "tcp"
)

func main() {
}

func New(port string) error{
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+port)
    if err != nil {
	return fmt.Errorf("error listening:", err.Error())
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + port)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }

    return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}

func checkPort(port string) bool {
    conn, err := net.Dial("tcp", "localhost:"+port)
    if err != nil {
        // Port is not open
        return false
    }
    defer conn.Close()

    // Port is open
    return true
}
