#!/bin/bash

#########################################################################
# Comic Downloader Script
#########################################################################
# 
# DESCRIPTION:
#   This script downloads random comics from various webcomic sources.
#   It supports multiple sources and can be easily extended to add more.
#
# USAGE:
#   ./comic_downloader.sh [--source xkcd|softer|existential|all]
#
# OPTIONS:
#   --source    Specify which comic source to use:
#               - xkcd: Download from XKCD
#               - softer: Download from A Softer World
#               - existential: Download from Existential Comics
#               - all: Randomly select one source (default)
#
# ADDING NEW SOURCES:
#   To add a new comic source:
#   1. Add the source name to the SOURCES array
#   2. Add the source to either PNG_SOURCES or JPG_SOURCES based on image type
#   3. Create a new parse_<source>() function with the parsing logic
#   4. Add the source to the case statement in download_comic()
#
# EXAMPLES:
#   ./comic_downloader.sh                  # Random source
#   ./comic_downloader.sh --source xkcd    # XKCD only
#   ./comic_downloader.sh --source softer  # A Softer World only
#
#########################################################################

# Define all available sources
SOURCES=("xkcd" "softer" "existential")

# Base pictures folder
BASE_PICTURES_FOLDER=$(xdg-user-dir PICTURES)

# Define extension for each source type
declare -A EXTENSIONS
# PNG sources
PNG_SOURCES=("xkcd" "existential")
# JPG sources
JPG_SOURCES=("softer")

# Automatically set extensions based on source lists
for source in "${PNG_SOURCES[@]}"; do
    EXTENSIONS["$source"]="png"
done

for source in "${JPG_SOURCES[@]}"; do
    EXTENSIONS["$source"]="jpg"
done

# Default is to download from all available sources
SOURCE="all"

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --source)
            SOURCE="$2"
            shift
            ;;
        *)
            echo "Unknown parameter: $1"
            echo "Usage: $0 [--source xkcd|softer|existential|all]"
            exit 1
            ;;
    esac
    shift
done

# Create necessary directories
for source in "${SOURCES[@]}"; do
    mkdir -p "$BASE_PICTURES_FOLDER/${source}_strips" 2>/dev/null
done

# Parser functions for each source

# XKCD parser
parse_xkcd() {
    local temp_file=$1
    
    # Download and parse
    wget -q "https://c.xkcd.com/random/comic/" -O "$temp_file"
    local title=$(grep "og:title" "$temp_file" | sed 's/.*content="\([^"]*\)".*/\1/')
    local imageurl=$(grep "og:image" "$temp_file" | sed 's/.*content="\([^"]*\)".*/\1/')
    
    echo "$title"
    echo "$imageurl"
}

# A Softer World parser
parse_softer() {
    local temp_file=$1
    
    # Generate a random ID and download
    local ID=$(shuf -i 1-1248 -n 1)  # As of 2023-03-11, there are 1248 comics and the project is not updated anymore
    wget -q "https://www.asofterworld.com/index.php?id=$ID" -O "$temp_file"
    
    # Parse
    local imageurl=$(grep jpg "$temp_file" | grep clean | awk -F"\"" '{print $2}')
    local title="ASofterWorld_$ID"
    
    echo "Comic #$ID"
    echo "$title"
    echo "$imageurl"
}

# Existential Comics parser
parse_existential() {
    local temp_file=$1
    
    # Download and parse
    curl -s -L "https://www.existentialcomics.com/comic/random" -o "$temp_file"
    local comic_num=$(grep -o "comic/[0-9]*" "$temp_file" | head -1 | cut -d'/' -f2)
    local imageurl=$(grep png "$temp_file" | grep comicImg | awk -F "//" '{print "https://"$2}' | sed 's/".*//')
    local title="ExistentialComics_$comic_num"
    
    echo "Comic #$comic_num"
    echo "$title"
    echo "$imageurl"
}

# Generic function to handle common operations
download_comic() {
    local source=$1
    local temp_file="/tmp/${source}_index.html"
    local folder="$BASE_PICTURES_FOLDER/${source}_strips"
    local extension="${EXTENSIONS[$source]}"
    local parse_result=""
    local title=""
    local imageurl=""
    
    echo "[+] Downloading random $source comic..."
    
    # Call the appropriate parser function
    case $source in
        "xkcd")
            parse_result=$(parse_xkcd "$temp_file")
            title=$(echo "$parse_result" | sed -n '1p')
            imageurl=$(echo "$parse_result" | sed -n '2p')
            echo "$title"
            ;;
            
        "softer")
            parse_result=$(parse_softer "$temp_file")
            echo "$(echo "$parse_result" | sed -n '1p')"
            title=$(echo "$parse_result" | sed -n '2p')
            imageurl=$(echo "$parse_result" | sed -n '3p')
            ;;
            
        "existential")
            parse_result=$(parse_existential "$temp_file")
            echo "$(echo "$parse_result" | sed -n '1p')"
            title=$(echo "$parse_result" | sed -n '2p')
            imageurl=$(echo "$parse_result" | sed -n '3p')
            ;;
    esac
    
    echo "$imageurl"
    
    echo "[+] Downloading the image..."
    wget -q "$imageurl" -O "$folder/$title.$extension"
    
    echo "[+] Saving the image to a file..."
    echo "[Ok] Done ($folder/$title.$extension)!"
    
    # Clean up
    rm -f "$temp_file"
    
    # Display if possible
    kitten icat "$folder/$title.$extension" 2>/dev/null
}

# Main execution based on source parameter
if [[ "$SOURCE" == "all" ]]; then
    # Choose one source randomly
    RANDOM_SOURCE=${SOURCES[$RANDOM % ${#SOURCES[@]}]}
    echo "Randomly selected source: $RANDOM_SOURCE"
    download_comic "$RANDOM_SOURCE"
elif [[ " ${SOURCES[*]} " =~ " $SOURCE " ]]; then
    # If source is valid, download from it
    download_comic "$SOURCE"
else
    echo "Invalid source: $SOURCE"
    echo "Valid options are: ${SOURCES[*]} or all"
    exit 1
fi 