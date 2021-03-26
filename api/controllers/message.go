package controllers

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/jsdaniell/whats_api/api/models"
	"github.com/jsdaniell/whats_api/api/responses"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func Message(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if len(auth) == 0{
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing 'Authorization' on header"))
	}

	var messageBody models.MessageModel

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &messageBody)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cmd := exec.Command( "./whats-cli", "send", messageBody.Number, messageBody.Message, auth)

	stdout, _ := cmd.StdoutPipe()
	err = cmd.Start()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()

		_, err = w.Write([]byte(m))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
		}
		return
	}
	err = cmd.Wait()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
}
