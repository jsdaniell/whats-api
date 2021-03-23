package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jsdaniell/whats_api/api/models"
	"github.com/jsdaniell/whats_api/api/responses"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func Message(w http.ResponseWriter, r *http.Request) {
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

	cmd := exec.Command("npx", "whats-cli", "send", messageBody.Number, messageBody.Message)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()

		fmt.Println(m)

		w.Write([]byte(m))
		return
	}
	cmd.Wait()
}
