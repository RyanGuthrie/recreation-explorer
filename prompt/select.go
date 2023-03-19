package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"sort"
	"strings"
)

var ExitError = fmt.Errorf("Exit")

func Select[T any](promptLabel string, items []T, cursorPos *int, selectTemplates promptui.SelectTemplates, searchValue func(item T) string, less func(l, r int) bool, exitT T) (*T, error) {
	sort.Slice(items, less)
	items = append([]T{exitT}, items...)

	listHeight := 15
	prompt := promptui.Select{
		Label:             promptLabel,
		Items:             items,
		Size:              listHeight,
		CursorPos:         *cursorPos,
		Templates:         &selectTemplates,
		StartInSearchMode: *cursorPos == 0,
		Searcher: func(input string, index int) bool {
			return strings.Contains(strings.ToLower(searchValue(items[index])), strings.ToLower(input))
		},
	}

	pos, _, err := prompt.RunCursorAt(*cursorPos, *cursorPos-(listHeight/3))
	if pos == 0 {
		return nil, ExitError
	} else if err != nil {
		return nil, err
	}
	*cursorPos = pos

	return &items[pos], err
}
