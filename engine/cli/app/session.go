package app

import (
	pb "engine/lib/structs"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"google.golang.org/grpc/metadata"
	"log"

	"os"
)

func (app *App) SignInWithRequest(admin bool) error {
	var name, pass string
	fmt.Print("Введите логин: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите пароль: ")
	fmt.Fscan(os.Stdin, &pass)

	return app.StartSession(name, pass, admin)
}

func (app *App) StartSession(name, pass string, admin bool) error {
	var err error
	if app.Session, err = app.srvc_auth.StartSession(app.ctx, &pb.Authenticate{Name: name, Password: pass}); err != nil {
		return fmt.Errorf("StartSession: %s ", err)
	}
	log.Printf("Токен сессии: %s", app.Session.Token)

	decoded, err := app.srvc_auth.DecodeSession(app.ctx, app.Session)
	if err != nil {
		return fmt.Errorf("StartSession: %s ", err)
	}

	app.ctx = metadata.AppendToOutgoingContext(app.ctx, "pid", decoded.Id)

	if admin {
		app.ctx = metadata.AppendToOutgoingContext(app.ctx, "admin-id", decoded.Id)
	}

	app.U = decoded

	return nil
}

func (app *App) CognitoSignUp(email, name, password string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		return err
	}

	cg := cognitoidentityprovider.New(sess, aws.NewConfig())

	signUpInput := cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("APP_CLIENT_ID")),
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: aws.String("name"), Value: aws.String(name)},
			{Name: aws.String("nickname"), Value: aws.String(name)},
			{Name: aws.String("picture"), Value: aws.String("---")},
			{Name: aws.String("birthdate"), Value: aws.String("01.11.2018")},
		},
	}

	signUpOut, err := cg.SignUp(&signUpInput)
	if err != nil {
		return err
	}
	log.Printf("%#v", signUpOut)

	return nil
}

//func (app *App) CognitoSession(email, password string) error  {
//	sess, err := session.NewSession(&aws.Config{
//		Region: aws.String("eu-central-1"),
//		Credentials: credentials.NewEnvCredentials(),
//	})
//	if err != nil {
//		return err
//	}
//	cg := cognitoidentityprovider.New(sess, aws.NewConfig())
//
//	signInInput := cognitoidentityprovider.AdminInitiateAuthInput{
//		AuthFlow: aws.String("ADMIN_NO_SRP_AUTH"),
//		UserPoolId: aws.String(os.Getenv("COGNITO_USER_POOL_ID")),
//		ClientId: aws.String(os.Getenv("APP_CLIENT_ID")),
//		AuthParameters: map[string]*string{
//			"USERNAME": aws.String(email),
//			"PASSWORD": aws.String(password),
//		},
//
//	}
//	out, err := cg.AdminInitiateAuth(&signInInput)
//	if err != nil {
//		return err
//	}
//
//	log.Println("Use congito user pool")
//	_cognito_region := os.Getenv("COGNITO_REGION")
//	_cognito_userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
//	// 1. Download and store the JSON Web Key (JWK) for your user pool.
//	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", _cognito_region, _cognito_userPoolID)
//	fmt.Println(jwkURL)
//	_cognito_jwk := verifytoken.GetJWK(jwkURL)
//	//log.Printf("Cognito JWK: \n %+v", _cognito_jwk)
//
//	tokenStr := out.AuthenticationResult.AccessToken
//	jwt, err := verifytoken.ValidateToken(aws.StringValue(tokenStr),os.Getenv("COGNITO_REGION"), os.Getenv("COGNITO_USER_POOL_ID"),_cognito_jwk)
//	if err != nil {
//		return err
//	}
//	log.Printf("JWT: %#v", jwt)
//
//
//	//log.Printf("%#v", out)
//	return nil
//}
