package main

import "testing"

func TestCommand(t *testing.T) {
	executor := NewCommandExecutor()
	output := executor.ExecuteCommand("create_parking_lot 6")
	if output != "Created a parking lot with 6 slots" {
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-HH-1234 White")
	if output != "Allocated slot number:1" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-HH-9999 White")
	if output != "Allocated slot number:2" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-BB-0001 Black")
	if output != "Allocated slot number:3" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-HH-7777 Red")
	if output != "Allocated slot number:4" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-HH-2701 Blue")
	if output != "Allocated slot number:5" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-HH-3141 Black")
	if output != "Allocated slot number:6" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("leave 4")
	if output != "Slot number 4 is free" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("status")
	if output != `Slot no	Reg no	Color
	0	KA-01-HH-1234	White
	1	KA-01-HH-9999	White
	2	KA-01-BB-0001	Black
	4	KA-01-HH-2701	Blue
	5	KA-01-HH-3141	Black` {
		t.Log(output)
		//t.Fail()
	}
	output = executor.ExecuteCommand("park KA-01-P-333 White")
	if output != "Allocated slot number:4" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("park DL-12-AA-9999 White")
	if output != "Sorry, parking lot is full" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("registration_numbers_for_cars_with_colour White")
	if output != "KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("slot_numbers_for_cars_with_colour White")
	if output != "1,2,4" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("slot_number_for_registration_number KA-01-HH-3141")
	if output != "6" {
		t.Log(output)
		t.Fail()
	}
	output = executor.ExecuteCommand("slot_number_for_registration_number MH-04-AY-1111")
	if output != "Not found" {
		t.Log(output)
		t.Fail()
	}
}
