package data

type Data struct {
	contentType string
}

func New(contentType string) *Data {
	return &Data{
		contentType: contentType,
	}
}

func (d *Data) GetContentType() string {
	return d.contentType
}
