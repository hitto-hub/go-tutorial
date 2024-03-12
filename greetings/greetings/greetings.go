package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
//         ↓変数名 ↓型       ↓ returnの型指定
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        //     ↓ ここにname ↓ ここにerror
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
// var message string
// message = fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    //     ↓ name    ↓ error
    return message, nil
}
