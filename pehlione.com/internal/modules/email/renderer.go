package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

//go:embed templates/html/*.tmpl
var htmlTemplatesFS embed.FS

//go:embed templates/text/*.tmpl
var textTemplatesFS embed.FS

type Renderer struct {
	baseURL      string
	brandName    string
	supportEmail string
	html         map[string]*template.Template
	text         map[string]*template.Template
}

type RenderResult struct {
	Subject string
	HTML    []byte
	Text    []byte
}

func NewRenderer(baseURL, brandName, supportEmail string) *Renderer {
	funcMap := template.FuncMap{}
	r := &Renderer{
		baseURL:      strings.TrimRight(baseURL, "/"),
		brandName:    brandName,
		supportEmail: supportEmail,
		html:         map[string]*template.Template{},
		text:         map[string]*template.Template{},
	}

	funcMap["trackURL"] = r.trackURL

	r.html = mustLoadTemplates(htmlTemplatesFS, "templates/html", "layout", funcMap)
	r.text = mustLoadTemplates(textTemplatesFS, "templates/text", "layout_text", funcMap)
	return r
}

func mustLoadTemplates(fsys embed.FS, dir, layoutName string, funcMap template.FuncMap) map[string]*template.Template {
	out := map[string]*template.Template{}
	layoutPath := path.Join(dir, "base.tmpl")
	layoutBytes, err := fs.ReadFile(fsys, layoutPath)
	if err != nil {
		panic(err)
	}

	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if entry.Name() == "base.tmpl" {
			continue
		}

		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		contentPath := path.Join(dir, entry.Name())
		contentBytes, err := fs.ReadFile(fsys, contentPath)
		if err != nil {
			panic(err)
		}

		tpl := template.New(name).Funcs(funcMap)
		if _, err := tpl.Parse(string(layoutBytes)); err != nil {
			panic(err)
		}
		if _, err := tpl.Parse(string(contentBytes)); err != nil {
			panic(err)
		}
		parsed := tpl.Lookup(layoutName)
		if parsed == nil {
			panic(fmt.Sprintf("missing layout %s for %s", layoutName, name))
		}
		out[name] = parsed
	}
	return out
}

func (r *Renderer) Render(templateName string, payload map[string]any) (RenderResult, error) {
	htmlTpl, ok := r.html[templateName]
	if !ok {
		return RenderResult{}, fmt.Errorf("unknown HTML template: %s", templateName)
	}
	textTpl, ok := r.text[templateName]
	if !ok {
		return RenderResult{}, fmt.Errorf("unknown text template: %s", templateName)
	}

	data := r.enrichPayload(templateName, payload)
	subject := r.subjectFor(templateName, data)
	data["EmailTitle"] = subject

	var htmlBuf bytes.Buffer
	if err := htmlTpl.Execute(&htmlBuf, data); err != nil {
		return RenderResult{}, err
	}

	var textBuf bytes.Buffer
	if err := textTpl.Execute(&textBuf, data); err != nil {
		return RenderResult{}, err
	}

	return RenderResult{
		Subject: subject,
		HTML:    htmlBuf.Bytes(),
		Text:    textBuf.Bytes(),
	}, nil
}

func (r *Renderer) enrichPayload(templateName string, payload map[string]any) map[string]any {
	data := map[string]any{
		"BrandName":    r.brandName,
		"SupportEmail": r.supportEmail,
		"BaseURL":      r.baseURL,
		"StatusLabel":  "",
		"PreviewText":  defaultPreview(templateName),
	}
	for k, v := range payload {
		data[k] = v
	}
	return data
}

func defaultPreview(templateName string) string {
	switch templateName {
	case TemplateVerifyEmail:
		return "Verify your email to finish setting up your account."
	case TemplateOrderPaid:
		return "Thanks for your purchase! Here are the details."
	case TemplateOrderShipped:
		return "Your order is on the way."
	case TemplateOrderDelivered:
		return "Your package was delivered."
	case TemplateOrderRefunded:
		return "We processed your refund."
	case TemplatePasswordReset:
		return "Reset your password securely."
	default:
		return ""
	}
}

func (r *Renderer) subjectFor(templateName string, data map[string]any) string {
	orderID := shortOrderID(data["OrderID"])
	switch templateName {
	case TemplateVerifyEmail:
		return "Verify your email"
	case TemplateOrderPaid:
		if orderID != "" {
			return fmt.Sprintf("Order %s confirmed", orderID)
		}
		return "Your order is confirmed"
	case TemplateOrderShipped:
		if orderID != "" {
			return fmt.Sprintf("Order %s is on the way", orderID)
		}
		return "Your order is on the way"
	case TemplateOrderDelivered:
		if orderID != "" {
			return fmt.Sprintf("Order %s delivered", orderID)
		}
		return "Your order was delivered"
	case TemplateOrderRefunded:
		if orderID != "" {
			return fmt.Sprintf("Refund processed for %s", orderID)
		}
		return "Refund processed"
	case TemplatePasswordReset:
		return "Reset your password"
	default:
		return "Notification"
	}
}

func shortOrderID(v any) string {
	s, _ := v.(string)
	if len(s) <= 8 {
		return s
	}
	return s[:8]
}

func (r *Renderer) trackURL(raw string, campaign string) string {
	if raw == "" {
		return r.baseURL
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	if parsed.Scheme == "" {
		base, err := url.Parse(r.baseURL)
		if err != nil {
			return raw
		}
		if strings.HasPrefix(parsed.Path, "/") {
			base.Path = parsed.Path
		} else {
			base.Path = "/" + parsed.Path
		}
		base.RawQuery = parsed.RawQuery
		parsed = base
	}

	q := parsed.Query()
	if _, ok := q["utm_source"]; !ok {
		q.Set("utm_source", "email")
	}
	if _, ok := q["utm_medium"]; !ok {
		q.Set("utm_medium", "transactional")
	}
	if campaign != "" {
		q.Set("utm_campaign", campaign)
	}
	parsed.RawQuery = q.Encode()
	return parsed.String()
}
