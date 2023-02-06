# GLPIAnarchy

![hakker.png](/hakerr.png)

GLPIAnarchy is a tool to scan a GLPI to find its version and quickly identify if it is vulnerable to the most critical vulnerability.
It also allows to exploit these vulnerabilities.

## Current xploit supported

- CVE-2020-5248
- CVE-2020-15175
- CVE-2021-43778
- CVE-2022-31061
- CVE-2022-35914
- CVE-2022-39323

## Upcoming xploit

- CVE-2020-15176
- CVE-2021-39211
- CVE-2023-22500

## Install

- With go install
```
go install github.com/Feals-404/GLPIAnarchy@latest
```
- From source 
```
git clone https://github.com/Feals-404/GLPIAnarchy
cd GLPIAnarchy
go build
```

## Usage

```
Usage:
  GLPIAnarchy [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     Decrypt retrieved Passwd string. The target must be below version 9.4.6
  help        Help about any command
  scan        Scans the GLPI : Recovers version and potential vulnerability
  xploit      Exploit the GLPI with a specified vulnerability

Flags:
  -h, --help   help for GLPIAnarchy

Use "GLPIAnarchy [command] --help" for more information about a command.
```

## Legal Disclaimer

Using this tool to attack a target without mutual consent is illegal. It is the responsibility of the end user to obey all applicable laws for their location. The developers assume no responsibility and are not liable for any misuse or damage caused by this program.

I used a modified version of semver (https://github.com/Masterminds/semver)
