package telegram

import (
	"errors"

	"github.com/nickrisaro/freezer-bot/encargade"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Configurar(urlPublica string, urlPrivada string, token string, encargade *encargade.Encargade) (*tb.Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.Webhook{Listen: urlPrivada, Endpoint: &tb.WebhookEndpoint{PublicURL: urlPublica, Cert: ""}},
	})

	if err != nil {
		return nil, err
	}

	b.Handle("/ping", func(m *tb.Message) {
		b.Send(m.Chat, "Pong!")
	})

	b.Handle("/help", func(m *tb.Message) {
		ayuda := "Hola soy Andrew, el encargado de tu freezer, te puedo decir que hay en él, poner cosas nuevas y sacar las que ya están ahí\n"
		ayuda += "Si querés agregar algo tenés que respetar el formato nombre,cantidad,unidad de medida. Por ejemplo: /agregar ñoquis,500,gramo\n"
		ayuda += "Si querés quitar algo tenés que respetar el formato nombre,cantidad. Por ejemplo: /quitar ñoquis,250\n"
		ayuda += "Las unidades de medida pueden ser unidad, kilo, gramo, litro, mililitro u otra\n"
		ayuda += "Empezá creando tu freezer con el comando /crear"
		b.Send(m.Chat, ayuda)
	})

	b.Handle("/start", func(m *tb.Message) {
		ayuda := "Hola soy Andrew, el encargado de tu freezer, te puedo decir que hay en él, poner cosas nuevas y sacar las que ya están ahí\n"
		ayuda += "Me podés hablar en privado o agregar a un grupo para que varias personas sepan que hay en el freezer!\n"
		ayuda += "Si tenés alguna duda me podés pedir ayuda con el comando /help\n"
		ayuda += "Empezá creando tu freezer con el comando /crear"
		b.Send(m.Chat, ayuda)
	})

	b.Handle("/crear", func(m *tb.Message) {
		nombreDelFreezer := m.Chat.Title
		if len(nombreDelFreezer) == 0 {
			nombreDelFreezer = m.Chat.FirstName + " " + m.Chat.LastName
		}
		err := encargade.NuevoFreezer(m.Chat.ID, nombreDelFreezer)
		if err != nil {
			b.Send(m.Chat, "Ups, no pude crear tu freezer, probá más tarde")
		} else {
			b.Send(m.Chat, "Listo, ya creé tu freezer, agregale comida con /agregar")
		}
	})

	b.Handle("/agregar", func(m *tb.Message) {
		productoAAgregar := m.Payload

		err := encargade.MeterEnFreezer(m.Chat.ID, productoAAgregar)
		if err != nil {
			b.Send(m.Chat, "Ups, no pude agregar el producto a tu freezer, revisá el formato o probá más tarde\nUso correcto /agregar milanesas,1,unidad")
		} else {
			b.Send(m.Chat, "Listo, ya agregué comida a tu freezer, fijate lo que hay con /listar")
		}
	})

	b.Handle("/quitar", func(m *tb.Message) {
		productoAQuitar := m.Payload

		err := encargade.SacarDelFreezer(m.Chat.ID, productoAQuitar)
		if err != nil {
			if errors.Is(err, errors.New("no existe ese producto")) {
				b.Send(m.Chat, "Ups, no encontré ese producto en tu freezer, revisá el nombre")
			} else {
				b.Send(m.Chat, "Ups, no pude sacar el producto de tu freezer, revisá el formato o probá más tarde\nUso correcto /quitar milanesas,1")
			}
		} else {
			b.Send(m.Chat, "Listo, ya saqué comida de tu freezer, fijate lo que hay con /listar")
		}
	})

	b.Handle("/listar", func(m *tb.Message) {
		b.Send(m.Chat, encargade.QueCosasHayEnEsteFreezer(m.Chat.ID))
	})

	return b, nil
}
