package main

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Process struct {
	Exe string
	PID uint32
}

type AV string

const (
	AVAcronis           AV = "Acronis"
	AVAdaware           AV = "Adaware"
	AVAhnLab            AV = "AhnLab"
	AVAlibaba           AV = "Alibaba"
	AVALYac             AV = "ALYac"
	AVAntiy             AV = "Antiy"
	AVArcabit           AV = "Arcabit"
	AVAvast             AV = "Avast"
	AVAVG               AV = "AVG"
	AVAvira             AV = "Avira"
	AVBaidu             AV = "Baidu"
	AVBitdefender       AV = "Bitdefender"
	AVBullGuard         AV = "BullGuard"
	AVClamAV            AV = "Clam AV"
	AVClamWin           AV = "ClamWin"
	AVCMC               AV = "CMC"
	AVCrowdStrike       AV = "CrowdStrike"
	AVCybereason        AV = "Cybereason"
	AVCylance           AV = "Cylance"
	AVCynet             AV = "Cynet"
	AVCyren             AV = "Cyren"
	AVDrWeb             AV = "DrWeb"
	AVElastic           AV = "Elastic"
	AVEmsisoft          AV = "Emsisoft"
	AVEScan             AV = "EScan"
	AVESET              AV = "ESET"
	AVETrustEZ          AV = "ETrustEZ"
	AVFortinet          AV = "Fortinet"
	AVFSecure           AV = "F-Secure"
	AVGData             AV = "G Data"
	AVGoogle            AV = "Google"
	AVGridinsoft        AV = "Gridinsoft"
	AVHBEDV             AV = "HBEDV"
	AVIkarus            AV = "Ikarus"
	AVJiangmin          AV = "Jiangmin"
	AVK7                AV = "K7"
	AVKaspersky         AV = "Kaspersky"
	AVLionic            AV = "Lionic"
	AVMalwarebytes      AV = "Malwarebytes"
	AVMaxSecure         AV = "Max Secure"
	AVMcAfee            AV = "McAfee"
	AVMicrosoftDefender AV = "Microsoft Defender"
	AVNano              AV = "Nano"
	AVNorton            AV = "Norton"
	AVPaloAltoNetworks  AV = "Palo Alto Networks"
	AVPanda             AV = "Panda"
	AVQuickHeal         AV = "Quick Heal"
	AVRising            AV = "Rising"
	AVSangfor           AV = "Sangfor"
	AVSecureAge         AV = "SecureAge"
	AVSentinelOne       AV = "SentinelOne"
	AVSophos            AV = "Sophos"
	AVSuperAntiSpyware  AV = "SuperAntiSpyware"
	AVTachyon           AV = "Tachyon"
	AVTehtris           AV = "Tehtris"
	AVTencent           AV = "Tencent"
	AVTrapmine          AV = "Trapmine"
	AVTrellix           AV = "Trellix"
	AVTrendMicro        AV = "Trend Micro"
	AVTrustlook         AV = "Trustlook"
	AVVba32             AV = "VBA32"
	AVVIPRE             AV = "VIPRE"
	AVVirIt             AV = "VirIT"
	AVViRobot           AV = "ViRobot"
	AVWebroot           AV = "Webroot"
	AVXcitium           AV = "Xcitium"
	AVYandex            AV = "Yandex"
	AVZillya            AV = "Zillya"
	AVZoneAlarm         AV = "ZoneAlarm"
	AVZoner             AV = "Zoner"
)

var avProcesses = map[string]AV{
	"adawareservice.exe":          AVAdaware,
	"advchk.exe":                  AVNorton,
	"afwserv.exe":                 AVAvast,
	"ahnsd.exe":                   AVAhnLab,
	"alertsvc.exe":                AVNorton,
	"alunotify.exe":               AVNorton,
	"arestore.exe":                AVNorton,
	"ashwebsv.exe":                AVAvast,
	"asoelnch.exe":                AVNorton,
	"aswengsrv.exe":               AVAvast,
	"aswidsagent.exe":             AVAvast,
	"aswtoolssvc.exe":             AVAvast,
	"aswupdsv.exe":                AVAvast,
	"avastbrowser.exe":            AVAvast,
	"avastnm.exe":                 AVAvast,
	"avastsvc.exe":                AVAvast,
	"avastui.exe":                 AVAvast,
	"avgsvc.exe":                  AVAVG,
	"avguard.exe":                 AVAvira,
	"avmaisrv.exe":                AVAvast,
	"avp.exe":                     AVKaspersky,
	"avpcc.exe":                   AVKaspersky,
	"avpm.exe":                    AVKaspersky,
	"avsched32.exe":               AVHBEDV,
	"avwupsrv.exe":                AVHBEDV,
	"bdagent.exe":                 AVBitdefender,
	"bdmcon.exe":                  AVBitdefender,
	"bdnagent.exe":                AVBitdefender,
	"bdoesrv.exe":                 AVBitdefender,
	"bdss.exe":                    AVBitdefender,
	"bdswitch.exe":                AVBitdefender,
	"bitdefender_p2p_startup.exe": AVBitdefender,
	"bullguardcore.exe":           AVBullGuard,
	"buvss.exe":                   AVNorton,
	"cavrid.exe":                  AVETrustEZ,
	"cavtray.exe":                 AVETrustEZ,
	"cltlmh.exe":                  AVNorton,
	"cltrt.exe":                   AVNorton,
	"cmgrdian.exe":                AVMcAfee,
	"coinst.exe":                  AVNorton,
	"conathst.exe":                AVNorton,
	"cpysnpt.exe":                 AVNorton,
	"cyneteps.exe":                AVCynet,
	"cynetms.exe":                 AVCynet,
	"doscan.exe":                  AVNorton,
	"dvpapi.exe":                  AVCyren,
	"efainst64.exe":               AVNorton,
	"ekrn.exe":                    AVESET,
	"elaminst.exe":                AVNorton,
	"enterpriseservice.exe":       AVVIPRE,
	"fldghost.exe":                AVNorton,
	"frameworkservic.exe":         AVMcAfee,
	"frameworkservice.exe":        AVMcAfee,
	"freshclam.exe":               AVClamWin,
	"fshoster32.exe":              AVFSecure,
	"gdscan.exe":                  AVGData,
	"icepack.exe":                 AVNorton,
	"instca.exe":                  AVNorton,
	"k7crvsvc.exe":                AVK7,
	"kavfs.exe":                   AVKaspersky,
	"kavfsgt.exe":                 AVKaspersky,
	"kavfsmui.exe":                AVKaspersky,
	"kavfsrcn.exe":                AVKaspersky,
	"kavfsscs.exe":                AVKaspersky,
	"kavfswh.exe":                 AVKaspersky,
	"kavfswp.exe":                 AVKaspersky,
	"kavshell.exe":                AVKaspersky,
	"kavtray.exe":                 AVKaspersky,
	"mbam.exe":                    AVMalwarebytes,
	"mcapexe.exe":                 AVMcAfee,
	"mcui32.exe":                  AVNorton,
	"mfemms.exe":                  AVMcAfee,
	"mgavrtcl.exe":                AVMcAfee,
	"mghtml.exe":                  AVMcAfee,
	"mgui.exe":                    AVBullGuard,
	"msmpeng.exe":                 AVMicrosoftDefender,
	"navapsvc.exe":                AVNorton,
	"navapw32.exe":                AVNorton,
	"navw32.exe":                  AVNorton,
	"ncolow.exe":                  AVNorton,
	"nod32krn.exe":                AVESET,
	"nod32kui.exe":                AVESET,
	"nortonsecurity.exe":          AVNorton,
	"npfmntor.exe":                AVNorton,
	"nsc.exe":                     AVNorton,
	"nsmdtr.exe":                  AVNorton,
	"nswscsvc.exe":                AVNorton,
	"ntrtscan.exe":                AVTrendMicro,
	"nuperfscan.exe":              AVNorton,
	"ofcdog.exe":                  AVTrendMicro,
	"patch.exe":                   AVTrendMicro,
	"pavfires.exe":                AVPanda,
	"pavfnsvr.exe":                AVPanda,
	"pavkre.exe":                  AVPanda,
	"pavmail.exe":                 AVPanda,
	"pavprot.exe":                 AVPanda,
	"pavprsrv.exe":                AVPanda,
	"pavsched.exe":                AVPanda,
	"pavsrv50.exe":                AVPanda,
	"pavsrv51.exe":                AVPanda,
	"pavsrv52.exe":                AVPanda,
	"pavupg.exe":                  AVPanda,
	"pcscan.exe":                  AVTrendMicro,
	"pntiomon.exe":                AVTrendMicro,
	"pop3pack.exe":                AVTrendMicro,
	"pop3trap.exe":                AVTrendMicro,
	"poproxy.exe":                 AVNorton,
	"prevsrv.exe":                 AVPanda,
	"realmon.exe":                 AVETrustEZ,
	"ruleup.exe":                  AVNorton,
	"savscan.exe":                 AVNorton,
	"savservice.exe":              AVSophos,
	"sbserv.exe":                  AVNorton,
	"scan32.exe":                  AVMcAfee,
	"sefinst.exe":                 AVNorton,
	"sevntx64.exe":                AVNorton,
	"spider.exe":                  AVDrWeb,
	"srtsp_ca.exe":                AVNorton,
	"symdgnhc.exe":                AVNorton,
	"symerr.exe":                  AVNorton,
	"symvtcatalogdb.exe":          AVNorton,
	"tmproxy.exe":                 AVTrendMicro,
	"trayicos.exe":                AVEScan,
	"tuih.exe":                    AVNorton,
	"uistub.exe":                  AVNorton,
	"uiwnsnotificationapp.exe":    AVNorton,
	"updaterui.exe":               AVMcAfee,
	"updtnv28.exe":                AVNorton,
	"upgrade.exe":                 AVNorton,
	"vet32.exe":                   AVETrustEZ,
	"vetmsg.exe":                  AVETrustEZ,
	"vettray.exe":                 AVETrustEZ,
	"vpnca.exe":                   AVNorton,
	"vptray.exe":                  AVNorton,
	"vsserv.exe":                  AVBitdefender,
	"wa_3rd_party_host_32.exe":    AVNorton,
	"wa_3rd_party_host_64.exe":    AVNorton,
	"webproxy.exe":                AVPanda,
	"webscanx.exe":                AVMcAfee,
	"wfpunins.exe":                AVNorton,
	"wpinstca.exe":                AVNorton,
	"wrsa.exe":                    AVWebroot,
	"wsc_proxy.exe":               AVAvast,
	"wscstub.exe":                 AVNorton,
	"xcommsvr.exe":                AVBitdefender,
	"zaprivacyservice.exe":        AVZoneAlarm,
}

func printAV() {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		fmt.Printf("Error: CreateToolhelp32Snapshot failed: %v\n\n", err)
		return
	}
	defer windows.CloseHandle(snapshot)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	if err = windows.Process32First(snapshot, &entry); err != nil {
		fmt.Printf("Error: Process32First failed: %v\n\n", err)
		return
	}

	avFound := make(map[AV][]Process)

	for {
		exe := strings.ToLower(syscall.UTF16ToString(entry.ExeFile[:]))

		if av, ok := avProcesses[exe]; ok {
			avFound[av] = append(avFound[av], Process{
				Exe: exe,
				PID: entry.ProcessID,
			})
		}

		// get next process
		if err = windows.Process32Next(snapshot, &entry); err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				break
			}
			fmt.Printf("Error: Process32Next failed: %v\n\n", err)
			return
		}
	}

	if len(avFound) == 0 {
		return
	}

	fmt.Print("## Antiviruses\n\n")
	for name, ps := range avFound {
		fmt.Printf("- %s\n", name)
		for _, p := range ps {
			fmt.Printf("  - %s (PID %d)\n", p.Exe, p.PID)
		}
	}
	fmt.Print("\n\n")
}

// Analysis tools that may be avoid.
var procsBlacklist = map[string]struct{}{
	"dfsserv":        {},
	"fiddler":        {},
	"gemu-ga":        {},
	"httpdebuggerui": {},
	"ida64":          {},
	"joeboxcontrol":  {},
	"joeer":          {},
	"ksdumper":       {},
	"ksdumperclient": {},
	"ollydbg":        {},
	"pestudio":       {},
	"pr1_tools":      {},
	"prl_cc":         {},
	"processhacker":  {},
	"regedit":        {},
	"taskmgr":        {},
	"vboxservice":    {},
	"vboxtray":       {},
	"vgauthservice":  {},
	"vmacthlp":       {},
	"vmsrvc":         {},
	"vmtoolsd":       {},
	"vmusrvc":        {},
	"vmwaretray":     {},
	"vmwareuser":     {},
	"wireshark":      {},
	"x32dbg":         {},
	"x96dbg":         {},
}
