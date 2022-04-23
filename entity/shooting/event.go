package shooting

import "time"
const (
	PLAYER_A_TAKEN_ACTION string = "PLAYER_A_TAKEN_ACTION"
	PLAYER_B_TAKEN_ACTION = "PLAYER_B_TAKEN_ACTION"
	PLAYER_A_TAKEN_ACTION_COZ_TIMEOUT = "PLAYER_A_TAKEN_ACTION_COZ_TIMEOUT"
	PLAYER_B_TAKEN_ACTION_COZ_TIMEOUT = "PLAYER_B_TAKEN_ACTION_COZ_TIMEOUT"
	MESSAGE_FROM_A = "MESSAGE_FROM_A"
	MESSAGE_FROM_B = "MESSAGE_FROM_B"
)
type Content struct {
	Action string	`json:"action"`
	Message string	`json:"message"`
}
type Event struct {
	ID          string 		`json:"id"`
	Type      string 		`json:"status,omitempty"`
	Content Content `json:"content,omitempty"`
	PlayerABoard 	[10][10] int	`json:"player_a_board,omitempty"`
	PlayerBBoard	[10][10] int 	`json:"player_b_board,omitempty"`
	CreatedAt 	time.Time   `json:"is_requested"`
}