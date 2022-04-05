package pub_sub

import (
	"fmt"
)

type PersonalSubscriber struct {
	ch   chan string
	name string
}

func NewPersonalSubscriber(name string) *PersonalSubscriber {
	return &PersonalSubscriber{
		ch:   make(chan string),
		name: name,
	}
}

func (p *PersonalSubscriber) Receiver(message string) {
	p.ch <- message
}

func (p *PersonalSubscriber) ListenOnChannel() {
	for msg := range p.ch {
		fmt.Printf("%s님 %s이(가) 발생했어요. 가까운 피난소로 대피하거나, 119에 전화하세요!\n", p.name, msg)
	}
}

func (p *PersonalSubscriber) GetName() string {
	return p.name
}
