# PNG to ASCII Converter

Convert PNG images into ASCII art for display in the terminal, with optional ANSI color output.

---

## Features

- Converts PNG images to ASCII art
- Optional ANSI color output
- Adjustable output width and brightness multiplier
- Batch conversion of all PNG files in a directory
- Interactive prompts for input/output settings

---

![Example Output](example/Zro_Logo.png)

```bash
Pick a directory with one or more PNGs inside
to start coversions to ascii text files to an output directory.


Using settings:
Input Directory: ./example
Output Directory: ./example/ASCII
Output Base Name: ascii_
Width: 60
Brightness: 3
Colored: true

Continue with these settings? (Y/N)
(default: Y): Y
```

                                        =========
                                    .=@@@@@@@@@@@@@@=
                                  =@@@@@@@@@@@@@@@@@@@@=
                                 =@@@@@@@@@@@@@@@@@@@@@@@
                                @@@@@@@@@@@@@@@@@@@@@@@@@@
                               @@@@@@@@@@@@@@@@@@@@@@@@@@@@
                               @@@@@@@@@@@@@@@@@@@@@@@@@@@@
                              @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
                              @@@@=@@@@@@@@@@@@@@@@@@@@@@@@@
                               @@   @@@@@@@@@@@@@    @@@@@@@
                               @@    @@@@===@@@@@    @@@@@@@
                               =@=   @@=     =@@@@  @@@@@@@@
                                @@=@*@@       =@@@@@@@@@@@@@
                                %@@@@@@        @@@@@@@@@@@@@
                                 @@@@@@        @@@@@@@@@@@@@
                                @@@@@@@   @*   %@@@@@@@@@@@@
                               =@@@@@@@   @@=  @@@@@@@@@@@@@
                             ==@@@@@@@@@@@@@@@@@@@@@@@@@@@@@=
                         ==+@@@@@@@@@@@@@@@@@@@@@@@@ @@@@@@@@
                     ==@@@@@@@@@@@@@@@@@@@@@@@@@@@==@@@@@@@@@
                   =@@@@@@@@@@@@@@@@@@@@@@@@@@@@==@@@@@@@@@@=
                  +@@@@@@@@@@@@@@@@@@@@@@@@@@@==@@@@=@@@@@@@
                 =@@@@=  =@@@@==   @@@@@+@@@@ @@@@@ @@@@@@@=
                 =@@@=  @@@==      @@@=   @@ @@@@@ @@@@@@@.
                                   ==     @= @@@@. @@@@@
                                            @@@@@ :@@@@
                                            @@@@   @@@=
                                            ===

---

## Installation

You need **Go 1.20+** installed.

```bash
git clone https://github.com/projectz-ro/png-to-ascii.git
cd png-to-ascii
make
```

This will produce an executable `png-to-ascii` in `png-to-ascii/build/` folder.

### Running the program:

#### Linux/macOS:

You can run it directly from the build folder:

```bash
./build/png-to-ascii
```

Or move it to a directory in your PATH to access from anywhere on
your system using just the name of the program:

```bash
sudo mv build/png-to-ascii /usr/local/bin/
```

then, whenever you want to run it..

```bash
png-to-ascii
```

#### Windows (PowerShell):

```powershell
.\build\png-to-ascii.exe
```

After this, the program is ready to use. You don’t need to modify anything else—just follow the interactive prompts to convert PNG images to ASCII art.---

## Usage

From the directory containing your PNG files, run:

```bash
png-to-ascii
```

Then follow the interactive prompts.

### Example session

```text
Enter input directory (default: .):
Enter output directory (default: ./ASCII):
Enter output file base name (default: ascii_):
Enter ASCII character width (default: 100):
Enter brightness multiplier (default: 3):
Output with ANSI colors? (Y/N) (default: Y):
```

Converted files will be saved as:

```
<output_dir>/<output_base_name><original_file>.txt
```

For example, if you convert `first.png` with default options:

```
./ASCII/ascii_first.txt
```

---

## License

MIT
