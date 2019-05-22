package gui

import (
	"fmt"
	"github.com/ivahaev/russian-time"
	"github.com/jroimartin/gocui"
	"stash.mail.ru/qafeta/feta-media-tools/wedding_invitation/server"
	"strconv"
	"time"
)

const (
	PLACE_URL = "https://www.google.com/maps/place/%D0%A0%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD+%D0%A6%D0%B5%D0%BD%D1%82%D1%80%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D0%B9/@55.6771685,37.8953379,16z/data=!4m12!1m6!3m5!1s0x0:0xe6203f90b40c821b!2z0KDQtdGB0YLQvtGA0LDQvSDQptC10L3RgtGA0LDQu9GM0L3Ri9C5!8m2!3d55.677168!4d37.8953371!3m4!1s0x0:0xe6203f90b40c821b!8m2!3d55.677168!4d37.8953371"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func EventPlace(g *gocui.Gui) error {
	x0, _, x1, y1, err := g.ViewPosition("Description")
	check(err)

	if v, err := g.SetView("Event", x0, y1+2, x1, y1+13); err != nil {
		if err != gocui.ErrUnknownView {
			check(err)
		}

		v.Title = "Место"
		v.Wrap = true
		v.FgColor = gocui.ColorYellow

		fmt.Fprintln(v, "\n")
		fmt.Fprintln(v, " \x1b[0;35mОктябрьский пр., 198, Люберцы, Московская обл., Россия\x1b[0m")
		fmt.Fprintln(v, "\n")
		fmt.Fprintln(v, " "+PLACE_URL)
	}
	return nil
}

func Controls(g *gocui.Gui) error {

	x0, _, x1, y1, err := g.ViewPosition("Time")
	check(err)

	if v, err := g.SetView("Controls", x0, y1+2, x1, y1+14); err != nil {
		if err != gocui.ErrUnknownView {
			check(err)
		}

		v.Title = "Навигация"
		v.Wrap = true

		fmt.Fprintln(v, "\n")
		fmt.Fprintln(v, " \x1b[0;33mm\x1b[0m : Открывает ссылку с местом проведения  свадьбы в браузере")
		fmt.Fprintln(v, "\n")
		fmt.Fprintln(v, " \x1b[0;33mi\x1b[0m : Открывает приглашение в браузере")
		fmt.Fprintln(v, "\n")
		fmt.Fprintln(v, " \x1b[0;33mq\x1b[0m : Выход")
	}

	return nil
}

func WeddingTime(g *gocui.Gui) error {

	weddingDate := time.Date(2019, time.August, 23, 17, 0, 0, 0, time.Local)

	_, y0, x1, y1, err := g.ViewPosition("Description")
	check(err)

	if v, err := g.SetView("Time", x1+3, y0, x1+47, y1+5); err != nil {
		if err != gocui.ErrUnknownView {
			check(err)
		}

		fmt.Fprintln(v, "")
		v.Title = "Время"
		v.Wrap = true

		v.FgColor = gocui.ColorCyan
		v.Highlight = true
	}

	go func() {
		timeView, err := g.View("Time")
		check(err)

		g.Update(func(g *gocui.Gui) error {
			timeView.Clear()
			fmt.Fprint(timeView, "\n\n")
			fmt.Fprintln(timeView, " Свадьба начнется: "+fmt.Sprintf("%d %s %d в %02d:%02d", weddingDate.Day(), rtime.Time(weddingDate).Month().String()+"a", weddingDate.Year(), weddingDate.Hour(), weddingDate.Minute()))
			fmt.Fprintln(timeView, "")
			fmt.Fprintln(timeView, " До свадьбы осталось:")
			fmt.Fprintln(timeView, "")
			fmt.Fprintln(timeView, getWeddingEventTime(weddingDate))
			fmt.Fprintln(timeView, "")

			art := `             ,-==-,
             \/\/\/
              \\//
             .-''-.
           .'.-""-.'.---.
          / /    .'\ \--.'.
          | |   / /| |   \ \
          \ \  | | / /    | |
           '.'-:\ \.'    / /
             '---; '-..-'.'
                   '----'`

			fmt.Fprintln(timeView, art)
			return nil
		})
	}()

	return nil
}

func getWeddingEventTime(date time.Time) string {

	totalSeconds := int(date.Sub(time.Now()) / time.Second)

	days := int(totalSeconds / (24 * 60 * 60))
	totalSeconds -= days * (24 * 60 * 60)

	hours := int(totalSeconds / (60 * 60))
	totalSeconds -= hours * (60 * 60)

	minutes := int(totalSeconds / 60)
	totalSeconds -= minutes * 60

	var daysStr string
	var hoursStr string
	var minutesStr string
	var secondsStr string

	if days < 0 || hours < 0 || minutes < 0 || totalSeconds < 0 {
		return "\t\t\x1b[0;32mМероприятие прошло!!!\x1b[0m"
	}

	d := []byte(strconv.Itoa(days))
	lastNum, _ := strconv.Atoi(string(d[len(d)-1]))
	var before int
	if len(d) > 1 {
		before, _ = strconv.Atoi(string(d[len(d)-2]))
	} else {
		before = 0
	}

	if lastNum == 1 && before != 1 {
		daysStr = fmt.Sprintf("%d день", days)
	} else if lastNum >= 2 && lastNum <= 4 && before != 1 {
		daysStr = fmt.Sprintf("%d дня", days)
	} else {
		daysStr = fmt.Sprintf("%d дней", days)
	}

	h := []byte(strconv.Itoa(hours))
	lastNum, _ = strconv.Atoi(string(h[len(h)-1]))
	if len(h) > 1 {
		before, _ = strconv.Atoi(string(h[len(h)-2]))
	} else {
		before = 0
	}

	if lastNum == 1 && before != 1 {
		hoursStr = fmt.Sprintf("%d час", hours)
	} else if lastNum >= 2 && lastNum <= 4 && before != 1 {
		hoursStr = fmt.Sprintf("%d часа", hours)
	} else {
		hoursStr = fmt.Sprintf("%d часов", hours)
	}

	m := []byte(strconv.Itoa(minutes))
	lastNum, _ = strconv.Atoi(string(m[len(m)-1]))
	if len(m) > 1 {
		before, _ = strconv.Atoi(string(m[len(m)-2]))
	} else {
		before = 0
	}

	if lastNum == 1 && before != 1 {
		minutesStr = fmt.Sprintf("%d минута", minutes)
	} else if lastNum >= 2 && lastNum <= 4 && before != 1 {
		minutesStr = fmt.Sprintf("%d минуты", minutes)
	} else {
		minutesStr = fmt.Sprintf("%d минут", minutes)
	}

	s := []byte(strconv.Itoa(totalSeconds))
	lastNum, _ = strconv.Atoi(string(s[len(s)-1]))
	if len(s) > 1 {
		before, _ = strconv.Atoi(string(s[len(s)-2]))
	} else {
		before = 0
	}

	if lastNum == 1 && before != 1 {
		secondsStr = fmt.Sprintf("%d секунда", totalSeconds)
	} else if lastNum >= 2 && lastNum <= 4 && before != 1 {
		secondsStr = fmt.Sprintf("%d секунды", totalSeconds)
	} else {
		secondsStr = fmt.Sprintf("%d секунд", totalSeconds)
	}

	return fmt.Sprintf("\x1b[0;32m   %s %s %s %s\x1b[0m", daysStr, hoursStr, minutesStr, secondsStr)
}

func Description(g *gocui.Gui) error {

	maxX, _ := g.Size()

	if v, err := g.SetView("Description", 1, 2, maxX/2+18, 17); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Общее"
		v.Wrap = true

		v.FgColor = gocui.ColorGreen

		text := ` 
          Свой жизненный путь                            _
          Мы вдвоём начинаем.                     mMm  _[_]_
          На свадьбу родных                      /(")\  (")
          И друзей приглашаем.                  //)^(\\//:\\
          Мы очень вас просим                  /(/&@&\\/|~|/
          К нам в гости прийти,               / /-~'~-\ |||
          Отпраздновать вместе                '/       \|||
          Начало пути.                        '----------'--
          Вдвоём решили мы навечно 
          Свою судьбу соединить. 
          И приглашаем вас сердечно, 
          Чтоб радость нашу разделить!`

		fmt.Fprintln(v, text)
	}
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {

	server.Stop()

	return gocui.ErrQuit
}
