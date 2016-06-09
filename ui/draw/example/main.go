package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/egonelbre/exp/qpc"

	"github.com/egonelbre/spector/ui/draw"
	"github.com/egonelbre/spector/ui/draw/render-gl21"

	"net/http"
	_ "net/http/pprof"
)

func init() { runtime.LockOSThread() }

func main() {
	flag.Parse()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	go func() {
		for {
			runtime.GC()
			time.Sleep(1)
		}
	}()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False) // do not steal focus
	glfw.WindowHint(glfw.Samples, 4)

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(800, 600, "Spector", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.Restore()
	window.SetPos(32, 64)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	if err := gl.GetError(); err != 0 {
		fmt.Println("INIT", err)
	}

	startnano := time.Now().UnixNano()

	DrawList := draw.NewList()
	for !window.ShouldClose() {
		start := qpc.Now()
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			return
		}

		now := float64(time.Now().UnixNano()-startnano) / 1e9
		width, height := window.GetSize()

		{ // reset window
			gl.MatrixMode(gl.MODELVIEW)
			gl.LoadIdentity()

			gl.Viewport(0, 0, int32(width), int32(height))
			gl.Ortho(0, float64(width), float64(height), 0, 30, -30)
			gl.ClearColor(1, 1, 1, 1)
			gl.Clear(gl.COLOR_BUFFER_BIT)
		}

		DrawList.Reset()

		DrawList.AddRectFill(&draw.Rectangle{
			draw.Vector{10, 10},
			draw.Vector{50, 50},
		}, draw.Red)

		CircleRadius := float32(50.0 * math.Sin(now*1.3))
		DrawList.AddCircle(
			draw.Vector{100, 100}, CircleRadius, draw.Red)
		DrawList.AddArc(
			draw.Vector{200, 100}, CircleRadius/2+50,
			float32(now),
			float32(math.Sin(now)*10), draw.Blue)

		LineWidth := float32(math.Sin(now*2.1)*5 + 5)

		LineCount := int(width / 8)
		line := make([]draw.Vector, LineCount)
		for i := range line {
			r := float64(i) / float64(LineCount-1)
			line[i].X = float32(r) * float32(width)
			line[i].Y = float32(height)*0.5 + float32(math.Sin(r*11.8+now)*100)
		}
		DrawList.AddLine(line[:], LineWidth, draw.Blue)

		CircleCount := int(width / 8)
		circle := make([]draw.Vector, CircleCount)
		for i := range circle {
			p := float64(i) / float64(CircleCount)
			a := now + p*math.Pi*2
			w := math.Sin(p*62)*20.0 + 100.0
			circle[i].X = float32(width)*0.5 + float32(math.Cos(a)*w)
			circle[i].Y = float32(height)*0.5 + float32(math.Sin(a)*w)
		}

		// DrawList.PushClip(draw.Rect(0, 0, float32(width)/2, float32(height)/2))
		DrawList.AddClosedLine(circle[:], LineWidth, draw.Green)
		// DrawList.PopClip()

		render.List(width, height, DrawList)
		if err := gl.GetError(); err != 0 {
			fmt.Println(err)
		}
		stop := qpc.Now()

		window.SwapBuffers()
		runtime.GC()
		glfw.PollEvents()

		fmt.Printf("%-10.3f\n", stop.Sub(start).Duration().Seconds()*1000)
	}

}
