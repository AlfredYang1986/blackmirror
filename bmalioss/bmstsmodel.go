package bmalioss

type BmSTS struct {
	RequestId string	`json:"RequestId"`
	AssumedRoleUser map[string]interface{} `json:"AssumedRoleUser"`
	Credentials map[string]interface{} `json:"Credentials"`
}

func (b BmSTS) GetAccessKeyId() string {
	return b.Credentials["AccessKeyId"].(string)
}

func (b BmSTS) GetAccessKeySecret() string {
	return b.Credentials["AccessKeySecret"].(string)
}

func (b BmSTS) GetSecurityToken() string {
	return b.Credentials["SecurityToken"].(string)
}

func (b *BmSTS) ResetSecurityProp(mp map[string]string) error {
	tm := make(map[string]interface{}, 3)
	for k, v := range mp {
		tm[k] = v
	}
	b.Credentials = tm
	return nil
}
