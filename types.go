package main

type Application struct {
	Set	string
	PWDB	[]Password
}

type Password struct {
	Key		string
	Len		int
	StepLeft	int
	StepRight	int
}
