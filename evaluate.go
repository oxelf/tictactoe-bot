package main

type EvalBoardResult struct {
	Won int `json:"won"` // 0 = nobody 1 = player 1, 2 = player 2
}

func (board *Board) EvaluateBoard() EvalBoardResult {
	return EvalBoardResult{Won: 0}
}
