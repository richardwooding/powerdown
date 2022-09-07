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
CLI interface to the EskomSePush API

Usage:
  powerdown [command]

Available Commands:
  allowance   Retrieves information about the allowance of your EskomSePush token
  area        Retrieve information for your area
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  search      Perform a search

Flags:
      --config string   config file (default is $HOME/.powerdown.yaml)
  -h, --help            help for powerdown
      --token string    EskomSePush token

Use "powerdown [command] --help" for more information about a command.
```

## Install with homebrew
```
brew install richardwooding/repo/powerdown
```

## Run with docker
```
docker run --volume $HOME/.powerdown.yaml:/.powerdown.yaml ghcr.io/richardwooding/powerdown
```
## Usage examples

### Check your API allowance

```
$ powerdown allowance
Using config file: /Users/richardwooding/.powerdown.yaml

Count  Limit  Type
20     50     daily
```

### Search areas by text

```
$ powerdown search text --query plett
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
Using config file: /Users/richardwooding/.powerdown.yaml

Count  Id
344    westerncape-3-oudtshoorn
116    eskde-3-oudtshoornoudtshoornwesterncape
54     westerncape-12-george
20     westerncape-7-mosselbay
8      westerncape-7-knysna
```
