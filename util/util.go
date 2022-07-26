package util

import "unsafe"

const DeviceIMEI = "imei"
const DeviceOAID = "oaid"
const DeviceIDFA = "idfa"

func StringToBytes(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer(s))
}
