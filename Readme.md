## Overview

This project implements a blockchain in Go. It includes core blockchain functionalities, a CLI for interaction, networking for peer-to-peer communication, a React.js client, wallet management, and tests.

## Directory Structure

- **blockchain**: Core blockchain logic including block creation, validation, and chain management.
- **cli**: Command Line Interface for interacting with the blockchain.
- **network**: Peer-to-peer communication logic.
- **reactjs-client**: React.js front-end client for web interaction.
- **test**: Test cases for various components.
- **wallet**: Cryptographic operations for managing wallets and transactions.

## Getting Started

### Prerequisites

- Go (1.18.4)
- Node.js and npm (for React.js client)

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/AminMortezaie/golang-blockchain.git
   cd golang-blockchain
   ```

2. **Build the Application**
   ```bash
   go build -o main.exe main.go
   ```

3. **Run the Application**
   ```bash
   ./main.exe
   ```

4. **Running the React Client**
   Navigate to the `reactjs-client` directory and run:
   ```bash
   npm install
   npm start
   ```

### Configuration

- **main.go**: Entry point for the Go application. Initializes the blockchain and starts the CLI and network.
- **blockchain/blockchain.go**: Core blockchain functionalities.
- **cli/cli.go**: CLI commands.
- **network/network.go**: Networking functionalities.
- **wallet/wallet.go**: Wallet functionalities.

### Usage

- **CLI Commands**: Use the CLI to interact with the blockchain. Examples:
  ```bash
  ./main.exe createblockchain -address ADDRESS
  ./main.exe addblock -data "Some data"
  ./main.exe printchain
  ```

- **React Client**: Access the React.js client at `http://localhost:3000` after starting the React server.

### Testing

Run tests using:
```bash
go test ./test/...
```
