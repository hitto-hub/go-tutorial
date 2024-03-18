package greetings

import (
    "errors"
    "fmt"
    "math/rand"
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
    message := fmt.Sprintf(randomFormat(), name)
    //     ↓ name    ↓ error
    return message, nil
}
// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
