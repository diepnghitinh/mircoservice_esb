package test

import (
	"github.com/grokify/html-strip-tags-go"
	"github.com/labstack/gommon/log"
	"regexp"
	"strings"
	"testing"
)

func TestStriptags(t *testing.T) {

	content := `
		<p class="description" style="margin: 10px 0px; padding: 0px 0px 5px; text-rendering: geometricprecision; color: rgb(68, 68, 68); font-variant-numeric: normal; font-variant-east-asian: normal; font-weight: 700; font-stretch: normal; line-height: 18px; font-family: arial;">Ba ngày qua, nhiều người dân Đà Nẵng thức cả đêm để hứng từng giọt nước sinh hoạt nhưng số nước này cũng chỉ đủ tắm, giặt ở mức tối thiểu.</p><article class="content_detail fck_detail width_common block_ads_connect" style="margin: 0px; padding: 0px 0px 10px; text-rendering: geometricprecision; width: 500px; float: left; color: rgb(51, 51, 51); font-family: arial;"><p class="Normal" style="margin-right: 0px; margin-bottom: 1em; margin-left: 0px; padding: 0px; text-rendering: geometricprecision; font-variant-numeric: normal; font-variant-east-asian: normal; font-stretch: normal; line-height: 18px;">4h sáng ngày 21/8, vợ chồng ông Đặng Em (57 tuổi) thức giấc chuẩn bị đồ nấu bánh canh và xôi bán buổi sáng, nhưng khi vặn vòi lấy nước để chế biến, ông thấy nước chảy nhỏ giọt. Một tiếng sau, nước bị mất hoàn toàn.&nbsp;</p><p class="Normal" style="margin-right: 0px; margin-bottom: 1em; margin-left: 0px; padding: 0px; text-rendering: geometricprecision; font-variant-numeric: normal; font-variant-east-asian: normal; font-stretch: normal; line-height: 18px;">Quán bánh canh buổi sáng mở hàng muộn hơn bình thường khoảng một giờ đồng hồ nên vắng khách. Ông Em bán hết lượt bát đũa chuẩn bị sẵn song không có nước để rửa.&nbsp;"Đây là lần đầu tiên nhà tôi lâm vào tình trạng này", ông Em nói.&nbsp;</p><p class="Normal" style="margin-right: 0px; margin-bottom: 1em; margin-left: 0px; padding: 0px; text-rendering: geometricprecision; font-variant-numeric: normal; font-variant-east-asian: normal; font-stretch: normal; line-height: 18px;">Nhà ông ở ngõ số 7 đường Pasteur, phường Hải Châu 1 (quận Hải Châu)&nbsp;- khu vực trung tâm thành phố Đà Nẵng. Hôm 18/8, cũng như hàng trăm nghìn&nbsp;người dân thành phố, ông nhận được tin nhắn từ Công ty Cấp nước Đà Nẵng (Dawaco) thông báo việc&nbsp;giảm công suất cấp nước&nbsp;do nguồn nước thô bị nhiễm mặn. Đây là lần thứ ba kể từ tháng 4 đến nay, Dawaco gửi thông báo như trên. Hai lần trước, khu vực trung tâm thành phố ít bị ảnh hưởng.</p><table align="center" border="0" cellpadding="3" cellspacing="0" class="tplCaption" style="margin: 10px auto; padding: 0px; text-rendering: geometricprecision; max-width: 100%; width: 500px;"><tbody style="margin: 0px; padding: 0px; text-rendering: geometricprecision;"><tr style="margin: 0px; padding: 0px; text-rendering: geometricprecision;"><td style="margin: 0px; padding: 0px; text-rendering: geometricprecision; line-height: 0;"><img alt="Ông Đặng Em chờ cả buổi sáng vẫn không hứng được nước thuỷ cục. Ảnh: Nguyễn Đông." data-natural-h="350" data-natural-width="500" src="https://i-vnexpress.vnecdn.net/2019/08/21/3-3309-1566376487.jpg" data-width="500" data-pwidth="500" style="margin: 0px; padding: 0px; text-rendering: geometricprecision; border: 0px; font-size: 0px; line-height: 0; max-width: 100%;"></td></tr><tr style="margin: 0px; padding: 0px; text-rendering: geometricprecision;"><td style="margin: 0px; padding: 0px; text-rendering: geometricprecision; line-height: 0;"><p class="Image" style="margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 10px; text-rendering: geometricprecision; font-variant-numeric: normal; font-variant-east-asian: normal; font-stretch: normal; font-size: 13px; line-height: 16px; background: rgb(245, 245, 245);">Ông Đặng Em chờ cả buổi sáng vẫn không hứng được nước. Ảnh:&nbsp;<em style="margin: 0px; padding: 0px; text-rendering: geometricprecision;">Nguyễn Đông.</em></p></td></tr></tbody></table><p class="Normal" style="margin-right: 0px; margin-bottom: 1em; margin-left: 0px; padding: 0px; text-rendering: geometricprecision; font-variant-numeric: normal; font-variant-east-asian: normal; font-stretch: normal; line-height: 18px;">Ở bên kia sông Hàn, nhiều hộ dân ở quận Sơn Trà và Ngũ Hành Sơn đã phải chịu cảnh thiếu nước sinh hoạt suốt ba ngày nay. Bà Đặng Thị Anh (57 tuổi, đường Hồ Ngọc Lãm, phường Thọ Quang, quận Sơn Trà) cho biết máy giặt của gia đình bà không thể hoạt động vì không đủ nước, sinh hoạt hàng ngày bị đảo lộn.</p></article>
	`
	content = strip.StripTags(content)
	space := regexp.MustCompile(`\s+`)

	s := space.ReplaceAllString(content, " ")
	s = strings.Replace(s, "&nbsp;", " ", -1)

	//log.Print(strings.Trim(s, " "))

	log.Print(s)

}