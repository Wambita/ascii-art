## ASCII-Art Generator README
### Overview

The ASCII-Art program converts input text into a graphic representation using ASCII characters. It supports various fonts and special characters, including spaces, newlines, and punctuation. This tool is ideal for creating stylized text for use in command line applications, documentation, or simply for artistic purposes.

### Features

- Multiple Fonts: Supports 'shadow', 'standard', and 'thinkertoy' fonts.
- Special Character Handling: Processes and displays numbers, letters, spaces, special characters, and \n for newlines.


## Funtionality
### `OpenFile`

Reads ASCII art from a file specified by filename.
Returns a string representing the ASCII art.

### `DisplayArt`

Prints a string using ASCII art, where words are separated by spaces and lines are separated by \n.
Uses PrintWord to print individual words.
Returns an error if encountered during printing.

### `MapFile`

This go function maps the rune of printable ascii characters from `' '` to `~` that is 95 printable ascii characters to their corresponding art character lines from the art file that is represented as a string
It returns a 2D array of `[rune][]string`

### [`main()`](https://github.com/Dudimath/ascii_art/blob/main/programme/main.go)
Parses command-line flags -shadow and -thinkertoy to determine which ASCII art file to use.
Prints the ASCII art based on the input string and chosen ASCII art file.
This code uses command-line flags to select different ASCII art files (shadow.txt, thinkertoy.txt, or standard.txt) and prints the input string using ASCII art.

## Usage

To run the program, use the following commands:

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```
## Flags
- `shadow`: Use shadow.txt for ASCII art.
- `thinkertoy`: Use thinkertoy.txt for ASCII art.
```bash
go run . "Hello" -shadow
go run . "Hello" -thinkertoy


### Installation

Ensure that Go is installed on your system. You can install Go from the official Go website. Once Go is installed, you can clone and run the program directly from the source code.

    Clone the repository:

   ``` bash

    git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art
    cd ascii-art
```
### Usage

Run the program by passing the text string as an argument. You can also specify a font style optionally. If no font is specified, 'standard' is used by default.

**Basic Command:**

``` bash

go run . "Hello" | cat -e
```
Output:

```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$                                                  
```

**Command with Font Selection:**

``` bash

go run . "Hello" -shadow  | cat -e
```
Output :
```
                                 $
_|                _| _|          $
_|_|_|     _|_|   _| _|   _|_|   $
_|    _| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
```
```bash
go run . "hello" -thinkertoy | cat -e
```
```
                 $
o        o o     $
|        | |     $
O--o o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```
### Special Characters:

To handle special characters or avoid shell interpretation issues, especially in bash, enclose the input in single quotes:

```bash

go run . '!@#$%^&*()_+' | cat -e

```
Output: 
```
 _               _  _      _    _   __  /\               _       __ __                    $
| |    ____    _| || |_   | |  (_) / / |/\|   ___     /\| |/\   / / \ \              _    $
| |   / __ \  |_  __  _| / __)    / /        ( _ )    \ ` ' /  | |   | |           _| |_  $
| |  / / _` |  _| || |_  \__ \   / /         / _ \/\ |_     _| | |   | |          |_   _| $
|_| | | (_| | |_  __  _| (   /  / / _       | (_>  <  / , . \  | |   | |            |_|   $
(_)  \ \__,_|   |_||_|    |_|  /_/ (_)       \___/\/  \/|_|\/  | |   | |                  $
      \____/                                                    \_\ /_/   ______          $
                                                                         |______|         $
```
#### Handling Line Breaks

To incorporate line breaks in your ASCII art, use \n within the input text. For example:

```bash

go run . "Hello\nWorld" | cat -e

```

Output:
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
__          __                 _       _  $
\ \        / /                | |     | | $
 \ \  /\  / /    ___    _ __  | |   __| | $
  \ \/  \/ /    / _ \  | '__| | |  / _` | $
   \  /\  /    | (_) | | |    | | | (_| | $
    \/  \/      \___/  |_|    |_|  \__,_| $
                                          $
                                          $
```


```bash

go run . "Hello\n\nWorld" | cat -e

```

Output:
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
__          __                 _       _  $
\ \        / /                | |     | | $
 \ \  /\  / /    ___    _ __  | |   __| | $
  \ \/  \/ /    / _ \  | '__| | |  / _` | $
   \  /\  /    | (_) | | |    | | | (_| | $
    \/  \/      \___/  |_|    |_|  \__,_| $
                                          $
                                          $
```


```bash

go run . "Hello\n" | cat -e

```

Output:
```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```


```bash

go run . "\nWorld" | cat -e

```

Output
```
$
__          __                 _       _  $
\ \        / /                | |     | | $
 \ \  /\  / /    ___    _ __  | |   __| | $
  \ \/  \/ /    / _ \  | '__| | |  / _` | $
   \  /\  /    | (_) | | |    | | | (_| | $
    \/  \/      \___/  |_|    |_|  \__,_| $
                                          $
                                          $
```


```bash

go run . "\n" | cat -e

```

Output:
```
$
```


```bash
go run . "\n\n" | cat -e
```

Output:
```
$
$
```


```bash
go run . "Hello
> World" | cat -e
```
Output
```
 _              _   _                $
| |            | | | |               $
| |__     ___  | | | |   ___         $
|  _ \   / _ \ | | | |  / _ \        $
| | | | |  __/ | | | | | (_) |       $
|_| |_|  \___| |_| |_|  \___/        $
                                     $
                                     $
                           _       _  $
                          | |     | | $
__      __   ___    _ __  | |   __| | $
\ \ /\ / /  / _ \  | '__| | |  / _` | $
 \ V  V /  | (_) | | |    | | | (_| | $
  \_/\_/    \___/  |_|    |_|  \__,_| $
                                      $
                                      $
```


### Fonts / Banner Format
This project comes with three predefined ASCII fonts, located in banner files:

  - [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  - [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  - [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)


**- Each character is represented over 8 lines, providing a clear and sizable output.*
**- Characters are separated by a new line `\n`.*

### Development Notes

### Testing

Unit tests are highly recommended to ensure the integrity and functionality of the code. Test cases can be executed with:

```bash
cd toany folder

go test -v
 ```

### Allowed Packages

Only standard Go packages are allowed in this project. This restriction is to ensure that the utility remains lightweight and dependency-free.

### Contribution

Contributions are welcome. Please adhere to the existing coding standards and include unit tests for any new features or changes. Ensure to thoroughly test the code before pushing any updates.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

### Author
This project was build and maintained by  [shfana](https://learn.zone01kisumu.ke/git/shfana)