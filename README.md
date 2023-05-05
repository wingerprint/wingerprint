# Wingerprint

Windows system info and escalation paths:

- [x] Environment
  - [x] VM detection
- [x] Users
- [x] Privileges
  - [x] ​AlwaysInstallElevated
  - [ ] https://github.com/gtworek/Priv2Admin
- [x] Antivirus
- [x] Unattended files
- [ ] Vulnerable applications
- [ ] DLL hijacking
- [ ] Passwords in PowerShell history

Output is standard Markdown:

```mkd
# WINGERPRINT


## Environment

- VM detected (QUEMU ARM Virtual Machine)


## Users

- Username:           Username
  Privilege level:    Administrator
  Password age:       1h1m1s
  Last logon:         2020-01-01 01:01:01 +0000 CEST
  Logon count:        1
  Flags:              [!] Password not required | Password never expires


## Privileges

- [!] AlwaysInstallElevated


## Antiviruses

- Microsoft Defender
  - msmpeng.exe (PID 1111)


## Unattended files

- c:\windows\panther\unattend.xml
```

## Compilation

1. Install [Go](https://go.dev)
2. `go build`

### Obfuscated

To be less detectable by antiviruses, instead of `go build`:

1. `go install mvdan.cc/garble@latest`
2. `garble -literals -seed random -tiny build`
