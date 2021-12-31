package redis

import (
	"espad_task/build/messages"
	"espad_task/domain/url"
	"espad_task/pkg/derrors"
	"espad_task/pkg/random"
	"fmt"
	"testing"
	"time"
)

func TestAddUrl(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	tests := []struct {
		name    string
		url     *url.Url
		wantErr error
	}{
		{
			name: "add shortUrl",
			url: &url.Url{
				OriginalLink: "https://www.educative.io/courses/grokking-the-system-design-interview/m2ygV4E81AR",
				ShortLink:    random.String(8),
				Expiration:   100 * time.Second,
			},
			wantErr: nil,
		},
		{
			name: "empty link",
			url: &url.Url{
				OriginalLink: "",
				ShortLink:    random.String(8),
				Expiration:   100 * time.Second,
			},
			wantErr: nil,
		},
		{
			name: "empty short link",
			url: &url.Url{
				OriginalLink: random.String(50),
				ShortLink:    "",
				Expiration:   100 * time.Second,
			},
			wantErr: nil,
		},
		{
			name: "duplicate empty short link",
			url: &url.Url{
				OriginalLink: random.String(50),
				ShortLink:    "",
				Expiration:   100 * time.Second,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := repoTest.AddUrl(tt.url); err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
		})
	}

}

func TestExistShortUrl(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	shortLink1 := random.String(9)
	err := repoTest.AddUrl(&url.Url{
		OriginalLink: "",
		ShortLink:    shortLink1,
		Expiration:   0,
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		shortUrl string
		exist    bool
		wantErr  error
	}{
		{
			name:     "exist short shortUrl",
			shortUrl: shortLink1,
			exist:    true,
			wantErr:  nil,
		},
		{
			name:     "empty short shortUrl",
			shortUrl: "",
			exist:    false,
			wantErr:  nil,
		},
		{
			name:     "not exist short shortUrl",
			shortUrl: random.String(9),
			exist:    false,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exist, err := repoTest.ExistShortUrl(tt.shortUrl)
			if err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			} else if exist != tt.exist {
				t.Fatalf("got result: %v, want result: %v", exist, tt.exist)
			}
		})
	}

}

func TestGetUrl(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	shortLink1 := random.String(9)
	link1 := random.String(50)
	err := repoTest.AddUrl(&url.Url{
		OriginalLink: link1,
		ShortLink:    shortLink1,
		Expiration:   0,
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		shortUrl string
		linkUrl  string
		wantErr  error
	}{
		{
			name:     "exist shortUrl",
			shortUrl: shortLink1,
			linkUrl:  link1,
			wantErr:  nil,
		},
		{
			name:     "not exist shortUrl",
			shortUrl: random.String(9),
			linkUrl:  "",
			wantErr:  derrors.New(derrors.NotFound, messages.UrlNotFound),
		},
		{
			name:     "empty shortUrl",
			shortUrl: "",
			linkUrl:  "",
			wantErr:  derrors.New(derrors.NotFound, messages.UrlNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkUrl, err := repoTest.GetUrl(tt.shortUrl)
			if err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			} else if linkUrl != tt.linkUrl {
				t.Fatalf("got result: %v, want result: %v", linkUrl, tt.linkUrl)
			}
		})
	}

	t.Run("expire key", func(t *testing.T) {
		url1 := &url.Url{
			OriginalLink: random.String(9),
			ShortLink:    random.String(50),
			Expiration:   1 * time.Second,
		}
		err := repoTest.AddUrl(url1)
		if err != nil {
			t.Fatal(err)
		}

		linkUrl, err := repoTest.GetUrl(url1.ShortLink)
		if err != nil {
			t.Fatal(err)
		} else if linkUrl != url1.OriginalLink {
			t.Fatalf("got result %v, want resutl %v", linkUrl, url1)
		}

		time.Sleep(2 * url1.Expiration)

		linkUrl, err = repoTest.GetUrl(url1.ShortLink)
		wantErr := derrors.New(derrors.NotFound, messages.UrlNotFound)
		if err != wantErr {
			t.Fatalf("got error: %v, want error: %v", err, wantErr)
		} else if linkUrl != "" {
			t.Fatalf("got result %v, want resutl %v", linkUrl, url1)
		}

	})

}

func TestDeleteUrlName(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	shortLink1 := random.String(9)
	err := repoTest.AddUrl(&url.Url{
		OriginalLink: random.String(50),
		ShortLink:    shortLink1,
		Expiration:   0,
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		shortUrl string
		wantErr  error
	}{
		{
			shortUrl: shortLink1,
			wantErr:  nil,
		},
		{
			shortUrl: random.String(9),
			wantErr:  nil,
		},
		{
			shortUrl: "",
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("delete %v", tt.shortUrl)
		t.Run(name, func(t *testing.T) {
			err := repoTest.DeleteUrl(tt.shortUrl)
			if err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
		})
	}

}
