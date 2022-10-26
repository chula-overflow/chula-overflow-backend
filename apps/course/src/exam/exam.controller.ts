import { Controller, Post, Get, Res, Body, Query } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CourseService } from 'src/course/course.service';
import { IS_MICROSERVICE } from 'src/main';
import {
  ExamCreateBody,
  ExamPropertyRequestBody,
  ExamRequestUpdateBody,
  ExamCourseIdRequestBody,
} from './exam.interface';
import { ExamService } from './exam.service';

@Controller('exam')
export class ExamController {
  constructor(
    private readonly ExamService: ExamService,
    private readonly CourseService: CourseService,
  ) { }

  @Post('/')
  @GrpcMethod('Exam')
  async createExam(
    @Res() response,
    @Body() reqBody: ExamCreateBody,
    data: ExamCreateBody,
    metadata: any,
  ) {
    const newExam = await this.ExamService.create(
      IS_MICROSERVICE ? data : reqBody,
    );
    const examId = newExam._id;
    const courseId = newExam.course_id;

    // add exam_id to course document
    await this.CourseService.addExamId(examId, courseId);

    return IS_MICROSERVICE ? newExam : response.status(201).json(newExam);
  }

  @Get('/')
  @GrpcMethod('Exam')
  async getExams(
    @Res() response,
    @Query() query: ExamPropertyRequestBody,
    data: ExamPropertyRequestBody,
    metadata: any,
  ) {
    console.log(data)

    if ((query && Object.keys(query).length != 0) || data) {
      const exam = await this.ExamService.findByCourseProperty(
        IS_MICROSERVICE ? data : query,
      );
      console.log(exam);
      return IS_MICROSERVICE
        ? { messages: exam }
        : response.status(200).json(exam);
    } else {
      const exams = await this.ExamService.find();
      console.log(exams);
      return IS_MICROSERVICE
        ? { messages: exams }
        : response.status(200).json(exams);
    }
  }

  @GrpcMethod('Exam')
  async deleteExamByCourseProperty(
    data: ExamPropertyRequestBody,
    metadata: any,
  ) {
    const examId = await this.ExamService.findIdByCourseProperty(data);

    const deletedExam = await this.ExamService.deleteById(examId);

    return deletedExam;
  }
}
