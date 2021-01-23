package ascii_art_web

import "os/exec"

func OpenBrowser() {

	exec.Command("xdg-open", Url).Start()

}
