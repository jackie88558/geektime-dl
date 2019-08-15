package cmds

import (
	"fmt"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/mmzou/geektime-dl/service"
	"github.com/urfave/cli"
)

//NewDownloadCommand login command
func NewDownloadCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "download",
			Usage:     "下载",
			UsageText: appName + " download",
			Action:    downloadAction,
			Before:    authorizationFunc,
		},
	}
}

func downloadAction(c *cli.Context) error {
	course, articles, err := application.CourseWithArticles(186)

	if err != nil {
		return err
	}

	fmt.Println(extractDownloadData(course, articles))

	return nil
}

func extractDownloadData(course *service.Course, articles []*service.Article) []downloader.Data {
	data := downloader.EmptyList
	if course.IsColumn() {
		key := "default"
		for _, article := range articles {
			if !article.IncludeAudio {
				continue
			}
			urls := []downloader.URL{
				{
					URL:  article.AudioDownloadURL,
					Size: article.AudioSize,
					Ext:  "mp3",
				},
			}

			streams := map[string]downloader.Stream{
				key: downloader.Stream{
					URLs:    urls,
					Size:    article.AudioSize,
					Quality: key,
				},
			}

			data = append(data, downloader.Data{
				Title:   article.ArticleTitle,
				Streams: streams,
				Type:    "audio",
			})
		}
	}

	return data
}
