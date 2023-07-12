package main

import (
	"log"
	"math"
	"time"

	. "github.com/alexellis/blinkt_go"
	"github.com/sixdouglas/suncalc"
)

func breakOut(colors []int) (int, int, int) {
	r := colors[0]
	g := colors[1]
	b := colors[2]
	return r, g, b
}

func getTwilightTimes() (time.Time, time.Time, error) {
	latitude := 41.5047
	longitude := -73.9696

	date := time.Now()
	times := suncalc.GetTimes(date, latitude, longitude)

	end := times[suncalc.Sunset].Value.Local()
	start := times[suncalc.Sunrise].Value.Local()

	log.Printf("twilight: %s -> %s", start, end)

	return start, end, nil
}

func getCurrentBrightness() (float64, error) {
	start, end, err := getTwilightTimes()
	if err != nil {
		return 0, err
	}

	// Get the current time
	now := time.Now()

	// Calculate the equidistant point between start and end twilight
	equidistant := end.Add(start.Sub(end) / 2)

	// Calculate the duration between the equidistant point and the current time
	duration := now.Sub(equidistant)

	// Convert the duration to hours
	sec := duration.Seconds()
	log.Printf("now: %s -> %s", now, duration)

	// Calculate the brightness using a sine function
	brightness := math.Sin(2 * math.Pi * sec / 86400)

	// Adjust the brightness range from [-1, 1] to [0, 0.5]
	brightness = (brightness + 1) / 4

	return brightness, nil
}

func main() {
	brightness, err := getCurrentBrightness()
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(false)

	colors := [10][3]int{
		{0, 0, 0},       // 0 black
		{139, 69, 19},   // 1 brown
		{255, 0, 0},     // 2 red
		{255, 69, 0},    // 3 orange
		{255, 255, 0},   // 4 yellow
		{0, 255, 0},     // 5 green
		{0, 0, 255},     // 6 blue
		{128, 0, 128},   // 7 violet
		{255, 255, 100}, // 8 grey
		{255, 255, 255}, // 9 white
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

	r, g, b = breakOut(colors[hourten][:])
	blinkt.SetPixel(0, r, g, b)
	blinkt.SetPixel(1, r, g, b)

	r, g, b = breakOut(colors[hourunit][:])
	blinkt.SetPixel(2, r, g, b)
	blinkt.SetPixel(3, r, g, b)

	r, g, b = breakOut(colors[minuteten][:])
	blinkt.SetPixel(4, r, g, b)
	blinkt.SetPixel(5, r, g, b)

	r, g, b = breakOut(colors[minuteunit][:])
	blinkt.SetPixel(6, r, g, b)
	blinkt.SetPixel(7, r, g, b)

	log.Printf("brightness: %f", brightness)

	blinkt.Show()
	Delay(500)
}
