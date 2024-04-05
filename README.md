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
Initialize `whereami` with `whereami init` (applies migrations and configs).  

After installation, you can run WhereAmI using the following command:

```bash
./bin/whereami
```

This command will display your current geolocation information.

There are flags available to modify your query:
```
Flags:                                                                                                                                                                                
    -f, --full                   Display full output                                                                                                                                    
    -h, --help                   help for whereami                                                                                                                                      
        --ip string              Specify public IP to lookup info                                                                                                                       
    -l, --location-api string    Select ip location provider: [ipapi, ipdata]                                                                                                           
    -p, --public-ip-api string   Select public ip api provider: [ifconfig.me, ipinfo.io/ip, icanhazip.com]                                                                                 
    -v, --version                Display application version 
```

E.g. `whereami -l ipdata -p icanhazip.com` or `whereami --p 100.200.200.100` to query for particular IP 

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
dev ❯ whereami --location-api ipapi 
             _                          _ 
 __      __ | |__    _ __   _ __ ___   (_)
 \ \ /\ / / | '_ \  | '__| | '_ ` _ \  | |
  \ V  V /  | | | | | |    | | | | | | | |
   \_/\_/   |_| |_| |_|    |_| |_| |_| |_|
                                          
    ... getting your location data ...

Network Information
  ip: 38.54.75.103
  isp: Kaopu Cloud HK Limited
  flag: 🇦🇪

Geographical Information
  country: United Arab Emirates
  countryCode: AE
  region: Dubai
  regionCode: DU
  city: Dubai
  timezone: Asia/Dubai
  latitude: 25.2048
  longitude: 55.2708

Security Assessments
  vpn: false
  FraudScore: 75
  IsCrawler: false
  Host: 38.54.75.103
  Proxy: true
  VPN: true
  Tor: false
  RecentAbuse: false
  BotStatus: false

Miscellaneous
  map: https://www.google.com/maps?q=25.204800,55.270800
  date: Friday, April 5, 2024, 09:15
  comment: Fetched with ipapi provider. Updated with ipqualityscore provider. Using public ip provider: http://ifconfig.me
```
