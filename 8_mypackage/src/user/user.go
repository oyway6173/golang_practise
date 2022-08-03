package user

var userId int64 = 10 //если переменная/струтура/функция начинается с маленькой буквы, то она не экспортируемая

var Username = "Nikita" // экспортируемая

const (
	TypeAdmin = iota + 1
	TypeUser
)

const (
	noAccess = 1
)
