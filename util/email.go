package util

import (
	"log"
	"net/smtp"
	"sync"
	"time"

	"github.com/jordan-wright/email"
)

func Emails(VerificationCode string, emails string) {

	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 创建有5 个缓冲的通道，数据类型是  *email.Email
	ch := make(chan *email.Email, 5)
	// 连接池
	p, err := email.NewPool(
		"smtp.qq.com:100",
		3, // 数量设置成 3 个
		smtp.PlainAuth("", "2695009886@qq.com", "pndftuyzutbcdhff", "smtp.qq.com"),
	)

	if err != nil {
		log.Fatal("email.NewPool error : ", err)
	}

	// sync 包，控制同步
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			// 若 ch 无数据，则阻塞， 若 ch 关闭，则退出循环
			for e := range ch {
				// 超时时间 10 秒
				err := p.Send(e, 60*time.Second)
				if err != nil {
					log.Printf("p.Send error : %v , e = %v , i = %d\n", err, e, i)
				}
			}
		}()
	}
	for i := 0; i < 1; i++ {
		e := email.NewEmail()
		// 设置发送邮件的基本信息
		e.From = "表白墙注册 <2695009886@qq.com>"
		// e.To = []string{"2695009886@qq.com", "3493665801@qq.com", "2081212291@qq.com", "320028161@qq.com", "1598934549@qq.com", "1649336090@qq.com", "447909290@qq.com"}
		e.To = []string{emails}
		e.Subject = "表白墙注册验证码"
		e.Text = []byte("欢迎注册XXX校园表白墙, 您的验证码为: " + VerificationCode + "注意不要泄露给其他人")
		ch <- e
	}

	// 关闭通道
	close(ch)
	// 等待子协程退出
	wg.Wait()
	log.Println("send successfully ... ")
}
