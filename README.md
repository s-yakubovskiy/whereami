# WhereAmI

## Overview
WhereAmI is a command-line application that allows users to find their geolocation based on their public IP address. It's designed to provide quick and easy access to geolocation information, including country, region, city, and more.

## Features
- Retrieve geolocation information based on the public IP address.
- Display location details such as country, region, city, ISP, and more.
- Added Fraud Scores custom provider support (vpn, tor, fraud score, bot, etc)
- Easy-to-use command-line interface.
- Option to store location history in a local SQLite database.

## Prerequisites
Before you begin, ensure you have met the following requirements:
- Go (Golang) installed on your machine.
- Basic understanding of command-line operations.

## Installation
To install WhereAmI, follow these steps:

```bash
git clone https://github.com/s-yakubovskiy/whereami.git
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
  ./bin/whereami store
  ```

```bash
dev ❯ whereami show --provider ipapi --full
             _                          _ 
 __      __ | |__    _ __   _ __ ___   (_)
 \ \ /\ / / | '_ \  | '__| | '_ ` _ \  | |
  \ V  V /  | | | | | |    | | | | | | | |
   \_/\_/   |_| |_| |_|    |_| |_| |_| |_|
                                          

Network Information
  ip: 139.28.220.186
  isp: I-servers LTD
  flag: 🇫🇮

Geographical Information
  country: Finland
  countryCode: FI
  region: Uusimaa
  regionCode: 18
  city: Helsinki
  timezone: Europe/Helsinki
  zip: 00131
  latitude: 60.1797
  longitude: 24.9344

Security Assessments
  vpn: true
  FraudScore: 100
  IsCrawler: false
  Host: 139.28.220.186
  Proxy: true
  VPN: true
  Tor: true
  RecentAbuse: false
  BotStatus: true

Miscellaneous
  date: Saturday, March 23, 2024, 15:06
  comment: Fetched with ipapi provider. Updated with ipqualityscore provider
  map: https://www.google.com/maps?q=60.179700,24.934400

```
