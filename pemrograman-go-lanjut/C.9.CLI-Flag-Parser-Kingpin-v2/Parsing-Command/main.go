package main

import (
    "os"
    "fmt"
    "gopkg.in/alecthomas/kingpin.v2"
)

var app 					= kingpin.New("App", "Simple app")

var commandSomething 		= app.Command("something", "DO SOMETHING")
var commandSomethingArgX 	= commandSomething.Flag("x", "arg x").String()
var commandSomethingFlagY 	= commandSomething.Flag("y", "flag y").String()

// add
var commandAdd 				= app.Command("add", "MENAMBAH USER BARU")
var commandAddFlagOverride 	= commandAdd.Flag("override", "override existing user").Short('o').Bool()
var commandAddArgUser 		= commandAdd.Arg("user", "username").Required().String()

// update
var commandUpdate 			= app.Command("update", "UPDATE USER BARU")
var commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
var commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()

// delete
var commandDelete          	= app.Command("delete", "DELETE USER BARU")
var commandDeleteFlagForce 	= commandDelete.Flag("force", "force deletion").Short('f').Bool()
var commandDeleteArgUser   	= commandDelete.Arg("user", "username").Required().String()


func main() {
	// C.9.4. Parsing Command
		// commandAdd.Action(func(ctx *kingpin.ParseContext) error {
		//     user := *commandAddArgUser
		// 	override := *commandAddFlagOverride
		// 	fmt.Printf("ADDING USER %s, OVERRIDE %t \n", user, override)
		// 	return nil
		// })

		// commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
		//     oldUser := *commandUpdateArgOldUser
		// 	newUser := *commandUpdateArgNewUser
		// 	fmt.Printf("UPDATING USER FROM %s TO %s \n", oldUser, newUser)
		// 	return nil
		// })

		// commandDelete.Action(func(ctx *kingpin.ParseContext) error {
		//     user := *commandDeleteArgUser
		// 	force := *commandDeleteFlagForce
		// 	fmt.Printf("DELETING USER %s, FORCE %t \n", user, force)
		// 	return nil
		// })

		// kingpin.MustParse(app.Parse(os.Args[1:]))


	// C.9.5. Command Action Tanpa Menggunakan .Action()
		commandInString := kingpin.MustParse(app.Parse(os.Args[1:]))
		switch commandInString {
			case commandAdd.FullCommand(): // add user
				user := *commandAddArgUser
				override := *commandAddFlagOverride
				fmt.Printf("ADDING USER %s, OVERRIDE %t \n", user, override)

			case commandUpdate.FullCommand(): // update user
				oldUser := *commandUpdateArgOldUser
				newUser := *commandUpdateArgNewUser
				fmt.Printf("UPDATING USER FROM %s TO %s \n", oldUser, newUser)

			case commandDelete.FullCommand(): // delete user
				user := *commandDeleteArgUser
				force := *commandDeleteFlagForce
				fmt.Printf("DELETING USER %s, FORCE %t \n", user, force)
		}
}
