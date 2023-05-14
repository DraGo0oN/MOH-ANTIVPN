# MOHAA Anti-VPN

This is a simple implementation of an anti-VPN solution for the game "Medal of Honor: Allied Assault" (MOHAA) using Golang.

## How it works

The anti-VPN solution works by checking the player's IP address against a database API of known VPN IP addresses. If the player's IP address is found in the database, they are banned from the server.

## Installation

To use this anti-VPN solution on your MOHAA server, follow these steps:

1. Install Golang on your server. You can download it from the official website: https://golang.org/

2. Clone this repository:

   ````
   git clone https://github.com/DraGo0oN/MOH-ANTIVPN.git
   

3. Install the required dependencies:

   ````
   go mod download
   
4. Edit the `config.ini` file to include the following information:

   ````
   SERVER_IP=127.0.0.1     # MOHAA server IP
   SERVER_PORT=12203         # Server Port
   RCON_PASSWORD=123123  # RCON password
   WAIT_INTERVAL=5     # The number of seconds to wait before scanning for players. The default is 10.

Replace the values with your actual MOHAA server IP address, port number, RCON password and the number of seconds to wait before banning a player. The `WAIT_INTERVAL` field determines how often the program will scan for players to check their IP addresses against the database API of known VPN IP addresses. 

   Note: Make sure to save the changes to the `config.ini` file before running the program.


5. Run the program:

   ````
   go run main.go
   
If you want to compile Run:
   ````
   go build
````
6. Connect to your MOHAA server and test the anti-VPN solution by connecting through a VPN. You should be banned from the server.

## Configuration

The anti-VPN solution can be configured by editing the `config.ini` file. The following fields can be modified:

- `SERVER_IP`: The IP address of your MOHAA server.
- `SERVER_PORT`: The port number of your MOHAA server.
- `RCON_PASSWORD`: The RCON password for your MOHAA server.
- `WAIT_INTERVAL`: The number of seconds to wait before banning a player. The default is 5.

## Contributing

This project is open source and contributions are welcome! If you have any suggestions or improvements, please create a pull request or open an issue on Github.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
