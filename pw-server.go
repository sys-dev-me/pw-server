package main


import "log"
import "fmt"
import "os"
import "net"
import "bufio"
import "strconv"
import "strings"

const (
	CONN_TYPE = "tcp"
)


var pChar ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321{}()$-=!@#_/\\_&"

/*

func generatePassword ( pwSize int ) (string) {
	src := []rune(pChar)
	res := make([]rune, pwSize)

	for i := range res {
        res[i] = src[rand.Intn(len(src)-1)]
    }

	return string( res )
}

*/



func handleConnection ( conn net.Conn, app *Application ) {

	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))


	for {
        req, err := rw.ReadString('\n')
        if err != nil {
            rw.WriteString("failed to read input")
            rw.Flush()
            return
        }
		
		req = strings.Trim( req, "\n" )	
		pwLen, err := strconv.Atoi( req )
		if err !=nil {
			rw.WriteString( "Wrong data")
			return 
		}

	pw := app.Generate ( pwLen )

        rw.WriteString( fmt.Sprintf( "%v\n", pw )  )
        rw.Flush()
    }

}

func main () {


	app := new(Application)
	app.SetVoc( pChar )
	if len( os.Args ) < 2 {
		log.Println( "Not enough arguments. Usage example: pw-server [port]" ) 
		os.Exit( 1 )
	}

	
	port := os.Args[1]
	conn, err := net.Listen(CONN_TYPE, "localhost:"+port)

	// check if not listening
	if err !=nil {
		log.Println( "Unable to bind on port ", port )
	}

	log.Println( "Listening on: ", port )
	for {
		 res, err := conn.Accept()
		 if err !=nil {
		 	log.Println( "Something went wrong, can't accept connection" )
		 }

		go handleConnection(res, app)
		defer conn.Close()
	}
}
