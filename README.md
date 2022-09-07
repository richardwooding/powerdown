# powerdown 
CLI interface to the [EskomSePush](https://sepush.co.za/) [API](https://documenter.getpostman.com/view/1296288/UzQuNk3E)
## API Token
You need to provide your own API Token which you can [request here](https://docs.google.com/forms/d/e/1FAIpQLSeZhAkhDaQX_mLT2xn41TkVjLkOH3Py3YWHi_UqQP4niOY01g/viewform).

powerdown uses [viper](https://github.com/spf13/viper) for configuration and by default will load configuration
from `$HOME/.powerdown.yaml`

The contents of the file should be as follows
```
token: <your-api-token>
```
If you know your areas id you can use
```
token: <your-api-token>
id: <your-area-id>
```

Alternatively you can use the `--token string` and the `--id string` flag 
## Usage
```
$ powerdown
Powerdown 1.0.1

CLI interface to the EskomSePush API

Usage:
  powerdown [command]

Available Commands:
  allowance     Retrieves information about the allowance of your EskomSePush token
  area          Retrieve information for your area
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  is-event-soon Returns zero exit code if loadsheeding event is soon, non-zero if not or could not determine due to error
  search        Perform a search

Flags:
      --config string   config file (default is $HOME/.powerdown.yaml)
  -h, --help            help for powerdown
      --token string    EskomSePush token

Use "powerdown [command] --help" for more information about a command.
```

## Install/upgrade with homebrew on Mac OS X

### To install
```
brew install richardwooding/repo/powerdown
```
### To upgrade
```
brew upgrade powerdown
```

## Run with docker
```
docker run --volume $HOME/.powerdown.yaml:/.powerdown.yaml ghcr.io/richardwooding/powerdown
```
## Usage examples

### Check if a loadshedding event is imminent

How to interpret exit code
- 0 -> Shutting down is suggested
- Non-0 -> Do not shutdown

```
powerdown is-event-soon --id westerncape-7-knysna --suggest-shutdown-time 1h
Powerdown 1.0.1

Using config file: /Users/richardwooding/.powerdown.yaml
Retrieving information for area: capetown-7-oranjezicht

Stage 2 will start in 34h59m3.581748s at 2022-09-08 18:00:00 +0200 SAST

Do not recommend shutting down
```

### Check your API allowance

```
$ powerdown allowance
Powerdown 1.0.1

Using config file: /Users/richardwooding/.powerdown.yaml

Count  Limit  Type
11     50     daily
```

### Search areas by text

```
$ powerdown search text plett
Powerdown 1.0.1

Using config file: /Users/richardwooding/.powerdown.yaml
Search area matching: plett

Id                                              Name                          Region
eskme-7-plettenbergbaybitouwesterncape          Plettenberg Bay (7)           Eskom Municipal, Bitou, Western Cape
westerncape-7-plettenbergbay                    Plettenberg Bay (7)           Western Cape
eskde-7-plettenbergbaybitouwesterncape          Plettenberg Bay (7)           Eskom Direct, Bitou, Western Cape
eskde-7-plettenbergbayoutlyingbitouwesterncape  Plettenberg Bay Outlying (7)  Eskom Direct, Bitou, Western Cape
```
### Search area by latitude and longitude

```
$ powerdown search nearby --lat -33.6007 --lon 22.2026
Powerdown 1.0.1

Using config file: /Users/richardwooding/.powerdown.yaml

Count  Id
464    westerncape-3-oudtshoorn
146    eskde-3-oudtshoornoudtshoornwesterncape
76     westerncape-12-george
32     westerncape-7-mosselbay
7      westerncape-7-knysna
```
### Retrieve loadshedding information for a particular area

```
powerdown area --id westerncape-7-knysna
Powerdown 1.0.1

Using config file: /Users/richardwooding/.powerdown.yaml
Retrieving information for area: westerncape-7-knysna

Name        Region
Knysna (7)  Western Cape

Events
Note     Start                           End
Stage 2  2022-09-07 18:00:00 +0200 SAST  2022-09-07 20:30:00 +0200 SAST
Stage 2  2022-09-09 08:00:00 +0200 SAST  2022-09-09 10:30:00 +0200 SAST

Schedule via  http://www.stellenbosch.gov.za/news/notices/notices-engineering/6330-loadshedding-stage-1-4-schedule-2018/file
Date        Day of week  Stages
2022-09-07  Wednesday
                         18:00-20:30
                         10:00-12:30 18:00-20:30
                         02:00-04:30 10:00-12:30
                         02:00-04:30 10:00-12:30
                         02:00-04:30 10:00-12:30
                         02:00-04:30 10:00-14:30
                         02:00-06:30 10:00-14:30
2022-09-08  Thursday     02:00-04:30
                         02:00-04:30
                         02:00-04:30 18:00-20:30
                         02:00-04:30 10:00-12:30
                         02:00-06:30 10:00-12:30
                         02:00-06:30 10:00-12:30
                         02:00-06:30 10:00-12:30
                         02:00-06:30 10:00-14:30
2022-09-09  Friday       08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-12:30
                         00:00-04:30 08:00-12:30
                         00:00-04:30 08:00-12:30
                         00:00-04:30 08:00-12:30
2022-09-10  Saturday     16:00-18:30
                         08:00-10:30 16:00-18:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-12:30
                         00:00-04:30 08:00-12:30
                         00:00-04:30 08:00-12:30
2022-09-11  Sunday
                         16:00-18:30
                         08:00-10:30 16:00-18:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-10:30
                         00:00-02:30 08:00-12:30
                         00:00-04:30 08:00-12:30
2022-09-12  Monday       00:00-02:30
                         00:00-02:30
                         00:00-02:30 16:00-18:30
                         00:00-02:30 08:00-10:30
                         00:00-04:30 08:00-10:30
                         00:00-04:30 08:00-10:30
                         00:00-04:30 08:00-10:30
                         00:00-04:30 08:00-12:30
2022-09-13  Tuesday      06:00-08:30
                         06:00-08:30
                         06:00-08:30 22:00-00:30
                         06:00-08:30 14:00-16:30
                         06:00-10:30 14:00-16:30
                         06:00-10:30 14:00-16:30
                         06:00-10:30 14:00-16:30
                         06:00-10:30 14:00-18:30
```