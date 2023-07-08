package main

import (
	"math"
	"time"

	. "github.com/alexellis/blinkt_go"
	"github.com/sixdouglas/suncalc"
)

func breakOut(colours []int) (int, int, int) {
	r := colours[0]
	g := colours[1]
	b := colours[2]
	return r, g, b
}

func getTwilightTimes(latitude, longitude float64) (startTwilight, endTwilight time.Time, err error) {
	// Get the current date
	date := time.Now().UTC()

	// Calculate the times for astronomical twilight
	times := suncalc.GetTimes(date, latitude, longitude)

	startTwilight = times["astronomicalDawn"]
	endTwilight = times["astronomicalDusk"]

	return startTwilight, endTwilight, nil
}

func getCurrentBrightness(latitude, longitude float64) (brightness float64, err error) {
	startTwilight, endTwilight, err := getTwilightTimes(latitude, longitude)
	if err != nil {
		return 0, err
	}

	// Get the current time
	now := time.Now().UTC()

	// Calculate the equidistant point between start and end twilight
	equidistant := startTwilight.Add(endTwilight.Sub(startTwilight) / 2)

	// Calculate the duration between the equidistant point and the current time
	duration := now.Sub(equidistant)

	// Convert the duration to hours
	hours := duration.Hours()

	// Calculate the brightness using a sine function
	brightness = math.Sin(2 * math.Pi * hours / 24)

	// Adjust the brightness range from [-1, 1] to [0, 1]
	brightness = (brightness + 1) / 2

	return brightness, nil
}

func main() {
	latitude := 41.5047
	longitude := -73.9696

	brightness, err := getCurrentBrightness(latitude, longitude)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	colours := [10][3]int{
		{0, 0, 0},       //0 black
		{139, 69, 19},   //1 brown
		{255, 0, 0},     //2 red
		{255, 69, 0},    //3 orange
		{255, 255, 0},   //4 yellow
		{0, 255, 0},     //5 green
		{0, 0, 255},     //6 blue
		{128, 0, 128},   //7 violet
		{255, 255, 100}, //8 grey
		{255, 255, 255}, //9 white
	}

	blinkt.Setup()
	Delay(100)

	r := 0
	g := 0
	b := 0

	t := time.Now()
	hour := t.Hour()
	minute := t.Minute()

	hourten := hour / 10
	hourunit := hour % 10
	minuteten := minute / 10
	minuteunit := minute % 10

	r, g, b = breakOut(colours[hourten][:])
	blinkt.SetPixel(0, r, g, b)
	blinkt.SetPixel(1, r, g, b)

	r, g, b = breakOut(colours[hourunit][:])
	blinkt.SetPixel(2, r, g, b)
	blinkt.SetPixel(3, r, g, b)

	r, g, b = breakOut(colours[minuteten][:])
	blinkt.SetPixel(4, r, g, b)
	blinkt.SetPixel(5, r, g, b)

	r, g, b = breakOut(colours[minuteunit][:])
	blinkt.SetPixel(6, r, g, b)
	blinkt.SetPixel(7, r, g, b)

  blinkt.Show()
  Delay(500)
}
