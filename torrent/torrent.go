// Copyright 2020~2022 xgfone
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

package torrent

import (
	"io"
	"os"
	"time"

	"github.com/xgfone/go-bt/bencode"
	"github.com/xgfone/go-bt/metainfo"
)

// CreateTorrentConfig is the configuration information to create a .torrent file.
type CreateTorrentConfig struct {
	PieceLength int64
	Name        string
	RootDir     string
	Output      string
	Comment     string
	Announces   []string
	WebSeeds    []string
	NoDate      bool
}

// CreateTorrent creates a .torrent file.
func CreateTorrent(config CreateTorrentConfig) error {
	info, err := metainfo.NewInfoFromFilePath(config.RootDir, config.PieceLength)
	if err != nil {
		return err
	}

	if config.Name != "" {
		info.Name = config.Name
	}

	var mi metainfo.MetaInfo
	mi.Comment = config.Comment
	mi.InfoBytes, err = bencode.EncodeBytes(info)
	if err != nil {
		return err
	}

	switch len(config.Announces) {
	case 0:
	case 1:
		mi.Announce = config.Announces[0]
	default:
		mi.AnnounceList = metainfo.AnnounceList{config.Announces}
	}

	for _, seed := range config.WebSeeds {
		mi.URLList = append(mi.URLList, seed)
	}

	if !config.NoDate {
		mi.CreationDate = time.Now().Unix()
	}

	var out io.WriteCloser = os.Stdout
	if config.Output != "" {
		out, err = os.OpenFile(config.Output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	return mi.Write(out)
}
