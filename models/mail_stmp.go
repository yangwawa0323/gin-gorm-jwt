package models

// func Send() {

// 	auth = smtp.PlainAuth("", adminEmail, adminMailboxPassword, mailboxHost)

// 	templateData := GenerateTemplateData()

// 	req := NewMailDialer([]string{"yangwawa0323@163.com", "12238747@qq.com"},
// 		"Welcome to register 51cloudclass.com",
// 		"",
// 	)
// 	if err := req.ParseTemplate(mailTemplateFile, templateData); err != nil {
// 		utils.ErrorDebug( err, "can not read the template")
// 	}

// 	ok, _ := req.SendEmail()
// 	utils.Debug(strconv.FormatBool(ok))

// }

// func GenerateTemplateData() *TemplateData {
// 	var templateData TemplateData = make(TemplateData)
// 	templateData["ActivateString"] = fmt.Sprintf("%s%s", activateUrl, secret)
// 	return &templateData
// }

// func (req *MailDialer) SendEmail() (bool, error) {
// 	mime := "MIME-version: 1.0;\n"
// 	mime += "Content-Type: text/plain; charset=\"UTF-8\";\n\n"

// 	subject := "Subject: " + req.Subject + "!\n"
// 	msg := []byte(subject + mime + "\n" + req.Body)

// 	if err := smtp.SendMail(
// 		fmt.Sprintf("%s:%d", mailboxHost, mailboxPort),
// 		auth,
// 		adminEmail,
// 		req.To,
// 		msg); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// func (req *MailDialer) ParseTemplate(templateFileName string, data interface{}) error {
// 	t, err := template.ParseFiles(templateFileName)
// 	if err != nil {
// 		return err
// 	}
// 	buf := new(bytes.Buffer)
// 	if err = t.Execute(buf, data); err != nil {
// 		return err
// 	}

// 	req.Body = buf.String()
// 	return nil
// }
