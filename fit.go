package fit

import (
	"encoding/binary"
	"fmt"
	"github.com/jezard/fit/maps"
	"os"
	"time"
)

type File_id struct { //message number: 0
	Serial_number uint32
	Time_created  int64 //not equal to the uint32 found in the fit file
	Manufacturer  uint16
	Product       uint16
	Number        uint16
	File_type     byte
}
type File_creator struct { //message number: 1
	Software_version uint16
	Hardware_version uint8
}
type Device_info struct { //message number: 23
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
	Battery_status     uint8
}
type Event struct { //message number: 21
	Timestamp    int64 //not equal to the uint32 found in the fit file
	Time_trigger string
	Event        string
	Event_type   string
}
type Events struct {
	Events []Event
}
type Record struct { //message number: 20
	Timestamp                 uint32
	Position_lat              int32
	Position_long             int32
	Distance                  uint32
	Time_from_course          int32
	Compressed_speed_distance uint8
	Heart_rate                uint8
	Altitude                  uint16
	Speed                     uint16
	Power                     uint16
	Grade                     int16
	Cadence                   uint8
	Registance                uint8
	Cycle_length              uint8
	Temperature               int8
}
type Lap struct { //message number: 19
	Timestamp           uint32
	Start_time          uint32
	Start_position_lat  int32
	Start_position_long int32
	End_position_lat    int32
	End_position_long   int32
	Total_elapsed_time  uint32
	Total_timer_time    uint32
	Total_distance      uint32
	Total_cycles        uint32
	Message_index       uint16
	Total_calories      uint16
	Total_fat_calories  uint16
	Avg_speed           uint16
	Max_speed           uint16
	Avg_power           uint16
	Max_power           uint16
	Total_ascent        uint16
	Total_descent       uint16
	//event (0-1-ENUM): lap (9)
	//event_type (1-1-ENUM): stop (1)
	Avg_heart_rate uint8
	Max_heart_rate uint8
	Avg_cadence    uint8
	Max_cadence    uint8
	//intensity (23-1-ENUM): active (0)
	//lap_trigger (24-1-ENUM): manual (0)
	//sport (25-1-ENUM): cycling (2)
	Event_group uint8
}
type session struct { //message number 18
	timestamp           uint32
	start_time          uint32
	start_position_lat  int32
	start_position_long int32
	total_elapsed_time  uint32
	total_timer_time    uint32
	total_distance      uint32
	total_cycles        uint32
	nec_lat             int32
	nec_long            int32
	swc_lat             int32
	swc_long            int32
	message_index       uint16
	total_calories      uint16
	total_fat_calories  uint16
	avg_speed           uint16
	max_speed           uint16
	avg_power           uint16
	max_power           uint16
	total_ascent        uint16
	total_descent       uint16
	first_lap_index     uint16
	num_laps            uint16
	/*  event (0-1-ENUM): session (8)
	    event_type (1-1-ENUM): stop (1)
	    sport (5-1-ENUM): cycling (2)
	    sub_sport (6-1-ENUM): indoor_cycling (6)*/
	avg_heart_rate        uint8
	max_heart_rate        uint8
	avg_cadence           uint8
	max_cadence           uint8
	total_training_effect uint8
	event_group           uint8
	/*  trigger (28-1-ENUM): activity_end (0)*/
}
type activity struct { //message number 34
	timestamp        uint32
	total_timer_time uint32
	local_timestamp  uint32
	num_sessions     uint16
	/*  type (2-1-ENUM): manual (0)
	    event (3-1-ENUM): activity (26)
	    event_type (4-1-ENUM): stop (1)*/
	event_group uint8
}
type FitFile struct {
	FileId      File_id
	FileCreator File_creator
	DeviceInfo  Device_info
	Event       Event
	Record      Record //these need to be array but []Record not working
	Lap         Lap
	//Session  session
	//Activity activity
}

var crc uint16

var count int

func Parse(filename string) {
	fmt.Printf("FUNCTION Parse() called: %v\n", time.Now())
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

	fmt.Printf("%d (Header) %d (crc) %d (file size) bytes\n", n1, n2, fl)

	/*****
	*
	* Get the Fit file header information
	*
	*****/

	headerSize, _ := binary.Uvarint(b1[0:1]) //convert 1 byte to uint64

	//Indicates the length of this file header including header size. Minimum size is 12. This may be increased in future to add additional optional information.
	fmt.Printf("Header Size: %v\n", headerSize)

	//Protocol version number as provided in SDK
	fmt.Printf("protocol v.: %v\n", b1[1:2])

	//Profile version number as provided in SDK
	pv := binary.LittleEndian.Uint16(b1[2:4])
	fmt.Printf("profile v. : %v\n", pv)

	//Length of the data records section in bytes (not including Header or CRC)
	dataSize := binary.LittleEndian.Uint32(b1[4:8])
	fmt.Printf("Data Size  : %v bytes\n", dataSize)

	//ASCII values for “.FIT”. A FIT binary file opened with a text editor will contain a readable “.FIT” in the first line.
	fmt.Printf("Ascii      : %s%s%s%s\n", b1[8:9], b1[9:10], b1[10:11], b1[11:12])

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
		fmt.Printf("Expected CRC: %x\n", _crc) //still need to check this against a calculated value
	} else {
		_crc := binary.LittleEndian.Uint16(b2[0:2])
		fmt.Printf("Expected CRC: %x\n", _crc) //still need to check this against a calculated value
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
	fmt.Printf("Calculated CRC: %x \n", crc)

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

			//print out debug info
			fmt.Printf("\n[POS: %8d] ", uint64(dataSize)-k-1)
			fmt.Print("DEFINITION MESSAGE, ") //--not  accurate yet...!
			fmt.Printf("VAL: %b", rHead[0])
			fmt.Printf(" LOCAL MESSAGE TYPE: %#x", localMsgType)                   //last 4 bits
			fmt.Printf(" GLOB MESSAGE NUM: %d", def_message.global_message_number) //only the aligned correctly definitions value is correct
			fmt.Printf(" FIELDS: %d", def_message.number_of_fields)                //only the aligned correctly definitions value is correct

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
				fmt.Printf("\nFDN: %v\nSIZE: %v\nBASE_TYPE: %v\nOFFSET %v\n",
					def_contents.field_definition_number,
					def_contents.size,
					def_contents.base_type,
					def_contents.offset)

				//we will need a means of temporarily storing the all the fields definition data so that it can be used to retrieve the record data later
				def_message.field_defs = append(def_message.field_defs, def_contents)

			}

			rc_length = uint64(5 + uint64(nof[0])*3) //combined length of fixed and varible record content field

			k = skip(k, rc_length) //move the pointer to the end of the field definition

			continue
		} else { //10000000 (Data Message)
			var compHeader bool

			//check for compressed header
			if rHead[0]>>7 == 1 { //is compressed header
				compHeader = true
			}

			//set vars dependant on header type
			if compHeader {
				fmt.Printf("\n[POS: %8d] DATA MESSAGE, COMPRESSED HEADER, VAL: %b, LOCAL MESSAGE TYPE: %#x (%d), TIME OFFSET %02ds, GLOB_MEG_NUM %d",
					uint64(dataSize)-k-1,
					rHead[0],
					rHead[0]&0x60>>5, //LMT is bits 5-6,
					rHead[0]&0x60>>5, //LMT is bits 5-6,
					rHead[0]&0x1F,    //time offset 0-4
					def_message.global_message_number)

				localMsgType = rHead[0] & 0x60 >> 5 //LMT is bits 5-6

			} else {
				localMsgType = rHead[0] & 0x1f //LMT is bits 0-3
			}

			fitFile.FileId.Serial_number = 123

			//process data
			fmt.Println("\n") //debug
			switch def_message.global_message_number {

			case 0: //file_id
				//TODO extract contents into our data stucture

				var sumRecordsDataSize int
				for _, val := range def_message.field_defs {
					sumRecordsDataSize += val.size

					if !glob_msge_num_0_read {
						v := make([]byte, val.size)
						r.ReadAt(v, int64(headerSize+k+val.offset+1))
						switch val.field_definition_number {
						case 0:
							fitFile.FileId.File_type = v[0]

							fmt.Printf("FILE TYPE: %d\n", v[0]) //debug
							break
						case 1:
							fitFile.FileId.Manufacturer = binary.LittleEndian.Uint16(v[0:val.size])

							fmt.Printf("MANUFACTURER: %d (%s)\n", binary.LittleEndian.Uint16(v[0:val.size]), maps.Manufacturer(uint16(binary.LittleEndian.Uint16(v[0:val.size])))) //debug
							break
						case 2:
							fitFile.FileId.Product = binary.LittleEndian.Uint16(v[0:val.size])

							fmt.Printf("PRODUCT: %d (%s)\n", binary.LittleEndian.Uint16(v[0:val.size]), maps.Product(uint16(binary.LittleEndian.Uint16(v[0:val.size])))) //debug
							break
						case 3:
							fitFile.FileId.Serial_number = binary.LittleEndian.Uint32(v[0:val.size])

							fmt.Printf("SERIAL NUMBER: %d\n", binary.LittleEndian.Uint32(v[0:val.size])) //debug
							break
						case 4:
							fitFile.FileId.Time_created = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)

							t := time.Unix(fitFile.FileId.Time_created, 0)

							fmt.Printf("TIME CREATED: %d (rectified) %v\n", fitFile.FileId.Time_created, t) //debug
							break
						case 5:
							fitFile.FileId.Number = binary.LittleEndian.Uint16(v[0:val.size])

							fmt.Printf("NUMBER: %d\n", binary.LittleEndian.Uint16(v[0:val.size])) //debug
						}
					}

				}
				glob_msge_num_0_read = true
				def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize)) //move the reader to the end of the record data
				break

			case 21: //event
				//TODO extract contents into our data stucture

				var sumRecordsDataSize int
				for _, val := range def_message.field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 253:
						fitFile.Event.Timestamp = int64(binary.LittleEndian.Uint32(v[0:val.size])) + 631065600 //need to add on unix timestamp for 31/12/1989 to get up to correct date (We can still get up to 2038)
						t := time.Unix(fitFile.Event.Timestamp, 0)                                             //debug
						fmt.Printf("EVENT TIMESTAMP: %d (rectified) %v\n", fitFile.Event.Timestamp, t)         //debug
						break
					case 4:
						temp, _ := binary.Uvarint(v[0:1])
						fitFile.Event.Time_trigger = maps.Time_trigger(uint64(temp))
						fmt.Printf("EVENT TIME TRIGGER: %s\n", fitFile.Event.Time_trigger) //debug
						break
					case 1:
						temp, _ := binary.Uvarint(v[0:1])
						fitFile.Event.Event_type = maps.Event_type(uint64(temp))
						fmt.Printf("EVENT TYPE: %s\n", fitFile.Event.Event_type) //debug
						break
					case 0:
						temp, _ := binary.Uvarint(v[0:1])
						fitFile.Event.Event = maps.Event(uint64(temp))
						fmt.Printf("EVENT: %s\n", fitFile.Event.Event) //debug
					}

				}
				def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize))
				break

			case 49: //file_creator
				//TODO extract contents into our data stucture

				var sumRecordsDataSize int
				for _, val := range def_message.field_defs {
					sumRecordsDataSize += val.size
					v := make([]byte, val.size)
					r.ReadAt(v, int64(headerSize+k+val.offset+1))
					switch val.field_definition_number {
					case 0:
						fitFile.FileCreator.Software_version = binary.LittleEndian.Uint16(v[0:val.size])
						fmt.Printf("SOFTWARE VERSION: %d\n", binary.LittleEndian.Uint16(v[0:val.size])) //debug
						break
					case 1:
						fitFile.FileCreator.Hardware_version = v[0]
						fmt.Printf("HARDWARE VERSION: %d\n", v[0]) //debug
					}
				}
				def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize))
				break

			//TODO: need to add the remaining cases
			default:
				var sumRecordsDataSize int
				for _, val := range def_message.field_defs {
					sumRecordsDataSize += val.size
				}
				glob_msge_num_0_read = true
				def_message.field_defs = nil
				k = skip(k, uint64(sumRecordsDataSize)) //move the reader to the end of the record data
			}
		}

	}
	fmt.Printf("\n\n%v", fitFile)

}
func skip(iter, inc uint64) uint64 {
	iter += inc
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
