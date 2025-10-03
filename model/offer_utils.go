package models

import (
	"strings"
)

// OfferUtils provides utility functions for working with offers

// AddTag adds a tag to the offer's tags
func (o *Offer) AddTag(tag string) {
	if o.Tags == nil {
		o.Tags = &JSON{}
	}

	// Convert to map if it's not already
	tagsMap := make(map[string]interface{})
	if o.Tags != nil {
		if tagsInterface, ok := o.Tags.Data.(map[string]interface{}); ok {
			for k, v := range tagsInterface {
				tagsMap[k] = v
			}
		}
	}

	// Add the tag (using tag as both key and value for uniqueness)
	cleanTag := strings.TrimSpace(strings.ToLower(tag))
	if cleanTag != "" {
		tagsMap[cleanTag] = cleanTag
	}

	o.Tags.Data = tagsMap
}

// RemoveTag removes a tag from the offer's tags
func (o *Offer) RemoveTag(tag string) {
	if o.Tags == nil {
		return
	}

	cleanTag := strings.TrimSpace(strings.ToLower(tag))
	tagsMap := make(map[string]interface{})

	if tagsInterface, ok := o.Tags.Data.(map[string]interface{}); ok {
		for k, v := range tagsInterface {
			if k != cleanTag {
				tagsMap[k] = v
			}
		}
	}

	o.Tags.Data = tagsMap
}

// HasTag checks if the offer has a specific tag
func (o *Offer) HasTag(tag string) bool {
	if o.Tags == nil {
		return false
	}

	cleanTag := strings.TrimSpace(strings.ToLower(tag))
	if tagsInterface, ok := o.Tags.Data.(map[string]interface{}); ok {
		_, exists := tagsInterface[cleanTag]
		return exists
	}
	return false
}

// GetTags returns all tags as a slice of strings
func (o *Offer) GetTags() []string {
	if o.Tags == nil {
		return []string{}
	}

	var tags []string
	if tagsInterface, ok := o.Tags.Data.(map[string]interface{}); ok {
		for k := range tagsInterface {
			tags = append(tags, k)
		}
	}
	return tags
}

// SetTags sets multiple tags at once
func (o *Offer) SetTags(tags []string) {
	if len(tags) == 0 {
		o.Tags = nil
		return
	}

	tagsMap := make(map[string]interface{})
	for _, tag := range tags {
		cleanTag := strings.TrimSpace(strings.ToLower(tag))
		if cleanTag != "" {
			tagsMap[cleanTag] = cleanTag
		}
	}

	o.Tags = &JSON{Data: tagsMap}
}

// IsValidLink checks if the link is valid (either https:// or relative path)
func (o *Offer) IsValidLink() bool {
	if o.Link == nil || *o.Link == "" {
		return true // null/empty links are considered valid
	}

	link := strings.TrimSpace(*o.Link)

	// Check if it's a relative path (starts with /)
	if strings.HasPrefix(link, "/") {
		return true
	}

	// Check if it's an https:// URL
	if strings.HasPrefix(link, "https://") {
		return true
	}

	// Check if it's an http:// URL (for backward compatibility)
	if strings.HasPrefix(link, "http://") {
		return true
	}

	return false
}

// NormalizeLink normalizes the link format
func (o *Offer) NormalizeLink() {
	if o.Link == nil || *o.Link == "" {
		return
	}

	link := strings.TrimSpace(*o.Link)

	// If it's a relative path, ensure it starts with /
	if !strings.HasPrefix(link, "/") && !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "/" + link
	}

	o.Link = &link
}

// GetLinkType returns the type of link (relative, https, http, or empty)
func (o *Offer) GetLinkType() string {
	if o.Link == nil || *o.Link == "" {
		return "empty"
	}

	link := strings.TrimSpace(*o.Link)

	if strings.HasPrefix(link, "/") {
		return "relative"
	}

	if strings.HasPrefix(link, "https://") {
		return "https"
	}

	if strings.HasPrefix(link, "http://") {
		return "http"
	}

	return "invalid"
}
