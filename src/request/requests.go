package request

import (
	"HostLoc-Daily-CheckIn/src/config"
	"fmt"
	"github.com/imroc/req/v3"
	"net/url"
)

func (r *Request) Member() (*req.Response, error) {
	u := url.Values{}

	u.Add("mod", "logging")
	u.Add("action", "login")
	u.Add("infloat", "yes")
	u.Add("handlekey", "login")
	u.Add("inajax", "1")
	u.Add("ajaxtarget", "fwin_content_login")

	return r.client.R().
		SetQueryString(u.Encode()).
		Get("https://hostloc.com/member.php")
}

func (r *Request) MainPage() (*req.Response, error) {
	return r.client.R().
		Get("https://hostloc.com")
}

func (r *Request) Login(account *config.Accounts, formhash string) (*req.Response, error) {
	u := url.Values{}

	u.Add("mod", "logging")
	u.Add("action", "login")
	u.Add("loginsubmit", "yes")
	u.Add("infloat", "yes")
	u.Add("lssubmit", "yes")
	u.Add("inajax", "1")

	data := map[string]string{
		"fastloginfield": "username",
		"username":       account.Username,
		"cookietime":     "2592000",
		"password":       account.Password,
		"formhash":       formhash,
		"quickforward":   "yes",
		"handlekey":      "ls",
	}

	return r.client.R().
		SetQueryString(u.Encode()).
		SetFormData(data).
		SetCookies(r.cookies...).
		Post("https://hostloc.com/member.php")

}

func (r *Request) CheckCoin() (*req.Response, error) {
	u := url.Values{}

	u.Add("mod", "spacecp")
	u.Add("ac", "credit")
	u.Add("showcredit", "1")
	u.Add("inajax", "1")

	return r.client.R().
		SetQueryString(u.Encode()).
		SetCookies(r.cookies...).
		Get("https://hostloc.com/home.php")
}

func (r *Request) Space(uid int) (*req.Response, error) {
	return r.client.R().
		SetCookies(r.cookies...).
		Get(fmt.Sprintf("https://hostloc.com/space-uid-%v.html", uid))
}
