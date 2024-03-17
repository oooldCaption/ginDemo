package reviewGo

import (
	"fmt"
	pb "github.com/cheggaaa/pb/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// downloadImage 下载图片并保存到指定的目录，同时显示进度条
func downloadImage(url, directory string, pool *pb.Pool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading:", url, err)
		return
	}
	defer resp.Body.Close()

	fileName := filepath.Base(url)
	filePath := filepath.Join(directory, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", filePath, err)
		return
	}
	defer file.Close()

	// 获取内容长度，创建进度条
	contentLength := resp.ContentLength
	bar := pb.Full.Start64(contentLength)
	bar.Set("prefix", fileName)
	pool.Add(bar)

	barReader := bar.NewProxyReader(resp.Body)

	// 写入文件，同时更新进度条
	io.Copy(file, barReader)
	bar.Finish()
}

func runDownload(dir string) {
	urls := []string{
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Mario/img/23_HBD_Mario_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Mario/img/23_HBD_Mario_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Mario/img/23_HBD_Mario_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Mario/img/23_HBD_Mario_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Zelda/img/23_HBD_Zelda_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Zelda/img/23_HBD_Zelda_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Zelda/img/23_HBD_Zelda_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Zelda/img/23_HBD_Zelda_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Animal/img/23_HBD_Animal_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Animal/img/23_HBD_Animal_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Animal/img/23_HBD_Animal_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Animal/img/23_HBD_Animal_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Splatoon/img/23_HBD_Splatoon_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Splatoon/img/23_HBD_Splatoon_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Splatoon/img/23_HBD_Splatoon_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Splatoon/img/23_HBD_Splatoon_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pikmin/img/23_HBD_Pikmin_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pikmin/img/23_HBD_Pikmin_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pikmin/img/23_HBD_Pikmin_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pikmin/img/23_HBD_Pikmin_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pokemon/img/23_HBD_Pokemon_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pokemon/img/23_HBD_Pokemon_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pokemon/img/23_HBD_Pokemon_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Pokemon/img/23_HBD_Pokemon_1200_1920.jpg",

		"https://www.nintendo.com/jp/wallpaper/23_HBD_Kirby/img/23_HBD_Kirby_1242_2688.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Kirby/img/23_HBD_Kirby_1080_1920.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Kirby/img/23_HBD_Kirby_1536_2048.jpg",
		"https://www.nintendo.com/jp/wallpaper/23_HBD_Kirby/img/23_HBD_Kirby_1200_1920.jpg",
	}

	directory := dir
	pool, err := pb.StartPool()
	if err != nil {
		fmt.Println("Error starting progress bar pool:", err)
		return
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			downloadImage(url, directory, pool)
		}(url)
	}

	wg.Wait()
	pool.Stop()
}

func RunDownScript() {
	runDownload("/Users/su/Desktop/test")
}
