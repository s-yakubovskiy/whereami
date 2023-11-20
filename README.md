# WhereAmI

## Overview
WhereAmI is a command-line application that allows users to find their geolocation based on their public IP address. It's designed to provide quick and easy access to geolocation information, including country, region, city, and more.

## Features
- Retrieve geolocation information based on the public IP address.
- Display location details such as country, region, city, ISP, and more.
- Easy-to-use command-line interface.
- Option to store location history in a local SQLite database.

## Prerequisites
Before you begin, ensure you have met the following requirements:
- Go (Golang) installed on your machine.
- Basic understanding of command-line operations.

## Installation
To install WhereAmI, follow these steps:

```bash
git clone https://github.com/[your-username]/whereami.git
cd whereami
go build -o bin/whereami
```

## Usage
After installation, you can run WhereAmI using the following command:

```bash
./bin/whereami
```

This command will display your current geolocation information.

### Commands
- `show`: Display current public IP address and fetch location information.

  Usage:
  ```bash
  ./bin/whereami show
  ```

- `store`: Store information in sqlite database (primitive checks to avoid store same records)

  Usage:
  ```bash
  ./bin/whereami show
  ```

