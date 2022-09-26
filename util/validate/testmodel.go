package validate

import "time"

type TestModel struct {
	CreateDate time.Time `json:"create_date" comment:"创建时间" validate:"beforeCurrentDate"`
	EndDate    time.Time `json:"end_date" comment:"结束时间" validate:"gtefield=CreateDate"`
	CardId     string    `json:"card_id" comment:"身份证号" validate:"isCardId"`
	Phone      string    `json:"phone" comment:"电话" validate:"isPhoneNo"`
	PassportNo string    `json:"passport_no" comment:"护照号码" validate:"isPassportNo"`
}
