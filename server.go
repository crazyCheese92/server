package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"github.com/nsf/termbox-go"
)

type coord struct {
	x int
	y int
}

func setCoord()coord {
	x = rand.Int31n(15)
	y = rand.Int31n(10)
	var c coord
	c.x, c.y = x, y
	return c
}

func main() {
	c = setCoord()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(c.x, c.y, '*', termbox.ColorRed, termbox.ColorBlue)
	termbox.Flush()
	
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.Littleendian, c)
	
	_, err = conn.write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
	
	termbox.Close()
}
