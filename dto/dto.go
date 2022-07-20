package dto

// DTO is the common struct for every service: they talk only through DTO, without
// actually knowing any other type. Clients will use their own payloads and they'll
// be sent and converted to the services by the channel type.
type DTO struct {
	Foo string
	Bar int
}
