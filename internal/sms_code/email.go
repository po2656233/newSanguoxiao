package sms_code

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	//// 邮件标题
	//Title string
	//// 邮件内容
	//Body string
	//// 收件人列表
	//RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

func New(sender, password, smtpAddr string, smtpPort int) *MailboxConf {
	if sender == "" && password == "" {

	}
	return &MailboxConf{
		// 发件人账号
		Sender: sender,
		// 发件人密码，QQ邮箱这里配置授权码
		SPassword: password,
		// SMTP 服务器地址， QQ邮箱是smtp.qq.com
		SMTPAddr: smtpAddr,
		// SMTP端口 QQ邮箱是25
		SMTPPort: smtpPort,
	}
}

// SendTo 发送邮件
func (self *MailboxConf) SendTo(title, body string, receiver []string) error {
	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, self.Sender, "天下游官方")
	m.SetHeader(`To`, receiver...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	// m.Attach("./Dockerfile") //添加附件
	err := gomail.NewDialer(self.SMTPAddr, self.SMTPPort, self.Sender, self.SPassword).DialAndSend(m)
	if err != nil {
		log.Fatalf("Send Email Fail, %s", err.Error())
		return err
	}
	log.Printf("Send Email Success")
	return nil
}

// SendCodeTo 示例:SendCodeTo("验证"",[]string{"2437854119@qq.com", "po2656233@qq.com"})
func (self *MailboxConf) SendCodeTo(receiver []string) error {
	//self.Title = "验证"
	//这里就是我们发送的邮箱内容，但是也可以通过下面的html代码作为邮件内容
	// mailConf.Body = "坚持才是胜利，奥里给"

	//这里支持群发，只需填写多个人的邮箱即可，我这里发送人使用的是QQ邮箱，所以接收人也必须都要是
	//QQ邮箱
	//self.RecipientList = receiver    //[]string{"2437854119@qq.com", "po2656233@qq.com"}
	//self.Sender = "po2656233@qq.com" //`邮箱账号`

	//这里QQ邮箱要填写授权码，网易邮箱则直接填写自己的邮箱密码，授权码获得方法在下面
	//self.SPassword = "doaqmadsooivbced" //"这里填写自己QQ邮箱授权码"

	//下面是官方邮箱提供的SMTP服务地址和端口
	// QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	// 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	// 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	// 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	// 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）

	//self.SMTPAddr = `smtp.qq.com`
	//self.SMTPPort = 25

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为<br>%s<br>为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, self.Sender, "天下游官方")
	m.SetHeader(`To`, receiver...)
	m.SetHeader(`Subject`, "验证码")
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	err := gomail.NewDialer(self.SMTPAddr, self.SMTPPort, self.Sender, self.SPassword).DialAndSend(m)
	if err != nil {
		log.Fatalf("Send Email Fail, %s", err.Error())
		return err
	}
	log.Printf("Send Email Success")
	return nil
}
