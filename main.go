package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {
	fmt.Println("Digite a URL do vídeo do YouTube:")
	var url string
	fmt.Scanln(&url)

	err := downloadVideo(url)
	if err != nil {
		fmt.Println("Erro ao fazer o download do vídeo:", err)
	}
}

func downloadVideo(url string) error {
	client := youtube.Client{}

	video, err := client.GetVideo(url)

	if err != nil {
		return err
	}

	format := getMP4WithAudio(video.Formats)

	stream, resp, err := client.GetStream(video, format)

	if err != nil {
		return err
	}

	defer stream.Close()

	serial := strconv.FormatInt(resp, 10)

	fileName := video.Title + serial +  ".mp4"

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, stream)

	if err != nil {
		return err
	}

	fmt.Println("Download do vídeo concluído:", fileName)

	return nil
}

func getMP4WithAudio(formats []youtube.Format) *youtube.Format {
	for _, f := range formats {
		if (strings.Contains(f.MimeType, "video/mp4") && f.AudioChannels > 0 && f.FPS == 30) {
			return &f
		}
	}

	return nil
}
