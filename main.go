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

	// Read and list all WAV files in the specified folder
	files, err := os.ReadDir(*folderPtr)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	// Loop through the files in the folder
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".wav" {
			filePath := filepath.Join(*folderPtr, file.Name())

			// Open the WAV file in read-write mode
			wavFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			defer wavFile.Close()

			// Read the 20th and 21st bytes of the file
			buffer := make([]byte, 21)
			_, err = wavFile.ReadAt(buffer, 20) // Read starting from the 20th byte (0-indexed)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			// Combine the values of the 20th and 21st bytes
			value := (uint16(buffer[0]) << 8) | uint16(buffer[1])

			// If the value is 0001 and not showing all, skip the file
			if value == 0100 && !*showAllPtr {
				continue
			}

			// If overwrite flag is true, update the file with 01 00 value
			if *overwritePtr {
				_, err := wavFile.WriteAt([]byte{1, 0}, 20) // Write 0001 at the correct position
				if err != nil {
					fmt.Println("Error writing to file:", err)
					continue
				}
				fmt.Printf("Updated \"%s\"\n", file.Name())
			} else {
				// Print filename and combined value as a hexadecimal string with leading zeros
				fmt.Printf("HEX: \"%04X\", file: \"%s\"\n", value, filePath)
			}
		}
	}
}
