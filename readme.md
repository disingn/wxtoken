部署之前请准备以下内容：
- 一台服务器（安装好 redis）
- 一个 oneapi 或者 newapi 的 Authorization 令牌（系统访问令牌）
- 一个微信公众号（个人或者认证的都行）
- 一个域名（可以是二级域名，也可以是一级域名）
- 略微有点代码基础

### 一、安装
#### 1.1 克隆代码
```shell
git clone 
```
#### 1.2 构建程序
```shell
cd wxtoken
export GOOS=linux                                                             
export GOARCH=amd64
go build -o wxtoken
```
#### 1.3 配置文件
```shell
cp config.example.yaml config.yaml
```
修改配置文件
```yaml
app:
  # 首次关注回复
  FirstSub: '嗨！这里是百晓生，一个智能助理的雏形。
  
请保持关注，在这里可以获取免费使用的全模型api。

便宜又好用的ai站：<a href="https://www.a0.chat/" target="_blank" style="color: red;">www.a0.chat</a>

如要领取每日免费 50万 Tokens 的全模型api，请回复文字：领取'
  #领取成功之后的回复
  UsedSub: "今日你已经领取过了～"
  #其他的回复
  OtherSub: "不要乱发消息哦～"
  #领取成功之后的回复
  TokenSub: '恭喜你打开 AI 世界的大门🎉

🎁成功获取到一枚 KEY: %s

🔑KEY 包含 50万Tokens 体验额度，有效期 1 天，每 24 小时可领一枚。

🚗配合接口地址使用：https://api.bxsapi.com/

😆不够清楚？👉<a href="https://xx7tl7g4rg.feishu.cn/docx/ImR1d2MI1oZvh1xz5DxcekWsnid/" target="_blank" style="color: red;">查看使用教程</a>

👉<a href="https://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=HJiFQMzB3DxCBas1uDTijpn0bs2Firn1&authKey=OB%2FCNRfKzzanEeue7EXGKosOsGrLosae4AeGP1mHqDhYwz4voBaIhCnCBiXfUklb&noverify=0&group_code=814702872/" target="_blank" style="color: red;">加群了解更多</a>

😎提示：该密钥同样支持在其他支持 OpenAI 服务的场景中使用，如沉浸式翻译，学术GPT等等，到期或用完将无法使用，作为免费服务本站不提供任何技术支持，敬请理解。'

sever:
  # 服务端口 如果使用 docker 部署请使用服务器本机 ip 地址
  RdsHost: "localhost"
  # 服务端口
  RdsPort: 6379
  # 请填写你的redis用户名，如果没有用户名请留空
  RdsUser: ""
    # 请填写你的redis数据库
  RdsDB: 0
  # 请填写你的redis密码，如果没有密码请留空
  RdsPass: ""
  # 请填写你的oneapi或者newapi的的 Authorization 令牌（系统访问令牌）
  Authorization: ""
  # 请填写你的oneapi或者newapi的的域名：例如：https://api.bxsapi.com/ 一定要以 / 结尾
  APIUrl: ""
  #微信公众号的令牌（Token）
  WxToken: ""
  #根据自己的需求修改 50000 是 1 刀
  Limit: 500000
  #过期时间 默认 1 天，最小是1天！不要写 0
  Expire: 1
```
嫌麻烦可以直接用实例的配置文件

#### 1.4 启动程序
```shell
./wxtoken
```
#### 1.5 配置Nginx反代
 这个自己使用宝塔或者 1panel 都可以，这里就不多说了
 
### 二、使用
#### 2.1 配置微信公众号
- 配置服务器地址
- 配置 Token（和你config.yaml 中的 WxToken 保持一致）
- 配置 EncodingAESKey（随便写）
- 配置消息加解密方式（明文模式）

上述的方式也适合本地修改二次开发等等

docker 部署：
#### 1.1准备工作
安装 docker 和 docker-compose （这里建议使用 docker-compose）
#### 1.2 部署
```shell
cd wxtoken
docker-compose up -d
```
这里注意修改 config.yaml 中的 RdsHost 为服务器的 ip 地址
切记！！！

### 三、其他
其他的就都和上面的一样了，如果有问题可以加群讨论

QQ: 814702872

### 四、演示公众号
[![pivOEkD.jpg](https://s11.ax1x.com/2024/01/04/pivOEkD.jpg)](https://imgse.com/i/pivOEkD)


