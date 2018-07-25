package controllers

import (
	"auth-service/models"
	"encoding/json"
	"errors"
	"net/rpc"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

// Регистрация нового пользователя
type RegistrationController struct {
	beego.Controller
}

type RegistrationUserModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type ActivationUser struct {
	Id       int
	Nickname string
}

// activation channel
var ac chan ActivationUser

var dataServiceRpc string

// TODO: this block is not tested
// Вызов функции активации на главном сервисе через RPC,
// а так же поддержание соединения и повторные попытки активации
func StartActivationCycle() {
	var client *rpc.Client
	var err error
	var lastuse time.Time
	connected := false
	for {
		select {
		case user, ok := <-ac:
			if !ok {
				beego.Critical("Activation channel is corrupted")
				panic("Activation channel is corrupted")
			} else {
				if !connected {
					if client, err = rpc.DialHTTP("tcp", dataServiceRpc); err != nil {
						beego.Critical("Can't connect to RPC Service:", err.Error())
					} else {
						connected = true
					}
				}
				added := false
				if connected {
					err = client.Call("UserActivationService.Activate", &user, nil)
					if err != nil {
						beego.Warning("User activation error on main-service", err.Error())
					} else {
						added = true
					}
				}
				// Активация в главном сервисе не удалась
				// записываем в бд, что бы попробовать в следующий раз
				if !added {
					go AddForActivationQueue(user)
				}
				lastuse = time.Now()
			}
		case t := <-time.After(time.Second * 30):
			if connected {
				// Close connection after 30 second of idling
				if t.Sub(lastuse) > time.Second*30 {
					client.Close()
					connected = false
				}
			}
		case <-time.After(time.Minute * 5):
			go LoadActivationQueue()
		}
	}
	client.Close()
}

func AddForActivationQueue(user ActivationUser) {
	u := models.NotActivatedOnMainServiceUser{
		Nickname: user.Nickname,
		UserId:   user.Id,
	}
	err := u.Insert()
	if err != nil {
		beego.Critical("Fatal error when activating new user", err.Error())
	}
}

// Блокирует остальные вызовы функции LoadActivationQueue
// Если большое количество пользователей или медленное соединение,
// то повторная попытка активации может пойти дольше, чем 5 минут => функция активации вызовется снова => рассинхрон данных
var activationMutex sync.Mutex

// Повторная попытка активации пользователей
func LoadActivationQueue() {
	activationMutex.Lock()
	defer activationMutex.Unlock()

	users, err := models.GetAllPendingForActivationUsers()
	if err != nil {
		beego.Critical("Can't read pending users", err.Error())
	} else {
		for _, user := range users {
			x := ActivationUser{
				Id:       user.UserId,
				Nickname: user.Nickname,
			}
			ac <- x
			err := user.Delete()
			if err != nil {
				beego.Critical(err.Error())
			}
		}
	}
}

func init() {
	go StartActivationCycle()
	ac = make(chan ActivationUser)
	dataServiceConf, err := config.NewConfig("ini", "conf/rpc.conf")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
	dataServiceRpc = dataServiceConf.String("data-service-rpc")
}

func MainServiceActivateUser(id int, nickname string) {
	ac <- ActivationUser{
		Nickname: nickname,
		Id:       id,
	}
}

var unactivated_users map[string]RegistrationUserModel

func init() {
	unactivated_users = make(map[string]RegistrationUserModel)
}

func NewUser(pass string, usr RegistrationUserModel) error {
	u := models.User{
		Login: usr.Login,
	}
	if u.FindByLogin() {
		return errors.New("User with this nickname is already registered")
	}
	if _, err := unactivated_users[pass]; err {
		return errors.New("Already in")
	}
	usr.Password = CustomStr(usr.Password).ToSHA1()
	unactivated_users[pass] = usr
	return nil
}

// TODO: add activation to main-service
func ActivateUser(pass string) error {
	if user, ok := unactivated_users[pass]; ok {
		delete(unactivated_users, pass)
		usr := models.User{
			Login:    user.Login,
			Password: user.Password,
		}
		id, err := usr.Insert()
		if err != nil {
			return err
		} else {
			go MainServiceActivateUser(id, user.Nickname)
			return nil
		}
	} else {
		return errors.New("Wrong pass")
	}
}

func (c *RegistrationController) Register() {
	var u RegistrationUserModel

	beego.Trace("Start Register")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err != nil {
		beego.Debug("Can't Unmarshal:", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		beego.Trace("Correct Unmarshal")

		if u.Login == "" || u.Password == "" || u.Nickname == "" {
			beego.Debug("login:", u.Login)
			beego.Debug("password:", u.Password)
			beego.Debug("nickname:", u.Nickname)
			c.Data["json"] = "Wrong Login/Password/Nickname"
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else if len(u.Password) < 6 {
			beego.Debug("login:", u.Login)
			beego.Debug("password:", u.Password)
			beego.Debug("nickname:", u.Nickname)
			c.Data["json"] = "Password is too short"
			c.Ctx.Output.SetStatus(HTTP_NOT_ACCEPTABLE)
		} else {
			beego.Trace("Valid combo login & password & nickname")

			// Создает токен для подтверждения регистрации, который должен отправляться на email
			pass := GenerateNewToken(15)
			beego.Trace(pass)
			err := NewUser(pass, u)
			if err != nil {
				beego.Debug("Register error:" + err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			} else {

				if err = models.SendingRegistrationToken(pass, u.Login); err != nil {
					beego.Debug("Register error:" + err.Error())
					c.Data["json"] = err.Error()
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				} else {
					c.Data["json"] = MakeMessageForSending("OK")
					beego.Trace("Register OK")
				}
			}
		}
	}
	c.ServeJSON()
}

func (c *RegistrationController) Activate() {
	pass := c.GetString("pass")
	if pass != "" {
		err := ActivateUser(pass)
		if err != nil {
			beego.Debug("Activation error: " + err.Error())
			c.Data["json"] = MakeMessageForSending(err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			beego.Trace("Activation OK")
			c.Data["json"] = MakeMessageForSending("Registered")
		}
	} else {
		beego.Debug("Empty pass")
		c.Data["json"] = MakeMessageForSending(HTTP_BAD_REQUEST_STR)
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	}
	c.ServeJSON()
}

func (c *RegistrationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// mock
// Post ...
// @Title Register
// @Description Проводит преварительную регистрацию пользователя, после требуется подтверждение
// @Param	body	body	controllers.RegistrationUserModel	true	"Никнейм, логин(email) и пароль обязательны" ""
// @Success	200 "OK"
// @Failure	400 "This user is already registered"
// @router / [post]
func (c *RegistrationController) Post() {
	c.Register()
}

// mock
// Get ...
// @Title Activate
// @Description Активирует аккаунт
// @Param	pass	query	string	false	"Код для активации аккаунта"
// @Failure 200 OK
// @router / [get]
func (c *RegistrationController) Get() {
	c.Activate()
}
