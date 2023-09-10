# MOHAA Anti-VPN

Welcome to the MOHAA Anti-VPN project, a simple yet effective solution designed to prevent VPN users from accessing your "Medal of Honor: Allied Assault" (MOHAA) game server, implemented in Golang.

## How It Works

Our anti-VPN solution operates by cross-referencing players' IP addresses with a database API containing known VPN IP addresses. If a player's IP address matches an entry in the database, they are promptly banned from your MOHAA server.

## Installation

To implement this anti-VPN solution on your MOHAA server, follow these straightforward steps:

1. **Install Golang**: Ensure you have Golang installed on your server. If not, you can download it from the official website: [Golang Official Website](https://golang.org/).

2. **Clone the Repository**:
   
   ```shell
   git clone https://github.com/ysdragon/MOH-ANTIVPN.git
   ```

3. **Install Dependencies**:
   
   ```shell
   go mod download
   ```

4. **Configure `config.ini`**:
   
   Edit the `config.ini` file to provide the following details:

   ```ini
   SERVER_IP=127.0.0.1     # MOHAA server IP
   SERVER_PORT=12203       # Server Port
   RCON_PASSWORD=123123    # RCON password
   WAIT_INTERVAL=10         # Seconds to wait before scanning for players (default is 10).
   ```

   Replace these values with your actual MOHAA server IP address, port number, RCON password, and the desired interval (in seconds) for player IP checks. The `WAIT_INTERVAL` determines how frequently the program scans for players and checks their IP addresses against the VPN database.

   **Note**: Ensure you save the changes to the `config.ini` file before running the program.

5. **Run the Program**:
   
   To run the program, execute the following command:

   ```shell
   go run main.go
   ```

   If you prefer to compile it, use:

   ```shell
   go build
   ```

6. **Test Your Setup**:
   
   Connect to your MOHAA server and test the anti-VPN solution by connecting through a VPN. You should observe the immediate banning of VPN users from your server.

## Configuration

Customize the anti-VPN solution to your preferences by modifying the `config.ini` file. Here are the configurable fields:

- `SERVER_IP`: The IP address of your MOHAA server.
- `SERVER_PORT`: The port number of your MOHAA server.
- `RCON_PASSWORD`: The RCON password for your MOHAA server.
- `WAIT_INTERVAL`: The interval (in seconds) before banning a player (default is 5).

## Contributing

This project is open source, and we encourage contributions. If you have any suggestions, improvements, or bug fixes, please create a pull request or open an issue on [Github](https://github.com/ysdragon/MOH-ANTIVPN).

## License

The MOHAA Anti-VPN project is licensed under the MIT License. For detailed licensing information, please refer to the [LICENSE](LICENSE) file.