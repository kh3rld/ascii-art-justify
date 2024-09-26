# ASCII Art Justify

This project allows you to display ASCII art banners with customizable text alignment. You can align your text to the left, right, center, or justify it based on terminal size.

## Features

- **Text Alignment**: Use the `--align` flag to specify how you want your text aligned:
  - `center`
  - `left`
  - `right`
  - `justify`
  
- **Terminal Size Adaptation**: The output adapts to your terminal size. Reducing the terminal window will adjust the graphical representation accordingly.

- **Checksum Validation**: Each banner file is validated for integrity using SHA-256 checksums to ensure the files are not corrupted.

- **Multiple Banner Styles**: You can choose from different ASCII art styles such as `standard`, `shadow`, or `thinkertoy`.

## Usage

To run the program, use the following command:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Options

- **--align=<type>**: Specify the alignment type. Must be one of the following:
  - `left`
  - `right`
  - `center`
  - `justify`
  
- **STRING**: The text you want to display in ASCII art.

- **BANNER**: The banner style you want to use. Available options are:
  - `standard`
  - `shadow`
  - `thinkertoy`

### Examples

1. **Center Aligned Text with Standard Banner**
   ```bash
   go run . --align=center "Hello World" standard
   ```
   ![Center Aligned Example](/screenshots/center.png)
2. **Left Aligned Text with Shadow Banner**
   ```bash
   go run . --align=left "Hello There" shadow
   ```
   ![Left Aligned Example](/screenshots/left.png)
3. **Right Aligned Text with Thinkertoy Banner**
   ```bash
   go run . --align=right "Greetings" thinkertoy
   ```
   ![Right Aligned Example](/screenshots/right.png)
4. **Justified Text with Standard Banner**
   ```bash
   go run . --align=justify "How are you doing today?" standard
   ```
   ![Justified Example](/screenshots/justified.png)
5. **Single Argument without Alignment**
   ```bash
   go run . "Just a simple string"
   ```
   ![Single Argument Example](/screenshots/simple.png)
### Error Handling

If you provide an invalid alignment option or banner, the program will display a usage message:
```
Usage: go run . [OPTION] [STRING] [BANNER]
Example: go run . --align=right something standard
```

## Implementation

The project is structured into several packages:

- **`main`**: Entry point that handles command-line arguments and invokes other packages.
- **`reading`**: Reads banner files and validates them using checksums.
- **`printart`**: Handles the rendering of ASCII art based on alignment specifications.
- **`check`**: Implements checksum validation for the banner files.

## Testing

You can run tests using: Navigate to respective directories and run the following command.

```bash
go test -v
```

## Contributions

Feel free to submit issues or pull requests if you have suggestions for improvements or new features!
