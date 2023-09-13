package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// See https://learn.microsoft.com/en-us/windows/win32/api/iads/ne-iads-ads_user_flag_enum
const (
	ADS_UF_SCRIPT                                 = 0x1
	ADS_UF_ACCOUNTDISABLE                         = 0x2
	ADS_UF_HOMEDIR_REQUIRED                       = 0x8
	ADS_UF_LOCKOUT                                = 0x10
	ADS_UF_PASSWD_NOTREQD                         = 0x20
	ADS_UF_PASSWD_CANT_CHANGE                     = 0x40
	ADS_UF_ENCRYPTED_TEXT_PASSWORD_ALLOWED        = 0x80
	ADS_UF_TEMP_DUPLICATE_ACCOUNT                 = 0x100
	ADS_UF_NORMAL_ACCOUNT                         = 0x200
	ADS_UF_INTERDOMAIN_TRUST_ACCOUNT              = 0x800
	ADS_UF_WORKSTATION_TRUST_ACCOUNT              = 0x1000
	ADS_UF_SERVER_TRUST_ACCOUNT                   = 0x2000
	ADS_UF_DONT_EXPIRE_PASSWD                     = 0x10000
	ADS_UF_MNS_LOGON_ACCOUNT                      = 0x20000
	ADS_UF_SMARTCARD_REQUIRED                     = 0x40000
	ADS_UF_TRUSTED_FOR_DELEGATION                 = 0x80000
	ADS_UF_NOT_DELEGATED                          = 0x100000
	ADS_UF_USE_DES_KEY_ONLY                       = 0x200000
	ADS_UF_DONT_REQUIRE_PREAUTH                   = 0x400000
	ADS_UF_PASSWORD_EXPIRED                       = 0x800000
	ADS_UF_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION = 0x1000000

	USER_PRIV_GUEST = 0
	USER_PRIV_USER  = 1
	USER_PRIV_ADMIN = 2
)

// See https://learn.microsoft.com/en-us/windows/win32/api/lmaccess/ns-lmaccess-user_info_2
type USER_INFO_2 struct {
	Usri2_name           *uint16
	Usri2_password       *uint16
	Usri2_password_age   uint32
	Usri2_priv           uint32
	Usri2_home_dir       *uint16
	Usri2_comment        *uint16
	Usri2_flags          uint32
	Usri2_script_path    *uint16
	Usri2_auth_flags     uint32
	Usri2_full_name      *uint16
	Usri2_usr_comment    *uint16
	Usri2_parms          *uint16
	Usri2_workstations   *uint16
	Usri2_last_logon     uint32
	Usri2_last_logoff    uint32
	Usri2_acct_expires   uint32
	Usri2_max_storage    uint32
	Usri2_units_per_week uint32
	Usri2_logon_hours    uintptr
	Usri2_bad_pw_count   uint32
	Usri2_num_logons     uint32
	Usri2_logon_server   *uint16
	Usri2_country_code   uint32
	Usri2_code_page      uint32
}

func printUsers() {
	fmt.Print("## Users\n\n")

	modNetapi32 := windows.NewLazySystemDLL("netapi32.dll")
	procNetUserEnum := modNetapi32.NewProc("NetUserEnum") // https://learn.microsoft.com/en-us/windows/win32/api/lmaccess/nf-lmaccess-netuserenum
	procNetApiBufferFree := modNetapi32.NewProc("NetApiBufferFree")

	var (
		dataPtr      uintptr
		entriesRead  uint32
		entiresTotal uint32
		resumeHandle uintptr
		sizeTest     USER_INFO_2
	)

	// https://learn.microsoft.com/en-us/windows/win32/api/lmaccess/nf-lmaccess-netuserenum
	r, _, _ := procNetUserEnum.Call(
		uintptr(0),                             // servername: local
		uintptr(2),                             // level: USER_INFO_2
		uintptr(2),                             // filter: FILTER_NORMAL_ACCOUNT
		uintptr(unsafe.Pointer(&dataPtr)),      // bufptr
		uintptr(0xFFFFFFFF),                    // prefmaxlen
		uintptr(unsafe.Pointer(&entriesRead)),  // entriesread
		uintptr(unsafe.Pointer(&entiresTotal)), // totalentries
		uintptr(unsafe.Pointer(&resumeHandle)), // resume_handle
	)
	if r != uintptr(windows.ERROR_SUCCESS) {
		fmt.Print("Error: fetching users failed\n\n")
		return
	} else if dataPtr == uintptr(0) {
		fmt.Print("Error: null pointer while fetching users\n\n")
		return
	}

	iter := dataPtr
	for i := uint32(0); i < entriesRead; i++ {
		data := (*USER_INFO_2)(unsafe.Pointer(iter))

		if (data.Usri2_flags & ADS_UF_ACCOUNTDISABLE) == ADS_UF_ACCOUNTDISABLE {
			iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
			continue
		}

		fmt.Println("- Username:          ", windows.UTF16PtrToString(data.Usri2_name))
		if fullname := windows.UTF16PtrToString(data.Usri2_full_name); fullname != "" {
			fmt.Println("  Full name:         ", windows.UTF16PtrToString(data.Usri2_full_name))
		}
		fmt.Print("  Privilege level:    ")
		switch data.Usri2_priv {
		case USER_PRIV_GUEST:
			fmt.Print("Guest")
		case USER_PRIV_USER:
			fmt.Print("User")
		case USER_PRIV_ADMIN:
			fmt.Print("Administrator")
		default:
			fmt.Print("Unknown")
		}
		fmt.Print("\n")
		if passwordAge := time.Duration(data.Usri2_password_age) * time.Second; passwordAge > 0 {
			fmt.Println("  Password age:      ", passwordAge)
		}
		if lastLogon := time.Unix(int64(data.Usri2_last_logon), 0); lastLogon.Unix() > 0 {
			fmt.Println("  Last logon:        ", time.Unix(int64(data.Usri2_last_logon), 0))
		}
		if logonCount := data.Usri2_num_logons; logonCount > 0 {
			fmt.Println("  Logon count:       ", data.Usri2_num_logons)
		}
		if badPasswordCount := data.Usri2_bad_pw_count; badPasswordCount > 0 {
			fmt.Println("  Bad password count:", data.Usri2_bad_pw_count)
		}

		var flags []string
		if (data.Usri2_flags & ADS_UF_LOCKOUT) == ADS_UF_LOCKOUT {
			flags = append(flags, "Locked")
		}
		if (data.Usri2_flags & ADS_UF_PASSWD_NOTREQD) == ADS_UF_PASSWD_NOTREQD {
			flags = append(flags, "[!] Password not required")
		}
		if (data.Usri2_flags & ADS_UF_PASSWD_CANT_CHANGE) == ADS_UF_PASSWD_CANT_CHANGE {
			flags = append(flags, "Password can't change")
		}
		if (data.Usri2_flags & ADS_UF_DONT_EXPIRE_PASSWD) == ADS_UF_DONT_EXPIRE_PASSWD {
			flags = append(flags, "Password never expires")
		}
		if (data.Usri2_flags & ADS_UF_PASSWORD_EXPIRED) == ADS_UF_PASSWORD_EXPIRED {
			flags = append(flags, "Password expired")
		}
		if len(flags) > 0 {
			fmt.Println("  Flags:             ", strings.Join(flags, " | "))
		}

		fmt.Print("\n")

		iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
	}

	procNetApiBufferFree.Call(dataPtr)

	fmt.Print("\n")
}
