package nvms

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/tebruno99/nvms/nvdb"
	"github.com/tebruno99/nvms/scanner"
	"log"
)

type MediaManager struct {
	db *sql.DB
}

func NewMediaManager(db *sql.DB) *MediaManager {
	return &MediaManager{db: db}
}

func (mm *MediaManager) ScanLibrary(id int) error {
	rows, err := mm.db.Query("SELECT lp.id,path FROM library_path lp JOIN library l ON lp.library_id=l.id WHERE l.id = ?", id)
	if err != nil {
		return err
	}
	lp := make([]nvdb.LibraryPath, 0)

	for rows.Next() {
		var id int64
		var path string
		if err := rows.Scan(&id, &path); err == nil {
			lp = append(lp, nvdb.LibraryPath{ID: id, Path: path})
		}
	}
	rows.Close()

	for _, path := range lp {
		fs, err := scanner.NewLocalScanner(scanner.WithPath(path.Path), scanner.WithFilter(scanner.NewExtensionFilter(".mp4", ".mp3", ".m4a", ".m4v")))
		if err != nil {
			return err
		}
		resChan := fs.Scan(context.Background())

		for res := range resChan {
			mm.processFile(context.Background(), res, path.ID)
		}
	}

	return nil

}

func (mm *MediaManager) processFile(ctx context.Context, event scanner.MetaEvent, libraryPathID int64) {
	if event.Error != nil {
		log.Printf("file error %s: %s", event.Path, event.Error)
		return
	}
	var id int64
	var path string
	if err := mm.db.QueryRowContext(ctx, "SELECT id,path FROM media_item mi WHERE mi.path = ?", event.Path).Scan(id, path); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("db path lookup error %s: %s", event.Path, event.Error)
			return
		}
	}

	ffprobeInfo, err := scanner.FFprobeProbe(event.Path)
	if err != nil {
		log.Printf("ffprobe error %s: %s", event.Path, event.Error)
		return
	}
	mediaInfo, err := scanner.MediaInfo(event.Path)
	if err != nil {
		log.Printf("mediainfo error %s: %s", event.Path, event.Error)
		return
	}
	fmt.Printf("%v\n", mediaInfo)

	res, err := mm.db.ExecContext(ctx, "INSERT INTO media_item (path,library_path_id, mime_type,file_modified_at) VALUES (?, ?, ?, ?)", event.Path, libraryPathID, ffprobeInfo.Format.FormatName, event.ModifiedAt)
	if err != nil {
		log.Printf("db insert error %s: %s", event.Path, err)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Printf("db error lastInertId Error %s: %s", event.Path, err)
		return
	}
	for _, stream := range ffprobeInfo.Streams {
		_, err := mm.db.ExecContext(ctx, "INSERT INTO media_item_stream (media_item_id,stream_index,type,codec,height,width) VALUES (?, ?,?,?,?,?)",
			id, stream.Index, stream.CodecType, stream.CodecName, stream.Height, stream.Width)
		if err != nil {
			log.Printf("db insert stream error %s: %s", event.Path, err)
			continue
		}
	}

}
