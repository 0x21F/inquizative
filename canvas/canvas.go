package canvas

import "net/http"

type CanvasConn struct {
	token  string
	client http.Client
}

type CanvasError struct {
	why string
}

func (c *CanvasError) Error() string {
	return c.why
}

func (c CanvasConn) baseReq(op, url string) (*http.Request, error) {
	if c.token == "" {
		return nil, &CanvasError{
			why: "no token",
		}
	}

	req, err := http.NewRequest(op, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)

	return req, nil
}
