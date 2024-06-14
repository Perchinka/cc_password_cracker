# Password cracker

## Overview

This repository contains a password cracker designed for a coding challenge by Jhon Crickett. The project is written in Go

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Perchinka/cc_password_cracker.git
   ```
2. Navigate to the project directory:
   ```bash
   cd cc_password_cracker
   ```
3. Install the dependencies:
   ```bash
   go mod tidy
   ```

## Usage

Run the password cracker with the following command:
```bash
go run cmd/password_cracker/main.go
```

## Project Structure

- `cmd/password_cracker`: Main application code.
- `internal/cracker`: Internal package for the cracking logic
- `go.mod`: Go module file.
