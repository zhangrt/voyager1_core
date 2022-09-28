package validate

import "time"

type TestModel struct {
	CreateDate time.Time `json:"create_date" comment:"创建时间" en_comment:"CreateDate" validate:"beforeCurrentDate"`
	EndDate    time.Time `json:"end_date" comment:"结束时间" en_comment:"EndDate" validate:"gtefield=CreateDate"`
	CardId     string    `json:"card_id" comment:"身份证号" en_comment:"CardId" validate:"isCardId"`
	Phone      string    `json:"phone" comment:"电话" en_comment:"Phone" validate:"isPhoneNo"`
	PassportNo string    `json:"passport_no" comment:"护照号码" en_comment:"PassportNo" validate:"isPassportNo"`
}
