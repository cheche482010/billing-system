package utils

var Regex = map[string]string{
	"cedula":     `^([0-9]{7,9})$`,
	"caracteres": `^[a-zA-Z]`,
	"rif":        `^([vejpgVEJPG]{1})([0-9]{9})$`,
	"enteros":    `^[0-9]{1,4}$`,
	"dinero":     `^\d{1,3}(?:\.\d\d\d)*(?:,\d{1,2})?$`,
	"password":   `(?=^.{4,}$)((?=.*\\d)|(?=.*\\W+))(?![.\\n])(?=.*[A-Z])(?=.*[a-z]).*$`,
	"telefono":   `0{0,2}([\\+]?[\\d]{1,3}?)?([\\(][\\d]{2,3}\\)]?)?[0-9][0-9 \\-]{6,}(?([xX]|([eE]xt[\\.]?))?([\\d]{1,5}))?`,
	"fechas":     `^(19|20)(((([02468][048])|([13579][26]))-02-29)|(\d{2})-((02-((0[1-9])|1\\d|2[0-8]))|((((0[13456789])|1[012]))-((0[1-9])|((1|2)\\d)|30))|(((0[13578])|(1[02]))-31)))$`,
	"correo":     `^(([A-Za-z0-9]+_+)|([A-Za-z0-9]+\\-+)|([A-Za-z0-9]+\\.+)|([A-Za-z0-9]+\\+\\+))*[A-Za-z0-9]+@((\\w+\\-+)|(\w+\\.))*\\w{1,63}\\.[a-zA-Z]{2,6}$`,
}
