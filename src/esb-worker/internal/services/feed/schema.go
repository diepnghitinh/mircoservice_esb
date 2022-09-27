package feed

type PushManageMsgContractReminderData struct {
	PropertyId string `json:"property_id"`
	PropertyName string `json:"property_name"`
	ContractCount int `json:"contract_count"`
}

type PushManageMsgContractReminder struct {
	Type string `json:"type"`
	RefId string `json:"ref_id"`
	Data *PushManageMsgContractReminderData `json:"data"`
	UserIds []string `json:"user_ids"`
	Company string `json:"company_title"`
}

type PushManageMsgData struct {
	Title string `json:"title"`
	Data interface{} `json:"data"`
	ApiSource string `json:"api_source"`
	Tag string `json:"tag"`
	NameSpace string `json:"namespace"`
}

type PushManageMsg struct {
	Type string `json:"type"`
	Data *PushManageMsgData `json:"data"`
}