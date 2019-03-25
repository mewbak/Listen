package gui

import (
	"listen/logic"

	"github.com/gotk3/gotk3/gtk"
)

// InitWidgets calls and initalizes the starting values of all widge.ts
func InitWidgets(window Elements) Elements {
	window = setAppWindowValues(window)
	window = setHeaderValues(window)
	window = setTrackBoxValues(window)
	//	window = SetPopMenuValues(window)
	return window
}

// setAppWindowValues initializes all widget values that belong to GUI.AppWindow.
func setAppWindowValues(w Elements) Elements {
	// AppWindow
	w.AppWindow.Add(w.TrackBox)
	w.AppWindow.SetTitlebar(w.Header)
	w.AppWindow.SetDefaultSize(logic.ArtSize, logic.ArtSize)
	w.AppWindow.SetResizable(false) // may consider changing this one day
	// 								   don't think this application is ever
	// 								   supposed to go full screen tho
	w.AppWindow.Show()
	return w
}

// setHeaderValues initalizes all widget values that belong to GUI.Header.
func setHeaderValues(w Elements) Elements {
	// Buttons
	//w.ImgMenu.SetFromIconName("open-menu-symbolic", gtk.ICON_SIZE_BUTTON)
	w.ImgFile.SetFromIconName("document-open-symbolic", gtk.ICON_SIZE_BUTTON)
	w.ImgPlay.SetFromIconName("media-playback-start-symbolic",
		gtk.ICON_SIZE_BUTTON)
	//w.MenuButt.SetImage(w.ImgMenu)
	w.PlayButt.SetImage(w.ImgPlay)
	w.PlayButt.SetSensitive(false)
	w.FileButt.SetImage(w.ImgFile)
	w.VolButt.SetValue(1)
	w.VolButt.SetSensitive(false)
	//w.MenuButt.SetPopover(w.PopMenu)

	// ProgScale
	w.ProgScale.SetDrawValue(false)
	w.ProgScale.SetHExpand(true)

	// Header
	// *Remember: Order Matters!*
	w.Header.SetCustomTitle(w.ProgScale)
	w.Header.PackStart(w.PlayButt)
	//w.Header.PackEnd(w.MenuButt)
	w.Header.PackEnd(w.FileButt)
	w.Header.PackEnd(w.VolButt)
	w.Header.SetShowCloseButton(true)
	w.Header.ShowAll()

	return w
}

// setTrackBoxValues initializes all widget values that belong to GUI.TrackBox.
func setTrackBoxValues(w Elements) Elements {
	// TrackBox
	w.TrackBox.Add(w.ImgTrack)
	w.ImgTrack.SetHExpand(true)
	w.ImgTrack.SetVExpand(true)
	w.ImgTrack.SetHAlign(gtk.ALIGN_CENTER)
	w.TrackBox.ShowAll()

	return w
}

/*// setPopMenuValues initializes all widget values that belong to GUI.PopMenu.
func setPopMenuValues(w gui.GUI) gui.GUI {

	return w
} */
