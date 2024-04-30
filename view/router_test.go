package view_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	// "net/http/httptest"
	"testing"

	// "github.com/Samurai1986/auth-service/model"
	// "github.com/gin-gonic/gin"
)

// func TestRegister(t *testing.T){
// 	//test empty register
// 	requestbody :=`
// 	{
// 		"email":"",
//         "password":""
//         "first_name":"",
//         "last_name":""
// 	}`
// 	ctx, r := gin.CreateTestContext([]byte)
// }

//autogen by tabnine (not working yet)
// i think this should be: tests := []struct {} where body is a json string
// then in test function we send the json string to the server 
//and check the result by the code and error message
//TODO: rewrite and check
func TestRouter(t *testing.T) {
    // r := gin.Default()

    tests := []struct {
        name    string
        method  string
        path    string
        body    string
        want    int
        wantErr bool
    }{
        {
            name:   "Create User",
            method: http.MethodPost,
            path:   "http://localhost:8000/api/v1/auth-service/sign-up",
            body: `{
                "email":"test123@test",
                "password":"password",
				"first_name":"test",
                "last_name":"test"
            }`,
            want:    http.StatusCreated,
            wantErr: false,
        },
		{
			name: "Empty request",
			method: http.MethodPost,
            path: "/api/v1/auth-service/sign-up",
            body: `{
				"email": "",
				"password": "",
				"first_name": "",
				"last_name": ""
			}`,
			want:    http.StatusBadRequest,
            wantErr: true,
		},
        {
            name:   "Wrong password",
            method: http.MethodPost,
            path:   "/api/v1/auth-service/sign-in",
            body: `{
                "email":    "test123@test",
                "password": "wrongpassword",
            }`,
            want:    http.StatusUnauthorized,
            wantErr: true,
        },        
        {
            name:   "Login User",
            method: http.MethodPost,
            path:   "/api/v1/auth-service/sign-in",
            body: `{
                "email":    "test123@test",
                "password": "password",
            }`,
            want:    http.StatusOK,
            wantErr: false,
        },
        // {
        //     name:   "Read User",
        //     method: http.MethodGet,
        //     path:   "/api/v1/auth-service/me",
        //     body: &model.UserDTO{
        //         Email: "test@test",
        //     },
        //     want:    http.StatusOK,
        //     wantErr: false,
        // },
        // {
        //     name:   "Update User",
        //     method: http.MethodPut,
        //     path:   "/api/v1/auth-service/update",
        //     body: &model.RegisterDTO{
        //         Email:    "test@test",
        //         Password: "password",
        //     },
        //     want:    http.StatusOK,
        //     wantErr: false,
        // },
        // {
        //     name:   "Delete User",
        //     method: http.MethodDelete,
        //     path:   "/api/v1/auth-service/delete",
        //     body: &model.RegisterDTO{
        //         Email: "test@test",
        //     },
        //     want:    http.StatusOK,
        //     wantErr: false,
        // },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, err := http.NewRequest(tt.method, tt.path, nil)
            if err!= nil {
                t.Fatal(err)
            }

            if tt.body!= "" {
                body, err := json.Marshal(tt.body) //why?
                if err!= nil {
                    t.Fatal(err)
                }
                req.Body = io.NopCloser(bytes.NewReader(body))
                req.ContentLength = int64(len(body))
                req.Header.Set("Content-Type", "application/json")
            }

            // w := httptest.NewRecorder()
            if tt.method == http.MethodPost {
               res, err := http.Post(tt.path, req.Header.Get("Content-Type"), bytes.NewBufferString(tt.body))
               if err!= nil {
                   t.Fatal(err)
               }
               if res.StatusCode != tt.want {
                    t.Errorf("got wrong status code: want=%v, got=%v", tt.want, res.StatusCode)
                    // t.Errorf("error: %v", )
               }
            }
            // r.ServeHTTP(w,req)
			// if w.Code!= tt.want {
            //     t.Errorf("got wrong status code: want=%v, got=%v", tt.want, w.Code)
            //     t.Errorf("error: %v", w.Body.String())
            // }
        })
    }
}