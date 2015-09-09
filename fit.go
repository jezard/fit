package fit

import (
	"encoding/binary"
	"fmt"
	"github.com/jezard/fit/maps"
	"math"
	"os"
	"time"
)

type File_id struct { //message number: 0
	Serial_number uint32
	Time_created  int64
	Manufacturer  string
	Product       string
	Number        uint16
	File_type     byte
}
type File_creator struct { //message number: 1
	Software_version float64
	Hardware_version uint8
}
type Device_info struct { //message number: 23
	Timestamp          int64
	Serial_number      uint32
	Cum_operating_time uint32
	Manufacturer       string
	Product            string
	Software_version   float64
	Battery_voltage    float64
	Device_index       uint8
	Device_type        string
	Hardware_version   uint8
	Battery_status     uint8
	Source_type        string
	Ant_network        string
}
type Event struct { //message number: 21
	Timestamp    int64
	Time_trigger string
	Event        string
	Event_type   string
}
type Record struct { //message number: 20
	Timestamp                 int64
	Position_lat              float64
	Position_long             float64
	Distance                  float64
	Time_from_course          int32
	Compressed_speed_distance uint8
	Heart_rate                uint8
	Altitude                  float64
	Speed                     float64
	Power                     uint16
	Grade                     int16
	Cadence                   uint8
	Registance                uint8
	Cycle_length              uint8
	Temperature               int8
}
type Lap struct { //message number: 19
	Timestamp              int64
	Start_time             int64
	Start_position_lat     float64
	Start_position_long    float64
	End_position_lat       float64
	End_position_long      float64
	Total_elapsed_time     float64
	Total_timer_time       float64
	Total_distance         float64
	Total_cycles           uint32
	Total_work             uint32
	Message_index          uint16
	Total_calories         uint16
	Total_fat_calories     uint16
	Avg_speed              float64
	Max_speed              float64
	Avg_power              uint16
	Max_power              uint16
	Norm_power             uint16
	Left_right_balance_100 float64
	Total_ascent           uint16
	Total_descent          uint16
	Avg_heart_rate         uint8
	Max_heart_rate         uint8
	Avg_cadence            uint8
	Max_cadence            uint8
	Event_group            uint8
	Event                  string
	Event_type             string
	Intensity              uint8
}
type Session struct { //message number 18
	timestamp             uint32
	start_time            uint32
	start_position_lat    int32
	start_position_long   int32
	total_elapsed_time    uint32
	total_timer_time      uint32
	total_distance        uint32
	total_cycles          uint32
	nec_lat               int32
	nec_long              int32
	swc_lat               int32
	swc_long              int32
	message_index         uint16
	total_calories        uint16
	total_fat_calories    uint16
	avg_speed             uint16
	max_speed             uint16
	avg_power             uint16
	max_power             uint16
	total_ascent          uint16
	total_descent         uint16
	first_lap_index       uint16
	num_laps              uint16
	avg_heart_rate        uint8
	max_heart_rate        uint8
	avg_cadence           uint8
	max_cadence           uint8
	total_training_effect uint8
	event_group           uint8
}
type activity struct { //message number 34
	timestamp        uint32
	total_timer_time uint32
	local_timestamp  uint32
	num_sessions     uint16
	event_group      uint8
}
type FitFile struct {
	FileId      File_id
	FileCreator File_creator
	DeviceInfo  []Device_info
	Events      []Event
	Records     []Record
	Laps        []Lap
	//Session  session
	//Activity activity
}

var crc uint16

var count int

var section string //verbose_mode - flag for record type

var verbose_mode bool

func Parse(filename string, show_verbose_mode bool) FitFile {
	verbose_mode = show_verbose_mode
	if verbose_mode {
		fmt.Printf("FUNCTION Parse() called: %v\n", time.Now())
	} //verbose_mode
	const FIT_HDR_TYPE_MASK uint8 = 0x0F
	crc = 0 //reset CRC

	//data structures
	var fitFile FitFile

	f, err := os.Open(filename)
	check(err)

	fi, err := f.Stat() //file info
	check(err)

	b1 := make([]byte, 14) //for header (14 may need to be made a dynamic value)
	n1, err := f.Read(b1)
	check(err)

	fl := fi.Size()
	b2 := make([]byte, 2)         //for crc
	n2, err := f.ReadAt(b2, fl-2) //last 2 bytes of file

	if verbose_mode {
		fmt.Printf("%d (Header) %d (crc) %d (file size) bytes\n", n1, n2, fl)
	} //verbose_mode

	/*****
	*
	* Get the Fit file header information
	*
	*****/

	headerSize, _ := binary.Uvarint(b1[0:1]) //convert 1 byte to uint64

	//Indicates the length of this file header including header size. Minimum size is 12. This may be increased in future to add additional optional information.
	if verbose_mode {
		fmt.Printf("Header Size: %v\n", headerSize)
	} //verbose_mode

	//Protocol version number as provided in SDK
	if verbose_mode {
		fmt.Printf("protocol v.: %v\n", b1[1:2])
	} //verbose_mode

	//Profile version number as provided in SDK
	pv := binary.LittleEndian.Uint16(b1[2:4])
	if verbose_mode {
		fmt.Printf("profile v. : %v\n", pv)
	} //verbose_mode

	//Length of the data records section in bytes (not including Header or CRC)
	dataSize := binary.LittleEndian.Uint32(b1[4:8])
	if verbose_mode {
		fmt.Printf("Data Size  : %v bytes\n", dataSize)
	} //verbose_mode

	//ASCII values for “.FIT”. A FIT binary file opened with a text editor will contain a readable “.FIT” in the first line.
	if verbose_mode {
		fmt.Printf("Ascii      : %s%s%s%s\n", b1[8:9], b1[9:10], b1[10:11], b1[11:12])
	} //verbose_mode

	/*****
	*
	* Get the next CRCs
	*
	*****/

	if headerSize > 12 {
		_crc := binary.LittleEndian.Uint16(b1[12:14]) //try to get the CRC from the header
		if _crc == 0x0000 {
			_crc = binary.LittleEndian.Uint16(b2[0:2]) //otherwise calcuate it from the final 2 bytes
		}
		if verbose_mode {
			fmt.Printf("Expected CRC: %x\n", _crc)
		} //verbose_mode
	} else {
		_crc := binary.LittleEndian.Uint16(b2[0:2])
		if verbose_mode {
			fmt.Printf("Expected CRC: %x\n", _crc)
		} //verbose_mode
	}
	//close the file
	f.Close()

	//RE-READ the file to get the correct calculated CRC (the contents of the file excluding the CRCs final 2 bytes)
	nf, err := os.Open(filename)
	check(err)

	b3 := make([]byte, fl-2) //file length - 2 byte final crc
	nf.Read(b3)

	for i := 0; i < len(b3); i++ {
		calc_crc(b3[i])

	}
	nf.Close()
	if verbose_mode {
		fmt.Printf("Calculated CRC: %x \n", crc)
	} //verbose_mode

	/*****
	*
	* Get the next bits (file records!!!)
	*
	*****/

	r, err := os.Open(filename)
	check(err)

	//test to read first record header
	rHead := make([]byte, 1)

	type Field_def struct {
		field_definition_number int
		size                    int
		base_type               int
		offset                  uint64 //length of data (calculated using field def size) within data record preceding the data field
	}
	type Def_message struct {
		arch                  int
		global_message_number uint16
		number_of_fields      uint64
		field_defs            []Field_def
	}
	//map definition info to a local message type and global message number !important
	definition := make(map[uint64]Def_message)

	var def_message Def_message

	var localMsgType byte

	var glob_msge_num_0_read bool //flag when first global message num = 0 read as any that follow are probably errors

	var k uint64
	for k = 0; k < uint64(dataSize); k++ { //loop through file byte by byte

		var rc_length uint64
		r.ReadAt(rHead, int64(headerSize+k))

		//see section 4.3 of SDK referring to definition header/record content,
		//4.21 describes fixed content of first 5 bytes, followed by variable number of field definitions @ 3 bytes/field

		//Is definition message? message type is bit 6
		if rHead[0]>>6 == 1 { //01000000 -> 01
			section = "DEFINITION"
			localMsgType = rHead[0] & 0x1f

			//get record content 4.2.1 of fit SDK

			//skip arch for now!
			arch := make([]byte, 1)
			r.ReadAt(arch, int64(headerSize+k+1))
			def_message.arch = int(arch[0])

			//Global Message Number
			gmn := make([]byte, 2)
			r.ReadAt(gmn, int64(headerSize+k+3))
			def_message.global_message_number = binary.LittleEndian.Uint16(gmn[0:2])

			//number of fields
			nof := make([]byte, 1)
			r.ReadAt(nof, int64(headerSize+k+5))

			def_message.number_of_fields = uint64(nof[0])

			//THIS verbose_mode INFO SEEMS PRETTY ACCURATE - I'M GETTING CONFUSED WITH GLOBAL AND LOCAL MESSAGE TYPES...
			if verbose_mode {
				fmt.Printf("\n[POS: %8d] ", uint64(dataSize)-k-1) //verbose_mode
				fmt.Print("DEFINITION MESSAGE HEADER, ")          //verbose_mode

				fmt.Printf("VAL: %b", rHead[0]) //verbose_mode

				fmt.Printf(" LOCAL MESSAGE TYPE: %d", localMsgType) //verbose_mode

				fmt.Printf(" GLOB MESSAGE NUM: %d", def_message.global_message_number)

				fmt.Printf(" FIELDS: %d", def_message.number_of_fields) //verbose_mode
			}

			//field definitions
			var f uint64

			const DEF_MSG_RECORD_HEADER_SIZE = 1 //bytes
			const DEF_MSG_FIXED_CONTENT_SIZE = 5 //bytes
			const DEF_MSG_FIELD_DEF_SIZE = 3     //bytes

			var cumulative_size uint64

			//loop through each of the fields defs
			for f = 0; f < def_message.number_of_fields; f++ {
				var def_contents Field_def
				var r_offset uint64 //byte offset for file reader

				// 1 byte field definition number. 4.2.1.4.1 in FIT SDK
				fdn := make([]byte, 1)
				r_offset = 0
				r.ReadAt(fdn, int64(headerSize+k+DEF_MSG_RECORD_HEADER_SIZE+DEF_MSG_FIXED_CONTENT_SIZE+(DEF_MSG_FIELD_DEF_SIZE*f)+r_offset))

				//...size
				size := make([]byte, 1)
				r_offset = 1
				r.ReadAt(size, int64(headerSize+k+DEF_MSG_RECORD_HEADER_SIZE+DEF_MSG_FIXED_CONTENT_SIZE+(DEF_MSG_FIELD_DEF_SIZE*f)+r_offset))

				//...base type
				baseType := make([]byte, 1)
				r_offset = 2
				r.ReadAt(baseType, int64(headerSize+k+DEF_MSG_RECORD_HEADER_SIZE+DEF_MSG_FIXED_CONTENT_SIZE+(DEF_MSG_FIELD_DEF_SIZE*f)+r_offset))

				//cumulative size for calulating data field offset
				cumulative_size += uint64(size[0])

				//def_contents.field_definition_number - for an activity file, see 9.1 FIT Messages in FIT File Types Desription
				def_contents.field_definition_number = int(fdn[0])
				def_contents.size = int(size[0])
				def_contents.base_type = int(baseType[0])
				def_contents.offset = cumulative_size - uint64(size[0]) //start, not end of field data
				if verbose_mode {
					fmt.Printf("\n\tFIELD DEF NUMBER: %v\n\tSIZE: %v\n\tBASE_TYPE: %v\n\tOFFSET %v\n",
						def_contents.field_definition_number,
						def_contents.size,
						def_contents.base_type,
						def_contents.offset)
				} //verbose_mode

				//we will need a means of temporarily storing the all the fields definition data so that it can be used to retrieve the record data later
				def_message.field_defs = append(def_message.field_defs, def_contents)

			}
			//store field definitions against local message type
			definition[uint64(localMsgType)] = def_message //of course this gets overwritten if localMsgType has been used before
			if verbose_mode {
				fmt.Printf("\t---------------\n\t%d BYTES \n", cumulative_size)
			} //verbose_mode
			rc_length = uint64(5 + uint64(nof[0])*3) //combined length of fixed and varible record content field

			k = skip(k, rc_length) //move the pointer to the end of the field definition

			continue
		} else { //10000000 (Data Message)
			section = "DATA"
			var compHeader bool
			def_message.field_defs = nil

			//check for compressed header
			if rHead[0]>>7 == 1 { //is compressed header
				compHeader = true
			}

			//set vars dependant on header type
			if compHeader {

				//timeOffset := rHead[0]&0x1F,    //time offset 0-4

				localMsgType = rHead[0] & 0x60 >> 5 //LMT is bits 5-6

			} else {
				localMsgType = rHead[0] & 0x1f //LMT is bits 0-3
			}

			if verbose_mode {
				fmt.Printf("\n[POS: %8d] DATA MESSAGE HEADER, VAL: %8b LOCAL MESSAGE TYPE: %d GLOB MESSAGE NUMBER: %d", uint64(dataSize)-k-1, rHead[0], localMsgType, definition[uint64(localMsgType)].global_message_number)
				//process data
				fmt.Println("\n")
			} //verbose_mode

			global_message_number := definition[uint64(localMsgType)].global_message_number

			//Here's where we extract the data from the .fit activity file and add it to our FitFile data structure
			switch global_message_number { //look up the global message number using the local message type

			case 0: //file_id
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size

					if !glob_msge_num_0_read {
						v := make([]byte, val.size)
						r.ReadAt(v, int64(headerSize+k+val.offset+1))
						switch val.field_definition_number {
						case 0:
							fitFile.FileId.File_type = v[0]
							if verbose_mode {
								fmt.Printf("\tFILE TYPE: %d\n", v[0])
							} //verbose_mode
							break
						case 1:
							fitFile.FileId.Manufacturer = maps.Manufacturer(uint64(binary.LittleEndian.Uint16(v[0:val.size])))
							if verbose_mode {
								fmt.Printf("\tMANUFACTURER: %s\n", fitFile.FileId.Manufacturer)
							} //verbose_mode
							break
						case 2:
							fitFile.FileId.Product = maps.Product(uint64(binary.LittleEndian.Uint16(v[0:val.size])))
							if verbose_mode {
								fmt.Printf("\tPRODUCT: %s\n", fitFile.FileId.Product)
							} //verbose_mode
							break
						case 3:
							fitFile.FileId.Serial_number = binary.LittleEndian.Uint32(v[0:val.size])
							if verbose_mode {
								fmt.Printf("\tSERIAL NUMBER: %d\n", binary.LittleEndian.Uint32(v[0:val.size]))
							} //verbose_mode
							break
						case 4:
							fitFile.FileId.Time_created = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
							t := time.Unix(fitFile.FileId.Time_created, 0)
							if verbose_mode {
								fmt.Printf("\tTIME CREATED: %d (rectified) %v\n", fitFile.FileId.Time_created, t)
							} //verbose_mode
							break
						case 5:
							fitFile.FileId.Number = binary.LittleEndian.Uint16(v[0:val.size])
							if verbose_mode {
								fmt.Printf("\tNUMBER: %d\n", binary.LittleEndian.Uint16(v[0:val.size]))
							} //verbose_mode
						}
					}

				}
				glob_msge_num_0_read = true
				//def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize)) //move the reader to the end of the record data
				break

			case 19: //lap
				var lap Lap
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 253:
						lap.Timestamp = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(lap.Timestamp, 0)
						if verbose_mode {
							fmt.Printf("\tLAP TIMESTAMP: %d (rectified) %v\n", lap.Timestamp, t)
						} //verbose_mode
						break
					case 2:
						lap.Start_time = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(lap.Start_time, 0)
						if verbose_mode {
							fmt.Printf("\tLAP START TIME: %d (rectified) %v\n", lap.Start_time, t)
						} //verbose_mode
						break
					case 3:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						lap.Start_position_lat = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLAP START LAT: %f°\n", lap.Start_position_lat)
						} //verbose_mode
						break
					case 4:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						lap.Start_position_long = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLAP START LON: %f°\n", lap.Start_position_long)
						} //verbose_mode
						break
					case 5:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						lap.End_position_lat = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLAP END LAT: %f°\n", lap.End_position_lat)
						} //verbose_mode
						break
					case 6:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						lap.End_position_long = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLAP END LON: %f°\n", lap.End_position_long)
						} //verbose_mode
						break
					case 7:
						lap.Total_elapsed_time = float64(binary.LittleEndian.Uint32(v[0:val.size])) / 1000
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL ELAPSED TIME: %fs\n", lap.Total_elapsed_time)
						} //verbose_mode
						break
					case 8:
						lap.Total_timer_time = float64(binary.LittleEndian.Uint32(v[0:val.size])) / 1000
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL TIMER TIME: %fs\n", lap.Total_timer_time)
						} //verbose_mode
						break
					case 9:
						lap.Total_distance = float64(binary.LittleEndian.Uint32(v[0:val.size])) / 100
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL DISTANCE: %f M\n", lap.Total_distance)
						} //verbose_mode
						break
					case 10:
						lap.Total_cycles = binary.LittleEndian.Uint32(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL CYCLES: %d Cycles\n", lap.Total_cycles)
						} //verbose_mode
						break
					case 41:
						lap.Total_work = binary.LittleEndian.Uint32(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL WORK: %d J\n", lap.Total_work)
						} //verbose_mode
						break
					case 254:
						lap.Message_index = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP MESSAGE INDEX: %d\n", lap.Message_index)
						} //verbose_mode
						break
					case 11:
						lap.Total_calories = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL CALORIES: %d Kcal\n", lap.Total_calories)
						} //verbose_mode
						break
					case 12:
						lap.Total_fat_calories = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL FAT CALORIES: %d Kcal\n", lap.Total_fat_calories)
						} //verbose_mode
						break
					case 13:
						lap.Avg_speed = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 1000
						if verbose_mode {
							fmt.Printf("\tLAP AVERAGE SPEED: %f M/S\n", lap.Avg_speed)
						} //verbose_mode
						break
					case 14:
						lap.Max_speed = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 1000
						if verbose_mode {
							fmt.Printf("\tLAP AVERAGE SPEED: %f M/S\n", lap.Max_speed)
						} //verbose_mode
						break
					case 19:
						lap.Avg_power = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP AVERAGE POWER %d W\n", lap.Avg_power)
						} //verbose_mode
						break
					case 20:
						lap.Max_power = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP MAX POWER %d W\n", lap.Max_power)
						} //verbose_mode
						break
					case 21:
						lap.Total_ascent = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL ASCENT %d M\n", lap.Total_ascent)
						} //verbose_mode
						break
					case 22:
						lap.Total_descent = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP TOTAL DESCENT %d M\n", lap.Total_descent)
						} //verbose_mode
						break
					case 33:
						lap.Norm_power = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tLAP NORMALIZED POWER %d W\n", lap.Norm_power)
						} //verbose_mode
						break
					case 34:
						lap.Left_right_balance_100 = (float64(binary.LittleEndian.Uint16(v[0:val.size])) / 65535) * 100
						if verbose_mode {
							fmt.Printf("\tLAP LEFT RIGHT BALANCE %1.2f Percent (0 = left, 50 = center, 100 = right) \n", lap.Left_right_balance_100) //needs verifing!
						} //verbose_mode
						break
					case 0:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Event = maps.Event(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tLAP EVENT: %s\n", lap.Event)
						} //verbose_mode
						break
					case 1:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Event_type = maps.Event_type(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tLAP EVENT TYPE: %s\n", lap.Event_type)
						} //verbose_mode
						break
					case 15:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Avg_heart_rate = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tLAP AVERAGE HEART RATE: %d BPM\n", lap.Avg_heart_rate)
						} //verbose_mode
						break
					case 16:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Max_heart_rate = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tLAP MAX HEART RATE: %d BPM\n", lap.Max_heart_rate)
						} //verbose_mode
						break
					case 17:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Avg_cadence = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tLAP AVERAGE CADENCE: %d RPM\n", lap.Avg_cadence)
						} //verbose_mode
						break
					case 18:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Max_cadence = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tLAP MAX CADENCE: %d RPM\n", lap.Max_cadence)
						} //verbose_mode
						break
					case 23:
						temp, _ := binary.Uvarint(v[0:1])
						lap.Intensity = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tLAP INTENSITY: %d \n", lap.Intensity)
						} //verbose_mode
						break

					}

				}
				fitFile.Laps = append(fitFile.Laps, lap)
				//def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize))
				break

			case 20: //Record!!!
				var record Record
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 253:
						record.Timestamp = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(record.Timestamp, 0)                                             //verbose_mode
						if verbose_mode {
							fmt.Printf("\tRECORD TIMESTAMP: %d (rectified) %v\n", record.Timestamp, t)
						} //verbose_mode
						break
					case 0:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						record.Position_lat = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLAT: %f°\n", record.Position_lat)
						} //verbose_mode
						break
					case 1:
						semicircles := float64(binary.LittleEndian.Uint32(v[0:val.size])) //convert from semicircles to degrees
						record.Position_long = semicircles_to_degrees(semicircles)
						if verbose_mode {
							fmt.Printf("\tLON: %f°\n", record.Position_long)
						} //verbose_mode
						break
					case 2:
						record.Altitude = (float64(binary.LittleEndian.Uint16(v[0:val.size])) / 5) - 500
						if verbose_mode {
							fmt.Printf("\tALTITUDE: %f M\n", record.Altitude)
						} //verbose_mode
						break
					case 3:
						temp, _ := binary.Uvarint(v[0:val.size])
						record.Heart_rate = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tHEART RATE: %d BPM\n", record.Heart_rate)
						} //verbose_mode
						break
					case 4:
						temp, _ := binary.Uvarint(v[0:val.size])
						record.Cadence = uint8(temp)
						if verbose_mode {
							fmt.Printf("\tCADENCE: %d RPM\n", record.Cadence)
						} //verbose_mode
						break
					case 5:
						record.Distance = float64(binary.LittleEndian.Uint32(v[0:val.size])) / 100
						if verbose_mode {
							fmt.Printf("\tDISTANCE: %f M\n", record.Distance)
						} //verbose_mode
						break
					case 6:
						record.Speed = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 1000
						if verbose_mode {
							fmt.Printf("\tSPEED: %f M/S\n", record.Speed)
						} //verbose_mode
						break
					case 7:
						record.Power = binary.LittleEndian.Uint16(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tPOWER: %d W\n", record.Power)
						} //verbose_mode
						break
					case 13:
						temp, _ := binary.Uvarint(v[0:val.size])
						record.Temperature = int8(temp)
						if verbose_mode {
							fmt.Printf("\tTEMP: %d°C\n", record.Temperature)
						} //verbose_mode
						break
						//not yet implemented but found in struct - see profile.xslx of fit sdk
						/*
							11 - Time_from_course          int32
							8  - Compressed_speed_distance uint8
							9  - Grade                     int16
							10 - Resistance                uint8
							12 - Cycle_length              uint8
						*/

					}

				}
				fitFile.Records = append(fitFile.Records, record)
				//def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize))
				break

			case 21: //event
				var event Event
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 253:
						event.Timestamp = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(event.Timestamp, 0)
						if verbose_mode {
							fmt.Printf("\tEVENT TIMESTAMP: %d (rectified) %v\n", event.Timestamp, t)
						} //verbose_mode
						break
					case 4:
						temp, _ := binary.Uvarint(v[0:1])
						event.Time_trigger = maps.Time_trigger(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tEVENT TIME TRIGGER: %s\n", event.Time_trigger)
						} //verbose_mode
						break
					case 1:
						temp, _ := binary.Uvarint(v[0:1])
						event.Event_type = maps.Event_type(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tEVENT TYPE: %s\n", event.Event_type)
						} //verbose_mode
						break
					case 0:
						temp, _ := binary.Uvarint(v[0:1])
						event.Event = maps.Event(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tEVENT: %s\n", event.Event)
						} //verbose_mode
					}

				}
				fitFile.Events = append(fitFile.Events, event)
				k = skip(k, uint64(sumRecordsDataSize))
				break

			case 23: //device info

				/*
					Timestamp          uint32
					Serial_number      uint32
					Cum_operating_time uint32
					Manufacturer       uint16
					Product            uint16
					Software_version   uint16
					Battery_voltage    uint16
					Device_index       uint8
					Device_type        uint8
					Hardware_version   uint8
					Battery_status     uint8*
				*/
				var device_info Device_info
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 253:
						device_info.Timestamp = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(device_info.Timestamp, 0)
						if verbose_mode {
							fmt.Printf("\tDEVICE INFO TIMESTAMP: %d (rectified) %v\n", device_info.Timestamp, t)
						} //verbose_mode
						break
					case 1:
						temp, _ := binary.Uvarint(v[0:1])
						device_info.Device_type = maps.Device_type(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tDEVICE TYPE: %s\n", device_info.Device_type)
						} //verbose_mode
						break
					case 2:
						device_info.Manufacturer = maps.Manufacturer(uint64(binary.LittleEndian.Uint16(v[0:val.size])))
						if verbose_mode {
							fmt.Printf("\tMANUFACTURER: %s\n", device_info.Manufacturer)
						} //verbose_mode
						break
					case 3:
						device_info.Serial_number = binary.LittleEndian.Uint32(v[0:val.size])
						if verbose_mode {
							fmt.Printf("\tSERIAL NUMBER: %d\n", device_info.Serial_number)
						} //verbose_mode
						break
					case 4:
						device_info.Product = maps.Product(uint64(binary.LittleEndian.Uint16(v[0:val.size])))
						if verbose_mode {
							fmt.Printf("\tPRODUCT: %s\n", device_info.Product)
						} //verbose_mode
						break
					case 5:
						device_info.Software_version = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 100
						if verbose_mode {
							fmt.Printf("\tSOFTWARE VERSION: %f\n", device_info.Software_version)
						} //verbose_mode
						break
					case 6:
						device_info.Hardware_version = v[0]
						if verbose_mode {
							fmt.Printf("\tHARDWARE VERSION: %d\n", device_info.Hardware_version)
						} //verbose_mode
					case 10:
						device_info.Battery_voltage = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 256
						if verbose_mode {
							fmt.Printf("\tBATTERY VOLTAGE: %f V\n", device_info.Battery_voltage)
						} //verbose_mode
						break
					case 22:
						temp, _ := binary.Uvarint(v[0:1])
						device_info.Ant_network = maps.Ant_network(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tANT NETWORK: %s\n", device_info.Ant_network)
						} //verbose_mode
						break
					case 25:
						temp, _ := binary.Uvarint(v[0:1])
						device_info.Source_type = maps.Source_type(uint64(temp))
						if verbose_mode {
							fmt.Printf("\tSOURCE TYPE: %s\n", device_info.Source_type)
						} //verbose_mode
						break
					}
				}
				fitFile.DeviceInfo = append(fitFile.DeviceInfo, device_info)
				k = skip(k, uint64(sumRecordsDataSize))

				break

			case 49: //file_creator
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 0:
						fitFile.FileCreator.Software_version = float64(binary.LittleEndian.Uint16(v[0:val.size])) / 100
						if verbose_mode {
							fmt.Printf("\tSOFTWARE VERSION: %f\n", binary.LittleEndian.Uint16(v[0:val.size]))
						} //verbose_mode
						break
					case 1:
						fitFile.FileCreator.Hardware_version = v[0]
						if verbose_mode {
							fmt.Printf("\tHARDWARE VERSION: %d\n", v[0])
						} //verbose_mode
					}
				}
				//def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize))
				break

			//TODO: add the remaining cases - the default allows the appropriate number bytes to be skipped for unknown data records
			default:
				if verbose_mode {
					fmt.Println("\t>> UNKNOWN RECORD")
				} //verbose_mode
				var sumRecordsDataSize int
				for _, val := range definition[uint64(localMsgType)].field_defs {
					sumRecordsDataSize += val.size
				}
				glob_msge_num_0_read = true
				def_message.global_message_number = 0   //reset
				k = skip(k, uint64(sumRecordsDataSize)) //move the reader to the end of the record data
			}
		}
	}
	return fitFile

}
func skip(iter, inc uint64) uint64 {
	iter += inc
	if verbose_mode {
		fmt.Printf("\nSKIPPING %s RECORD LENGTH OF %d BYTES >>\n=====================================================================\n", section, inc)
	} //verbose_mode
	return iter
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//function to calculate CRC
func calc_crc(char uint8) {
	crc_table := [16]uint16{
		0x0000, 0xCC01, 0xD801, 0x1400, 0xF001, 0x3C00, 0x2800, 0xE401,
		0xA001, 0x6C00, 0x7800, 0xB401, 0x5000, 0x9C01, 0x8801, 0x4400,
	}
	var tmp uint16

	tmp = crc_table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ crc_table[char&0xF]

	tmp = crc_table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ crc_table[(char>>4)&0xF]
}

//converts semicircles to degrees
func semicircles_to_degrees(semicircles float64) float64 {
	semicircles = semicircles * (180 / math.Pow(2, 31))
	return semicircles
}
