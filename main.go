package main

import (
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
	
	"github.com/joho/godotenv"
	"github.com/fatih/color"
)

func main() {

	err := godotenv.Load("config.ini")
	if err != nil {
		color.Red("Error loading config.ini file")
		fmt.Println()
		return
	}

	serverIp := os.Getenv("SERVER_IP")
	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		color.Red("Invalid server port")
		fmt.Println()
		return
	}

	rconPassword := os.Getenv("RCON_PASSWORD")
	waitInterval := os.Getenv("WAIT_INTERVAL")
	waitDuration, err := time.ParseDuration(waitInterval + "s")
	if err != nil {
		color.Red("Invalid wait interval")
		fmt.Println()
		return
	}


	year := color.YellowString(fmt.Sprintf("%d", time.Now().Year()))
	fmt.Printf("\n%s - %s\n\n", color.RedString("- ANTI VPN BY DraGoN"), year)

	for {
		query := "\xFF\xFF\xFF\xFF\x02rcon " + rconPassword + " status \n"
	
		socket, err := net.Dial("udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
		if err != nil {
			color.Red("Error connecting to server! Retrying in " + waitDuration.String() + " seconds...")
			fmt.Println()
			time.Sleep(waitDuration)
			continue
		}
		
		socket.Write([]byte(query))

		socket.SetReadDeadline(time.Now().Add(3 * time.Second))
		data := make([]byte, 1500)
		_, err = socket.Read(data)

		if err != nil {
	        color.Red("Server is offline! Retrying in " + waitDuration.String() + " seconds...")
	        fmt.Println()
	        time.Sleep(waitDuration)
	        continue
		}

		dataReturned := string(data)
		regexIpAddress := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?:\/\d{2})?`)
		ipAddresses := regexIpAddress.FindAllString(dataReturned, -1)

		socket.Close()

		for _, ipAddress := range ipAddresses {
			fmt.Printf("%s ", color.GreenString(ipAddress))
			fmt.Printf("- %s: ", color.MagentaString("Using VPN/Proxy"))
			detectVPN(ipAddress, serverIp, serverPort, rconPassword)
		}

		if len(ipAddresses) == 0 {
			color.Red("- No Players!")
			fmt.Println()
		}

		time.Sleep(waitDuration)
	}
}

func detectVPN(ipAddress string, serverIp string, serverPort int, rconPassword string) {
	client := &http.Client{}
	url := fmt.Sprintf("https://blackbox.ipinfo.app/lookup/%s", ipAddress)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		color.Red("Could not get VPN information.")
		fmt.Println()
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if strings.TrimSpace(string(body)) == "Y" {
		color.Yellow("YES!!")
		fmt.Println()

		socket, _ := net.Dial("udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
		banCommand := fmt.Sprintf("\xFF\xFF\xFF\xFF\x02rcon %s banipr %s Using VPN\n", rconPassword, ipAddress)
		socket.Write([]byte(banCommand))

		socket.SetReadDeadline(time.Now().Add(3 * time.Second))
		data := make([]byte, 2048)
		_, err := socket.Read(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		response := string(data)
		response = strings.Replace(response, "....print\n", "", -1)
		fmt.Println(response)
		socket.Close()
	} else if strings.TrimSpace(string(body)) == "N" {
		color.Red("No")
		fmt.Println()
	}
}