package main

import (
	"log"
	"os"

	authHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/auth"
	courseHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/course"
	examHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/exam"
	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/auth"
	courseSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/course"
	examSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/exam"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *App) RegisterRoute() {
	auth, course, exam := GetHandler(app.config)

	metadata := MetaData{
		url:          "/auth/login",
		requiredAuth: false,
		method:       POST,
	}
	app.AddHdr(auth.Login, metadata)

	metadata = MetaData{
		url:          "/auth/revoke",
		requiredAuth: true,
		method:       GET,
	}
	app.AddHdr(auth.Revoke, metadata)

	metadata = MetaData{
		url:          "/auth/me",
		requiredAuth: true,
		method:       GET,
	}
	app.AddHdr(auth.Me, metadata)

	metadata = MetaData{
		url:          "/course",
		requiredAuth: false,
		method:       GET,
	}
	app.AddHdr(course.GetAllCourseSummary, metadata)

	metadata = MetaData{
		url:          "/course/:course_id",
		requiredAuth: false,
		method:       GET,
	}
	app.AddHdr(course.GetCourse, metadata)

	metadata = MetaData{
		url:          "/exam/:exam_id",
		requiredAuth: false,
		method:       GET,
	}
	app.AddHdr(exam.GetExam, metadata)
}

func GetHandler(conf *config.Config) (authHdr.Handler, courseHdr.Handler, examHdr.Handler) {
	conn, err := grpc.Dial(conf.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	authClient := proto.NewAuthClient(conn)
	authService := authSrv.NewService(authClient)
	auth := authHdr.NewHandler(&authService)

	conn, err = grpc.Dial(conf.Course.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	courseClient := proto.NewCourseClient(conn)
	courseService := courseSrv.NewService(courseClient)
	course := courseHdr.NewHandler(&courseService)

	conn, err = grpc.Dial(conf.Exam.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	examClient := proto.NewExamClient(conn)
	examService := examSrv.NewService(examClient)
	exam := examHdr.NewHandler(&examService)

	return auth, course, exam
}
