package audio

import (
	"listen/util"
//	"github.com/dhowden/tag"
	"github.com/gotk3/gotk3/gdk"
	"listen/gui"
	"os"

) 

func findArt() *gdk.Pixbuf {

	// process picture
	mus, meta := openMusic()	
	defer mus.Close()

	pic := meta.Picture()

	println((*pic).String())

	f, err := os.Create("./.temp_cover")
	if err != nil {
		util.SendError(err, "creating file")
	}
	defer f.Close()
	defer os.Remove("./.temp_cover")


	_, err = f.Write((*pic).Data)
	if err != nil {
		util.SendError(err, "writing to cover")
	}

	pix, err := gdk.PixbufNewFromFileAtScale("./.temp_cover",
	gui.ArtSize,
	gui.ArtSize,
	true)
	
	if err != nil {
		util.SendError(err, "setting album art pixbuf")
	}

	return pix
}