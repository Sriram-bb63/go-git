package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	initFlag := flag.Bool("init", false, "Initialize go-git application in this directory")
	snapshotFlag := flag.Bool("snap", false, "Snapshot of the current code")
	flag.Parse()

	if *initFlag {
		InitializeTracker()
	}

	if *snapshotFlag {
		fmt.Print("Snapshot name: ")
		reader := bufio.NewReader(os.Stdin)
		snapShotName, _ := reader.ReadString('\n')
		ProcessSnapshotName(&snapShotName)
		tracker := Track()
		WriteJsonFile(&snapShotName, tracker)
		fmt.Printf("Snapshot created at ./.go-git/snapshots/%v.json\n", snapShotName)
	}
}
