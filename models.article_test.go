package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()
	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}
	// Check that each member is identical
	for i, v := range alist {
		if v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title ||
			v.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}
}

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(getAllArticles())
	a, err := createNewArticle("New test title", "New test content")
	allArticles := getAllArticles()
	newLength := len(allArticles)
	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {
			t.Fail()
		}
}