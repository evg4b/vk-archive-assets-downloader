package loader

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"github.com/evg4b/vk-archive-assets-downloader/contract"
	"github.com/evg4b/vk-archive-assets-downloader/utils"
)

func (p *Loader) downloadFile(logger *log.Logger, attachemt contract.Attachemt) (downloadError error) {
	defer utils.PanicInterceptor(&downloadError)
	defer p.attachemtPb.Increment()

	logger.Printf("started downloading %s:%s %s", attachemt.DialogName, attachemt.Type, attachemt.Url)
	dir, filename, err := getFilePath(p.dest, attachemt)
	if err != nil {
		return fmt.Errorf("failed create path to attachemt %s: %v", attachemt.Url, err)
	}

	if attachemt.IsLink {
		file := filepath.Join(dir, filename)
		logger.Printf("Created link file %s to url %s\n", file, attachemt.Url)
		return makeLinkFile(file, attachemt.Url)
	}

	path, err := downloadToFile(logger, dir, filename, attachemt.Url)
	if err != nil {
		return fmt.Errorf("failed download attachemt %s: %v", attachemt.Url, err)
	}

	logger.Printf("attachemt  %s:%s %s downloaded to file %s", attachemt.DialogName, attachemt.Type, attachemt.Url, path)

	return nil
}

func downloadToFile(logger *log.Logger, dir, filename, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if len(filepath.Ext(filename)) == 0 {
			contentType := getContentType(resp)
			if len(contentType) > 0 {
				extentions, err := mime.ExtensionsByType(contentType)
				if err == nil && len(extentions) > 0 {
					filename = fmt.Sprintf("%s%s", filename, extentions[len(extentions)-1])
					logger.Printf("Automatically detect file extension for file %s\n", filename)
				}
			}
		}

		if filepath.Ext(filename) == ".html" || filepath.Ext(filename) == ".htm" {
			document, err := goquery.NewDocumentFromResponse(resp)
			if err != nil {
				return "", fmt.Errorf("failed to download file %s: %v", url, err)
			}

			if document.Find("#page_layout #content #msg_back_button").Length() != 0 {
				return "", fmt.Errorf("failed to download file %s: not found", url)
			}
		}

		resultPath := filepath.Join(dir, filename)
		out, err := os.Create(resultPath)
		if err != nil {
			return "", err
		}

		defer out.Close()

		_, err = io.Copy(out, resp.Body)

		return resultPath, err
	}

	return "", fmt.Errorf("failed to download file %s: %v", url, resp.Status)
}

func getContentType(r *http.Response) string {
	contentType := r.Header.Get("Content-type")
	if len(contentType) > 0 {
		return contentType
	}

	return ""
}
