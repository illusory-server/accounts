package vo

import (
	"encoding/json"
	"github.com/illusory-server/accounts/internal/domain/vo"
	"github.com/illusory-server/accounts/pkg/errors/codex"
	"github.com/illusory-server/accounts/pkg/errors/errx"
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestVoLink(t *testing.T) {
	t.Run("Should correct constructor", func(t *testing.T) {
		l := "https://joska.com"
		link, err := vo.NewLink(l)
		assert.NoError(t, err)
		assert.Equal(t, l, link.Value())
	})

	t.Run("Should correct error with incorrect url", func(t *testing.T) {
		link, err := vo.NewLink("no url xdddd random text")
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Link{}, link)

		link, err = vo.NewLink("empty")
		assert.Error(t, err)
		assert.Equal(t, codex.InvalidArgument, errx.Code(err))
		assert.Equal(t, vo.Link{}, link)

	})

	t.Run("Should correct marshal json", func(t *testing.T) {
		l := "https://joska.com"
		link, err := vo.NewLink(l)
		assert.NoError(t, err)
		assert.Equal(t, l, link.Value())

		res, err := json.Marshal(link)
		assert.NoError(t, err)
		assert.Equal(t, "\""+l+"\"", string(res))
	})

	t.Run("Should correct empty", func(t *testing.T) {
		l := "https://joska.com"
		link, err := vo.NewLink(l)
		assert.NoError(t, err)
		assert.False(t, link.Empty())

		link = vo.Link{}
		assert.True(t, link.Empty())
	})
}

func FuzzVoLink(f *testing.F) {
	// Добавляем начальные seed-значения для фаззинга
	seedURLs := []string{
		"http://example.com",
		"https://example.com",
		"ftp://example.com",
		"invalid-url",
		"",
		"http://",
		"http://localhost",
		"http://192.168.1.1",
		"http://[::1]",
		"http://example.com/path?query=param",
		"http://example.com:8080",
		"http://example.com/with spaces", // URL с пробелами
		"http://example.com/with%20encoded%20spaces", // URL с закодированными пробелами
		"http://example.com/with\nnewline",           // URL с управляющими символами
		"http://example.com/with\ttab",
		"http://example.com/with\"quote",
		"http://example.com/with<lessthan",
		"http://example.com/with>greaterthan",
		"http://example.com/with`backtick",
		"http://example.com/with{brace}",
		"http://example.com/with|pipe",
		"http://example.com/with^caret",
		"http://example.com/with~tilde",
		"http://example.com/with#hash",
		"http://example.com/with%percent",
		"http://example.com/with\\backslash",
		"http://example.com/with/slash",
		"http://example.com/with?question",
		"http://example.com/with&ampersand",
		"http://example.com/with=equals",
		"http://example.com/with+plus",
		"http://example.com/with,comma",
		"http://example.com/with;semcolon",
		"http://example.com/with'apostrophe",
		"http://example.com/with!exclamation",
		"http://example.com/with@at",
		"http://example.com/with$dollar",
		"http://example.com/with*asterisk",
		"http://example.com/with(parentheses)",
		"http://example.com/with[brackets]",
		"http://example.com/with{braces}",
		"http://example.com/with`backticks`",
		"http://example.com/with|pipes|",
		"http://example.com/with^carets^",
		"http://example.com/with~tildes~",
		"http://example.com/with#hashes#",
		"http://example.com/with%percents%",
		"http://example.com/with\\backslashes\\",
		"http://example.com/with/slashes/",
		"http://example.com/with?questions?",
		"http://example.com/with&ampersands&",
		"http://example.com/with=equals=",
		"http://example.com/with+pluses+",
		"http://example.com/with,commas,",
		"http://example.com/with;semcolons;",
		"http://example.com/with'apostrophes'",
		"http://example.com/with!exclamations!",
		"http://example.com/with@ats@",
		"http://example.com/with$dollars$",
		"http://example.com/with*asterisks*",
		"http://example.com/with(parentheses)",
		"http://example.com/with[brackets]",
		"http://example.com/with{braces}",
		"http://example.com/with`backticks`",
		"http://example.com/with|pipes|",
		"http://example.com/with^carets^",
		"http://example.com/with~tildes~",
		"http://example.com/with#hashes#",
		"http://example.com/with%percents%",
		"http://example.com/with\\backslashes\\",
		"http://example.com/with/slashes/",
		"http://example.com/with?questions?",
		"http://example.com/with&ampersands&",
		"http://example.com/with=equals=",
		"http://example.com/with+pluses+",
		"http://example.com/with,commas,",
		"http://example.com/with;semcolons;",
		"http://example.com/with'apostrophes'",
		"http://example.com/with!exclamations!",
		"http://example.com/with@ats@",
		"http://example.com/with$dollars$",
		"http://example.com/with*asterisks*",
	}

	for _, url := range seedURLs {
		f.Add(url)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Проверяем, что входная строка является валидной UTF-8
		if !utf8.ValidString(input) {
			t.Skip("Input is not valid UTF-8, skipping")
		}

		link, err := vo.NewLink(input)
		if err != nil {
			// Если URL некорректен, ожидаем ошибку
			if link.Validate() == nil {
				t.Errorf("Expected error for invalid URL: %s", input)
			}
		} else {
			// Если URL корректен, проверяем, что значение сохранено правильно
			if link.Value() != input {
				t.Errorf("Expected value %s, got %s", input, link.Value())
			}
			// Проверяем, что URL проходит валидацию
			if err := link.Validate(); err != nil {
				t.Errorf("Expected valid URL, got error: %v", err)
			}
		}
	})
}
