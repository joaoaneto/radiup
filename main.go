package main
import (
	"fmt"
	"time"
	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/playlist"
	"github.com/joaoaneto/radiup/dbconf"
)

func main(){
	// Declarations without meaning only for build the main.go

	var music []cycle.Music
	var cycle1 cycle.Cycle

	usuario := cycle.User{"Usuario", "usuario", "123456", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), "usuario@gmail.com", 'M'}
	fmt.Println(usuario)

	streamer := streamer.Streamer{"Streamer"}
	fmt.Println(streamer)

	playlist := playlist.Playlist{1, music, cycle1}
	fmt.Println(playlist)

	fmt.Print(dbconf.GetConnectionData().Database)

}