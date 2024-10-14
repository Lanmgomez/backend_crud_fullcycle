package user

type USERSCRUD struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	ActiveUser string `json:"activeUser"`
}

type USERS struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type USERLOGIN struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LOGINLOGS struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	UserAgent string `json:"userAgent"`
	LoginTime string `json:"loginTime"`
	Status    string `json:"status"`
	IpAddress string `json:"ipAddress"`
}

type USER_IDENTIFICATION_CONTACT struct {
	Id            int64  `json:"id"`
	FullName      string `json:"fullName"`
	CpfOrCnpj     string `json:"cpfOrCnpj"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	AddressNumber string `json:"addressNumber"`
	Complement    string `json:"complement"`
	Neighborhood  string `json:"neighborhood"`
	City          string `json:"city"`
	Uf            string `json:"uf"`
	ZipCode       string `json:"zipCode"`
}

type PAYMENT_METHOD struct {
	Id                     int64  `json:"id"`
	PaymentUserId          int64  `json:"paymentUserId"`
	PaymentFormInstallment string `json:"paymentFormInstallment"`
	Token                  string `json:"token"`
	DateTime               string `json:"dateTime"`
}

type PAYMENTS struct {
	UserIdentification USER_IDENTIFICATION_CONTACT
	PaymentForm        PAYMENT_METHOD
}

type UF_STATES struct {
	Id    int    `json:"id"`
	State string `json:"state"`
	Uf    string `json:"uf"`
}