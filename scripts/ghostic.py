#!/usr/bin/env python3

import os
import sys
import base64
import argparse
from urllib.request import urlopen
from urllib.error import HTTPError, URLError

def download_and_decode_base64(url: str, output_files: list[str]) -> None:
    """
    Downloads base64-encoded content from a URL, decodes it, and saves to a file.

    Args:
        url: The URL containing base64-encoded data as plain text.
        output_file: Path to save the decoded binary file.

    Raises:
        Various exceptions on errors (network, decoding, etc.).
    """
    try:
        with urlopen(url) as response:
            if response.status != 200:
                raise HTTPError(
                    url, response.status, response.reason, response.headers, None
                )

            # Read the response as text (base64 is ASCII-safe)
            base64_content = response.read().decode("utf-8").strip()

        # Decode the base64 string to bytes
        try:
            decoded_bytes = base64.b64decode(base64_content)
        except base64.binascii.Error as e:
            print(f"Base64 decoding error: {e}", file=sys.stderr)
            print("The content may not be valid base64 or may contain whitespace/issues.", file=sys.stderr)
            sys.exit(1)

        for x in output_files:
            if os.path.exists(x):
                print(f"Error: {x} already exists, skipping.")
                continue
            # Write the decoded binary data to the output file
            with open(x, "wb") as f:
                f.write(decoded_bytes)
            print(f"Successfully decoded and saved file to: {x}")
            print(f"   Size: {len(decoded_bytes)} bytes")
            continue

    except HTTPError as e:
        print(f"HTTP Error: {e.code} {e.reason} for URL: {url}", file=sys.stderr)
        sys.exit(1)
    except URLError as e:
        print(f"URL Error: {e.reason} for URL: {url}", file=sys.stderr)
        sys.exit(1)
    except UnicodeDecodeError:
        print("Error: Response could not be decoded as UTF-8 text (expected base64 string).", file=sys.stderr)
        sys.exit(1)
    except OSError as e:
        print(f"File writing error: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    data_url = "https://github.com/joshland/local_policy/raw/refs/heads/master/scripts/data/ghostty"
    #outfiles = [ "ghostty", "xterm-ghostty" ]
    outfiles = [ "/usr/share/terminfo/g/ghostty", "/usr/share/terminfo/x/xterm-ghostty" ]
    #parser = argparse.ArgumentParser(
    #    description="Download base64-encoded content from a URL and decode it to a binary file."
    #)
    #parser.add_argument("url", help="URL containing base64-encoded data (as plain text).")
    #parser.add_argument(
    #    "output",
    #    help="Output file path to save the decoded binary data (e.g., image.png, document.pdf).",
    #)
    #args = parser.parse_args()

    download_and_decode_base64(data_url, outfiles)

