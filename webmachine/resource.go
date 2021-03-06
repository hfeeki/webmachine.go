package webmachine

import (
    "io"
    "net/http"
    "time"
)

type DefaultRequestHandler struct{}

func NewDefaultRequestHandler() *DefaultRequestHandler {
    return new(DefaultRequestHandler)
}

func (p *DefaultRequestHandler) StartRequest(req Request, cxt Context) (Request, Context) {
    return req, cxt
}

func (p *DefaultRequestHandler) ServiceAvailable(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) AllowedMethods(req Request, cxt Context) ([]string, Request, Context, int, error) {
    return []string{GET, HEAD}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) URITooLong(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) MalformedRequest(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) IsAuthorized(req Request, cxt Context) (bool, string, Request, Context, int, error) {
    return true, "", req, cxt, 0, nil
}

func (p *DefaultRequestHandler) Forbidden(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ValidContentHeaders(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) KnownContentType(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ValidEntityLength(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) Options(req Request, cxt Context) ([]string, Request, Context, int, error) {
    return []string{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ContentTypesProvided(req Request, cxt Context) ([]MediaTypeHandler, Request, Context, int, error) {
    return []MediaTypeHandler{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ContentTypesAccepted(req Request, cxt Context) ([]MediaTypeInputHandler, Request, Context, int, error) {
    return []MediaTypeInputHandler{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) IsLanguageAvailable(language []string, req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) CharsetsProvided(charsets []string, req Request, cxt Context) ([]CharsetHandler, Request, Context, int, error) {
    return []CharsetHandler{NewStandardCharsetHandler("utf-8")}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) EncodingsProvided(encodings []string, req Request, cxt Context) ([]EncodingHandler, Request, Context, int, error) {
    return []EncodingHandler{NewIdentityEncoder(), NewDeflateEncoder()}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) Variances(req Request, cxt Context) ([]string, Request, Context, int, error) {
    return []string{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ResourceExists(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) LastModified(req Request, cxt Context) (time.Time, Request, Context, int, error) {
    return time.Time{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) PreviouslyExisted(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) GenerateETag(req Request, cxt Context) (string, Request, Context, int, error) {
    return "", req, cxt, 0, nil
}

func (p *DefaultRequestHandler) MovedTemporarily(req Request, cxt Context) (string, Request, Context, int, error) {
    return "", req, cxt, 0, nil
}

func (p *DefaultRequestHandler) MovedPermanently(req Request, cxt Context) (string, Request, Context, int, error) {
    return "", req, cxt, 0, nil
}

func (p *DefaultRequestHandler) AllowMissingPost(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) DeleteResource(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) DeleteCompleted(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) PostIsCreate(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) CreatePath(req Request, cxt Context) (string, Request, Context, int, error) {
    return "", req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ProcessPost(req Request, cxt Context) (Request, Context, int, http.Header, io.WriterTo, error) {
    return req, cxt, 0, nil, nil, nil
}

func (p *DefaultRequestHandler) IsConflict(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) Expires(req Request, cxt Context) (time.Time, Request, Context, int, error) {
    return time.Time{}, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) MultipleChoices(req Request, cxt Context) (bool, http.Header, Request, Context, int, error) {
    return false, nil, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) FinishRequest(req Request, cxt Context) (bool, Request, Context, int, error) {
    return true, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) ResponseIsRedirect(req Request, cxt Context) (bool, Request, Context, int, error) {
    return false, req, cxt, 0, nil
}

func (p *DefaultRequestHandler) HasRespBody(req Request, cxt Context) bool {
    return false
}
