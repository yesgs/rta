package domain

type TaobaoTbkRtaConsumerMatchData struct {
	/*
	   返回结果封装对象     */
	ResultList         *[]string                           `json:"result_list,omitempty" `
	StrategyResultList *[]TaobaoTbkRtaConsumerMatchMapData `json:"strategy_result_list,omitempty" `
}

type TaobaoTbkRtaConsumerMatchMapData struct {
	Status     string `json:"status"`
	StrategyId string `json:"strategy_id"`
}
