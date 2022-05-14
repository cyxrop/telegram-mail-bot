package mail

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseType(t *testing.T) {
	testCases := []struct {
		mail string
		t    Type
	}{
		{
			mail: "test123@gmail.com",
			t:    TypeGoogle,
		},
		{
			mail: "test123@mail.ru",
			t:    TypeMailRu,
		},
		{
			mail: "test123@yandex.ru",
			t:    TypeYandex,
		},
		{
			mail: "test123@outlook.com",
			t:    TypeUnknown,
		},
		{
			mail: "invalid",
			t:    TypeUnknown,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s - %d", tc.mail, tc.t), func(t *testing.T) {
			assert.Equal(t, tc.t, ParseType(tc.mail))
		})
	}
}
