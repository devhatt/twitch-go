package helpers

var users = []string{"k1nnha", "alecell", "hxsggsz", "zoldyzdk", "PiluVitu"}

func ValidUsers(name string) bool {
	for _, user := range users {
		if user == name {
			return true
		}
	}
	return false
}
