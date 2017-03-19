package web

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/ngenerio/instantly/pkg/models"
	"github.com/ngenerio/instantly/pkg/web/payloads"
	log "github.com/sirupsen/logrus"
)

var ErrEmailExists error = errors.New("Email already exists")
var ErrInvalidCredentials error = errors.New("Invalid email or password")
var ErrInternal error = errors.New("Something happened. Please try again")

func HomeHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Home - Instantly"
	session := c.Get("session").(*sessions.Session)
	id := session.Values["id"].(int)

	transactions, err := models.GetUserTransactions(id)

	if err != nil {
		log.WithFields(log.Fields{
			"user_id": id,
		}).Error("Error retrieving user transactions")
		return err
	}

	var moneyIn, moneyOut float64
	var totalTransactions, totalFailedTransactions int
	var totalMTN, totalVodafone, totalTigo, totalAirtel float64

	for _, trx := range transactions {
		if trx.Status == models.StatusFailed || trx.Status == models.StatusPending {
			if trx.Status == models.StatusFailed {
				totalFailedTransactions += 1
			}
			continue
		}

		if trx.Type == models.Credit {
			moneyIn += trx.Amount
		} else {
			moneyOut += trx.Amount
		}

		switch trx.MNO {
		case "MTN":
			totalMTN += 1
		case "VODAFONE":
			totalVodafone += 1
		case "TIGO":
			totalTigo += 1
		case "AIRTEL":
			totalAirtel += 1
		}
		totalTransactions += 1
	}

	params.Data = make(map[string]interface{})
	params.Data["MoneyIn"] = moneyIn
	params.Data["MoneyOut"] = moneyOut
	params.Data["TotalTransactions"] = totalTransactions
	params.Data["TotalFailedTransactions"] = totalFailedTransactions
	params.Data["Page"] = "home"
	if totalMTN == float64(0) {
		params.Data["TotalMTN"] = 0
	} else {
		params.Data["TotalMTN"] = (totalMTN / float64(totalTransactions)) * 100
	}

	if totalVodafone == float64(0) {
		params.Data["TotalVodafone"] = 0
	} else {
		params.Data["TotalVodafone"] = (totalVodafone / float64(totalTransactions)) * 100
	}

	if totalTigo == float64(0) {
		params.Data["TotalTigo"] = 0
	} else {
		params.Data["TotalTigo"] = (totalTigo / float64(totalTransactions)) * 100
	}

	if totalAirtel == float64(0) {
		params.Data["TotalAirtel"] = 0
	} else {
		params.Data["TotalAirtel"] = (totalAirtel / float64(totalTransactions)) * 100
	}
	return c.Render(http.StatusOK, "index", params)
}

func TransactionsHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Transactions - Instantly"
	session := c.Get("session").(*sessions.Session)
	id := session.Values["id"].(int)

	transactions, err := models.GetUserTransactions(id)

	if err != nil {
		log.WithFields(log.Fields{
			"user_id": id,
		}).Error("Error retrieving user transactions")
		return err
	}

	params.Data = make(map[string]interface{})
	params.Data["Page"] = "transactions"
	params.Data["Transactions"] = transactions

	return c.Render(http.StatusOK, "transactions", params)
}

func SettingsHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Settings - Instantly"
	params.Data = make(map[string]interface{})
	params.Data["Page"] = "settings"
	session := c.Get("session").(*sessions.Session)
	params.Flashes = session.Flashes()
	session.Save(c.Request(), c.Response())
	user := c.Get("user").(*models.User)
	params.Data["User"] = user
	return c.Render(http.StatusOK, "settings", params)
}

func SaveSettings(c echo.Context) error {
	sett := new(payloads.Settings)
	if err := c.Bind(sett); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/settings")
	}

	user := c.Get("user").(*models.User)
	user.CallbackURL = sett.CallbackURL
	user.Token = sett.Token
	user.NetworkOperator = sett.NetworkOperator
	user.MobileNumber = sett.MobileNumber
	err := user.Update()

	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/settings")
	}

	SetFlash(c, c.Response().Writer(), c.Request(), "success", "Settings have been saved")
	return c.Redirect(http.StatusFound, "/settings")
}

func LoginHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Login - Instantly"
	session := c.Get("session").(*sessions.Session)
	params.Flashes = session.Flashes()
	session.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "login", params)
}

func RegisterHandler(c echo.Context) error {
	params := new(Params)
	params.Title = "Register - Instantly"
	session := c.Get("session").(*sessions.Session)
	params.Flashes = session.Flashes()
	session.Save(c.Request(), c.Response().Writer())
	return c.Render(http.StatusOK, "signup", params)
}

func RegisterUser(c echo.Context) error {
	user := new(payloads.User)
	if err := c.Bind(user); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	exists, err := models.DoesUserExist(map[string]interface{}{"email_address": user.Email})
	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	if exists {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrEmailExists.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	dbUser, err := models.CreateUser(user)
	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/register")
	}

	log.Info(fmt.Sprintf("New user has been created %v", dbUser))
	session := c.Get("session").(*sessions.Session)
	session.Values["id"] = dbUser.ID
	session.Save(c.Request(), c.Response().Writer())
	c.Redirect(http.StatusFound, "/")
	return nil
}

func LoginUser(c echo.Context) error {
	user := new(payloads.User)
	if err := c.Bind(user); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInternal.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	dbUser := new(models.User)
	if err := dbUser.GetUser(map[string]interface{}{"email_address": user.Email}); err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInvalidCredentials.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))

	if err != nil {
		SetFlash(c, c.Response().Writer(), c.Request(), "error", ErrInvalidCredentials.Error())
		return c.Redirect(http.StatusFound, "/login")
	}

	log.Info(fmt.Sprintf("User with email address %s has logged in", user.Email))
	session := c.Get("session").(*sessions.Session)
	session.Values["id"] = dbUser.ID
	session.Save(c.Request(), c.Response().Writer())
	c.Redirect(http.StatusFound, "/")
	return nil
}

func Logout(c echo.Context) error {
	session := c.Get("session").(*sessions.Session)
	delete(session.Values, "id")
	SetFlash(c, c.Response().Writer(), c.Request(), "success", "You have successfully logged out")
	c.Redirect(http.StatusFound, "/")
	return nil
}
