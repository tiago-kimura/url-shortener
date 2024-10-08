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

type MaxLengthRule struct {
	MaxLength int
}

func (r *MaxLengthRule) ProcessRule(urlShortener UrlShortener) error {
	if len(urlShortener.UrlOriginal) > r.MaxLength {
		return errors.New("URL exceeds maximum length") // TODO: check
	}
	return nil
}

func ProcessRule(urlShortener UrlShortener) error {
	if len(strings.Trim(urlShortener.UrlOriginal, " ")) == 0 {
		return errors.New("URL cannot be null")
	}
	return nil
}
