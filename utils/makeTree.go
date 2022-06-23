package utils

import "strings"

type Branch interface {
	Add(name string) Branch
	AddBranch(branch Branch)
	Text() string
	Title() string
	Branches() []Branch
}

type Tree struct {
	branches []Branch
	title    string
}

func MakeTree(title string) Branch {
	return &Tree{
		title:    title,
		branches: []Branch{},
	}
}

func (tree *Tree) Title() string {
	return tree.title
}

func (tree *Tree) Add(name string) Branch {
	created := MakeTree(name)
	tree.branches = append(tree.branches, created)
	return created
}

func (tree *Tree) AddBranch(branch Branch) {
	tree.branches = append(tree.branches, branch)
}

func (tree *Tree) Branches() []Branch {
	return tree.branches
}

func (tree *Tree) Text() string {
	return tree.Title() + "\n" + createTreeBranches(tree.Branches(), []bool{})
}

func createText(text string, spaces []bool, end bool) string {
	const (
		emptySpace     = "    "
		middleBranch   = "├── "
		continueBranch = "│   "
		endBranch      = "└── "
	)

	var (
		treeText  string
		indicator string
		output    string
	)

	lines := strings.Split(text, "\n")

	if end {
		indicator = endBranch
	} else {
		indicator = middleBranch
	}

	for _, space := range spaces {
		if space {
			treeText += emptySpace
		} else {
			treeText += continueBranch
		}
	}

	for i := range lines {
		text := lines[i]

		if i == 0 {
			output += treeText + indicator + text + "\n"
			continue
		}

		if end {
			indicator = emptySpace
		} else {
			indicator = continueBranch
		}

		output += treeText + indicator + text + "\n"
	}

	return output
}

func createTreeBranches(branches []Branch, spaces []bool) string {
	var output string

	for i, branch := range branches {
		end := i == len(branches) - 1
		output += createText(branch.Title(), spaces, end)

		if len(branch.Branches()) > 0 {
			_spaces := append(spaces, end)
			output += createTreeBranches(branch.Branches(), _spaces)
		}
	}

	return output
}