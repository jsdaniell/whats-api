package controllers

import (
	"bufio"
	"github.com/gorilla/mux"
	"github.com/jsdaniell/whats_api/api/responses"
	"github.com/jsdaniell/whats_api/api/utils/shell_commands"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os/exec"
)

func Connect(w http.ResponseWriter, r *http.Request) {

	sessionID := mux.Vars(r)["sessionID"]

	// create new WhatsApp connection
	cmd := exec.Command("./whats-cli", "connect", sessionID)

	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

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

func Disconnect(w http.ResponseWriter, r *http.Request) {

	sessionID := mux.Vars(r)["sessionID"]

	// create new WhatsApp connection
	err := shell_commands.ExecuteShellCommand("./whats-cli", "version")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}


	err = shell_commands.ExecuteShellCommand("./whats-cli", "disconnect", sessionID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte("disconnected"))
	return
}
