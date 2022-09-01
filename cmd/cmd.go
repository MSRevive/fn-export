package cmd

import (
	"fmt"
	"net"
	"flag"
	"os"
	"strings"
)

type flags struct {
	address string
	port string
	password string
	steamid string
	slot string
}

func doFlags(args []string) *flags {
	flgs := &flags{}

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.StringVar(&flgs.address, "addr", "127.0.0.1", "The address of the server.")
	flagSet.StringVar(&flgs.port, "port", "1337", "The port the server is on.")
	flagSet.StringVar(&flgs.password, "password", "", "The password for the FN server.")
	flagSet.StringVar(&flgs.steamid, "steamid", "STEAM_0:0_1838", "SteamID of the character.")
	flagSet.StringVar(&flgs.slot, "slot", "1", "The slot the character.")
	flagSet.Parse(args[1:])

	return flgs
}

func Run(args []string) error {
	flgs := doFlags(args)
	filename := fmt.Sprintf("./chars/%s_%s.char", strings.Replace(flgs.steamid, ":", "_", -1), flgs.slot)

	fmt.Printf("Attempting to open TCP socket to %s:%s...\n", flgs.address, flgs.port)
	conn, err := net.Dial("tcp", flgs.address+":"+flgs.port)
	if err != nil {
		fmt.Println("Failed to connect to server")
		return err
	}
	defer conn.Close()

	send := fmt.Sprintf("retr %s %s")
	_, err = conn.Write([]byte(send))
	if err != nil {
		fmt.Println("Failed to write data to via server")
		return err
	}
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read data from via server")
		return err
	}

	fmt.Printf("Received %i bytes, writing bytes to file %s\n", len, filename)
	if err := os.WriteFile(filename, buffer, 0755); err != nil {
		return err
	}

	return nil
}