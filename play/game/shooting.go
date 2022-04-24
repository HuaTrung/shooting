package game

import "time"

const PLAYERA="playerA"
const PLAYERB="playerB"

type ShootingReceivingEvent struct{
	Receiver string
	Sender string
	EnemyShot string
	EnemyBoard [10][10]int
	PreYourBoard [10][10]int
	NxtYourBoard [10][10]int
	CreatedAt 	time.Time
}

type ShootingSendingEvent struct{
	Receiver string
	Sender string
	YourShot string
	EnemyBoard [10][10]int
	YourBoard [10][10]int
	CreatedAt 	time.Time
}