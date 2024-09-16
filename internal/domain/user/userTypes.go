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
