package dto

type LoginResponseDto struct {
    Userid int    `json:"userid"`
    Token  string `json:"token"`
}