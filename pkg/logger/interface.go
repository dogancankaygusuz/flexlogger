package logger

// Logger interface'i, uygulamamızın kullanacağı metodları belirler.
// Dependency Injection yaparken bu interface'i kullanacağız.
type Logger interface {
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}
