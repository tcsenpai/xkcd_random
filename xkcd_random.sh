#!/bin/bash

BASE_PICTURES_FOLDER=$(xdg-user-dir PICTURES)
PICTURES_FOLDER="$BASE_PICTURES_FOLDER/xkcd_strips"

mkdir $PICTURES_FOLDER 2>/dev/null

echo "[+] Downloading random xkcd comic..."
# Use curl to fetch the random comic page
wget -q "https://c.xkcd.com/random/comic/" -O /tmp/comics_index.html

echo "[*] Parsing the comic..."
# Extract title and image URL using grep and sed
title=$(grep "og:title" /tmp/comics_index.html | sed 's/.*content="\([^"]*\)".*/\1/')
imageurl=$(grep "og:image" /tmp/comics_index.html | sed 's/.*content="\([^"]*\)".*/\1/')

echo "$title"
echo "$imageurl"

echo "[+] Downloading the image..."
# Download the image using curl
curl -s "$imageurl" -o "$PICTURES_FOLDER/$title.png"

echo "[+] Saving the image to a file..."
echo "[Ok] Done ($PICTURES_FOLDER/$title.png)!"

rm -rf /tmp/comics_index.html

# If `kitten icat` exists, we execute it
kitten icat "$PICTURES_FOLDER/$title.png" 2>/dev/null

