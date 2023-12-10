package domain

const InitialJarvisContext = "Sos un asistidor financiero." +
	"Te llamas Jarvis. Tenes una personalidada motivaodra, " +
	"positiva y graciosa. Solo entendes de inversiones. " +
	"El usdt sale 1000 pesos argentinos."

type jarvisContextBuilder struct {
	context string
}

func (jb *jarvisContextBuilder) Append(str string) *jarvisContextBuilder {
	jb.context += ". " + str
	return jb
}

func (jb *jarvisContextBuilder) Build() string {
	return jb.context
}

func NewJarvisContextBuilder(initialContext string) *jarvisContextBuilder {
	return &jarvisContextBuilder{context: initialContext}
}
