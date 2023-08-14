package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define a flag for the folder path
	folderPtr := flag.String("folder", ".", "Path to the folder containing WAV files")
	overwritePtr := flag.Bool("overwrite", false, "Overwrite files with non-0100 20-21 bytes value, i.e. fix the main issue")
	showAllPtr := flag.Bool("list", false, "Show all files and 20-21 bytes value")

	// Parse the command-line flags
	flag.Parse()

	// Display help if -h or -help flags are provided
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	}

	// Recursively search for all WAV files in the specified folder
	err := filepath.WalkDir(*folderPtr, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".wav" {
			// Open the WAV file in read-write mode
			wavFile, err := os.OpenFile(path, os.O_RDWR, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return nil // Continue to next file
			}
			defer wavFile.Close()

			// Read the 20th and 21st bytes of the file
			buffer := make([]byte, 21)
			_, err = wavFile.ReadAt(buffer, 20) // Read starting from the 20th byte (0-indexed)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return nil // Continue to next file
			}

			// Combine the values of the 20th and 21st bytes
			value := (uint16(buffer[0]) << 8) | uint16(buffer[1])

			// Convert value to hex "0100"
			newHexValue := []byte{1, 0}

			// If overwrite flag is true, update the file with "0100" value
			if *overwritePtr && value != 256 {
				_, err := wavFile.WriteAt(newHexValue, 20) // Write "0100" at the correct position
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return nil // Continue to next file
				}
				fmt.Printf("Updated \"%s\"\n", path)
			} else if *showAllPtr || value != 256 {
				// Print value and full path of the filename as a hexadecimal string
				fmt.Printf("HEX: \"%04X\", file: \"%s\"\n", value, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}
