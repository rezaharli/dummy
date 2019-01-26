package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kharism/dummy/dblayer"
	"github.com/kharism/dummy/models"
)

func NewCommandExecutor() CommandExecutor {
	executor := CommandExecutor{}
	executor["PARK"] = Park
	executor["LEAVE"] = Leave
	executor["EXIT"] = Exit
	executor["STATUS"] = Status
	executor["CREATE_PARKING_LOT"] = CreateParking
	executor["REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR"] = QueryRegByColor
	executor["SLOT_NUMBERS_FOR_CARS_WITH_COLOUR"] = QuerySlotByColor
	executor["SLOT_NUMBER_FOR_REGISTRATION_NUMBER"] = QuerySlotByReg
	executor["COLOR_FOR_REGISTRATION_NUMBER"] = QuerySlotByColor
	return executor
}

type CommandExecutor map[string]Executable

func (c *CommandExecutor) ExecuteCommand(str string) string {
	str = strings.Trim(str, "\n")
	parsed := strings.Fields(str)
	UprCase := strings.ToUpper(parsed[0])
	if _, ok := (*c)[UprCase]; !ok {
		return "Command Not Found"
	}
	if len(parsed) == 1 {
		return (*c)[UprCase](nil)
	} else {
		params := parsed[1:]
		return (*c)[UprCase](params)
	}
}

type Executable func(params []string) string

// this bellow is commmands implementing Executable
func QuerySlotByColor(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <color>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"
	}
	cFilter := &dblayer.Filter{Type: "eq", FieldName: "Color", FieldValue: params[0]}
	datas := CarStorage.Filter(cFilter)
	ss := []string{}
	if len(datas) == 0 {
		return "Not found"
	}
	for _, dd := range datas {
		ss = append(ss, strconv.Itoa(dd.(models.ParkedCar).Slot+1))
	}
	buff.WriteString(strings.Join(ss, ","))
	return buff.String()
}
func QueryRegByColor(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <Color>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"
	}
	cFilter := &dblayer.Filter{Type: "eq", FieldName: "Color", FieldValue: params[0]}
	datas := CarStorage.Filter(cFilter)
	ss := []string{}
	if len(datas) == 0 {
		return "Not found"
	}
	for _, dd := range datas {
		ss = append(ss, dd.(models.ParkedCar).RegNumber)
	}
	buff.WriteString(strings.Join(ss, ", "))
	return buff.String()
}
func QuerySlotByReg(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <reg>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"
	}
	cFilter := &dblayer.Filter{Type: "eq", FieldName: "RegNumber", FieldValue: params[0]}
	datas := CarStorage.Filter(cFilter)
	ss := []string{}
	if len(datas) == 0 {
		return "Not found"
	}
	for _, dd := range datas {
		ss = append(ss, strconv.Itoa(dd.(models.ParkedCar).Slot+1))
	}
	buff.WriteString(strings.Join(ss, ","))
	return buff.String()
}
func QueryColorByReg(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <RegNumber>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"
	}
	cFilter := &dblayer.Filter{Type: "eq", FieldName: "RegNumber", FieldValue: params[0]}
	datas := CarStorage.Filter(cFilter)
	if len(datas) == 0 {
		return "Not found"
	}
	ss := []string{}
	for _, dd := range datas {
		ss = append(ss, dd.(models.ParkedCar).Color)
	}
	buff.WriteString(strings.Join(ss, ","))
	return buff.String()
}
func Status(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	iter := CarStorage.GetIterator()
	buff.WriteString(fmt.Sprint("Slot no\tReg no\tColor\n"))
	for iter.HasNext() {
		v := iter.Next()
		uu := v.(models.ParkedCar)
		buff.WriteString(uu.String() + "\n")
	}
	return buff.String()
}
func CreateParking(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <capacity>"
	}
	capacity, err := strconv.Atoi(params[0])
	if err != nil {
		return "Not a Number"
	}
	CarStorage = dblayer.NewMemStorage(capacity)
	buff.WriteString(fmt.Sprintf("Created a parking lot with %d slots", capacity))
	return buff.String()
}
func Park(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 2 {
		return "Wrong parameter length, need : <reg number> <color>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"
	}
	car := models.ParkedCar{}
	idx := CarStorage.GetEmptyIndex()
	if idx == -1 {
		return "Sorry, parking lot is full"
	}
	car.Slot = idx
	car.Color = params[1]
	car.RegNumber = params[0]
	CarStorage.Save(car)
	buff.WriteString(fmt.Sprint("Allocated slot number:", idx+1))
	return buff.String()
}
func Leave(params []string) string {
	buff := bytes.NewBuffer([]byte{})
	if len(params) != 1 {
		return "Wrong parameter length, need : <slot id>"
	}
	if CarStorage == nil {
		return "Car Storage is not initialized yet"

	}
	capStr := params[0]
	idx, err := strconv.Atoi(capStr)
	if err != nil {
		return "Not a Number"

	}
	err = CarStorage.Delete(idx - 1)
	if err == nil {
		buff.WriteString(fmt.Sprint("Slot number ", idx, " is free"))
	} else {
		return err.Error()

	}
	return buff.String()
}
func Exit(params []string) string {
	os.Exit(0)
	return ""
}
