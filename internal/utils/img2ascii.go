package utils

import (
	"bytes"
	"image"
	"io"
	"net/http"
	"strings"

	_ "image/png"

	"github.com/charmbracelet/lipgloss"
	"github.com/nfnt/resize"
)

func ConvertToAscii(url string, res uint) (string, error) {
	var output strings.Builder
	resp, err := http.Get(url)
	if err != nil {
		return "error", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error", err
	}
	
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "error", err
	}
	newImg := resize.Resize(res,res, img,resize.Lanczos2)
	
	bounds := newImg.Bounds()

	style := lipgloss.NewStyle()
	cache := make(map[string]string)

	for y:=bounds.Min.Y; y<bounds.Max.Y;y+=2 {
		for x:=bounds.Min.X; x<bounds.Max.X;x++ {
			
			r1, g1, b1, _ := newImg.At(x,y).RGBA()
			r2,g2,b2:=r1,g1,b1
			
			if y+1 < bounds.Max.Y {
				r2,g2,b2,_ = newImg.At(x,y+1).RGBA()
			}

			key:= rgbToHex(r1,g1,b1) + rgbToHex(r2,g2,b2)
			if val, ok := cache[key]; ok {
				output.WriteString(val)
				continue
			}
			block := style.
				Foreground(lipgloss.Color(rgbToHex(r1, g1, b1))).
				Background(lipgloss.Color(rgbToHex(r2, g2, b2))).
				Render("▀")
			cache[key] = block
			output.WriteString(block)
		}
		output.WriteString("\n")
	}	
	return output.String(), nil
}

func rgbToHex(r, g, b uint32) string {
	return "#" + hex8(r>>8) + hex8(g>>8) + hex8(b>>8)
	
}

func hex8(v uint32) string {
	const hex = "0123456789ABCDEF"
	return string([]byte{hex[v>>4], hex[v&0xF]})
}
