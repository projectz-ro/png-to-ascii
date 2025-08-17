package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"png-to-ascii/utils"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

const chars = " .:-=+*#%@"

var brightMult = 3

var (
	scanner        *bufio.Scanner = bufio.NewScanner(os.Stdin)
	inputDir       string         = ""
	outputDir      string         = ""
	outputBaseName string         = ""
	width          int            = 100
	colored        bool           = true
)

func configure() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	utils.ClearScreen(3)
	inputDir = utils.PromptWithDefault(scanner, "Enter input directory", cwd)

	utils.ClearScreen(3)
	outputDir = utils.PromptWithDefault(scanner, "Enter output directory", filepath.Join(cwd, "ASCII"))

	utils.ClearScreen(3)
	outputBaseName = utils.PromptWithDefault(scanner, "Enter output file base name", "ascii_")

	utils.ClearScreen(3)
	widthStr := utils.PromptWithDefault(scanner, "Enter ASCII character width", "100")

	width, err = strconv.Atoi(widthStr)
	if err != nil {
		log.Fatalf("Invalid width: %v", err)
	}

	utils.ClearScreen(3)
	brightMultStr := utils.PromptWithDefault(scanner, "Enter brightness multiplier", "3")

	brightMult, err = strconv.Atoi(brightMultStr)
	if err != nil {
		log.Fatalf("Invalid brightness multiplier: %v", err)
	}

	utils.ClearScreen(3)
	coloredStr := utils.PromptWithDefault(scanner, "Output with ANSI colors? (Y/N)", "Y")

	if strings.ToLower(coloredStr) != "y" {
		colored = false
	}

	utils.ClearScreen(3)
	fmt.Printf("\nUsing settings:\n")
	fmt.Printf("Input Directory: %s\n", inputDir)
	fmt.Printf("Output Directory: %s\n", outputDir)
	fmt.Printf("Output Base Name: %s\n", outputBaseName)
	fmt.Printf("Width: %d\n", width)
	fmt.Printf("Brightness: %d\n", brightMult)
	fmt.Printf("Colored: %v\n\n", colored)
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		utils.ClearScreen(0)
		showHelp()
		os.Exit(0)
	}

	utils.ClearScreen(0)
	fmt.Println("Pick a directory with one or more PNGs inside")
	fmt.Println("to start coversions to ascii text files to an output directory.")
	fmt.Println("")

	for {
		configure()
		confirm := utils.PromptWithDefault(scanner, "Continue with these settings? (Y/N)", "Y")

		if strings.ToLower(confirm) == "y" {
			break
		}
	}

	utils.ClearScreen(3)
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		fmt.Printf("No files found in directory: %s\n", inputDir)
		return
	}

	index := 0
	processedCount := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".png" {
			fmt.Printf("Skipping %s (not a PNG file)\n", entry.Name())
			continue
		}

		fullpathInput := filepath.Join(inputDir, entry.Name())
		fmt.Printf("Processing: %s...", entry.Name())

		ascii, err := toASCII(fullpathInput, width)
		if err != nil {
			fmt.Printf(" ERROR: %v\n", err)
			continue
		}

		origName := strings.TrimSuffix(entry.Name(), ".png")
		fullpathOutput := filepath.Join(outputDir,
			fmt.Sprintf("%s%s.txt", outputBaseName, origName))
		err = os.WriteFile(fullpathOutput, []byte(ascii), 0644)
		if err != nil {
			fmt.Printf(" ERROR writing file: %v\n", err)
			continue
		}

		fmt.Printf(" ✓ -> %s\n", filepath.Base(fullpathOutput))
		index++
		processedCount++
	}

	fmt.Printf("\nProcessing complete! Converted %d PNG files to ASCII art.\n", processedCount)
}

func toASCII(filename string, newWidth int) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	img, err := png.Decode(f)
	if err != nil {
		return "", err
	}

	origBounds := img.Bounds()
	origWidth := origBounds.Dx()
	origHeight := origBounds.Dy()

	// Adjust height to maintain aspect
	newHeight := int(float64(origHeight) * float64(newWidth) / float64(origWidth))

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	var ascii strings.Builder
	for y := 0; y < newHeight-1; y += 2 {
		y2 := y + 1
		for x := range newWidth {
			c := dst.At(x, y)
			c2 := dst.At(x, y2)
			r, g, b, _ := c.RGBA()
			r2, g2, b2, _ := c2.RGBA()

			gray := toGrayscale(r, g, b)
			gray2 := toGrayscale(r2, g2, b2)
			avgGray := (gray + gray2) / 2

			idx := int(avgGray * float64(len(chars)-1))
			ch := string(chars[idx])
			if colored {
				rAvg := (r + r2) / 2
				gAvg := (g + g2) / 2
				bAvg := (b + b2) / 2

				ascii.WriteString(rgbToANSI(rAvg, gAvg, bAvg, ch))
			} else {
				ascii.WriteString(ch)
			}
		}
		ascii.WriteString("\n")
	}

	return ascii.String(), nil
}

func toGrayscale(r, g, b uint32) float64 {
	rf := float64(r) / 65535.0
	gf := float64(g) / 65535.0
	bf := float64(b) / 65535.0
	gray := (0.299*rf + 0.587*gf + 0.114*bf) * float64(brightMult)

	if gray > 1 {
		gray = 1
	} else if gray < 0 {
		gray = 0
	}
	return gray
}

func rgbToANSI(r, g, b uint32, ch string) string {
	r8 := r >> 8
	g8 := g >> 8
	b8 := b >> 8

	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r8, g8, b8, ch)
}
