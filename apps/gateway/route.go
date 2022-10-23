package main

import (
	"log"
	"os"

	authHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/auth"
	courseHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/course"
	examHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/exam"
	threadHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/thread"
	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/auth"
	courseSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/course"
	examSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/exam"
	threadSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/thread"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *App) RegisterRoute() {
	auth, course, exam, thread := GetHandler(app.config)

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

	metadata = MetaData{
		url:          "/thread/:thread_id",
		requiredAuth: false,
		method:       GET,
	}
	app.AddHdr(thread.GetThread, metadata)

	metadata = MetaData{
		url:          "/thread/post",
		requiredAuth: true,
		method:       GET,
	}
	app.AddHdr(thread.CreateThread, metadata)

	metadata = MetaData{
		url:          "/thread/:thread_id/reply",
		requiredAuth: true,
		method:       POST,
	}
	app.AddHdr(thread.CreateReply, metadata)
}

func GetHandler(conf *config.Config) (authHdr.Handler, courseHdr.Handler, examHdr.Handler, threadHdr.Handler) {
	conn, err := grpc.Dial(conf.AuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	authClient := proto.NewAuthClient(conn)
	authService := authSrv.NewService(authClient)
	auth := authHdr.NewHandler(&authService)

	conn, err = grpc.Dial(conf.CourseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	courseClient := proto.NewCourseClient(conn)
	courseService := courseSrv.NewService(courseClient)
	course := courseHdr.NewHandler(&courseService)

	examClient := proto.NewExamClient(conn)
	examService := examSrv.NewService(examClient)
	exam := examHdr.NewHandler(&examService)

	threadClient := proto.NewThreadClient(conn)
	threadService := threadSrv.NewService(threadClient)
	thread := threadHdr.NewHandler(&threadService)

	return auth, course, exam, thread
}
