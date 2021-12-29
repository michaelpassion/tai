package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
    switch lang {
    case "en":
        return fmt.Sprintf("Hi, %s!", name)
    case "pt":
        return fmt.Sprintf("Oi, %s!", name)
    case "es":
        return fmt.Sprintf("Hola, %s!", name)
    case "fr":
        return fmt.Sprintf("Bonjoury, %s!", name)
    default:
        return "", error.New("unknown language")
    }
}
