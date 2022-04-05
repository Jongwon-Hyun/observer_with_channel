package main

import pubSub "pugSub"

func main() {
	emergency := pubSub.NewEmergency()
	personalSub := pubSub.NewPersonalSubscriber("팅커벨")
	companySub := pubSub.NewCompanySubscriber("네버랜드")

	emergency.Subscribe(personalSub)
	emergency.Subscribe(companySub)
	emergency.Publish("태풍")
	emergency.Publish("화재")

	emergency.Unsubscribe(personalSub)

	emergency.Publish("산사태")
	emergency.Publish("해일")

	/*
		이슈: Publish 가 짝수일 경우만 제대로 동작함
		원인 : 모르겠음...
		조금 더 고 내공이 쌓이면 해결해 볼 생각중
	*/
}
