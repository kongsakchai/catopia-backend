package usecase

import (
	"fmt"
	"os"

	"github.com/kongsakchai/catopia-backend/domain"
)

type fileUsecase struct {
}

func NewFileUsecase() domain.FileUsecase {
	return &fileUsecase{}
}

func (u *fileUsecase) FileExit(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func (u *fileUsecase) RemoveFile(path string) {
	path = "uploads/" + path

	if u.FileExit(path) {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("File removed")
	} else {
		fmt.Println("File not found")
	}
}
