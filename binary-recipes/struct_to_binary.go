package binary_recipes

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

/**
convert structs to binary so that they can be stored or sent somewhere else
* package used :- encoding/gob
* only limited to golang
* useful for sending small packets of data when high bandwidth is not available
*/

type Meter struct {
	Id        uint32
	Voltage   uint8
	Current   uint8
	Energy    uint32
	Timestamp uint64
}

func generateMeterReadingsSlice() []Meter {
	readings := make([]Meter, 1)
	for _ = range 10 {
		meterReading := Meter{
			1001,
			12,
			14,
			1009783,
			uint64(time.Now().UnixNano()),
		}
		readings = append(readings, meterReading)
	}
	return readings
}

func Demo_for_Gob_Encoding() {
	readings := generateMeterReadingsSlice()
	file, err := os.Create("reading")
	if err != nil {
		log.Println("could not eopened the file")
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(readings)
	if err != nil {
		log.Println("could not encode to the file")
	}
}
