package controllers

import (
	"bufio"
	"errors"
	"github.com/jsdaniell/whats_api/api/responses"
	"github.com/jsdaniell/whats_api/api/utils/shell_commands"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os/exec"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0{
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing 'Authorization' on header"))
		return
	}

	// create new WhatsApp connection
	cmd := exec.Command("./whats-cli", "connect", auth)

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

		_, err = w.Write(png)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
		}

		return
	}
	err = cmd.Wait()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0{
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing 'Authorization' on header"))
	}

	err := shell_commands.ExecuteShellCommand("./whats-cli", "version")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}


	err = shell_commands.ExecuteShellCommand("./whats-cli", "disconnect", auth)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	_, err = w.Write([]byte("disconnected"))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
}
