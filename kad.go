package main

import (
	"encoding/json"
	"log"

	"github.com/remiberthoz/kad"
)

func main() {
	// you can define settings and the layout in JSON
	json_bytes := []byte(`{
		"switch-type":1,
		"stab-type":2,
		"layout":[
			["Esc",{"x":0.5},"F1","F2","F3","F4",{"x":0.5},"F5","F6","F7","F8",{"x":0.5},"F9","F10","F11","F12",{"x":1,"a":7},""],
			[{"y":0.25,"a":4},"~","!\n1","@\n2","#\n3","$\n4","%\n5","^\n6","&\n7","*\n8","(\n9",")\n0","_\n-","+\n=",{"w":2},"Backspace"],
			[{"y":-0.5,"x":15.5},"Delete"],
			[{"y":-0.5,"w":1.5},"Tab","Q","W","E","R","T","Y","U","I","O","P","{\n[","}\n]",{"w":1.5},"|\n\\"],
			[{"y":-0.5,"x":15.5},"Home"],
			[{"y":-0.5,"w":1.75},"Caps Lock","A","S","D","F","G","H","J","K","L",":\n;","\"\n'",{"w":2.25},"Enter"],
			[{"w":2.25},"Shift","Z","X","C","V","B","N","M","<\n,",">\n.","?\n/",{"w":2.25},"Shift",{"a":7},"↑"],
			["Fn","Ctrl","Alt",{"w":1.25},"Win",{"w":6.25},"",{"w":1.25},"Alt",{"w":1.25},"Ctrl",{"x":0.5},"←","↓","→"]
		],
		"case": {
			"case-type":"sandwich",
			"usb-location":128.5,
			"usb-width": 20
		},
		"top-padding":5,
		"left-padding":15,
		"right-padding":15,
		"bottom-padding":5,
		"kerf":0.2,
		"custom": [
			{
				"layers": [
					"top",
					"switch",
					"open",
					"closed"
				],
				"op": "cut",
				"polygon": "custom-circle",
				"diameter": 4,
				"rel_to": "[0,0]",
				"points": "[-x+7.5,-y+7.5];[-x+7.5,y-7.5];[x-7.5,-y+7.5];[x-7.5,y-7.5]"
			},
			{
				"layers": [
					"bottom"
				],
				"op": "cut",
				"polygon": "custom-hexagon",
				"radius": 4,
				"angle": 180,
				"rel_to": "[0,0]",
				"points": "[-x+7.5,-y+7.5];[-x+7.5,y-7.5];[x-7.5,-y+7.5];[x-7.5,y-7.5]"
			},
			{
				"layers": [
					"switch"
				],
				"op": "cut",
				"polygon": "custom-circle",
				"diameter": 2,
				"rel_to": "[0,0]",
				"points": "[-52.5,-16.5];[-57,21.5];[47.5,21.5];[43,-16.5]"
			},
			{
				"layers": [
					"bottom"
				],
				"op": "cut",
				"polygon": "custom-hexagon",
				"radius": 2,
				"rel_to": "[0,0]",
				"points": "[-52.5,-16.5];[-57,21.5];[47.5,21.5];[43,-16.5]"
			}
		]
	}`)

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err := json.Unmarshal(json_bytes, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = "Black_and_Yellow"      // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./"        // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// here are some more settings defined for this case
	//cad.Case.UsbWidth = 12 // all dimension are in 'mm'
	cad.Fillet = 5         // 3mm radius on the rounded corners of the case

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}