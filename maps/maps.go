package maps

import (
	"strconv"
)

//maps
func Product(id uint64) string {
	product := make(map[uint64]string)

	product[id] = "UNKNOWN PRODUCT ID=" + strconv.Itoa(int(id))

	//products list
	product[1] = "HRM1"
	product[2] = "XH01"
	product[3] = "AXB01"
	product[4] = "AXB02"
	product[5] = "HRM2SS"
	product[6] = "DSI_ALF02"
	product[7] = "HRM3SS"
	product[8] = "HRM_RUN_SINGLE_BYTE_PRODUCT_ID"
	product[9] = "BSM"
	product[10] = "BCM"
	product[11] = "AXS01"
	product[473] = "FR301_CHINA"
	product[474] = "FR301_JAPAN"
	product[475] = "FR301_KOREA"
	product[494] = "FR301_TAIWAN"
	product[717] = "FR405"
	product[782] = "FR50"
	product[987] = "FR405_JAPAN"
	product[988] = "FR60"
	product[1011] = "DSI_ALF01"
	product[1018] = "FR310XT"
	product[1036] = "EDGE500"
	product[1124] = "FR110"
	product[1169] = "EDGE800"
	product[1199] = "EDGE500_TAIWAN"
	product[1213] = "EDGE500_JAPAN"
	product[1253] = "CHIRP"
	product[1274] = "FR110_JAPAN"
	product[1325] = "EDGE200"
	product[1328] = "FR910XT"
	product[1333] = "EDGE800_TAIWAN"
	product[1334] = "EDGE800_JAPAN"
	product[1341] = "ALF04"
	product[1345] = "FR610"
	product[1360] = "FR210_JAPAN"
	product[1380] = "VECTOR_SS"
	product[1381] = "VECTOR_CP"
	product[1386] = "EDGE800_CHINA"
	product[1387] = "EDGE500_CHINA"
	product[1410] = "FR610_JAPAN"
	product[1422] = "EDGE500_KOREA"
	product[1436] = "FR70"
	product[1446] = "FR310XT_4T"
	product[1461] = "AMX"
	product[1482] = "FR10"
	product[1497] = "EDGE800_KOREA"
	product[1499] = "SWIM"
	product[1537] = "FR910XT_CHINA"
	product[1551] = "FENIX"
	product[1555] = "EDGE200_TAIWAN"
	product[1561] = "EDGE510"
	product[1567] = "EDGE810"
	product[1570] = "TEMPE"
	product[1600] = "FR910XT_JAPAN"
	product[1623] = "FR620"
	product[1632] = "FR220"
	product[1664] = "FR910XT_KOREA"
	product[1688] = "FR10_JAPAN"
	product[1721] = "EDGE810_JAPAN"
	product[1735] = "VIRB_ELITE"
	product[1736] = "EDGE_TOURING"
	product[1742] = "EDGE510_JAPAN"
	product[1752] = "HRM_RUN"
	product[1765] = "FR920XT"
	product[1821] = "EDGE510_ASIA"
	product[1822] = "EDGE810_CHINA"
	product[1823] = "EDGE810_TAIWAN"
	product[1836] = "EDGE1000"
	product[1837] = "VIVO_FIT"
	product[1853] = "VIRB_REMOTE"
	product[1885] = "VIVO_KI"
	product[1903] = "FR15"
	product[1907] = "VIVO_ACTIVE"
	product[1918] = "EDGE510_KOREA"
	product[1928] = "FR620_JAPAN"
	product[1929] = "FR620_CHINA"
	product[1930] = "FR220_JAPAN"
	product[1931] = "FR220_CHINA"
	product[1936] = "APPROACH_S6"
	product[1956] = "VIVO_SMART"
	product[1967] = "FENIX2"
	product[1988] = "EPIX"
	product[2050] = "FENIX3"
	product[2052] = "EDGE1000_TAIWAN"
	product[2053] = "EDGE1000_JAPAN"
	product[2061] = "FR15_JAPAN"
	product[2070] = "EDGE1000_CHINA"
	product[2072] = "FR620_RUSSIA"
	product[2073] = "FR220_RUSSIA"
	product[2079] = "VECTOR_S"
	product[2100] = "EDGE1000_KOREA"
	product[2130] = "FR920XT_TAIWAN"
	product[2131] = "FR920XT_CHINA"
	product[2132] = "FR920XT_JAPAN"
	product[2134] = "VIRBX"
	product[2135] = "VIVO_SMART_APAC"
	product[2150] = "VIVO_FIT2"
	product[2153] = "FR225"
	product[2160] = "VIVO_ACTIVE_APAC"
	product[2161] = "VECTOR_2"
	product[2162] = "VECTOR_2S"
	product[2172] = "VIRBXE"
	product[2173] = "FR620_TAIWAN"
	product[2188] = "FENIX3_CHINA"
	product[2189] = "FENIX3_TWN"
	product[10007] = "SDM4"
	product[10014] = "EDGE_REMOTE"
	product[20119] = "TRAINING_CENTER"
	product[65532] = "ANDROID_ANTPLUS_PLUGIN"
	product[65534] = "CONNECT"

	return product[id]
}

func Manufacturer(id uint64) string {
	manufacturer := make(map[uint64]string)

	manufacturer[id] = "UNKNOWN MANUFACTURER ID=" + strconv.Itoa(int(id))

	manufacturer[1] = "GARMIN"
	manufacturer[2] = "GARMIN_FR405_ANTFS"
	manufacturer[3] = "ZEPHYR"
	manufacturer[4] = "DAYTON"
	manufacturer[5] = "IDT"
	manufacturer[6] = "SRM"
	manufacturer[7] = "QUARQ"
	manufacturer[8] = "IBIKE"
	manufacturer[9] = "SARIS"
	manufacturer[10] = "SPARK_HK"
	manufacturer[11] = "TANITA"
	manufacturer[12] = "ECHOWELL"
	manufacturer[13] = "DYNASTREAM_OEM"
	manufacturer[14] = "NAUTILUS"
	manufacturer[15] = "DYNASTREAM"
	manufacturer[16] = "TIMEX"
	manufacturer[17] = "METRIGEAR"
	manufacturer[18] = "XELIC"
	manufacturer[19] = "BEURER"
	manufacturer[20] = "CARDIOSPORT"
	manufacturer[21] = "A_AND_D"
	manufacturer[22] = "HMM"
	manufacturer[23] = "SUUNTO"
	manufacturer[24] = "THITA_ELEKTRONIK"
	manufacturer[25] = "GPULSE"
	manufacturer[26] = "CLEAN_MOBILE"
	manufacturer[27] = "PEDAL_BRAIN"
	manufacturer[28] = "PEAKSWARE"
	manufacturer[29] = "SAXONAR"
	manufacturer[30] = "LEMOND_FITNESS"
	manufacturer[31] = "DEXCOM"
	manufacturer[32] = "WAHOO_FITNESS"
	manufacturer[33] = "OCTANE_FITNESS"
	manufacturer[34] = "ARCHINOETICS"
	manufacturer[35] = "THE_HURT_BOX"
	manufacturer[36] = "CITIZEN_SYSTEMS"
	manufacturer[37] = "MAGELLAN"
	manufacturer[38] = "OSYNCE"
	manufacturer[39] = "HOLUX"
	manufacturer[40] = "CONCEPT2"
	manufacturer[42] = "ONE_GIANT_LEAP"
	manufacturer[43] = "ACE_SENSOR"
	manufacturer[44] = "BRIM_BROTHERS"
	manufacturer[45] = "XPLOVA"
	manufacturer[46] = "PERCEPTION_DIGITAL"
	manufacturer[47] = "BF1SYSTEMS"
	manufacturer[48] = "PIONEER"
	manufacturer[49] = "SPANTEC"
	manufacturer[50] = "METALOGICS"
	manufacturer[51] = "4IIIIS"
	manufacturer[52] = "SEIKO_EPSON"
	manufacturer[53] = "SEIKO_EPSON_OEM"
	manufacturer[54] = "IFOR_POWELL"
	manufacturer[55] = "MAXWELL_GUIDER"
	manufacturer[56] = "STAR_TRAC"
	manufacturer[57] = "BREAKAWAY"
	manufacturer[58] = "ALATECH_TECHNOLOGY_LTD"
	manufacturer[59] = "MIO_TECHNOLOGY_EUROPE"
	manufacturer[60] = "ROTOR"
	manufacturer[61] = "GEONAUTE"
	manufacturer[62] = "ID_BIKE"
	manufacturer[63] = "SPECIALIZED"
	manufacturer[64] = "WTEK"
	manufacturer[65] = "PHYSICAL_ENTERPRISES"
	manufacturer[66] = "NORTH_POLE_ENGINEERING"
	manufacturer[67] = "BKOOL"
	manufacturer[68] = "CATEYE"
	manufacturer[69] = "STAGES_CYCLING"
	manufacturer[70] = "SIGMASPORT"
	manufacturer[71] = "TOMTOM"
	manufacturer[72] = "PERIPEDAL"
	manufacturer[73] = "WATTBIKE"
	manufacturer[76] = "MOXY"
	manufacturer[77] = "CICLOSPORT"
	manufacturer[78] = "POWERBAHN"
	manufacturer[79] = "ACORN_PROJECTS_APS"
	manufacturer[80] = "LIFEBEAM"
	manufacturer[81] = "BONTRAGER"
	manufacturer[82] = "WELLGO"
	manufacturer[83] = "SCOSCHE"
	manufacturer[84] = "MAGURA"
	manufacturer[85] = "WOODWAY"
	manufacturer[86] = "ELITE"
	manufacturer[87] = "NIELSEN_KELLERMAN"
	manufacturer[88] = "DK_CITY"
	manufacturer[89] = "TACX"
	manufacturer[90] = "DIRECTION_TECHNOLOGY"
	manufacturer[91] = "MAGTONIC"
	manufacturer[92] = "1PARTCARBON"
	manufacturer[93] = "INSIDE_RIDE_TECHNOLOGIES"
	manufacturer[94] = "SOUND_OF_MOTION"
	manufacturer[95] = "STRYD"
	manufacturer[255] = "DEVELOPMENT"
	manufacturer[257] = "HEALTHANDLIFE"
	manufacturer[258] = "LEZYNE"
	manufacturer[259] = "SCRIBE_LABS"
	manufacturer[260] = "ZWIFT"
	manufacturer[261] = "WATTEAM"
	manufacturer[262] = "RECON"
	manufacturer[263] = "FAVERO_ELECTRONICS"
	manufacturer[264] = "DYNOVELO"
	manufacturer[265] = "STRAVA"
	manufacturer[5759] = "ACTIGRAPHCORP"

	return manufacturer[id]
}

func Sport(id uint64) string {
	sport := make(map[uint64]string)

	sport[id] = "UNKNOWN SPORT ID=" + strconv.Itoa(int(id))

	sport[0] = "GENERIC"
	sport[1] = "RUNNING"
	sport[2] = "CYCLING"
	sport[3] = "TRANSITION"
	sport[4] = "FITNESS_EQUIPMENT"
	sport[5] = "SWIMMING"
	sport[6] = "BASKETBALL"
	sport[7] = "SOCCER"
	sport[8] = "TENNIS"
	sport[9] = "AMERICAN_FOOTBALL"
	sport[10] = "TRAINING"
	sport[11] = "WALKING"
	sport[12] = "CROSS_COUNTRY_SKIING"
	sport[13] = "ALPINE_SKIING"
	sport[14] = "SNOWBOARDING"
	sport[15] = "ROWING"
	sport[16] = "MOUNTAINEERING"
	sport[17] = "HIKING"
	sport[18] = "MULTISPORT"
	sport[19] = "PADDLING"
	sport[20] = "FLYING"
	sport[21] = "E_BIKING"
	sport[254] = "ALL"

	return sport[id]
}

func Sub_sport(id uint64) string {
	sub_sport := make(map[uint64]string)

	sub_sport[id] = "UNKNOWN SUB SPORT ID=" + strconv.Itoa(int(id))

	sub_sport[0] = "GENERIC"
	sub_sport[1] = "TREADMILL"
	sub_sport[2] = "STREET"
	sub_sport[3] = "TRAIL"
	sub_sport[4] = "TRACK"
	sub_sport[5] = "SPIN"
	sub_sport[6] = "INDOOR_CYCLING"
	sub_sport[7] = "ROAD"
	sub_sport[8] = "MOUNTAIN"
	sub_sport[9] = "DOWNHILL"
	sub_sport[10] = "RECUMBENT"
	sub_sport[11] = "CYCLOCROSS"
	sub_sport[12] = "HAND_CYCLING"
	sub_sport[13] = "TRACK_CYCLING"
	sub_sport[14] = "INDOOR_ROWING"
	sub_sport[15] = "ELLIPTICAL"
	sub_sport[16] = "STAIR_CLIMBING"
	sub_sport[17] = "LAP_SWIMMING"
	sub_sport[18] = "OPEN_WATER"
	sub_sport[19] = "FLEXIBILITY_TRAINING"
	sub_sport[20] = "STRENGTH_TRAINING"
	sub_sport[21] = "WARM_UP"
	sub_sport[22] = "MATCH"
	sub_sport[23] = "EXERCISE"
	sub_sport[24] = "CHALLENGE"
	sub_sport[25] = "INDOOR_SKIING"
	sub_sport[26] = "CARDIO_TRAINING"
	sub_sport[27] = "INDOOR_WALKING"
	sub_sport[28] = "E_BIKE_FITNESS"
	sub_sport[254] = "ALL"

	return sub_sport[id]
}

func Time_trigger(id uint64) string {
	time_trigger := make(map[uint64]string)

	time_trigger[id] = "UNKNOWN TIME TRIGGER ID=" + strconv.Itoa(int(id))

	time_trigger[0] = "MANUAL"
	time_trigger[1] = "AUTO"
	time_trigger[2] = "FITNESS_EQUIPMENT"

	return time_trigger[id]
}

func Event_type(id uint64) string {
	event_type := make(map[uint64]string)

	event_type[id] = "UNKNOWN EVENT TYPE ID=" + strconv.Itoa(int(id))

	event_type[0] = "START"
	event_type[1] = "STOP"
	event_type[2] = "CONSECUTIVE_DEPRECIATED"
	event_type[3] = "MARKER"
	event_type[4] = "STOP_ALL"
	event_type[5] = "BEGIN_DEPRECIATED"
	event_type[6] = "END_DEPRECIATED"
	event_type[7] = "END_ALL_DEPRECIATED"
	event_type[8] = "STOP_DISABLE"
	event_type[9] = "STOP_DISABLE_ALL"

	return event_type[id]
}

func Event(id uint64) string {
	event := make(map[uint64]string)

	event[id] = "UNKNOWN EVENT ID=" + strconv.Itoa(int(id))

	event[0] = "TIMER"
	event[3] = "WORKOUT"
	event[4] = "WORKOUT_STEP"
	event[5] = "POWER_DOWN"
	event[6] = "POWER_UP"
	event[7] = "OFF_COURSE"
	event[8] = "SESSION"
	event[9] = "LAP"
	event[10] = "COURSE_POINT"
	event[11] = "BATTERY"
	event[12] = "VIRTUAL_PARTNER_PACE"
	event[13] = "HR_HIGH_ALERT"
	event[14] = "HR_LOW_ALERT"
	event[15] = "SPEED_HIGH_ALERT"
	event[16] = "SPEED_LOW_ALERT"
	event[17] = "CAD_HIGH_ALERT"
	event[18] = "CAD_LOW_ALERT"
	event[19] = "POWER_HIGH_ALERT"
	event[20] = "POWER_LOW_ALERT"
	event[21] = "RECOVERY_HR"
	event[22] = "BATTERY_LOW"
	event[23] = "TIME_DURATION_ALERT"
	event[24] = "DISTANCE_DURATION_ALERT"
	event[25] = "CALORIE_DURATION_ALERT"
	event[26] = "ACTIVITY"
	event[27] = "FITNESS_EQUIPMENT"
	event[28] = "LENGTH"

	return event[id]
}

func Device_type(id uint64) string {
	device_type := make(map[uint64]string)

	device_type[id] = "UNKNOWN DEVICE TYPE ID=" + strconv.Itoa(int(id))

	device_type[1] = "ANTFS"
	device_type[11] = "BIKE_POWER"
	device_type[12] = "ENVIRONMENT_SENSOR_LEGACY"
	device_type[15] = "MULTI_SPORT_SPEED_DISTANCE"
	device_type[16] = "CONTROL"
	device_type[17] = "FITNESS_EQUIPMENT"
	device_type[18] = "BLOOD_PRESSURE"
	device_type[19] = "GEOCACHE_NODE"
	device_type[20] = "LIGHT_ELECTRIC_VEHICLE"
	device_type[25] = "ENV_SENSOR"
	device_type[26] = "RACQUET"
	device_type[119] = "WEIGHT_SCALE"
	device_type[120] = "HEART_RATE"
	device_type[121] = "BIKE_SPEED_CADENCE"
	device_type[122] = "BIKE_CADENCE"
	device_type[123] = "BIKE_SPEED"
	device_type[124] = "STRIDE_SPEED_DISTANCE"

	return device_type[id]
}

func Source_type(id uint64) string {
	source_type := make(map[uint64]string)

	source_type[id] = "UNKNOWN SOURCE TYPE ID=" + strconv.Itoa(int(id))

	source_type[0] = "ANT"                  // External device connected with ANT
	source_type[1] = "ANTPLUS"              // External device connected with ANT+
	source_type[2] = "BLUETOOTH"            // External device connected with BT
	source_type[3] = "BLUETOOTH_LOW_ENERGY" // External device connected with BLE
	source_type[4] = "WIFI"                 // External device connected with Wifi
	source_type[5] = "LOCAL"                // Onboard device

	return source_type[id]
}

func Ant_network(id uint64) string {
	ant_network := make(map[uint64]string)

	ant_network[id] = "UNKNOWN SOURCE TYPE ID=" + strconv.Itoa(int(id))

	ant_network[0] = "PUBLIC"
	ant_network[1] = "ANTPLUS"
	ant_network[2] = "ANTFS"
	ant_network[3] = "PRIVATE"

	return ant_network[id]
}

func Session_trigger(id uint64) string {
	session_trigger := make(map[uint64]string)

	session_trigger[id] = "UNKNOWN SESSION TRIGGER ID=" + strconv.Itoa(int(id))

	session_trigger[0] = "ACTIVITY_END"
	session_trigger[1] = "MANUAL"
	session_trigger[2] = "AUTO_MULTI_SPORT"
	session_trigger[3] = "FITNESS_EQUIPMENT"

	return session_trigger[id]
}

func Activity(id uint64) string {
	activity := make(map[uint64]string)

	activity[id] = "UNKNOWN ACTIVITY ID=" + strconv.Itoa(int(id))

	activity[0] = "MANUAL"
	activity[1] = "AUTO_MULTI_SPORT"

	return activity[id]
}

func Global_message_type(id uint16) string {
	global_message_type := make(map[uint16]string)

	global_message_type[id] = "UNKNOWN GLOBAL MESSAGE TYPE ID=" + strconv.Itoa(int(id))

	global_message_type[0] = "FILE_ID"
	global_message_type[18] = "SESSION"
	global_message_type[19] = "LAP"
	global_message_type[20] = "RECORD"
	global_message_type[21] = "EVENT"
	global_message_type[23] = "DEVICE_INFO"
	global_message_type[34] = "ACTIVITY"
	global_message_type[49] = "FILE_CREATOR"

	return global_message_type[id]
}
