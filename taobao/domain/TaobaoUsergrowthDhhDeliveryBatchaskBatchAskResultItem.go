package domain

type TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem struct {
	/*在巨浪平台可投放的任务ID列表     */
	TaskIdList *[]string `json:"task_id_list,omitempty" `

	/*oaid的md5值， 32位小写     */
	OaidMd5 *string `json:"oaid_md5,omitempty" `

	/*idfa的md5值， 32位小写     */
	IdfaMd5 *string `json:"idfa_md5,omitempty" `

	/*imei的md5值， 32位小写     */
	ImeiMd5 *string `json:"imei_md5,omitempty" `

	/*该设备要做的大航海的任务id     */
	TaskId *string `json:"task_id,omitempty" `
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) SetTaskIdList(v []string) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem {
	s.TaskIdList = &v
	return s
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) SetOaidMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem {
	s.OaidMd5 = &v
	return s
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) SetIdfaMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem {
	s.IdfaMd5 = &v
	return s
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) SetImeiMd5(v string) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem {
	s.ImeiMd5 = &v
	return s
}

func (s *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem) SetTaskId(v string) *TaobaoUsergrowthDhhDeliveryBatchaskBatchAskResultItem {
	s.TaskId = &v
	return s
}
