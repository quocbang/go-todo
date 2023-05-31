package utils

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type requestIDKeyType int

const requestIDKey requestIDKeyType = 1

// func Logging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		requestID := xid.New().String()

// 		r = r.WithContext(context.WithValue(r.Context(), requestIDKey, requestID))

// 		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
// 		var respBuf bytes.Buffer
// 		ww.Tee(&respBuf)

// 		defer zap.L().Sync()

// 		if r.Method == "" {
// 			r.Method = http.MethodGet
// 		}

// 		reqLogFields := []zap.Field{zap.String("rid", requestID)}

// 		if !strings.Contains(r.URL.Path, "user/login") {
// 			request, err := httputil.DumpRequest(r, r.Method != http.MethodGet)
// 			reqLogFields = append(reqLogFields,
// 				zap.String("request", string(request)),
// 				zap.NamedError("request_dump_error", err),
// 			)
// 			// 取得db查詢結果時，不寫入response，避免server效能浪費
// 		} else if strings.Contains(r.URL.Path, "/data") {
// 			respBuf.Reset()
// 		}
// 		zap.L().Info("start request", reqLogFields...)

// 		defer func(start time.Time) {
// 			zap.L().Info("server response",
// 				zap.String("remote_address", r.RemoteAddr),
// 				zap.String("rid", requestID),
// 				zap.Int("status", ww.Status()),
// 				zap.String("request_method", r.Method),
// 				zap.String("protocol", r.Proto),
// 				zap.Duration("duration", time.Since(start)),
// 			)
// 		}(time.Now())

// 		next.ServeHTTP(ww, r)
// 	})
// }

// Logging adds request ids and logs responses.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: get request ID from incoming context
		requestID := xid.New().String()
		r = r.WithContext(context.WithValue(r.Context(), requestIDKey, requestID))

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		var respBuf bytes.Buffer
		ww.Tee(&respBuf)

		defer zap.L().Sync()

		if r.Method == "" {
			r.Method = http.MethodGet
		}

		reqLogFields := []zap.Field{zap.String("rid", requestID)}
		// 登入畫面不需要印出請求資訊
		if !strings.Contains(r.URL.Path, "user/login") {
			request, err := httputil.DumpRequest(r, r.Method != http.MethodGet)
			reqLogFields = append(reqLogFields,
				zap.String("request", string(request)),
				zap.NamedError("request_dump_error", err),
			)
			// 取得db查詢結果時，不寫入response，避免server效能浪費
		} else if strings.Contains(r.URL.Path, "/data") {
			respBuf.Reset()
		}
		zap.L().Info("start request", reqLogFields...)

		defer func(start time.Time) {
			zap.L().Info("server responses",
				zap.String("request_method", r.Method),
				zap.String("request_url", r.URL.Path),
				zap.Int("status_code", ww.Status()),
				zap.String("remote_address", r.RemoteAddr),
				zap.String("x-forwarded-for", r.Header.Get("x-forwarded-for")),
				zap.String("rid", requestID),
				zap.String("response", respBuf.String()),
				zap.String("duration", time.Since(start).String()))
		}(time.Now())

		next.ServeHTTP(ww, r)
	})
}
