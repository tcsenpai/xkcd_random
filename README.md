# Comic Downloader

This is a simple script to download random comics from various webcomic sources and, if `kitten icat` is installed, display the comic in the terminal.

It saves the comics in your Pictures directory in subfolders named after each source (e.g., `xkcd_strips`, `softer_strips`, etc.).

![Screenshot](screenshot.png)

## Supported Sources

Currently, the script supports the following webcomic sources:

- **XKCD** - The classic webcomic about romance, sarcasm, math, and language
- **A Softer World** - Emily Horne and Joey Comeau's photographic comic strip
- **Existential Comics** - A comic about philosophy and philosophers

## Usage

```
./comic_downloader.sh [--source xkcd|softer|existential|all]
```

Options:

- `--source`: Specify which comic source to use (default: all)
  - `xkcd`: Download from XKCD
  - `softer`: Download from A Softer World
  - `existential`: Download from Existential Comics
  - `all`: Randomly select one source

## Installation

Just clone the repository and run the script.

```
git clone https://github.com/yourusername/comic_downloader.git
cd comic_downloader
chmod +x comic_downloader.sh
./comic_downloader.sh
```

## Details

This is a simple bash script that:

1. Downloads a random comic from the specified source
2. Saves it to your Pictures directory
3. Displays it in the terminal if you have `kitten icat` installed

The script is designed to be easily extensible, so you can add more comic sources in the future.

## License

This project is licensed under the [The Unlicense](LICENSE).
