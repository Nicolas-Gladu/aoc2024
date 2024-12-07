package pkg

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type DayChallenge struct {
	Path   string
	Inputs []string
}

func NewDayChallenge(path string) (*DayChallenge, error) {
	dayChallenge := DayChallenge{
		Path: path,
	}

	err := dayChallenge.getDataInput()
	if err != nil {
		return nil, err
	}

	return &dayChallenge, nil
}

func (d *DayChallenge) getDataInput() error {
	file, err := os.Open(d.Path)
	if errors.Is(err, os.ErrNotExist) {
		// handle the case where the file doesn't exist
		log.Println(d.Path)
		return errors.New("input file from day %d doesn't exist")
	}
	if err != nil {
		return err
	}

	err = d.readLines(file)
	if err != nil {
		return err
	}

	return nil
}

func (d *DayChallenge) readLines(f *os.File) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d.Inputs = append(d.Inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {

		return err
	}

	return nil
}
