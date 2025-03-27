package main
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
    "strconv"
    "math/rand"
)

const (
	MaxGuess = 300
	MinGuess = 0
	MaxGuessAmount = 10
)

func main () {
	myApp := app.New() //создаем приложение
	myWindow := myApp.NewWindow("Searching!") //создаем окно
	
	var computerGuess, guessesLeft = 0, -1 //объявляем переменные для загаданного числа и кол-ва попыток
	computerGuessLabel := widget.NewLabel("Нажмите на кнопку ниже, чтобы компьютер загадал число от 0 до 300") //лейбл - кол-во попыток
    //кнопка старта
	startButton := widget.NewButton("Загадать число!", func() {
		guessesLeft = MaxGuessAmount
		computerGuess = rand.Intn(301)
		computerGuessLabel.SetText("Компьютер загадал число! Осталось попыток: " + strconv.Itoa(guessesLeft))
	})
	guessDisplay := widget.NewLabel("Введите первое число!") //лейбл-подсказка
	userInput := widget.NewEntry() //поле для ввода
	userInput.SetPlaceHolder("Введите число: ")
	tryButton := widget.NewButton("Попробовать!", func() { //кнопка попытки
		if guessesLeft == -1 {
			guessDisplay.SetText("Для начала начните игру!")
			computerGuessLabel.SetText("Нажмите на кнопку ниже, чтобы компьютер загадал число от 0 до 300")
			return
		}
		if guessesLeft == 1 {
			guessDisplay.SetText("К сожалению, вы не угадали число " + strconv.Itoa(computerGuess) + ". Попробуйте сыграть еще раз!")
			computerGuessLabel.SetText("Нажмите на кнопку ниже, чтобы компьютер загадал число от 0 до 300")
			return
		}
		userGuess, _ := strconv.Atoi(userInput.Text) //получаем число из инпута
		if userGuess < MinGuess || userGuess > MaxGuess {
			guessDisplay.SetText("Неправильное число! Введите число между 0 и 300")
			return
		}
		if userGuess == computerGuess {
			guessDisplay.SetText("Поздравляем! Вы отгадали число " + strconv.Itoa(computerGuess) + " за " + strconv.Itoa(MaxGuessAmount-guessesLeft+1) + " попыток!")
			computerGuessLabel.SetText("Нажмите на кнопку ниже, чтобы компьютер загадал число от 0 до 300")
            return
		}
		guessesLeft-- //уменьшаем попытки
		computerGuessLabel.SetText("Попыток осталось: " + strconv.Itoa(guessesLeft))
		if userGuess < computerGuess {
			guessDisplay.SetText("Ваше число меньше загаданного!")
		} else {
			guessDisplay.SetText("Ваше число больше загаданного!")
		}
	})
	myWindow.SetContent( //все созданные виджеты засовываем в окно в нужном нам порядке
		container.NewVBox(
			computerGuessLabel,
			startButton,
			guessDisplay,
			userInput,
			tryButton,
		),
	)
	myWindow.Resize(fyne.NewSize(300,200)) //изменение размера окна
	myWindow.ShowAndRun() //запуск программы
}