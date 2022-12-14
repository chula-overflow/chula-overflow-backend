package main

import (
	"log"
	"os"

	authHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/auth"
	courseHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/course"
	examHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/exam"
	nlpHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/nlp"
	threadHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/thread"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/middleware"
	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/auth"
	courseSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/course"
	examSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/exam"
	nlpSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/nlp"
	threadSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/thread"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RouteManager struct {
	app        *App
	middleware func(*fiber.Ctx) error
}

type Handler = func(*context.Ctx) error

type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

func (app *App) RegisterRoute() {
	validator := validator.GetValidator()
	authS, courseS, examS, threadS, nlpS := GetService(app.config)
	auth, course, exam, thread, nlp := GetHandler(&authS, &courseS, &examS, &threadS, &nlpS, validator)
	router := NewRouteManager(app, &authS, validator)

	metadata := MetaData{
		url:          "/auth/login",
		requiredAuth: false,
		method:       POST,
	}
	router.AddRoute(auth.Login, metadata)

	metadata = MetaData{
		url:          "/auth/revoke",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(auth.Revoke, metadata)

	metadata = MetaData{
		url:          "/auth/me",
		requiredAuth: true,
		method:       GET,
	}
	router.AddRoute(auth.Me, metadata)

	// course
	metadata = MetaData{
		url:          "/course",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(course.GetAllCourses, metadata)

	metadata = MetaData{
		url:          "/course/:course_id",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(course.GetCourseByCourseId, metadata)

	metadata = MetaData{
		url:          "/course",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(course.CreateCourse, metadata)

	metadata = MetaData{
		url:          "/course/:course_id",
		requiredAuth: true,
		method:       PUT,
	}
	router.AddRoute(course.UpdateCourse, metadata)

	metadata = MetaData{
		url:          "/course/:course_id",
		requiredAuth: true,
		method:       DELETE,
	}
	router.AddRoute(course.DeleteCourse, metadata)

	// exam
	metadata = MetaData{
		url:          "/exam",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(exam.GetExam, metadata)

	metadata = MetaData{
		url:          "/exam",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(exam.CreateExam, metadata)

	metadata = MetaData{
		url:          "/exam",
		requiredAuth: true,
		method:       PUT,
	}
	router.AddRoute(exam.UpdateExamByCourseProperty, metadata)

	// thread
	metadata = MetaData{
		url:          "/thread",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(thread.GetThread, metadata)

	metadata = MetaData{
		url:          "/thread/:thread_id",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(thread.GetThreadById, metadata)

	metadata = MetaData{
		url:          "/thread",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.CreateThread, metadata)

	metadata = MetaData{
		url:          "/thread/:thread_id/upvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.UpvoteThread, metadata)

	metadata = MetaData{
		url:          "/thread/:thread_id/downvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.DownvoteThread, metadata)

	metadata = MetaData{
		url:          "/thread/:thread_id/answer",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.AddAnswer, metadata)

	metadata = MetaData{
		url:          "/answer/:answer_id/upvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.UpvoteAnswer, metadata)

	metadata = MetaData{
		url:          "/answer/:answer_id/downvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.DownvoteAnswer, metadata)

	metadata = MetaData{
		url:          "/problem/:problem_id/upvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.UpvoteProblem, metadata)

	metadata = MetaData{
		url:          "/problem/:problem_id/downvote",
		requiredAuth: true,
		method:       POST,
	}
	router.AddRoute(thread.DownvoteProblem, metadata)

	// Nlp
	metadata = MetaData{
		url:          "/tokenize",
		requiredAuth: false,
		method:       GET,
	}
	router.AddRoute(nlp.Tokenize, metadata)

}

func GetService(conf *config.Config) (authSrv.Service, courseSrv.Service, examSrv.Service, threadSrv.Service, nlpSrv.Service) {
	conn, err := grpc.Dial(conf.AuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	authClient := proto.NewAuthClient(conn)
	authService := authSrv.NewService(authClient)

	conn, err = grpc.Dial(conf.CourseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	courseClient := proto.NewCourseClient(conn)
	courseService := courseSrv.NewService(courseClient)

	examClient := proto.NewExamClient(conn)
	examService := examSrv.NewService(examClient)

	threadClient := proto.NewThreadClient(conn)
	threadService := threadSrv.NewService(threadClient)

	conn, err = grpc.Dial(conf.NlpURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	nlpClient := proto.NewNlpClient(conn)
	nlpService := nlpSrv.NewService(nlpClient)

	return authService, courseService, examService, threadService, nlpService
}

func GetHandler(authService *authSrv.Service, courseService *courseSrv.Service, examService *examSrv.Service, threadService *threadSrv.Service, nlpService *nlpSrv.Service, validator *validator.MyValidator) (authHdr.Handler, courseHdr.Handler, examHdr.Handler, threadHdr.Handler, nlpHdr.Handler) {
	auth := authHdr.NewHandler(authService, validator)

	course := courseHdr.NewHandler(courseService, validator)
	exam := examHdr.NewHandler(examService, validator)
	thread := threadHdr.NewHandler(threadService, validator)
	nlp := nlpHdr.NewHandler(nlpService, validator)

	return auth, course, exam, thread, nlp
}

func NewRouteManager(app *App, authService *authSrv.Service, validator *validator.MyValidator) RouteManager {
	middleware := middleware.CreateMiddleWare(authService, validator)
	return RouteManager{app, middleware}
}

func (r *RouteManager) AddRoute(handler Handler, metadata MetaData) {
	if metadata.requiredAuth {
		r.app.Use(metadata.url, r.middleware)
	}

	switch method := metadata.method; method {
	case GET:
		r.app.Get(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})

	case POST:
		r.app.Post(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})

	case PUT:
		r.app.Put(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})

	case DELETE:
		r.app.Delete(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})
	}
}
