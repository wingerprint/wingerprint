# Wingerprint

Windows system info and escalation paths:

- [x] Users
- [x] Privilieges
  - [x] ​AlwaysInstallElevated
- [x] Antivirus
- [x] Unattended files
- [x] VM detection
- [ ] Vulnerable applications
- [ ] DLL hijacking
- [ ] Passwords in PowerShell history

Output is standard Markdown:

```mkd
# WINGERPRINT

- VM environment


## Users

- COMPUTER\Username
  S-1-1-11-1111111111-111111111-1111111111-1111
  ⚠️ No password


## Privileges

- ⚠️ AlwaysInstallElevated


## Antiviruses

- Microsoft Defender
  - msmpeng.exe (PID 1111)


## Unattended files

- c:\windows\panther\unattend.xml
```