package mail

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

type Message struct {
	From    string
	Date    time.Time
	Subject string
}

type Poller struct {
	hosts map[Type]string
}

func NewPoller(hosts map[Type]string) Poller {
	return Poller{
		hosts: hosts,
	}
}

type PollResult struct {
	Messages      []Message
	Failed        int
	LastMessageID int64
}

func (p Poller) Poll(mail, pass string, since time.Time, lastSeq int64) (PollResult, error) {
	host, ok := p.hosts[ParseType(mail)]
	if !ok {
		return PollResult{}, fmt.Errorf("mail type is not supported")
	}

	c, err := p.connect(host, mail, pass)
	if err != nil {
		return PollResult{}, err
	}

	mbox, err := c.Select("INBOX", true)
	if err != nil {
		return PollResult{}, fmt.Errorf("select mailbox %q: %w", mail, err)
	}

	if mbox.Messages == 0 {
		return PollResult{}, nil
	}

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	criteria.Since = since
	seqNums, err := c.Search(criteria)
	if err != nil {
		return PollResult{}, fmt.Errorf("search unread messages %q: %w", mail, err)
	}

	seqNums = p.filterSequence(seqNums, uint32(lastSeq))
	if len(seqNums) == 0 {
		return PollResult{}, nil
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(seqNums...)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}
	messagesCh := make(chan *imap.Message)

	var fetchErr error
	go func() {
		if fetchErr = c.Fetch(seqSet, items, messagesCh); fetchErr != nil {
			log.Printf("fetch %q unread messages failed: %s\n", mail, err)
		}
	}()

	messages, failed, err := p.collectMessages(section, messagesCh)
	if err != nil {
		return PollResult{}, fmt.Errorf("collect %q messages: %w", mail, err)
	}

	if fetchErr != nil {
		return PollResult{}, fmt.Errorf("fetch %q messages: %w", mail, fetchErr)
	}

	return PollResult{
		Messages:      messages,
		Failed:        failed,
		LastMessageID: int64(p.maxSequence(seqNums)),
	}, nil
}

func (p Poller) connect(host, username, pass string) (*client.Client, error) {
	c, err := client.DialTLS(host, nil)
	if err != nil {
		return nil, fmt.Errorf("tls connect to %q: %w", host, err)
	}

	err = c.Login(username, pass)
	if err != nil {
		return nil, fmt.Errorf("login %q: %w", host, err)
	}

	return c, nil
}

func (p Poller) collectMessages(section imap.BodySectionName, ch chan *imap.Message) ([]Message, int, error) {
	messages := make([]Message, 0)
	failed := 0

	for msg := range ch {
		r := msg.GetBody(&section)
		if r == nil {
			log.Println("server didn't returned message body")
			failed++
			continue
		}

		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Printf("create reader: %s\n", err)
			failed++
			continue
		}

		// Print some info about the message
		header := mr.Header
		date, err := header.Date()
		if err != nil {
			log.Printf("get message date: %s\n", err)
			failed++
			continue
		}

		from, err := header.AddressList("From")
		if err != nil {
			log.Printf("get message from: %s\n", err)
			failed++
			continue
		}

		subject, err := header.Subject()
		if err != nil {
			log.Printf("get message subject: %s\n", err)
			failed++
			continue
		}

		messages = append(messages, Message{
			Date:    date,
			From:    p.formatAddressList(from),
			Subject: subject,
		})
	}

	return messages, failed, nil
}

func (p Poller) formatAddressList(l []*mail.Address) string {
	formatted := make([]string, len(l))
	for i, a := range l {
		formatted[i] = a.Address
	}
	return strings.Join(formatted, ", ")
}

func (p Poller) filterSequence(seq []uint32, lastSeq uint32) []uint32 {
	if len(seq) == 0 {
		return nil
	}

	filtered := make([]uint32, 0)
	for _, s := range seq {
		if s > lastSeq {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func (p Poller) maxSequence(seq []uint32) uint32 {
	if len(seq) == 0 {
		return 0
	}

	max := seq[0]
	for _, s := range seq {
		if max < s {
			max = s
		}
	}
	return max
}
