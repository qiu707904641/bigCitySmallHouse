package anjuke

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const SingleUrl = "https://apirent.anjuke.com/zufang/wechat/rent/api_get_detail?infoId=2671816932039689&openId=ocS7q0JWrakjtDislkipP_WP-OeU&signature=51757b539deef4e21361bbb9ba0f3099&city_id=12&isAuction=2&source_type=1&legoAuction=%7B%22isauction%22%3A%222%22%2C%22lego%22%3A%7B%22lego_tid%22%3A%222d9c9757-2572-427c-ae2e-d93c5b6bd9b0%22%2C%22lego_ad_click_url%22%3A%22https%3A%5C%2F%5C%2Flegoclick.58.com%5C%2Fjump%3Ftarget%3DpZwY0ZnlsztdraOWUvYKuaYzmvD1mW7WuBdhuWK-sHwbrHEVmyNkPzYkrj6BmhcdPH0QnHDKP1TdP1b3njT1njmYPjmKTHcvP1D3nHmOn1ckn1bvrjbKrH0knTDOP1TkTHD_nHTKnkDQPWmzrH9LrHNkP10OTHDKwbnVHidKibfMGO4hBs4hbF1MV2s-BFxCCpWGCUNKnEDQTED1TyqWN1IQnDGg0h7opMwDpgP_pv-kN7qgNad5uRNKnHNvrHnkn1b3PHnknHNdn1DdnkDvTEDQTyD3P1c1mhDksHbQnjcVPjuhraY3rjNzsHm3uhcdP1Dzryw6m9DQPHmOn1T1rH9vP1cLPWT3n1ckTHDdPWb1njnOrjmvPjnLnjbOPHcKTHD_nHTKTEDKsEDKTEDKnHDOsWD1na3znjm8nH9OTHTKnTDKnikQnE7exEDQnjT1P9DQnjTQPWbzTHFbrynOP1NLsHcdP1cVPjcLmzd6uHF-syEOnvndmWuBuj-BnTDKPT7LXjTOryNkPWELuW-6Pj0QPvEKmyGoRvRWpA7YTHTKnzkvPHTQsjnLPHTkTHDKUMR_UTD1njE3uW9vuWEznAFbnvRB%22%2C%22lego_sid%22%3A%22127068352217971156682404352%22%2C%22platformCompany%22%3A%22ajk%22%2C%22lego_show_track_url_base%22%3A%22https%3A%5C%2F%5C%2Fadtrack.58.com%5C%2F1001692%5C%2F%22%2C%22lego_version%22%3A%221.0.1%22%2C%22legoHuiDu%22%3Atrue%7D%7D&unionId=o9PQht7TUVrcS-LSyJL5LLJ7LLUc&lego_appname=ajkWechat&lego_appid=wx099e0647f9a4717d&lego_tid=undefined&tid=2d9c9757-2572-427c-ae2e-d93c5b6bd9b0&dataSource=undefined&appname=wx&platform=ajkplugin&sidDict=%7B%22ab_test%22%3A%22%22%2C%22cate2%22%3A%2210%22%2C%22cate1%22%3A%221%22%2C%22GTID%22%3A%22127068352217971156682404352%22%2C%22sessionId%22%3A%22175652590217971156695500370%22%2C%22sid%22%3A%22127068352217971156682404352%22%2C%22cityid1_58%22%3A3%2C%22nameoflist%22%3A%2212_index-weinituijian-b%22%2C%22pagesource%22%3A%2212_index-weinituijian-b%22%2C%22recomshowlog%22%3A%22%22%2C%22houseid_list%22%3A%5B%222671816932039689%22%2C%222650512503811076%22%2C%222678726142290946%22%2C%222639429207600134%22%2C%222609742790077444%22%2C%222647480225047566%22%2C%222649082506798091%22%2C%222619486814120963%22%2C%222674271640644619%22%2C%222652096652893195%22%5D%2C%22page%22%3A%221%22%2C%22list_keywords%22%3A%7B%22search%22%3A%7B%22keyword%22%3A%22%22%7D%2C%22filter%22%3A%5B%5D%7D%2C%22os%22%3A%22android%22%7D&soj=%7B%22infoid%22%3A%222671816932039689%22%2C%22houseid%22%3A%222671816932039689%22%2C%22GTID%22%3A%22127068352217971156682404352%22%2C%22is_biz%22%3Atrue%2C%22is_down%22%3A%22%22%2C%22slot%22%3A%22ajk_rent_miniapp_index_search%22%2C%22sid%22%3A%22127068352217971156682404352%22%2C%22ad_type%22%3A%22gz%22%2C%22is_business%22%3A%22%22%2C%22gpos%22%3A%221%22%2C%22pos%22%3A%221%22%2C%22shengxinzu%22%3A%22%22%2C%22shidiheyanzhuangtai%22%3A%221%22%2C%22qingheyanzhuangtai%22%3A%22%22%2C%22recomshowlog%22%3A%22%22%2C%22tid%22%3A%222d9c9757-2572-427c-ae2e-d93c5b6bd9b0%22%2C%22anxuan%22%3A%221%22%2C%22qiyeanxuan%22%3A%22%22%2C%22isVr%22%3A%22%22%2C%22info_type%22%3A%226%22%7D&pageType=zufang_detail&inFrom=index-weinituijian-b&outFrom="

type SingleParser struct {
	*crawler.SingleParser
}

func NewSingleParser(param *crawler.SingleParam) *SingleParser {
	parser := crawler.NewSingleParser(param)
	return &SingleParser{
		SingleParser: parser,
	}
}

func (receiver *SingleParser) Parse() (*house.House, error) {
	single, err := receiver.fetch()
	if err != nil {
		return nil, err
	}
	log.Println(single)
	return nil, nil
}

func (receiver *SingleParser) fetch() (*Single, error) {
	req, err := http.NewRequest(http.MethodGet, SingleUrl, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("infoId", receiver.Param.Id)
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Cookie", "id58=CrIfoWMfMD59zgV6LZxVAg==")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var single Single
	err = json.Unmarshal(body, &single)
	if err != nil {
		return nil, err
	}

	return &single, nil
}
