package fingerprint

import (
	godbus "github.com/godbus/dbus/v5"
	systemhuawei "github.com/linuxdeepin/go-dbus-factory/system/com.huawei.fingerprint"
)

type Fingerprint = systemhuawei.Fingerprint

func NewFingerprint(conn *godbus.Conn) Fingerprint {
	return systemhuawei.NewFingerprint(conn)
}
