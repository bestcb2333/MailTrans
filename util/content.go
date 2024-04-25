package util

import (
	"errors"
	"io"
	"mime"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var section = &imap.BodySectionName{BodyPartName: imap.BodyPartName{
	Specifier: imap.EntireSpecifier,
}}

func GetContent(conn *client.Client, uids ...uint32) (chan *imap.Message, error) {

	// 创建一个UID序列组
	seqset := new(imap.SeqSet)
	seqset.AddNum(uids...)

	//获取完整的邮件体

	items := []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}
	messages := make(chan *imap.Message, len(uids))

	if err := conn.UidFetch(seqset, items, messages); err != nil {
		return nil, err
	}

	return messages, nil
}
func ExtractText(msg *imap.Message) (string, error) {
	if msg == nil {
		return "", errors.New("邮件是空的")
	}
	r := msg.GetBody(section)
	if r == nil {
		return "", errors.New("未找到邮件内容")
	}

	// 解析邮件体为多部分
	mr, err := mail.CreateReader(r)
	if err != nil {
		return "", err
	}

	var htmlBody string
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		contentTypeHeader := p.Header.Get("Content-Type")
		// 解析 Content-Type 头部来获取 MIME 类型和字符集
		mimeType, params, _ := mime.ParseMediaType(contentTypeHeader)
		if mimeType == "text/html" {
			charset := strings.ToLower(params["charset"]) // 获取字符集
			var reader io.Reader = p.Body
			if charset == "gbk" || charset == "gb2312" || charset == "gb18030" {
				// 对于 GBK 编码，使用转换器
				reader = transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder())
			}

			b, err := io.ReadAll(reader)
			if err != nil {
				return "", err
			}
			htmlBody = string(b)
			break // 假设只需要第一个 HTML 部分
		}
	}

	if htmlBody == "" {
		return "", errors.New("未找到HTML内容")
	}
	return htmlBody, nil

}
