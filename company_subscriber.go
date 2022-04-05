package pub_sub

import (
	"fmt"
)

type CompanySubscriber struct {
	ch   chan string
	name string
}

func NewCompanySubscriber(name string) *CompanySubscriber {
	return &CompanySubscriber{
		ch:   make(chan string),
		name: name,
	}
}

func (c *CompanySubscriber) Receiver(message string) {
	c.ch <- message
}

func (c *CompanySubscriber) ListenOnChannel() {
	for msg := range c.ch {
		fmt.Printf("%s님 %s이(가) 발생했어요. 회사 재난팀의 지시에 따라 주세요!\n", c.name, msg)
	}
}

func (c *CompanySubscriber) GetName() string {
	return c.name
}
