package locales

import (
	"log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
    "github.com/BurntSushi/toml"
)
var Bundle *i18n.Bundle 

func init() {
	Bundle = i18n.NewBundle(language.Hebrew)
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	Bundle.MustLoadMessageFile("locales/en.toml")
	Bundle.MustLoadMessageFile("locales/he.toml")


	_, err := Bundle.LoadMessageFile("locales/he.toml")
	if err != nil {
		log.Fatalf("failed to load locale file: %v", err)
	}
	
	_, err = Bundle.LoadMessageFile("locales/en.toml")
	if err != nil {
		log.Fatalf("failed to load locale file: %v", err)
	}	

	log.Println("Locale files loaded successfully")
}

// T translates a message key to the given language
func T(lang, messageID string, templateData ...map[string]interface{}) string {
    // Create localizer for the requested language
    localizer := i18n.NewLocalizer(Bundle, lang)
    
    // Configure the localization
    config := &i18n.LocalizeConfig{
        MessageID: messageID,
    }
    
    // Add template data if provided
    if len(templateData) > 0 {
        config.TemplateData = templateData[0]
    }
    
    // Localize the message
    msg, err := localizer.Localize(config)
    if err != nil {
        // Return the messageID if translation not found
        log.Printf("Translation not found for key '%s' in language '%s': %v", messageID, lang, err)
        return messageID
    }
    
    return msg
}

// GetLocalizer returns a localizer for the given language
func GetLocalizer(lang string) *i18n.Localizer {
    return i18n.NewLocalizer(Bundle, lang)
}