// Copyright 2011 Andrew Scherkus
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package id3

import (
	"bufio"
	"encoding/binary"
)

var ID3v23Tags = map[string]string{
	"TALB": "album",
	"TPE1": "artist",
	"COMM": "comments",
	"TCOM": "composer",
	"TCOP": "copyright",
	"TPOS": "disc",
	"TENC": "encodedby",
	"TSSE": "encoder",
	"TCON": "genre",
	"TLAN": "language",
	"TLEN": "length",
	"TMED": "media",
	"TOPE": "originalartist",
	"TIT2": "title",
	"TRCK": "track",
	"TEXT": "writer",
	"TYER": "year",
}

// ID3 v2.3 doesn't use sync-safe frame sizes: read in as a regular big endian number.
func parseID3v23FrameSize(reader *bufio.Reader) (int, error) {
	var size int32
	binary.Read(reader, binary.BigEndian, &size)
	return int(size), nil
}
