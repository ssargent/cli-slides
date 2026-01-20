# CLI Slides

A command-line presentation tool that renders markdown files as interactive slides in the terminal.

## Features

- Parse markdown files and convert headings (#, ##, ###) into slides
- Interactive CLI navigation with keyboard controls
- Clear screen rendering for each slide
- Slide numbering and progress display
- Simple and lightweight, no external dependencies beyond Go standard library

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ssargent/cli-slides.git
   cd cli-slides
   ```

2. Build the application:
   ```bash
   go build ./cmd/slides
   ```

## Usage

Create a markdown file with your presentation content. Use headings to define slides:

```markdown
# Introduction

Welcome to CLI Slides!

# Features

- Interactive navigation
- Markdown support

## Subsections

You can use ## and ### for subsections too.

# Conclusion

Thank you!
```

Run the presentation:

```bash
./slides show your-presentation.md
```

### Navigation

- `n` - Next slide
- `p` - Previous slide
- `q` - Quit presentation

## Requirements

- Go 1.16 or later
- Terminal that supports ANSI escape codes

## Project Structure

- `cmd/slides/` - Main application code
- `cmd/slides/cmd/` - Cobra CLI commands
- `test_slides.md` - Example presentation file

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.