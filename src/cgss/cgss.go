package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cgss/cg"
	"cgss/ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func Help(args []string) int {
	fmt.Println(`
        Commands:
            login <username><level><exp>
            logout <username>
            listplayer
            quit(q)
            help(h)
        `)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args []string) int {

}
