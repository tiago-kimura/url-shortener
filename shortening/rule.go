package shortening

import (
	"errors"
	"strings"
)

type URLRule interface {
	ProcessRule(urlShortener UrlShortener) error
}

type CompositeRule struct {
	rules []URLRule
}

func NewCompositeRule(rules ...URLRule) *CompositeRule {
	return &CompositeRule{rules: rules}
}

func (r *CompositeRule) ProcessRules(urlShortener UrlShortener) error {
	for _, rule := range r.rules {
		if err := rule.ProcessRule(urlShortener); err != nil {
			return err
		}
	}
	return nil
}

type HashExistsRule struct {
	Repository Repository
}

func (r *HashExistsRule) ProcessRule(urlShortener UrlShortener) error {
	existUrl, err := r.Repository.GetByUrlId(urlShortener.UrlId)
	if err != nil {
		return err
	}
	if existUrl.UrlOriginal != "" {
		return errors.New("URL hash already exists")
	}
	return nil
}

type MinLengthRule struct {
	MinLength int
}

func (r *MinLengthRule) ProcessRule(urlShortener UrlShortener) error {
	if len(urlShortener.UrlOriginal) < r.MinLength {
		return errors.New("URL does not have the minimum size")
	}
	return nil
}

type ValidUrl struct {
}

func (r *ValidUrl) ProcessRule(urlShortener UrlShortener) error {
	if len(strings.TrimSpace(urlShortener.UrlOriginal)) == 0 {
		return errors.New("URL cannot be null")
	}

	urlSlice := strings.Split(urlShortener.UrlOriginal, "https://")
	if len(strings.Trim(urlSlice[1], " ")) == 0 {
		return errors.New("URL cannot be null")
	}
	return nil
}
