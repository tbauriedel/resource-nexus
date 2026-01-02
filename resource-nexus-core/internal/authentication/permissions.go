package authentication

// GetPermissionForPath returns the permission for the given path / url.
func GetPermissionForPath(url string) (string, bool) {
	var permissionsMap = map[string]string{
		"": "",
	}

	perm, ok := permissionsMap[url]
	if !ok {
		return "", false
	}

	return perm, true
}

// BuildPermissionString builds a permission string from category, action and resource.
//
// Format: category:action:resource.
func BuildPermissionString(category, action, resource string) string {
	return category + ":" + action + ":" + resource
}
