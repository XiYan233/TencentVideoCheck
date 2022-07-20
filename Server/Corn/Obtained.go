package Corn

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ObtainedStruct struct {
	LscoreMonth            int    `json:"lscore_month"`
	LscoreMonthLimit       int    `json:"lscore_month_limit"`
	Msg                    string `json:"msg"`
	RankDesc               string `json:"rank_desc"`
	Ret                    int    `json:"ret"`
	UngiveCscore           int    `json:"ungive_cscore"`
	UngiveLscore           int    `json:"ungive_lscore"`
	WillExpireUngiveCscore int    `json:"will_expire_ungive_cscore"`
	WillExpireUngiveLscore int    `json:"will_expire_ungive_lscore"`
}

func Obtained(cookie string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vip.video.qq.com/fcgi-bin/comm_cgi?name=spp_vscore_user_mashup&type=2&otype=xjson", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "vip.video.qq.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("cookie", "pgv_pvid=3626964272; fqm_pvqid=754026fc-8b2e-4841-8eb4-b620c9fd9643; RK=l3X5ODzkW+; ptcz=69975641c5348f3dcaae3190f84c3b756252fbfd65270bacbee590c88920c2b6; pt_sms_phone=176******14; tvfe_boss_uuid=95c21b7213853706; video_platform=2; LW_sid=c156J4C9V8V4R836r3t6y2v1t8; LW_uid=11m624e9p8d4A8v6S3f6o2T1J9; eas_sid=31l6z4W9B854C8v6G3I6J410r9; pac_uid=0_e6932f22ce188; iip=0; tmeLoginType=2; wxopenid=; psrf_qqrefresh_token=E1F09FF2257066845F9F6592FD244C81; psrf_access_token_expiresAt=1661671383; euin=oKEq7eSsNK65oc**; psrf_qqaccess_token=2BD1129BF1BF967BB3EEA837CD3BD98C; wxunionid=; psrf_qqopenid=44016911A54F3CB1A27509744E7907F1; wxrefresh_token=; video_guid=7dc5fd513644863d; pgv_info=ssid=s592758924; _qpsvr_localtk=0.8438646674024592; main_login=qq; vqq_access_token=456DCCEF9B107CA1F75410FF84DBAF62; vqq_appid=101483052; vqq_openid=1ADB8583A2E271192633A6F50983B961; vqq_vuserid=139260750; vqq_refresh_token=F124F9C8F1EA9A8FC140F469E6EE7C66; login_time_init=2022-7-20 14:50:15; vqq_vusession=Bd0hRsxueH1yoLdB-lzAdw.N; vqq_next_refresh_time=6500; vqq_login_time_init=1658306518; login_time_last=2022-7-20 16:41:57")
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var obtainedStruct ObtainedStruct
	err = json.Unmarshal(bodyText, &obtainedStruct)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		log.Printf("本月已获得%v点V力值，%v", obtainedStruct.LscoreMonth, obtainedStruct.RankDesc)
	}
}
