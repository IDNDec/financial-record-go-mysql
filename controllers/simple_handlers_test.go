package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"financial-record/config"
)

// Helper: buat request dengan session login/guest
func reqWithSessionSimple(t *testing.T, method, target string, loggedIn bool, id string) *http.Request {
	t.Helper()
	r1 := httptest.NewRequest(method, target, nil)
	rr1 := httptest.NewRecorder()
	s, _ := config.Store.Get(r1, config.SESSION_ID)
	if loggedIn {
		s.Values["LOGGED_IN"] = true
		s.Values["ID"] = id
	}
	s.Save(r1, rr1)
	cookie := rr1.Header().Get("Set-Cookie")
	r2 := httptest.NewRequest(method, target, nil)
	if cookie != "" {
		r2.Header.Set("Cookie", cookie)
	}
	return r2
}

func TestSimple_Endpoints_Accessible(t *testing.T) {
	ac := NewAuthController(nil)
	fc := NewFinancialController(nil)
	uc := NewUserController(nil)

	tests := []struct {
		name    string
		method  string
		url     string
		handler http.HandlerFunc
		login   bool
	}{
		{"register page", http.MethodGet, "/register", config.GuestOnly(ac.Register), false},
		{"login page", http.MethodGet, "/login", config.GuestOnly(ac.Login), false},
		{"logout", http.MethodGet, "/logout", config.AuthOnly(ac.Logout), true},
		{"home", http.MethodGet, "/home", config.AuthOnly(fc.Home), true},
		{"financial download", http.MethodGet, "/financial/download_financial_record", config.AuthOnly(fc.DownloadFinancialRecord), true},
		{"financial add", http.MethodGet, "/financial/add_financial_record", config.AuthOnly(fc.AddFinacialRecord), true},
		{"financial edit", http.MethodGet, "/financial/edit_financial_record?id=1", config.AuthOnly(fc.EditFinancialRecord), true},
		{"financial delete", http.MethodGet, "/financial/delete_financial_record?id=1", config.AuthOnly(fc.DeleteFinancialRecord), true},
		{"profile", http.MethodGet, "/profile", config.AuthOnly(uc.Profile), true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := reqWithSessionSimple(t, tc.method, tc.url, tc.login, "userid-test")
			rr := httptest.NewRecorder()
			tc.handler.ServeHTTP(rr, req)
			if tc.name == "logout" || tc.name == "financial delete" {
				if rr.Code != http.StatusSeeOther {
					t.Fatalf("expected redirect (303) for %s, got %d", tc.url, rr.Code)
				}
			} else {
				if rr.Code != http.StatusOK {
					t.Fatalf("expected 200 for %s, got %d", tc.url, rr.Code)
				}
			}
		})
	}
}
