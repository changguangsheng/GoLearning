package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	Client   string `json:"client"`
	Fr       string `json:"fr"`
	NeedQc   int    `json:"needQc"`
	S        string `json:"s"`
	Uuid     string `json:"uuid"`
	Exchange bool   `json:"exchange"`
}
type DictResponse struct {
	Data struct {
		Kaoyan struct {
			ExamFreqInfo []struct {
				Pos            string  `json:"pos"`
				Chinese        string  `json:"chinese"`
				SenseTier      int     `json:"sense_tier"`
				SenseTierRatio float64 `json:"sense_tier_ratio"`
			} `json:"exam_freq_info"`
			Word          string `json:"word"`
			TotalExamFreq int    `json:"total_exam_freq"`
			DirectionInfo struct {
				Ranges      int    `json:"ranges"`
				Description string `json:"description"`
			} `json:"direction_info"`
			ImportantScore int `json:"important_score"`
			WordFamily     struct {
				Word               string `json:"word"`
				DerivativeInfoList []struct {
					Word               string        `json:"word"`
					DerivativeInfoList []interface{} `json:"derivative_info_list"`
					Mean               string        `json:"mean"`
				} `json:"derivative_info_list"`
				Mean string `json:"mean"`
			} `json:"word_family"`
			WordLevelPhraseInfo []struct {
				Phrase          string `json:"phrase"`
				Pos             string `json:"pos"`
				BriefDefinition string `json:"brief_definition"`
			} `json:"word_level_phrase_info"`
			Sense struct {
				SenseNum      int `json:"sense_num"`
				SenseInfoList []struct {
					Pos                string `json:"pos"`
					Zh                 string `json:"zh"`
					En                 string `json:"en"`
					Synonyms           string `json:"synonyms"`
					Antonyms           string `json:"antonyms"`
					ExaminationExample []struct {
						Sentence      string `json:"sentence"`
						SentenceTrans string `json:"sentence_trans"`
						Source        string `json:"source"`
						Year          int    `json:"year"`
						Month         int    `json:"month"`
						Day           int    `json:"day"`
						Type          string `json:"type"`
						Coll          string `json:"coll"`
					} `json:"examination_example,omitempty"`
				} `json:"sense_info_list"`
			} `json:"sense"`
		} `json:"kaoyan"`
		Gaokao struct {
			ExamFreqInfo []struct {
				SenseTier      int     `json:"sense_tier"`
				Pos            string  `json:"pos"`
				Chinese        string  `json:"chinese"`
				SenseTierRatio float64 `json:"sense_tier_ratio"`
			} `json:"exam_freq_info"`
			Word                string `json:"word"`
			WordLevelPhraseInfo []struct {
				BriefDefinition string `json:"brief_definition"`
				Phrase          string `json:"phrase"`
				Pos             string `json:"pos"`
			} `json:"word_level_phrase_info"`
			DirectionInfo struct {
				Ranges      int    `json:"ranges"`
				Description string `json:"description"`
			} `json:"direction_info"`
			ImportantScore int `json:"important_score"`
			WordFamily     struct {
				DerivativeInfoList []struct {
					DerivativeInfoList []interface{} `json:"derivative_info_list"`
					Word               string        `json:"word"`
					Mean               string        `json:"mean"`
				} `json:"derivative_info_list"`
				Word string `json:"word"`
				Mean string `json:"mean"`
			} `json:"word_family"`
			TotalExamFreq int `json:"total_exam_freq"`
			Sense         struct {
				SenseNum      int `json:"sense_num"`
				SenseInfoList []struct {
					En                 string `json:"en"`
					Zh                 string `json:"zh"`
					Pos                string `json:"pos"`
					Synonyms           string `json:"synonyms"`
					Antonyms           string `json:"antonyms"`
					ExaminationExample []struct {
						Sentence      string `json:"sentence"`
						Month         int    `json:"month"`
						Source        string `json:"source"`
						Coll          string `json:"coll"`
						SentenceTrans string `json:"sentence_trans"`
						Year          int    `json:"year"`
						Type          string `json:"type"`
						Day           int    `json:"day"`
					} `json:"examination_example"`
				} `json:"sense_info_list"`
			} `json:"sense"`
			WordUsage []struct {
				Usage string `json:"usage"`
			} `json:"word_usage"`
			ExchangeInfo struct {
				WordPl  []string `json:"word_pl"`
				WordEr  []string `json:"word_er"`
				WordEst []string `json:"word_est"`
			} `json:"exchange_info"`
		} `json:"gaokao"`
		Zhongkao struct {
			ExamFreqInfo []struct {
				Pos            string  `json:"pos"`
				Chinese        string  `json:"chinese"`
				SenseTier      int     `json:"sense_tier"`
				SenseTierRatio float64 `json:"sense_tier_ratio"`
			} `json:"exam_freq_info"`
			Word                string `json:"word"`
			WordLevelPhraseInfo []struct {
				Phrase          string `json:"phrase"`
				Pos             string `json:"pos"`
				BriefDefinition string `json:"brief_definition"`
			} `json:"word_level_phrase_info"`
			DirectionInfo struct {
				Ranges      int    `json:"ranges"`
				Description string `json:"description"`
			} `json:"direction_info"`
			ImportantScore int `json:"important_score"`
			WordFamily     struct {
				Word               string `json:"word"`
				DerivativeInfoList []struct {
					Word               string        `json:"word"`
					DerivativeInfoList []interface{} `json:"derivative_info_list"`
					Mean               string        `json:"mean"`
				} `json:"derivative_info_list"`
				Mean string `json:"mean"`
			} `json:"word_family"`
			TotalExamFreq int `json:"total_exam_freq"`
			Sense         struct {
				SenseNum      int `json:"sense_num"`
				SenseInfoList []struct {
					Pos                string `json:"pos"`
					Zh                 string `json:"zh"`
					En                 string `json:"en"`
					Synonyms           string `json:"synonyms"`
					Antonyms           string `json:"antonyms"`
					ExaminationExample []struct {
						Sentence      string `json:"sentence"`
						SentenceTrans string `json:"sentence_trans"`
						Source        string `json:"source"`
						Year          int    `json:"year"`
						Month         int    `json:"month"`
						Day           int    `json:"day"`
						Type          string `json:"type"`
						Coll          string `json:"coll"`
					} `json:"examination_example,omitempty"`
				} `json:"sense_info_list"`
			} `json:"sense"`
			WordUsage []struct {
				Usage string `json:"usage"`
			} `json:"word_usage"`
			Disambiguation []struct {
				SenseGroup         string `json:"sense_group"`
				DisambiguationList []struct {
					Word                string `json:"word"`
					SenseDisambiguation string `json:"sense_disambiguation"`
				} `json:"disambiguation_list"`
			} `json:"disambiguation"`
		} `json:"zhongkao"`
	} `json:"data"`
}

func query1(word string) {
	client := &http.Client{}
	dictRequest := &DictRequest{
		From:     "auto",
		To:       "zh-CHS",
		Text:     word,
		Client:   "pc",
		Fr:       "browser_pc",
		NeedQc:   1,
		S:        "60136718a425b8a5ce47bf7fdf253378",
		Uuid:     "07016eb4-528c-44ef-8bf7-4e9580a197c1",
		Exchange: false,
	}
	buf, err := json.Marshal(dictRequest)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewBuffer(buf)
	//var data = strings.NewReader(`{"from":"auto","to":"zh-CHS","text":"good","client":"pc","fr":"browser_pc","needQc":1,"s":"8dcdbc5f48b9b9f85a43f0307e471252","uuid":"59dd44f7-cb22-477c-be91-283b87107f49","exchange":false}`)
	req, err := http.NewRequest("POST", "https://fanyi.sogou.com/api/transpc/text/result", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "ABTEST=0|1683809836|v17; SNUID=E33C80FD63669909E5537A63631C0D6D; SUID=815EE29FEF53A00A00000000645CE62C; wuid=1683809836793; translate.sess=9c09e587-4292-410c-b4b7-756d97d93619; SUV=1683809837288; SGINPUT_UPSCREEN=1683809837297; FQV=a1b7ce7d9a5f7872ebe0bcdfc5b5fa90")
	req.Header.Set("Origin", "https://fanyi.sogou.com")
	req.Header.Set("Referer", "https://fanyi.sogou.com/text?keyword=&transfrom=auto&transto=zh-CHS&model=general")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}

	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("搜狗翻译：", word)
	for _, item := range dictResponse.Data.Kaoyan.ExamFreqInfo {
		fmt.Println(item.Pos, "."+item.Chinese)
	}
}

type DictRequest1 struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse1 struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query2(word string) {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	request := DictRequest1{
		TransType: "en2zh",
		Source:    word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewBuffer(buf)
	//创建请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}

	//设置请求头
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//读取响应
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode: ", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponse1
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("彩云小译：", word)
	for _, iterm := range dictResponse.Dictionary.Explanations {
		fmt.Println(iterm)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDic WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		query1(word)
		cancel()
	}()
	go func() {
		query2(word)
		cancel()
	}()
	<-ctx.Done()
}
