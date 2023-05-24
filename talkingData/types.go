package talkingData

//deviceIdType 设备ID类型如下：
//66代表 IMEI 明文；
//67代表 IDFA 明文；
//78代表 OAID 明文及 OAID_MD5（两者同时支持）； <==
//242代表 IMEI_MD5； <==
//243代表 IDFA_MD5； <==

const DeviceIdTypeIMEIMd5 = 242
const DeviceIdTypeOAIDMd5 = 78
const DeviceIdTypeIDFAMd5 = 243

const DeviceIdTypeOAID = 78
const DeviceIdTypeIDFA = 67
const DeviceIdTypeIMEI = 66
