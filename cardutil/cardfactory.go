package github.com/dropbot/offensive_bot/cardutil

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"errors"
)
func MakeCard(text string, template string, file string) (string, error) {

	//Replacing String in LaTeX template
	tex := strings.NewReplacer("{{text}}", text).Replace(template)
	//Debug Output
	fmt.Print(tex)

	//Constructing pdflatex command
	args := []string{"-halt-on-error"}
	dir, err := ioutil.TempDir("", "offensive_bot")

	//Checking if there are any TempDir Errors
	if err != nil {
		return nil, errors.New("Temporary Directory Failure")
	}
	topdf := exec.Cmd{Path: "pdflatex", Args: args, Dir: dir}
	//Redirecting StdinPipe
	std, err := topdf.StdinPipe()
	//Writing Tex to pdflatex
	fmt.Fprint(std, tex)
	//Checking if there are any Errors with the StdinPipe
	if err != nil {
		return nil, errors.New("StdinPipe Error (could be bufferoverflow)")
	}
	//Starting the pdflatex command
	topdf.Start()
	//Waiting for it to finish
	e := topdf.Wait()
	//Check if it executed without errors
	if e != nil {
		return nil, errors.New("pdflatex failure")
	}
	//Building convert command
	//For converting the pdf to png
	topng := exec.Command("convert", "texput.pdf", file+".png")

	//Starting the command
	topng.Start()
	//Wait for it to finish
	e = topng.Wait()
	//Check execution errors
	if e != nil {
		return nil, errors.New("convert failure")
	}
	//Success
	return (dir+file+".png"),nil
}
