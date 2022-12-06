//这个示例程序展示如何写一个基本的表组测试
package chapter9

import (
	"net/http"
	"testing"
)

const checkMark1 = "\u2713"
const ballotX1 = "\u2717"

// TestDownload确认http包的Get函数可以下载内容
// 并正确处理不同的状态
func TestDownload1(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://localhost:7009/interface/void/getLabelInfo",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}
	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX1, err)
				}
				t.Log("\t\tShould be able to Get the url",
					checkMark1)
				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark1)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v", u.statusCode, ballotX1, resp.StatusCode)
				}

			}
		}
	}
}
