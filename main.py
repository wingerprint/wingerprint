#!/usr/bin/env python

import os
import wmi as wmilib


print("# WINGERPRINT\n")

windir = os.path.normcase(os.environ["WINDIR"])
wmi = wmilib.WMI()


# VM environment

try:
    vm_manufacturers = ["oracle", "parallels", "virtualbox", "vmware"]
    bios_manufacturer = wmi.Win32_BIOS()[0].Manufacturer.lower()
    if any([x in bios_manufacturer for x in vm_manufacturers]):
        print("- VM environment")
except:
    pass


# Users

try:
    print("\n\n## Users\n")
    for user in wmi.Win32_UserAccount():
        if not user.Disabled:
            print(f"- {os.path.normpath(user.Caption)}")
            print(f"  {user.SID}")
            if not user.PasswordRequired:
                print("  ⚠️ No password")
            elif user.PasswordChangeable:
                print("  ⚠️ Password is changeable")
except:
    pass

# AV

av_processes = {
    "adawareservice.exe":          "Adaware",
    "advchk.exe":                  "Norton",
    "afwserv.exe":                 "Avast",
    "ahnsd.exe":                   "AhnLab",
    "alertsvc.exe":                "Norton",
    "alunotify.exe":               "Norton",
    "arestore.exe":                "Norton",
    "ashwebsv.exe":                "Avast",
    "asoelnch.exe":                "Norton",
    "aswengsrv.exe":               "Avast",
    "aswidsagent.exe":             "Avast",
    "aswtoolssvc.exe":             "Avast",
    "aswupdsv.exe":                "Avast",
    "avastbrowser.exe":            "Avast",
    "avastnm.exe":                 "Avast",
    "avastsvc.exe":                "Avast",
    "avastui.exe":                 "Avast",
    "avgsvc.exe":                  "AVG",
    "avguard.exe":                 "Avira",
    "avmaisrv.exe":                "Avast",
    "avp.exe":                     "Kaspersky",
    "avpcc.exe":                   "Kaspersky",
    "avpm.exe":                    "Kaspersky",
    "avsched32.exe":               "HBEDV",
    "avwupsrv.exe":                "HBEDV",
    "bdagent.exe":                 "Bitdefender",
    "bdmcon.exe":                  "Bitdefender",
    "bdnagent.exe":                "Bitdefender",
    "bdoesrv.exe":                 "Bitdefender",
    "bdss.exe":                    "Bitdefender",
    "bdswitch.exe":                "Bitdefender",
    "bitdefender_p2p_startup.exe": "Bitdefender",
    "bullguardcore.exe":           "BullGuard",
    "buvss.exe":                   "Norton",
    "cavrid.exe":                  "ETrustEZ",
    "cavtray.exe":                 "ETrustEZ",
    "cltlmh.exe":                  "Norton",
    "cltrt.exe":                   "Norton",
    "cmgrdian.exe":                "McAfee",
    "coinst.exe":                  "Norton",
    "conathst.exe":                "Norton",
    "cpysnpt.exe":                 "Norton",
    "cyneteps.exe":                "Cynet",
    "cynetms.exe":                 "Cynet",
    "doscan.exe":                  "Norton",
    "dvpapi.exe":                  "Cyren",
    "efainst64.exe":               "Norton",
    "ekrn.exe":                    "ESET",
    "elaminst.exe":                "Norton",
    "enterpriseservice.exe":       "VIPRE",
    "fldghost.exe":                "Norton",
    "frameworkservic.exe":         "McAfee",
    "frameworkservice.exe":        "McAfee",
    "freshclam.exe":               "ClamWin",
    "fshoster32.exe":              "F-Secure",
    "gdscan.exe":                  "G Data",
    "icepack.exe":                 "Norton",
    "instca.exe":                  "Norton",
    "k7crvsvc.exe":                "K7",
    "kavfs.exe":                   "Kaspersky",
    "kavfsgt.exe":                 "Kaspersky",
    "kavfsmui.exe":                "Kaspersky",
    "kavfsrcn.exe":                "Kaspersky",
    "kavfsscs.exe":                "Kaspersky",
    "kavfswh.exe":                 "Kaspersky",
    "kavfswp.exe":                 "Kaspersky",
    "kavshell.exe":                "Kaspersky",
    "kavtray.exe":                 "Kaspersky",
    "mbam.exe":                    "Malwarebytes",
    "mcapexe.exe":                 "McAfee",
    "mcui32.exe":                  "Norton",
    "mfemms.exe":                  "McAfee",
    "mgavrtcl.exe":                "McAfee",
    "mghtml.exe":                  "McAfee",
    "mgui.exe":                    "BullGuard",
    "msmpeng.exe":                 "Microsoft Defender",
    "navapsvc.exe":                "Norton",
    "navapw32.exe":                "Norton",
    "navw32.exe":                  "Norton",
    "ncolow.exe":                  "Norton",
    "nod32krn.exe":                "ESET",
    "nod32kui.exe":                "ESET",
    "nortonsecurity.exe":          "Norton",
    "npfmntor.exe":                "Norton",
    "nsc.exe":                     "Norton",
    "nsmdtr.exe":                  "Norton",
    "nswscsvc.exe":                "Norton",
    "ntrtscan.exe":                "Trend Micro",
    "nuperfscan.exe":              "Norton",
    "ofcdog.exe":                  "Trend Micro",
    "patch.exe":                   "Trend Micro",
    "pavfires.exe":                "Panda",
    "pavfnsvr.exe":                "Panda",
    "pavkre.exe":                  "Panda",
    "pavmail.exe":                 "Panda",
    "pavprot.exe":                 "Panda",
    "pavprsrv.exe":                "Panda",
    "pavsched.exe":                "Panda",
    "pavsrv50.exe":                "Panda",
    "pavsrv51.exe":                "Panda",
    "pavsrv52.exe":                "Panda",
    "pavupg.exe":                  "Panda",
    "pcscan.exe":                  "Trend Micro",
    "pntiomon.exe":                "Trend Micro",
    "pop3pack.exe":                "Trend Micro",
    "pop3trap.exe":                "Trend Micro",
    "poproxy.exe":                 "Norton",
    "prevsrv.exe":                 "Panda",
    "realmon.exe":                 "ETrustEZ",
    "ruleup.exe":                  "Norton",
    "savscan.exe":                 "Norton",
    "savservice.exe":              "Sophos",
    "sbserv.exe":                  "Norton",
    "scan32.exe":                  "McAfee",
    "sefinst.exe":                 "Norton",
    "sevntx64.exe":                "Norton",
    "spider.exe":                  "DrWeb",
    "srtsp_ca.exe":                "Norton",
    "symdgnhc.exe":                "Norton",
    "symerr.exe":                  "Norton",
    "symvtcatalogdb.exe":          "Norton",
    "tmproxy.exe":                 "Trend Micro",
    "trayicos.exe":                "EScan",
    "tuih.exe":                    "Norton",
    "uistub.exe":                  "Norton",
    "uiwnsnotificationapp.exe":    "Norton",
    "updaterui.exe":               "McAfee",
    "updtnv28.exe":                "Norton",
    "upgrade.exe":                 "Norton",
    "vet32.exe":                   "ETrustEZ",
    "vetmsg.exe":                  "ETrustEZ",
    "vettray.exe":                 "ETrustEZ",
    "vpnca.exe":                   "Norton",
    "vptray.exe":                  "Norton",
    "vsserv.exe":                  "Bitdefender",
    "wa_3rd_party_host_32.exe":    "Norton",
    "wa_3rd_party_host_64.exe":    "Norton",
    "webproxy.exe":                "Panda",
    "webscanx.exe":                "McAfee",
    "wfpunins.exe":                "Norton",
    "wpinstca.exe":                "Norton",
    "wrsa.exe":                    "Webroot",
    "wsc_proxy.exe":               "Avast",
    "wscstub.exe":                 "Norton",
    "xcommsvr.exe":                "Bitdefender",
    "zaprivacyservice.exe":        "ZoneAlarm",
}

av_found = {}

for proc in wmi.Win32_Process():
    proc_name = proc.Name.lower()
    if proc_name in av_processes:
        av_name = av_processes[proc_name]
        if av_name in av_found:
            av_found[av_name].append({"name": proc_name, "id": proc.ProcessId})
        else:
            av_found[av_name] = [{"name": proc_name, "id": proc.ProcessId}]

if len(av_found):
    print("\n\n## Antiviruses\n")
    for name, procs in av_found.items():
        print(f"- {name}")
        for proc in procs:
            print(f"  - {proc['name']} (PID {proc['id']})")


# Unattended files

unattended_filepaths = [
    os.path.normpath(windir + "\\..\\unattend.inf"),
    os.path.normpath(windir + "\\..\\unattend.txt"),
    os.path.normpath(windir + "\\panther\\unattend.xml"),
    os.path.normpath(windir + "\\panther\\unattend\\unattend.xml"),
    os.path.normpath(windir + "\\panther\\unattend\\unattended.xml"),
    os.path.normpath(windir + "\\panther\\unattended.xml"),
    os.path.normpath(windir + "\\sysprep.inf"),
    os.path.normpath(windir + "\\sysprep\\sysprep.inf"),
    os.path.normpath(windir + "\\sysprep\\sysprep.xml"),
    os.path.normpath(windir + "\\system32\\sysprep\\unattend.xml"),
    os.path.normpath(windir + "\\system32\\sysprep\\unattended.xml"),
]

unattended_files_found = []

for path in unattended_filepaths:
    if os.path.isfile(path):
        unattended_files_found.append(path)

if len(unattended_files_found):
    print("\n\n## Unattended files\n")
    for file in unattended_files_found:
        print(f"- {file}")

print()
