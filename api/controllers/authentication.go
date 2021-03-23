package controllers

import (
	"bufio"
	"github.com/jsdaniell/whats_api/api/responses"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os/exec"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	// create new WhatsApp connection
	cmd := exec.Command("npx", "whats-cli", "connect")

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()

		png, err := qrcode.Encode(m, qrcode.Medium, 256)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
		}

		w.Write(png)
		return
	}
	cmd.Wait()
}
