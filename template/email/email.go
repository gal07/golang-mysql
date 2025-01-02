package email

func ActivateUsers(link string) string {

	t := "<p>here to activation user <a href='http://localhost:8080/api/v1/auth/gotoactive?activation=" + link + "'>Activate</a></p>"
	return t

}
