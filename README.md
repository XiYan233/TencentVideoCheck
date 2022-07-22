# 腾讯视频会员V力值自动签到

<p align="center">
    <a href="https://github.com/Demontisa"><img alt="Author" src="https://img.shields.io/badge/author-Demontisa-blueviolet"/></a>
    <img alt="Go" src="https://img.shields.io/badge/code-Go-success"/>
</p>
通过调用官方接口，每天晚上 23:30 自动领取已完成任务的V力值，借此可以达到快速升级的目的。

------

## 准备
```shell
  # 克隆本项目
    git clone https://github.com/Demontisa/TencentVideoCheck.git
  # 进入项目目录
    cd TencentVideoCheck/Server
```
## 修改数据库配置
   
   * 导入数据库文件tencentvideocheck.sql
   * 进入Config目录下
   * 修改DBConfig.go文件内容

## 编译项目
```shell
   go build main.go
```

## 运行项目
```shell
   ./main
```

## Cookie获取步骤

1. 电脑打开浏览器访问`v.qq.com`，并登录账号
2. 访问`https://film.qq.com/vip/my/` 然后打开打开控制台(`F12`)、切换到Network，找到 `https://film.qq.com/vip/my/` ，把`Request URL(请求标头):`中的Cookie复制
```http request
例：
   pgv_pvid=644865514; fqm_pvqid=adegaed-8b2e-4841-8eb4-agadefa; RK=afafae+; ptcz=sgasedfaesdfaesukjdhfadef; pt_sms_phone=176******14; tvfe_boss_uuid=sgrsasefa; LW_sid=sgsgsgsgaswa; LW_uid=sfsgsgsg; eas_sid=sfsfsfsfsf; pac_uid=sfsfsfsf; iip=0; tmeLoginType=2; wxopenid=; psrf_qqrefresh_token=sfsfsfsf; psrf_access_token_expiresAt=sfsfsfsf; euin=sfsfsfsf**; psrf_qqaccess_token=asfsfsfsfsf; wxunionid=; psrf_qqopenid=agrsgs656s5f6sf; wxrefresh_token=; video_guid=sgsdg656sfs; main_login=qq; vqq_access_token=dsgwsagsg; vqq_vuserid=66466656; vqq_openid=afafuhaifhaif; vqq_appid=3544545; video_platform=2; bucket_id=644555; ts_refer=v.qq.com/; ts_uid=6464646; qq_nick=%E2%80%AA; QQLivePCVer=554455; pgv_info=ssid=5646646; uid=000000; ptag=v_qq_com; qq_head=https://tvpic.gtimg.cn/head/d4157ab5e91b0864e5182bsfsf; vqq_vusession=xxxxxxxx; _ga=GA1.1.1038321556.1658481124; _ga_2DYB6E8NZY=GS1.1.1658481124.1.1.1658482341.0; ts_last=film.qq.com/vip/my/
```
