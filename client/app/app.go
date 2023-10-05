package app

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gogapopp/gophkeeper/client/grpc_client"
	"github.com/gogapopp/gophkeeper/client/lib/file"
	"github.com/gogapopp/gophkeeper/client/lib/luhn"
	"github.com/gogapopp/gophkeeper/client/service"
	"github.com/gogapopp/gophkeeper/internal/jwt"
	"github.com/gogapopp/gophkeeper/models"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

var applicationErr error

type Application struct {
	userSecretPhrase string
	userID           int
	grpcClient       *grpc_client.GRPCClient
	getService       *service.GetService
	log              *zap.SugaredLogger
}

func NewApplication(grpcClient *grpc_client.GRPCClient, getService *service.GetService, log *zap.SugaredLogger) *Application {
	return &Application{
		grpcClient: grpcClient,
		getService: getService,
		log:        log,
	}
}

func (a *Application) CreateApp() error {
	app := tview.NewApplication()
	registerForm := a.registerForm(app)
	err := app.SetRoot(registerForm, true).EnableMouse(true).Run()
	if err != nil {
		applicationErr = fmt.Errorf("app.CreateApp.%s", err)
	}
	return applicationErr
}

func (a *Application) registerForm(app *tview.Application) *tview.Form {
	registerForm := tview.NewForm()
	registerForm.
		AddTextView("REGISTER:", "Registration form", 20, 2, true, true).
		AddInputField("Login", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Submit", func() {
			login := registerForm.GetFormItemByLabel("Login").(*tview.InputField).GetText()
			password := registerForm.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			user := models.User{
				Login:    login,
				Password: password,
			}
			err := a.grpcClient.Register(context.Background(), user)
			// если поля login и password пустые
			if login == "" || password == "" {
				// если кнопки нет то добавляем
				if registerForm.GetButtonIndex("Incorrect login or password") == -1 {
					registerForm.AddButton("Incorrect login or password", nil)
				}
			} else {
				// проверяем верный ли пароль и логин
				if err == nil {
					// проверяем кнопки "User already exists" и "Incorrect login or password" если есть меняем название
					if registerForm.GetButtonIndex("User already exists") != -1 {
						idx := registerForm.GetButtonIndex("User already exists")
						registerForm.RemoveButton(idx)
					}
					if registerForm.GetButtonIndex("Incorrect login or password") != -1 {
						idx := registerForm.GetButtonIndex("Incorrect login or password")
						registerForm.RemoveButton(idx)
					}
					// если кнопок нет создаём "User registered"
					if registerForm.GetButtonIndex("User registered") == -1 {
						registerForm.AddButton("User registered", nil)
					}
				} else {
					// проверяем кнопки "User already exists" и "Incorrect login or password" если их нет то добавляем
					if registerForm.GetButtonIndex("User already exists") == -1 {
						registerForm.AddButton("User already exists", nil)
					}
					if registerForm.GetButtonIndex("Incorrect login or password") == -1 {
						registerForm.AddButton("Incorrect login or password", nil)
					}
				}
			}
		}).
		AddButton("Go to login", func() {
			loginForm := a.loginForm(app)
			err := app.SetRoot(loginForm, true).EnableMouse(true).Run()
			if err != nil {
				applicationErr = fmt.Errorf("app.registerForm.%s", err)

			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	registerForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return registerForm
}

func (a *Application) loginForm(app *tview.Application) *tview.Form {
	loginForm := tview.NewForm()
	loginForm.
		AddTextView("LOGIN:", "Login form", 20, 2, true, true).
		AddInputField("Login", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddInputField("SecretPhrase", "", 20, nil, nil).
		AddButton("Submit", func() {
			login := loginForm.GetFormItemByLabel("Login").(*tview.InputField).GetText()
			password := loginForm.GetFormItemByLabel("Password").(*tview.InputField).GetText()
			a.userSecretPhrase = loginForm.GetFormItemByLabel("SecretPhrase").(*tview.InputField).GetText()
			user := models.User{
				Login:    login,
				Password: password,
			}
			if login == "" || password == "" {
				if loginForm.GetButtonIndex("Incorrect login or password") == -1 {
					loginForm.AddButton("Incorrect login or password", nil)
					return
				}
			}
			response, err := a.grpcClient.Login(context.Background(), user)
			if err == nil {
				userID, err := jwt.ParseToken(*response.Jwt)
				if err != nil {
					return
				}
				a.userID = userID
			}
			if err == nil {
				if idx := loginForm.GetButtonIndex("Incorrect login or password"); idx != -1 {
					loginForm.RemoveButton(idx)
				}
				if loginForm.GetButtonIndex("Go to Next Page") == -1 {
					loginForm.AddButton("Go to Next Page", func() {
						dataPagesForm := a.dataPagesForm(app)
						if err := app.SetRoot(dataPagesForm, true).EnableMouse(true).Run(); err != nil {
							applicationErr = fmt.Errorf("app.loginForm.%s", err)
						}
					})
				}
			} else {
				if loginForm.GetButtonIndex("Incorrect login or password") == -1 {
					loginForm.AddButton("Incorrect login or password", nil)
				}
			}
		}).
		AddButton("Go to register", func() {
			registerForm := a.registerForm(app)
			if err := app.SetRoot(registerForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.loginForm.%s", err)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	loginForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return loginForm
}

func (a *Application) dataPagesForm(app *tview.Application) *tview.Form {
	dataPagesForm := tview.NewForm()
	dataPagesForm.
		AddButton("Save", func() {
			saveDataPagesForm := a.saveDataPagesForm(app)
			if err := app.SetRoot(saveDataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.dataPagesForm.%s", err)
			}
		}).
		AddButton("Get", func() {
			getDataPagesForm := a.getDataPagesForm(app)
			if err := app.SetRoot(getDataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.dataPagesForm.%s", err)
			}
		}).
		AddButton("Return back", func() {
			loginForm := a.loginForm(app)
			if err := app.SetRoot(loginForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.dataPagesForm.%s", err)
			}
		}).
		AddButton("Sync", func() {
			err := a.grpcClient.SyncData(context.Background(), a.userID)
			if err != nil {
				if dataPagesForm.GetButtonIndex("Error") == -1 {
					dataPagesForm.AddButton("Error", nil)
				}
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	dataPagesForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return dataPagesForm
}

func (a *Application) getDataPagesForm(app *tview.Application) *tview.Form {
	getDataPagesForm := tview.NewForm()
	getDataPagesForm.
		AddButton("Text Data Page", func() {
			getTextDataForm := a.getTextDataForm(app)
			if err := app.SetRoot(getTextDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.getDataPagesForm.%s", err)
			}
		}).
		AddButton("Binary Data Page", func() {
			getBinaryDataForm := a.getBinaryDataForm(app)
			if err := app.SetRoot(getBinaryDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.getDataPagesForm.%s", err)
			}
		}).
		AddButton("Card Data Page", func() {
			getCardDataForm := a.getCardDataForm(app)
			if err := app.SetRoot(getCardDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.getDataPagesForm.%s", err)
			}
		}).
		AddButton("Return back", func() {
			dataPagesForm := a.dataPagesForm(app)
			if err := app.SetRoot(dataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.getDataPagesForm.%s", err)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	getDataPagesForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return getDataPagesForm

}

func (a *Application) saveDataPagesForm(app *tview.Application) *tview.Form {
	saveDataPagesForm := tview.NewForm()
	saveDataPagesForm.
		AddButton("Text Data Page", func() {
			textDataForm := a.textDataForm(app)
			if err := app.SetRoot(textDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.saveDataPagesForm.%s", err)
			}
		}).
		AddButton("Binary Data Page", func() {
			binaryDataForm := a.binaryDataForm(app)
			if err := app.SetRoot(binaryDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.saveDataPagesForm.%s", err)
			}
		}).
		AddButton("Bank Card Data Page", func() {
			cardDataForm := a.cardDataForm(app)
			if err := app.SetRoot(cardDataForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.saveDataPagesForm.%s", err)
			}

		}).
		AddButton("Return back", func() {
			dataPagesForm := a.dataPagesForm(app)
			if err := app.SetRoot(dataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.saveDataPagesForm.%s", err)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	saveDataPagesForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return saveDataPagesForm
}

func (a *Application) textDataForm(app *tview.Application) *tview.Form {
	textDataForm := tview.NewForm()
	textDataForm.
		AddTextArea("Text", "", 20, 10, 10000, nil).
		AddInputField("Metainfo", "", 20, nil, nil).
		AddButton("Submit", func() {
			metainfo := textDataForm.GetFormItemByLabel("Metainfo").(*tview.InputField).GetText()
			text := textDataForm.GetFormItemByLabel("Text").(*tview.TextArea).GetText()
			if text == "" {
				if textDataForm.GetButtonIndex("Saved") != -1 {
					idx := textDataForm.GetButtonIndex("Saved")
					textDataForm.RemoveButton(idx)
				}
				if textDataForm.GetButtonIndex("Incorrect text") == -1 {
					textDataForm.AddButton("Incorrect text", nil)
				}
				return
			}
			textdata := models.TextData{
				TextData: []byte(text),
				Metainfo: []byte(metainfo),
				UserID:   int64(a.userID),
			}
			err := a.grpcClient.AddTextData(context.Background(), textdata, a.userSecretPhrase)
			if err != nil {
				return
			}
			if textDataForm.GetButtonIndex("Saved") != -1 {
				idx := textDataForm.GetButtonIndex("Saved")
				textDataForm.RemoveButton(idx)
			}
			if textDataForm.GetButtonIndex("Saved") == -1 {
				textDataForm.AddButton("Saved", nil)
			}
		}).
		AddButton("Return back", func() {
			saveDataPagesForm := a.saveDataPagesForm(app)
			if err := app.SetRoot(saveDataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.textDataForm.%s", err)
			}
		})
	textDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return textDataForm
}

func (a *Application) binaryDataForm(app *tview.Application) *tview.Form {
	binaryDataForm := tview.NewForm()
	binaryDataForm.
		AddInputField("Binary path", "", 20, nil, nil).
		AddInputField("Metainfo", "", 20, nil, nil).
		AddButton("Submit", func() {
			metainfo := binaryDataForm.GetFormItemByLabel("Metainfo").(*tview.InputField).GetText()
			binaryFile := binaryDataForm.GetFormItemByLabel("Binary path").(*tview.InputField).GetText()
			if binaryFile == "" {
				if binaryDataForm.GetButtonIndex("Saved") != -1 {
					idx := binaryDataForm.GetButtonIndex("Saved")
					binaryDataForm.RemoveButton(idx)
				}
				if binaryDataForm.GetButtonIndex("Incorrect path") == -1 {
					binaryDataForm.AddButton("Incorrect path", nil)
				}
				return
			}
			file, err := file.ReadFile(binaryFile)
			if err != nil {
				binaryDataForm.GetButton(binaryDataForm.GetButtonIndex("Saved")).SetLabel("Error")
				time.Sleep(2 * time.Second)
				binaryDataForm.GetButton(binaryDataForm.GetButtonIndex("Error")).SetLabel("Saved")
				return
			}
			binarydata := models.BinaryData{
				BinaryData: file,
				Metainfo:   []byte(metainfo),
				UserID:     int64(a.userID),
			}
			err = a.grpcClient.AddBinaryData(context.Background(), binarydata, a.userSecretPhrase)
			if err != nil {
				return
			}
			if binaryDataForm.GetButtonIndex("Saved") != -1 {
				idx := binaryDataForm.GetButtonIndex("Saved")
				binaryDataForm.RemoveButton(idx)
			}
			if binaryDataForm.GetButtonIndex("Saved") == -1 {
				binaryDataForm.AddButton("Saved", nil)
			}
		}).
		AddButton("Return back", func() {
			saveDataPagesForm := a.saveDataPagesForm(app)
			if err := app.SetRoot(saveDataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.binaryDataForm.%s", err)
			}
		})
	binaryDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return binaryDataForm
}

func (a *Application) cardDataForm(app *tview.Application) *tview.Form {
	cardDataForm := tview.NewForm()
	cardDataForm.
		AddInputField("Card number", "", 20, nil, nil).
		AddInputField("Card name", "", 20, nil, nil).
		AddInputField("Card date", "", 20, nil, nil).
		AddInputField("CVV", "", 20, nil, nil).
		AddInputField("Metainfo", "", 20, nil, nil).
		AddButton("Submit", func() {
			metainfo := cardDataForm.GetFormItemByLabel("Metainfo").(*tview.InputField).GetText()
			cardNumber := cardDataForm.GetFormItemByLabel("Card number").(*tview.InputField).GetText()
			cardName := cardDataForm.GetFormItemByLabel("Card name").(*tview.InputField).GetText()
			cardDate := cardDataForm.GetFormItemByLabel("Card date").(*tview.InputField).GetText()
			cardCVV := cardDataForm.GetFormItemByLabel("CVV").(*tview.InputField).GetText()
			if cardNumber == "" || cardName == "" || cardDate == "" || cardCVV == "" || !luhn.CheckLuhn(cardNumber) {
				if cardDataForm.GetButtonIndex("Saved") != -1 {
					idx := cardDataForm.GetButtonIndex("Saved")
					cardDataForm.RemoveButton(idx)
				}
				if cardDataForm.GetButtonIndex("Incorrect") == -1 {
					cardDataForm.AddButton("Incorrect", nil)
				}
				return
			}
			carddata := models.CardData{
				CardNumberData: []byte(cardNumber),
				CardNameData:   []byte(cardName),
				CardDateData:   []byte(cardDate),
				CvvData:        []byte(cardCVV),
				Metainfo:       []byte(metainfo),
				UserID:         int64(a.userID),
			}
			err := a.grpcClient.AddCardData(context.Background(), carddata, a.userSecretPhrase)
			if err != nil {
				if cardDataForm.GetButtonIndex("Incorrect") == -1 {
					cardDataForm.AddButton("Incorrect", nil)
				}
				return
			}

			if cardDataForm.GetButtonIndex("Saved") != -1 {
				idx := cardDataForm.GetButtonIndex("Saved")
				cardDataForm.RemoveButton(idx)
			}
			if cardDataForm.GetButtonIndex("Saved") == -1 {
				cardDataForm.AddButton("Saved", nil)
			}
		}).
		AddButton("Return back", func() {
			saveDataPagesForm := a.saveDataPagesForm(app)
			if err := app.SetRoot(saveDataPagesForm, true).EnableMouse(true).Run(); err != nil {
				applicationErr = fmt.Errorf("app.cardDataForm.%s", err)
			}
		})
	cardDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return cardDataForm
}

func (a *Application) getTextDataForm(app *tview.Application) *tview.Form {
	getTextDataForm := tview.NewForm()
	getTextDataForm.
		AddTextView("GET TEXT DATA:", "you will get a .txt file with the selected text data", 20, 6, true, true)
	keys, err := a.getService.GetDatas(context.Background(), "textdata")
	if err != nil {
		if getTextDataForm.GetButtonIndex("Error") == -1 {
			getTextDataForm.AddButton("Error", nil)
		}
	}
	getTextDataForm.AddInputField("Unique key", "", 20, nil, nil)
	for k, v := range keys {
		getTextDataForm.AddTextView(fmt.Sprintf("%d", k), v, 20, 1, true, true)
	}
	getTextDataForm.AddButton("Get .txt", func() {
		strkey := getTextDataForm.GetFormItemByLabel("Unique key").(*tview.InputField).GetText()
		intkey, _ := strconv.Atoi(strkey)
		_, err := a.getService.GetTextData(context.Background(), intkey, a.userSecretPhrase)
		if err != nil {
			if errors.Is(err, errors.New("invalid hash key")) {
				if getTextDataForm.GetButtonIndex("Error Secret Phrase") == -1 {
					getTextDataForm.AddButton("Error Secret Phrase", nil)
				}
			}
			if getTextDataForm.GetButtonIndex("Error") == -1 {
				getTextDataForm.AddButton("Error", nil)
			}
		}
	})
	getTextDataForm.AddButton("Return back", func() {
		getDataPagesForm := a.getDataPagesForm(app)
		if err := app.SetRoot(getDataPagesForm, true).EnableMouse(true).Run(); err != nil {
			applicationErr = fmt.Errorf("app.getTextDataForm.%s", err)
		}
	})
	getTextDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return getTextDataForm
}

func (a *Application) getBinaryDataForm(app *tview.Application) *tview.Form {
	getBinaryDataForm := tview.NewForm()
	getBinaryDataForm.
		AddTextView("GET BINARY DATA:", "you will get a .txt file with the selected card data", 20, 6, true, true)
	keys, err := a.getService.GetDatas(context.Background(), "binarydata")
	if err != nil {
		if getBinaryDataForm.GetButtonIndex("Error") == -1 {
			getBinaryDataForm.AddButton("Error", nil)
		}
	}
	getBinaryDataForm.AddInputField("Unique key", "", 20, nil, nil)
	for k, v := range keys {
		getBinaryDataForm.AddTextView(fmt.Sprintf("%d", k), v, 20, 1, true, true)
	}
	getBinaryDataForm.AddButton("Get .txt", func() {
		strkey := getBinaryDataForm.GetFormItemByLabel("Unique key").(*tview.InputField).GetText()
		intkey, _ := strconv.Atoi(strkey)
		_, err := a.getService.GetBinaryData(context.Background(), intkey, a.userSecretPhrase)
		a.log.Info(err)
		if err != nil {
			if errors.Is(err, errors.New("invalid hash key")) {
				if getBinaryDataForm.GetButtonIndex("Error Secret Phrase") == -1 {
					getBinaryDataForm.AddButton("Error Secret Phrase", nil)
				}
			}
			if getBinaryDataForm.GetButtonIndex("Error") == -1 {
				getBinaryDataForm.AddButton("Error", nil)
			}
		}
	})
	getBinaryDataForm.AddButton("Return back", func() {
		getDataPagesForm := a.getDataPagesForm(app)
		if err := app.SetRoot(getDataPagesForm, true).EnableMouse(true).Run(); err != nil {
			applicationErr = fmt.Errorf("app.getCardDataForm.%s", err)
		}
	})
	getBinaryDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return getBinaryDataForm
}

func (a *Application) getCardDataForm(app *tview.Application) *tview.Form {
	getCardDataForm := tview.NewForm()
	getCardDataForm.
		AddTextView("GET CARD DATA:", "you will get a .txt file with the selected card data", 20, 6, true, true)
	keys, err := a.getService.GetDatas(context.Background(), "carddata")
	if err != nil {
		if getCardDataForm.GetButtonIndex("Error") == -1 {
			getCardDataForm.AddButton("Error", nil)
		}
	}
	getCardDataForm.AddInputField("Unique key", "", 20, nil, nil)
	for k, v := range keys {
		getCardDataForm.AddTextView(fmt.Sprintf("%d", k), v, 20, 1, true, true)
	}
	getCardDataForm.AddButton("Get .txt", func() {
		strkey := getCardDataForm.GetFormItemByLabel("Unique key").(*tview.InputField).GetText()
		intkey, _ := strconv.Atoi(strkey)
		_, err := a.getService.GetCardData(context.Background(), intkey, a.userSecretPhrase)
		a.log.Info(err)
		if err != nil {
			if errors.Is(err, errors.New("invalid hash key")) {
				if getCardDataForm.GetButtonIndex("Error Secret Phrase") == -1 {
					getCardDataForm.AddButton("Error Secret Phrase", nil)
				}
			}
			if getCardDataForm.GetButtonIndex("Error") == -1 {
				getCardDataForm.AddButton("Error", nil)
			}
		}
	})
	getCardDataForm.AddButton("Return back", func() {
		getDataPagesForm := a.getDataPagesForm(app)
		if err := app.SetRoot(getDataPagesForm, true).EnableMouse(true).Run(); err != nil {
			applicationErr = fmt.Errorf("app.getCardDataForm.%s", err)
		}
	})
	getCardDataForm.SetBorder(true).SetTitle("Gophkeeper").SetTitleAlign(tview.AlignCenter)
	return getCardDataForm
}
