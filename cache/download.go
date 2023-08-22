package cache

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yakiyo/tilde/utils"
	"github.com/Yakiyo/tilde/where"
	"github.com/charmbracelet/log"
)

const _ARCHIVE_URL = "https://tldr.sh/assets/tldr.zip"

// download archive
func Download() error {
	dir := where.Dir()
	if !utils.FsExists(dir) {
		log.Info("Creating root directory")
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(where.Zip())
	if err != nil {
		return err
	}
	// close and remove the temp file when done
	defer os.Remove(f.Name())
	defer f.Close()
	log.Info("Downloading archive file")
	// download the archive
	err = downloadArchive(f)
	if err != nil {
		return err
	}

	if utils.FsExists(where.Cache()) {
		log.Info("Moving legacy cache directory temporarily to `_cache`")
		// remove the old cache dir to `_cache`
		oldCache, err := swap(dir)
		if err != nil {
			return err
		}
		// remove it in the end
		defer func() {
			log.Info("Removing legacy cache")
			os.RemoveAll(oldCache)
		}()
	}

	log.Info("Extracting files")
	// unzip it
	err = unzip(f)
	if err != nil {
		return err
	}

	return nil
}

// the function that actually does the file downloading
// only download and save file to archive
func downloadArchive(file *os.File) error {
	resp, err := http.Get(_ARCHIVE_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// rename the old `caches` directory to `_cache`
func swap(dir string) (string, error) {
	newDir := filepath.Join(dir, "_cache")
	if utils.FsExists(newDir) {
		err := os.RemoveAll(newDir)
		if err != nil {
			return "", err
		}
	}
	return newDir, os.Rename(where.Cache(), newDir)
}

// this function does the job of unzipping the archive
func unzip(file *os.File) error {
	r, err := zip.OpenReader(file.Name())
	if err != nil {
		return err
	}
	dest := where.Cache()
	for _, f := range r.File {
		if err := extract(f, dest); err != nil {
			return err
		}
	}
	return nil
}

// extract individual files
func extract(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	path := filepath.Join(dest, f.Name)

	// if its a dir, make the directory and exit early
	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
		return nil
	}
	// otherwise we move to writing the file

	os.MkdirAll(filepath.Dir(path), f.Mode())
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, rc)
	if err != nil {
		return err
	}
	log.Debug("Extracting file", "file", f.Name)
	return nil
}
