// 这个示例程序展示如何内部模仿HTTP GET调用
// 与本书之前的例子有差别
package chapter9

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark2 = "\u2713"
const ballotX2 = "\u2717"

// feed模仿了我们期望接受的xml文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
<title>Going Go Programming</title>
<description>Golang : https://github.com/goinggo</description>
<link>http://www.goinggo.net/</link>
<item>
<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
<title>Object Oriented Programming Mechanics</title>
<description>Go is an object oriented language.</description>
<link>http://www.goinggo.net/2015/03/object-oriented</link>
</item>
</channel>
</rss>`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 并且内容可以被正确地反序列化并关闭
func TestDownload2(t *testing.T) {
	statusCode2 := http.StatusOK

	server2 := mockServer()
	defer server2.Close()
	t.Log("Given the need to test downloading content.")

	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server2.URL, statusCode2)
		{
			resp, err := http.Get(server2.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX2, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark2)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode2 {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode2, ballotX2, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode2, checkMark2)
		}
	}
}
