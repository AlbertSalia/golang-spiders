package model

type Position struct {
	Url          string //招聘详情页
	Type         string //类别  招聘网站
	Id           string //招聘岗位编号
	PositionName string //岗位名称
	Time         string //发布时间
	CompanyUrl   string //公司url
	CompanyName  string //公司名称
	Require      string //照片要求
	Labels       string //岗位标签
	Weal         string //岗位福利
	Address      string //工作地址
	Ctiy         string //城市
}
