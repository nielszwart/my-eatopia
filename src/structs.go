package main

type TemplateData struct {
	User    User
	Recipes []*Recipe
}

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
	IsAdmin  int
	Deleted  int
	Created  string
}

type Recipe struct {
	RecipeID       int
	RecipeParentID int
	UserID         int
	Title          string
	Description    string
	Instruction    string
	Ingredients    string
	Persons        int
	Image          string
	CookingTime    string
	Alias          string
	IsInspiration  int
	Deleted        int
	Edited         string
	Created        string
}
