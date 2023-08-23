// Date: 2023/6/5
// Author:
// Description：

package hztime_zone

type Timezone int

const (
	Timezone_8_CanadaPacific       Timezone = -8 // UTC-8
	Timezone_7_MexicoBajaSur       Timezone = -7 // UTC-7
	Timezone_6_AmericaMonterrey    Timezone = -6 // UTC-6
	Timezone_5_AmericaIndianapolis Timezone = -5 // UTC-5
	Timezone_4_AmericaKralendijk   Timezone = -4 // UTC-4
	Timezone_3_AmericaRecife       Timezone = -3 // UTC-3
	Timezone_2_BrazilDeNoronha     Timezone = -2 // UTC-2
	Timezone_1_AtlanticAzores      Timezone = -1 // UTC-1
	Timezone0_UTC                  Timezone = 0  // UTC
	Timezone1_EuropeParis          Timezone = 1  // UTC+1
	Timezone2_AfricaCairo          Timezone = 2  // UTC+2
	Timezone3_EuropeMoscow         Timezone = 3  // UTC+3
	Timezone4_AsiaDubai            Timezone = 4  // UTC+4
	Timezone5_AsiaAqtau            Timezone = 5  // UTC+5
	Timezone6_AsiaKashgar          Timezone = 6  // UTC+6
	Timezone7_AsiaPhnom_Penh       Timezone = 7  // UTC+7
	Timezone8_AsiaShanghai         Timezone = 8  // UTC+8
)

var (
	timezoneDesc = map[Timezone]string{
		Timezone_8_CanadaPacific:       "Canada/Pacific",       // UTC-8
		Timezone_7_MexicoBajaSur:       "Mexico/BajaSur",       // UTC-7
		Timezone_6_AmericaMonterrey:    "America/Monterrey",    // UTC-6
		Timezone_5_AmericaIndianapolis: "America/Indianapolis", // UTC-5
		Timezone_4_AmericaKralendijk:   "America/Kralendijk",   // UTC-4
		Timezone_3_AmericaRecife:       "America/Recife",       // UTC-3
		Timezone_2_BrazilDeNoronha:     "Brazil/DeNoronha",     // UTC-2
		Timezone_1_AtlanticAzores:      "Atlantic/Azores",      // UTC-1
		Timezone0_UTC:                  "UTC",                  // UTC+0
		Timezone1_EuropeParis:          "Europe/Paris",         // UTC+1
		Timezone2_AfricaCairo:          "Africa/Cairo",         // UTC+2
		Timezone3_EuropeMoscow:         "Europe/Moscow",        // UTC+3
		Timezone4_AsiaDubai:            "Asia/Dubai",           // UTC+4
		Timezone5_AsiaAqtau:            "Asia/Aqtau",           // UTC+5
		Timezone6_AsiaKashgar:          "Asia/Kashgar",         // UTC+6
		Timezone7_AsiaPhnom_Penh:       "Asia/Phnom_Penh",      // UTC+7
		Timezone8_AsiaShanghai:         "Asia/Shanghai",        // UTC+8
	}
)

func (t Timezone) String() string {
	return timezoneDesc[t]
}
