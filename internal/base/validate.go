package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

/*
Validate - должна проверить:
входной параметр req
поля req.UserID, req.Title, req.Description
Если какое-либо поле пустое, добавьте сообщение об ошибке в слайс res.
*/
func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, "ValidateRequest is nil!")
		return res
	}

	if req.UserID == "" {
		res = append(res, "ValidateRequest userID is empty!")
	}

	if req.Title == "" {
		res = append(res, "ValidateRequest title is empty!")
	}

	if req.Description == "" {
		res = append(res, "ValidateRequest description is empty!")
	}

	return res
}
