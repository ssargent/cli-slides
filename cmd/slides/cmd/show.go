/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Slide struct {
	Title   string
	Content string
}

func parseSlides(content string) []Slide {
	var slides []Slide
	lines := strings.Split(content, "\n")
	var currentTitle string
	var currentContent []string
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				if currentTitle != "" {
					slides = append(slides, Slide{Title: currentTitle, Content: strings.Join(currentContent, "\n")})
				}
				currentTitle = parts[1]
				currentContent = nil
			}
		} else {
			currentContent = append(currentContent, line)
		}
	}
	if currentTitle != "" {
		slides = append(slides, Slide{Title: currentTitle, Content: strings.Join(currentContent, "\n")})
	}
	return slides
}

func renderSlide(slide Slide, index int, total int) {
	fmt.Print("\033[2J\033[1;1H") // Clear screen
	fmt.Printf("Slide %d/%d\n\n", index+1, total)
	fmt.Println("=== " + slide.Title + " ===")
	fmt.Println()
	fmt.Println(slide.Content)
	fmt.Println()
	fmt.Println("Commands: 'n' next, 'p' previous, 'q' quit")
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [file]",
	Short: "Show slides from a markdown file",
	Long: `Render a markdown file as slides in the CLI. Each H1 header (#) starts a new slide.

Example:
  slides show presentation.md`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			return
		}
		content := string(data)
		slides := parseSlides(content)
		if len(slides) == 0 {
			fmt.Println("No slides found in the file.")
			return
		}
		currentIndex := 0
		for {
			renderSlide(slides[currentIndex], currentIndex, len(slides))
			var input string
			fmt.Scanln(&input)
			switch strings.TrimSpace(input) {
			case "n":
				if currentIndex < len(slides)-1 {
					currentIndex++
				}
			case "p":
				if currentIndex > 0 {
					currentIndex--
				}
			case "q":
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
