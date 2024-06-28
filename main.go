package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	server := g.Server()
	server.SetPort(80)
	server.Group("api", func(group *ghttp.RouterGroup) {
		group.ALL("login", login)
		group.ALL("cmd", cmd)
	})
	server.EnableAdmin()
	server.Run()
}

func login(request *ghttp.Request) {
	var cmdRequest *CmdRequest
	_ = request.Parse(&cmdRequest)
	g.Dump(cmdRequest)
	request.Response.Writeln("123")
}

func cmd(request *ghttp.Request) {
	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.Password("abc"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 仅用于示例，生产环境中应该验证主机密钥
	}

	// 连接到 SSH 服务器
	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// 创建新的 SSH 会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// 执行命令
	output, err := session.CombinedOutput("ls -l")
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}
	request.Response.Writeln(string(output))
}

type CmdRequest struct {
	Cmd string `json:"cmd"`
}
