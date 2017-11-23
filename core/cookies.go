package core

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type CookieJarType int8

const (
	JarMemory CookieJarType = iota
	JarJson
	JarGob
)

// JarOption used to configure how cookies data saved
//
type JarOption struct {
	// JarType specify the which way used to save the cookies
	//  JarMemory : just in memory without persist
	//	JarGob: persist by module encoding/gob
	JarType CookieJarType

	// Filename holds the file to use for storage of the cookies.
	// If it is empty, JarMemory will be used.
	Filename string
}

// SimpleJar implement http.CookieJar to handle cookies
//
type SimpleJar struct {
	filename string
	jarType  CookieJarType
	cookies  []*http.Cookie
}

// NewSimpleJar return SimpleJar object with sepecified option
// example:
//    jar := NewSimpleJar(JarOption{
//		JarType:  JarMemory,
//		Filename: "d:\\cookies.jar",
//	  })
//
func NewSimpleJar(option JarOption) *SimpleJar {
	if option.Filename == "" {
		option.JarType = JarMemory
	}

	return &SimpleJar{
		filename: option.Filename,
		jarType:  option.JarType,
		cookies:  make([]*http.Cookie, 0, 10),
	}
}

// SetCookies handles the receipt of the cookies in a reply for the
// given URL. It may or may not choose to save the cookies, depending
// on the jar's policy and implementation.
//
func (jar *SimpleJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if cookies == nil || len(cookies) == 0 {
		return
	}

	for _, newone := range cookies {
		var cookie *http.Cookie
		for _, old := range jar.cookies {
			if newone.Name == old.Name {
				cookie = old
				break
			}
		}

		if cookie == nil {
			cookie = new(http.Cookie)
			jar.cookies = append(jar.cookies, cookie)
		}

		*cookie = *newone
	}
}

// Cookies returns the cookies to send in a request for the given URL.
// It is up to the implementation to honor the standard cookie use
// restrictions such as in RFC 6265.
//
func (jar *SimpleJar) Cookies(u *url.URL) []*http.Cookie {
	s := len(jar.cookies)
	return jar.cookies[0:s]
}

// Load used to deserialization cookies data from file
//
func (jar *SimpleJar) Load() error {
	switch jar.jarType {
	case JarGob:
		fd, err := os.Open(jar.filename)
		if err == nil {
			err = gob.NewDecoder(fd).Decode(&jar.cookies)
		} else if os.IsNotExist(err) {
			err = nil
		}
		return err

	case JarJson:
		fd, err := os.Open(jar.filename)
		if err == nil {
			err = json.NewDecoder(fd).Decode(&jar.cookies)
		} else if os.IsNotExist(err) {
			err = nil
		}
		return err

	case JarMemory:
		return nil

	default:
		return fmt.Errorf("jar type %d not implement yet", jar.jarType)
	}
}

// Persist used to serialization cookies data into file
//
func (jar *SimpleJar) Persist() error {
	if len(jar.cookies) == 0 {
		return nil
	}

	switch jar.jarType {
	case JarGob:
		fd, err := os.Create(jar.filename)
		if err == nil {
			err = gob.NewEncoder(fd).Encode(jar.cookies)
		}
		return err

	case JarJson:
		fd, err := os.Create(jar.filename)
		if err == nil {
			enc := json.NewEncoder(fd)
			enc.SetIndent("", "    ")
			err = enc.Encode(jar.cookies)
		}
		return err

	case JarMemory:
		return nil

	default:
		return fmt.Errorf("jar type %d not implement yet", jar.jarType)
	}
}

// Clean cookies if not valid anymore
//
func (jar *SimpleJar) Clean() {
	jar.cookies = jar.cookies[0:0]
}

// Get cookie vlue by name
//
func (jar *SimpleJar) Get(name string) string {
	for _, v := range jar.cookies {
		if v.Name == name {
			return v.Value
		}
	}
	return ""
}
