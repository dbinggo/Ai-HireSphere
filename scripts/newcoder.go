package main

import (
	"Ai-HireSphere/common/thift/deepseek"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// main
//
//	@Description: 获取牛客数据
func main() {

	// 获取文章列表
	records, err := newcoder()
	if err != nil || records.Success != true || len(records.Data.Records) == 0 {
		panic(err)
	}
	prompt := fmt.Sprintf(Prompt, "字节", ExampleYes, ExampleNO)
	client := deepseek.NewDeepSeekClient("", "")
	// 控制并发
	var wg sync.WaitGroup

	for _, record := range records.Data.Records {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 获取文章主要内容的文章链接
			articleLink := record.MomentData.Uuid
			contentResp, err := getContent(articleLink)
			if err != nil || contentResp.Success != true {

				return
			}
			chatResp, err := client.Chat(prompt, contentResp.Data.Content)
			if err != nil {
				return
			}
			fmt.Println(chatResp)
		}()

	}
	wg.Wait()

}

const (
	Prompt = "你是一个资深帖子审核员，现在我会给你很多个面经帖子，请你输出他是否是%s公司的面经，如果是请按照如下格式给我相关json数据 %s，如果不是，请给我如下json数据 %s，注意：你只可以返回我json数据，请不要返回其他格式信息，返回的数据需要满足json格式"

	ExampleYes = `{
    isOK:"true",
	questions:[
		{
			index: 1,
			question: xxxxx,
		},
	]
}`
	ExampleNO = `{
    isOK:"false"
}`
)

// 获取牛客面经列表
func newcoder() (Record RecordResponse, err error) {

	// 进行请求
	// 以下代码为apifox自动生成 无需去编辑

	url := "https://gw-c.nowcoder.com/api/sparta/job-experience/experience/job/list"
	method := "POST"

	payload := strings.NewReader(`{
    "companyList": [
        665
    ],
    "jobId": 11002,
    "level": 3,
    "order": 3,
    "page": 1,
    "isNewJob": true
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("dnt", "1")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Cookie", "HMACCOUNT=CC29B2F3D6548445; gr_user_id=9cc3bad0-7245-4af5-8434-58dfca5cf23d; NOWCODERCLINETID=DA8313F295FA3D6509737730B6368457; NOWCODERUID=620EBA9EBF65983E2FF6C4ACC89D8443; Hm_lvt_a808a1326b6c06c437de769d1b85b870=1736171627; c196c3667d214851b11233f5c17f99d5_gr_session_id=ab8a7906-4024-4bd3-925d-60dacd1deaa6; acw_tc=1f1b8a366f9cbab3e78182ecdaf41acb840c84b295dba24c85681fa0ef201be3; t=9EF2D83D5E1D038962116D32CFB4104D; ls_sess_id=9EF2D83D5E1D038962116D32CFB4104D; c196c3667d214851b11233f5c17f99d5_gr_last_sent_sid_with_cs1=ab8a7906-4024-4bd3-925d-60dacd1deaa6; c196c3667d214851b11233f5c17f99d5_gr_last_sent_cs1=772926294; Hm_lpvt_a808a1326b6c06c437de769d1b85b870=1737994540; c196c3667d214851b11233f5c17f99d5_gr_cs1=772926294; c196c3667d214851b11233f5c17f99d5_gr_session_id_ab8a7906-4024-4bd3-925d-60dacd1deaa6=true")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "gw-c.nowcoder.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &Record)
	if err != nil {
		fmt.Println(err)
		return Record, err
	}
	return
}

type RecordResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		Current   int `json:"current"`
		Size      int `json:"size"`
		Total     int `json:"total"`
		TotalPage int `json:"totalPage"`
		Records   []struct {
			ContentId    string `json:"contentId"`
			ContentType  int    `json:"contentType"`
			InterviewExp *struct {
				Message     string `json:"message"`
				Icon        string `json:"icon"`
				ContentId   string `json:"contentId"`
				ContentType int    `json:"contentType"`
			} `json:"interviewExp"`
			UserBrief struct {
				UserId          int     `json:"userId"`
				Nickname        string  `json:"nickname"`
				Admin           bool    `json:"admin"`
				Followed        bool    `json:"followed"`
				HeadImgUrl      string  `json:"headImgUrl"`
				Gender          *string `json:"gender"`
				HeadDecorateUrl string  `json:"headDecorateUrl"`
				HonorLevel      int     `json:"honorLevel"`
				HonorLevelName  string  `json:"honorLevelName"`
				HonorLevelColor string  `json:"honorLevelColor"`
				WorkTime        string  `json:"workTime"`
				EducationInfo   *string `json:"educationInfo"`
				IdentityList    []struct {
					CompanyId     int    `json:"companyId"`
					IdentityClass string `json:"identityClass"`
					IdentityNo    int    `json:"identityNo"`
					Name          string `json:"name"`
					CompanyName   string `json:"companyName"`
					JobName       string `json:"jobName"`
				} `json:"identityList"`
				ActivityIconList []struct {
					ImgUrl      string      `json:"imgUrl"`
					ExpireTime  interface{} `json:"expireTime"`
					Name        string      `json:"name"`
					DiscussLink string      `json:"discussLink"`
					AppLink     string      `json:"appLink"`
					PcLink      string      `json:"pcLink"`
					Content     interface{} `json:"content"`
					Source      int         `json:"source"`
					Type        int         `json:"type"`
				} `json:"activityIconList"`
				ActivityIconListV2 []struct {
					ImgUrl      string      `json:"imgUrl"`
					ExpireTime  interface{} `json:"expireTime"`
					Name        string      `json:"name"`
					DiscussLink *string     `json:"discussLink"`
					AppLink     *string     `json:"appLink"`
					PcLink      *string     `json:"pcLink"`
					Content     interface{} `json:"content"`
					Source      *int        `json:"source"`
					Type        *int        `json:"type"`
				} `json:"activityIconListV2"`
				MemberIdentity  int         `json:"memberIdentity"`
				MemberStartTime interface{} `json:"memberStartTime"`
				MemberEndTime   interface{} `json:"memberEndTime"`
				Member          interface{} `json:"member"`
				AuthDisplayInfo string      `json:"authDisplayInfo"`
				EnterpriseInfo  interface{} `json:"enterpriseInfo"`
				BadgeIconUrl    interface{} `json:"badgeIconUrl"`
				NicknameStyle   struct {
					Direction string `json:"direction"`
					Colors    []struct {
						Light string      `json:"light"`
						Dark  interface{} `json:"dark"`
					} `json:"colors"`
				} `json:"nicknameStyle,omitempty"`
				CardActivityIcon struct {
					Icon         string `json:"icon"`
					IconDark     string `json:"iconDark"`
					ActivityName string `json:"activityName"`
					Router       string `json:"router"`
				} `json:"cardActivityIcon,omitempty"`
			} `json:"userBrief"`
			MomentData struct {
				Ip4             string        `json:"ip4"`
				Ip4Location     string        `json:"ip4Location"`
				Id              int           `json:"id"`
				Uuid            string        `json:"uuid"`
				UserId          int           `json:"userId"`
				Title           string        `json:"title"`
				NewTitle        interface{}   `json:"newTitle"`
				Content         string        `json:"content"`
				NewContent      interface{}   `json:"newContent"`
				Type            int           `json:"type"`
				Status          int           `json:"status"`
				HasEdit         bool          `json:"hasEdit"`
				IsAnonymousFlag bool          `json:"isAnonymousFlag"`
				BeMyOnly        bool          `json:"beMyOnly"`
				LinkMoment      interface{}   `json:"linkMoment"`
				ImgMoment       []interface{} `json:"imgMoment"`
				ClockMoment     interface{}   `json:"clockMoment"`
				VideoMoment     interface{}   `json:"videoMoment"`
				CreatedAt       int64         `json:"createdAt"`
				Circle          interface{}   `json:"circle"`
				EditTime        int64         `json:"editTime"`
				Edited          bool          `json:"edited"`
				ShowTime        int64         `json:"showTime"`
			} `json:"momentData,omitempty"`
			SubjectData []struct {
				Id          int         `json:"id"`
				Uuid        string      `json:"uuid"`
				TagId       int         `json:"tagId"`
				SubjectType int         `json:"subjectType"`
				Content     string      `json:"content"`
				CreatedAt   int64       `json:"createdAt"`
				IsFollow    interface{} `json:"isFollow"`
				Official    int         `json:"official"`
				HadVote     int         `json:"hadVote"`
			} `json:"subjectData"`
			VoteData struct {
				VoteId    int         `json:"voteId"`
				WithVote  bool        `json:"withVote"`
				VoteTitle interface{} `json:"voteTitle"`
				VoteType  interface{} `json:"voteType"`
			} `json:"voteData"`
			BlogZhuanlan *struct {
				Id            int         `json:"id"`
				HashId        string      `json:"hashId"`
				Title         string      `json:"title"`
				HeadUrl       string      `json:"headUrl"`
				ZhuanlanIntro string      `json:"zhuanlanIntro"`
				ItemId        interface{} `json:"itemId"`
				ItemType      interface{} `json:"itemType"`
				ArticlePrice  float64     `json:"articlePrice"`
			} `json:"blogZhuanlan"`
			FrequencyData struct {
				LikeCnt         int         `json:"likeCnt"`
				FollowCnt       int         `json:"followCnt"`
				CommentCnt      int         `json:"commentCnt"`
				TotalCommentCnt int         `json:"totalCommentCnt"`
				ViewCnt         int         `json:"viewCnt"`
				ShareCnt        int         `json:"shareCnt"`
				IsLike          bool        `json:"isLike"`
				IsFollow        bool        `json:"isFollow"`
				FlowerData      interface{} `json:"flowerData"`
			} `json:"frequencyData"`
			ExtraInfo struct {
				ContentTypeVar string `json:"contentType_var"`
				TrackIDVar     string `json:"trackID_var"`
				ContentIDVar   string `json:"contentID_var"`
				TrackId        string `json:"trackId"`
				DolphinVar     string `json:"dolphin_var"`
				EntityId       string `json:"entityId"`
				EntityIDVar    string `json:"entityID_var"`
			} `json:"extraInfo"`
			JobSubscript struct {
				JobIcon    string      `json:"jobIcon"`
				JobMessage string      `json:"jobMessage"`
				ShowFollow interface{} `json:"showFollow"`
			} `json:"jobSubscript,omitempty"`
			CommentExposure struct {
				CommentId       int    `json:"commentId"`
				HeadImgUrl      string `json:"headImgUrl"`
				Nickname        string `json:"nickname"`
				Content         string `json:"content"`
				CommentType     int    `json:"commentType"`
				ContentJsonList struct {
					Data []struct {
						Text   string      `json:"text"`
						Color  interface{} `json:"color"`
						Router interface{} `json:"router"`
					} `json:"data"`
				} `json:"contentJsonList"`
				Images                    []interface{} `json:"images"`
				CardActivityIcon          interface{}   `json:"cardActivityIcon"`
				CardActivityIconInContent interface{}   `json:"cardActivityIconInContent"`
			} `json:"commentExposure,omitempty"`
			InternalRecommend interface{} `json:"internalRecommend"`
			ContentData       struct {
				Id               string      `json:"id"`
				Uuid             string      `json:"uuid"`
				AuthorId         string      `json:"authorId"`
				Title            string      `json:"title"`
				NewTitle         interface{} `json:"newTitle"`
				RichText         interface{} `json:"richText"`
				Content          string      `json:"content"`
				NewContent       interface{} `json:"newContent"`
				TypeName         string      `json:"typeName"`
				BeMyOnly         bool        `json:"beMyOnly"`
				ContentImageUrls []struct {
					Src    string      `json:"src"`
					Alt    interface{} `json:"alt"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
				} `json:"contentImageUrls"`
				IsTop              interface{} `json:"isTop"`
				Hot                bool        `json:"hot"`
				IsGilded           bool        `json:"isGilded"`
				IsReward           bool        `json:"isReward"`
				Reward             float64     `json:"reward"`
				HasOfferCompareTag bool        `json:"hasOfferCompareTag"`
				StaffAnswer        bool        `json:"staffAnswer"`
				WithAnonymousOffer bool        `json:"withAnonymousOffer"`
				IsAnonymousFlag    bool        `json:"isAnonymousFlag"`
				IsWithAcceptFlag   interface{} `json:"isWithAcceptFlag"`
				PostCertify        int         `json:"postCertify"`
				EntityId           int         `json:"entityId"`
				EntityType         int         `json:"entityType"`
				NewReferral        interface{} `json:"newReferral"`
				CreateTime         int64       `json:"createTime"`
				EditTime           int64       `json:"editTime"`
				RecommendAd        bool        `json:"recommendAd"`
				Edited             bool        `json:"edited"`
				ShowTime           int64       `json:"showTime"`
			} `json:"contentData,omitempty"`
		} `json:"records"`
	} `json:"data"`
}

func getContent(uuid string) (content ContentResp, err error) {

	// generate by apifox

	url := "https://gw-c.nowcoder.com/api/sparta/detail/moment-data/detail/%s"
	url = fmt.Sprintf(url, uuid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("dnt", "1")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "gw-c.nowcoder.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &content)
	if err != nil {
		return
	}
	return
}

type ContentResp struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		Ip4         string `json:"ip4"`
		Ip4Location string `json:"ip4Location"`
		Id          int    `json:"id"`
		EntityId    int    `json:"entityId"`
		EntityType  int    `json:"entityType"`
		Uuid        string `json:"uuid"`
		UserId      int    `json:"userId"`
		UserBrief   struct {
			UserId           int         `json:"userId"`
			Nickname         string      `json:"nickname"`
			Admin            bool        `json:"admin"`
			Followed         bool        `json:"followed"`
			HeadImgUrl       string      `json:"headImgUrl"`
			Gender           string      `json:"gender"`
			HeadDecorateUrl  string      `json:"headDecorateUrl"`
			HonorLevel       int         `json:"honorLevel"`
			HonorLevelName   string      `json:"honorLevelName"`
			HonorLevelColor  string      `json:"honorLevelColor"`
			WorkTime         string      `json:"workTime"`
			EducationInfo    string      `json:"educationInfo"`
			IdentityList     interface{} `json:"identityList"`
			ActivityIconList []struct {
				ImgUrl      string      `json:"imgUrl"`
				ExpireTime  interface{} `json:"expireTime"`
				Name        string      `json:"name"`
				DiscussLink string      `json:"discussLink"`
				AppLink     string      `json:"appLink"`
				PcLink      string      `json:"pcLink"`
				Content     interface{} `json:"content"`
				Source      int         `json:"source"`
				Type        int         `json:"type"`
			} `json:"activityIconList"`
			ActivityIconListV2 []struct {
				ImgUrl      string      `json:"imgUrl"`
				ExpireTime  interface{} `json:"expireTime"`
				Name        string      `json:"name"`
				DiscussLink *string     `json:"discussLink"`
				AppLink     *string     `json:"appLink"`
				PcLink      *string     `json:"pcLink"`
				Content     interface{} `json:"content"`
				Source      *int        `json:"source"`
				Type        *int        `json:"type"`
			} `json:"activityIconListV2"`
			MemberIdentity  int         `json:"memberIdentity"`
			MemberStartTime interface{} `json:"memberStartTime"`
			MemberEndTime   interface{} `json:"memberEndTime"`
			Member          interface{} `json:"member"`
			AuthDisplayInfo string      `json:"authDisplayInfo"`
			EnterpriseInfo  interface{} `json:"enterpriseInfo"`
			BadgeIconUrl    interface{} `json:"badgeIconUrl"`
			NicknameStyle   struct {
				Direction string `json:"direction"`
				Colors    []struct {
					Light string      `json:"light"`
					Dark  interface{} `json:"dark"`
				} `json:"colors"`
			} `json:"nicknameStyle"`
			CardActivityIcon struct {
				Icon         string `json:"icon"`
				IconDark     string `json:"iconDark"`
				ActivityName string `json:"activityName"`
				Router       string `json:"router"`
			} `json:"cardActivityIcon"`
		} `json:"userBrief"`
		Title         string      `json:"title"`
		Content       string      `json:"content"`
		Type          int         `json:"type"`
		LinkMoment    interface{} `json:"linkMoment"`
		ImgMoment     interface{} `json:"imgMoment"`
		ClockMoment   interface{} `json:"clockMoment"`
		VideoMoment   interface{} `json:"videoMoment"`
		CreatedAt     int64       `json:"createdAt"`
		VoteId        int         `json:"voteId"`
		Circle        interface{} `json:"circle"`
		EditTime      int64       `json:"editTime"`
		ClientSystem  interface{} `json:"clientSystem"`
		SubjectData   interface{} `json:"subjectData"`
		FrequencyData struct {
			LikeCnt         int  `json:"likeCnt"`
			FollowCnt       int  `json:"followCnt"`
			CommentCnt      int  `json:"commentCnt"`
			TotalCommentCnt int  `json:"totalCommentCnt"`
			ViewCnt         int  `json:"viewCnt"`
			ShareCnt        int  `json:"shareCnt"`
			IsLike          bool `json:"isLike"`
			IsFollow        bool `json:"isFollow"`
			FlowerData      struct {
				RemainingFlowers    int  `json:"remainingFlowers"`
				SendFlowerUserCount int  `json:"sendFlowerUserCount"`
				SendFlowerCount     int  `json:"sendFlowerCount"`
				IsSendFlower        bool `json:"isSendFlower"`
				SendFlowerUsers     []struct {
					UserId  int    `json:"userId"`
					HeadImg string `json:"headImg"`
				} `json:"sendFlowerUsers"`
			} `json:"flowerData"`
		} `json:"frequencyData"`
		JobIds      interface{} `json:"jobIds"`
		CareerType  int         `json:"careerType"`
		ShowMessage struct {
			ShowContent bool        `json:"showContent"`
			Message     interface{} `json:"message"`
		} `json:"showMessage"`
		HasEdit           bool        `json:"hasEdit"`
		AllowAnonymous    bool        `json:"allowAnonymous"`
		CheckSalary       bool        `json:"checkSalary"`
		Edited            bool        `json:"edited"`
		ShowTime          int64       `json:"showTime"`
		InternalRecommend interface{} `json:"internalRecommend"`
		QuickReviews      []struct {
			Comment string `json:"comment"`
			Type    int    `json:"type"`
		} `json:"quickReviews"`
	} `json:"data"`
}
