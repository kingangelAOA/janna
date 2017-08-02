package performance

import "net/http"

//Fetch 请求总接口
type Fetch interface {
	request() interface{}
}

//Request http request
func (o *Interface) Request() (*http.Response, error) {
	client := &http.Client{}
	body, err := o.getContentByType()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(o.Method, o.URL, body)
	if err != nil {
		return nil, err
	}
	for k, v := range o.Headers {
		req.Header.Set(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
