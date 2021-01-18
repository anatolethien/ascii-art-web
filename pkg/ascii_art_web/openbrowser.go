package ascii_art_web

import "os/exec"

func OpenBrowser(url string) {

	_ = exec.Command("xdg-open", url).Start()

}
