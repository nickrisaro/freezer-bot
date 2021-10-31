package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func Configurar(urlPublica string, urlPrivada string, token string) (*tb.Bot, error) {
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

	return b, nil
}
