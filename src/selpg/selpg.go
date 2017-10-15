package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
)

type selpgArgs struct {
	StartPage     int
	EndPage       int
	InputFilename string
	PageLen       int
	PageType      int32
	PrintDest     string
}

//defaultPageLen = 72
const (
	defaultPageLen = 2
)

var progName string

func main() {
	sa := selpgArgs{}

	processArgs(&sa)

	processInput(sa)
}

func processArgs(sa *selpgArgs) {
	args := os.Args
	progName = args[0]

	if len(args) < 3 {
		log.Printf("%s: not enough arguments\n", progName)
		flag.Usage()
		os.Exit(1)
	}

	var fileType bool

	flag.IntVar(&sa.StartPage, "s", 0, "specify the start page")
	flag.IntVar(&sa.EndPage, "e", 0, "specify the end page")
	flag.IntVar(&sa.PageLen, "l", defaultPageLen, "specify the length of a page")
	flag.BoolVar(&fileType, "f", false, "default and 'l' for lines-delimited, 'f' for form-feed-delimited")
	flag.StringVar(&sa.PrintDest, "d", "", "specify the name of printer")

	flag.Parse()

	if fileType {
		sa.PageType = 'f'
	} else {
		sa.PageType = 'l'
	}

	if len(flag.Args()) == 1 {
		sa.InputFilename = flag.Args()[0]
	}

}

func processInput(sa selpgArgs) {
	var fin *os.File
	if len(sa.InputFilename) == 0 {
		fin = os.Stdin
	} else {
		var err error
		fin, err = os.Open(sa.InputFilename)

		if err != nil {
			log.Printf("%s: can't open file\n", progName)
			os.Exit(2)
		}
		defer fin.Close()
	}

	bufFin := bufio.NewReader(fin)

	var fout io.WriteCloser
	cmd := &exec.Cmd{}

	if len(sa.PrintDest) == 0 {
		fout = os.Stdout
	} else {
		cmd = exec.Command("lp", "-d", sa.PrintDest)
		//cmd = exec.Command("sed", "/hello/w 1.txt")
		//cmd = exec.Command("cat")
		//cmd = exec.Command("grep", "l")

		cmd.Stdout = os.Stdout

		var err error
		fout, err = cmd.StdinPipe()

		if err != nil {
			log.Printf("%s: can't open pipe\n", progName)
			os.Exit(3)

		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
			os.Exit(4)
		}
	}

	if sa.PageType == 'l' {
		lineCtr := 0
		pageCtr := 1
		for {
			line, crc := bufFin.ReadString('\n')

			if crc != nil {
				break
			}
			lineCtr++
			if lineCtr > sa.PageLen {
				pageCtr++
				lineCtr = 1
			}

			if (pageCtr >= sa.StartPage) && (pageCtr <= sa.EndPage) {
				_, err := fout.Write([]byte(line))
				if err != nil {
					//handle err
				}
			}
		}
	} else {
		pageCtr := 1
		for {
			line, crc := bufFin.ReadString('\f')
			if crc != nil {
				break
			}
			if (pageCtr >= sa.StartPage) && (pageCtr <= sa.EndPage) {
				_, err := fout.Write([]byte(line))

				if err != nil {
					//handle err

				}
			}
			pageCtr++
		}
	}

	fout.Close()

	if err := cmd.Wait(); err != nil {
		//handle err
	}
}
